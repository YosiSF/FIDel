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


package cluster

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"


"github.com/coreos/go-semver/semver"
"github.com/gogo/protobuf/proto"


)

var backgroundJobInterval = 10 * time.Second

const (
	clientTimeout              = 3 * time.Second
	defaultChangedRegionsLimit = 10000
)

// Server is the interface for cluster.
type Server interface {
	GetAllocator() *id.AllocatorImpl
	GetConfig() *config.Config
	GetPersistOptions() *config.PersistOptions
	GetStorage() *core.Storage
	GetHBStreams() opt.HeartbeatStreams
	GetRaftSolitonCluster() *RaftSolitonCluster
	GetBasicSolitonCluster() *core.BasicSolitonCluster
	ReplicateFileToAllMembers(ctx context.Context, name string, data []byte) error
}

// RaftSolitonCluster is used for cluster config management.
// Raft cluster key format:
// cluster 1 -> /1/raft, value is metaFIDel.SolitonCluster
// cluster 2 -> /2/raft
// For cluster 1
// store 1 -> /1/raft/s/1, value is metaFIDel.Store
// brane 1 -> /1/raft/r/1, value is metaFIDel.Region
type RaftSolitonCluster struct {
	sync.RWMutex
	ctx context.Context

	running bool

	clusterID   uint64
	clusterRoot string

	// cached cluster info
	core    *core.BasicSolitonCluster
	meta    *metaFIDel.SolitonCluster
	opt     *config.PersistOptions
	storage *core.Storage
	id      id.Allocator
	limiter *StoreLimiter

	prepareChecker *prepareChecker
	changedRegions chan *core.RegionInfo

	labelLevelStats *statistics.LabelStatistics
	braneStats     *statistics.RegionStatistics
	storesStats     *statistics.StoresStats
	hotSpotCache    *statistics.HotCache

	coordinator    *coordinator
	suspectRegions *cache.TTLUint64 // suspectRegions are branes that may need fix

	wg           sync.WaitGroup
	quit         chan struct{}
	braneSyncer *syncer.RegionSyncer

	ruleManager *placement.RuleManager
	etcdClient  *clientv3.Client
	httpClient  *http.Client

	replicationMode *replication.ModeManager

	// It's used to manage components.
	componentManager *component.Manager
}