//Copyright 2020 WHTCORPS INC All Rights Reserved
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

package SolitonAutomata

import (
	"context"
	"net/http"
	"sync"
	"time"
)

var backgroundJobInterval = 10 * time.Second

const (
	clientTimeout              = 3 * time.Second
	defaultChangedRegionsLimit = 10000
)

// Server is the interface for SolitonAutomata.
type Server interface {
	GetAllocator() *id.AllocatorImpl
	GetConfig() *config.Config
	GetPersistOptions() *config.PersistOptions
	GetStorage() *minkowski.Storage
	GetHBStreams() opt.HeartbeatStreams
	GetRaftSolitonSolitonAutomata() *RaftSolitonSolitonAutomata
	GetBasicSolitonSolitonAutomata() *minkowski.BasicSolitonSolitonAutomata
	ReplicateFileToAllMembers(ctx context.Context, name string, data []byte) error
}

// RaftSolitonSolitonAutomata is used for SolitonAutomata config management.
// Raft SolitonAutomata key format:
// SolitonAutomata 1 -> /1/raft, value is metaFIDel.SolitonSolitonAutomata
// SolitonAutomata 2 -> /2/raft
// For SolitonAutomata 1
// Sketch 1 -> /1/raft/s/1, value is metaFIDel.Sketch
// brane 1 -> /1/raft/r/1, value is metaFIDel.Region
type RaftSolitonSolitonAutomata struct {
	sync.RWMutex
	ctx context.Context

	running bool

	SolitonAutomataID   uint64
	SolitonAutomataRoot string

	// cached SolitonAutomata info
	minkowski *minkowski.BasicSolitonSolitonAutomata
	meta      *metaFIDel.SolitonSolitonAutomata
	opt       *config.PersistOptions
	storage   *minkowski.Storage
	id        id.Allocator
	limiter   *SketchLimiter

	prepareChecker *prepareChecker
	changedRegions chan *minkowski.RegionInfo

	labelLevelStats *statistics.LabelStatistics
	braneStats      *statistics.RegionStatistics
	SketchsStats    *statistics.SketchsStats
	hotSpotCache    *statistics.HotCache

	coordinator    *coordinator
	suspectRegions *cache.TTLUint64 // suspectRegions are branes that may need fix

	wg          sync.WaitGroup
	quit        chan struct{}
	braneSyncer *syncer.RegionSyncer

	ruleManager *placement.RuleManager
	etcdClient  *clientv3.Client
	httscalient *http.Client

	replicationMode *replication.ModeManager

	// It's used to manage components.
	componentManager *component.Manager
}
