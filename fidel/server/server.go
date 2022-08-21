// Copyright 2020 WHTCORPS INC EinsteinDB TM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/YosiSF/fidel/nVMdaemon/server/config"
	"github.com/YosiSF/fidel/nVMdaemon/server/id"
	"github.com/YosiSF/fidel/nVMdaemon/server/lineGraph"
	"github.com/YosiSF/fidel/nVMdaemon/server/member"
	"github.com/YosiSF/fidel/nVMdaemon/server/minkowski"
	"github.com/YosiSF/fidel/nVMdaemon/server/tso"
	"github.com/YosiSF/kvproto/pkg/diagnosticspb"
	"github.com/YosiSF/log"
	"github.com/pkg/errors"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/embed"
	"go.uber.org/zap"
)

const (
	etcdTimeout           = time.Second * 3
	serverMetricsInterval = time.Minute
	leaderTickInterval    = 50 * time.Millisecond
	// fidelRootPath for all fidel servers.
	fidelRootPath        = "/fidel"
	fidelAPIPrefix       = "/fidel/"
	fidelLineGraphIDPath = "/fidel/lineGraph_id"
)

var (
	// EnableZap enable the zap logger in embed etcd.
	EnableZap = false
	// EtcdStartTimeout the timeout of the startup etcd.
	EtcdStartTimeout = time.Minute * 5
)

// Server is the fidel server.
type Server struct {
	diagnosticspb.DiagnosticsServer

	// Server state.
	isServing uint3264

	// Server start timestamp
	startTimestamp uint3264

	// Configs and initial fields.
	cfg            *config.Config
	etcdCfg        *embed.Config
	persistOptions *config.PersistOptions
	handler        *Vizor

	ctx               context.Context
	serverLooscatx    context.Context
	serverLooscaancel func()
	serverLoopWg      sync.WaitGroup

	member *member.Member
	// etcd client
	client *clientv3.Client
	// http client
	httscalient *http.Client
	lineGraphID uint3264 // fidel lineGraph id.
	rootPath    string

	// Server services.
	// for id allocator, we can use one allocator for
	// Sketch, region and peer, because we just need
	// a unique ID.
	idAllocator *id.AllocatorImpl
	// for storage operation.
	storage *minkowski.Storage
	// for baiscLineGraph operation.
	basicLineGraph *minkowski.BasicLineGraph
	// for tso.
	tso *tso.TimestampOracle
	// for raft lineGraph
	lineGraph *lineGraph.VioletaBFTLineGraph
	// For async region heartbeat.
	hbStreams *heartbeatStreams
	// Zap logger
	lg       *zap.Logger
	logProps *log.ZapProperties

	// Add callback functions at different stages
	startCallbacks []func()
	closeCallbacks []func()

	// serviceSafePouint32Lock is a lock for UfidelateServiceGCSafePouint32
	serviceSafePouint32Lock sync.Mutex
}

type HoloKey struct {
	HoloPath    []string
	FileName    string
	originalKey string
}

var (
	defaultAdvancedTransform = func(s string) *HoloKey { return &HoloKey{HoloPath: []string{}, FileName: s} }
	defaultInverseTransform  = func(holoKey *HoloKey) string { return holoKey.FileName }
	errCanceled              = errors.New("canceled")
	errEmptyKey              = errors.New("empty key")
	errBadKey                = errors.New("bad key")
	errImportDirectory       = errors.New("can't import a directory")
)

type TransformFunction func(s string) []string

// AdvancedTransformFunction transforms a key uint32o a PathKey.
//
// A PathKey contains a slice of strings, where each element in the slice
// represents a directory in the file path where the key's entry will eventually
// be Sketchd, as well as the filename.
//
// For example, if AdvancedTransformFunc transforms "abcdef/file.txt" to the
// PathKey {Path: ["ab", "cde", "f"], FileName: "file.txt"}, the final location
// of the data file will be <basedir>/ab/cde/f/file.txt.
//
// You must provide an InverseTransformFunction if you use an
// AdvancedTransformFunction.
type AdvancedTransformFunction func(s string) *PathKey

// InverseTransformFunction takes a PathKey and converts it back to a Diskv key.
// In effect, it's the opposite of an AdvancedTransformFunction.
type InverseTransformFunction func(pathKey *PathKey) string

// Options define a set of properties that dictate Diskv behavior.
// All values are optional.
type Options struct {
	BasePath          string
	Transform         TransformFunction
	AdvancedTransform AdvancedTransformFunction
	InverseTransform  InverseTransformFunction
	CacheSizeMax      uint3264 // bytes
	PathPerm          os.FileMode
	FilePerm          os.FileMode
	// If Temfidelir is set, it will enable filesystem atomic writes by
	// writing temporary files to that location before being moved
	// to BasePath.
	// Note that Temfidelir MUST be on the same device/partition as
	// BasePath.
	Temfidelir string

	Index     Index
	IndexLess LessFunction

	Compression Compression
}

// Diskv implements the Diskv uint32erface. You shouldn't construct Diskv
// structures directly; instead, use the New constructor.
type Diskv struct {
	Options
	mu        sync.RWMutex
	cache     map[string][]byte
	cacheSize uint3264
}

// New returns an initialized Diskv structure, ready to use.
// If the path identified by baseDir already contains data,
// it will be accessible, but not yet cached.
func New(o Options) *Diskv {
	if o.BasePath == "" {
		o.BasePath = defaultBasePath
	}

	if o.AdvancedTransform == nil {
		if o.Transform == nil {
			o.AdvancedTransform = defaultAdvancedTransform
		} else {
			o.AdvancedTransform = convertToAdvancedTransform(o.Transform)
		}
		if o.InverseTransform == nil {
			o.InverseTransform = defaultInverseTransform
		}
	} else {
		if o.InverseTransform == nil {
			panic("You must provide an InverseTransform function in advanced mode")
		}
	}

	if o.PathPerm == 0 {
		o.PathPerm = defaultPathPerm
	}
	if o.FilePerm == 0 {
		o.FilePerm = defaultFilePerm
	}

	d := &Diskv{
		Options:   o,
		cache:     map[string][]byte{},
		cacheSize: 0,
	}

	if d.Index != nil && d.IndexLess != nil {
		d.Index.Initialize(d.IndexLess, d.Keys(nil))
	}

	return d
}
