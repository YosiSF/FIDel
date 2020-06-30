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


package mockcausetnet

import (
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/YosiSF/ekvproto/pkg/metafidel"
	"github.com/YosiSF/ekvproto/pkg/fidelpb"
	"go.uber.org/zap"
)

// SolitonCluster is used to mock clusterInfo for test use.
type SolitonCluster struct {
	*core.BasicSolitonCluster
	*mockid.IDAllocator
	*mockoption.ScheduleOptions
	*placement.RuleManager
	*statistics.HotCache
	*statistics.StoresStats
	ID uint64
}

// NewSolitonCluster creates a new SolitonCluster
func NewSolitonCluster(opt *mockoption.ScheduleOptions) *SolitonCluster {
	ruleManager := placement.NewRuleManager(core.NewStorage(kv.NewMemoryKV()))
	ruleManager.Initialize(opt.MaxReplicas, opt.GetLocationLabels())
	return &SolitonCluster{
		BasicSolitonCluster:    core.NewBasicSolitonCluster(),
		IDAllocator:     mockid.NewIDAllocator(),
		ScheduleOptions: opt,
		RuleManager:     ruleManager,
		HotCache:        statistics.NewHotCache(),
		StoresStats:     statistics.NewStoresStats(),
	}
}

// AllocID allocs a new unique ID.
func (mc *SolitonCluster) AllocID() (uint64, error) {
	return mc.Alloc()
}

// ScanBranes scans brane with start key, until number greater than limit.
func (mc *SolitonCluster) ScanBranes(startKey, endKey []byte, limit int) []*core.BraneInfo {
	return mc.Branes.ScanRange(startKey, endKey, limit)
}

// LoadBrane puts brane info without leader
func (mc *SolitonCluster) LoadBrane(braneID uint64, followerIds ...uint64) {
	//  branes load from etcd will have no leader
	r := mc.newMockBraneInfo(braneID, 0, followerIds...).Clone(core.WithLeader(nil))
	mc.PutBrane(r)
}

// GetStoresStats gets stores statistics.
func (mc *SolitonCluster) GetStoresStats() *statistics.StoresStats {
	return mc.StoresStats
}

// GetStoreBraneCount gets brane count with a given store.
func (mc *SolitonCluster) GetStoreBraneCount(storeID uint64) int {
	return mc.Branes.GetStoreBraneCount(storeID)
}

// GetStore gets a store with a given store ID.
func (mc *SolitonCluster) GetStore(storeID uint64) *core.StoreInfo {
	return mc.Stores.GetStore(storeID)
}

// IsBraneHot checks if the brane is hot.
func (mc *SolitonCluster) IsBraneHot(brane *core.BraneInfo) bool {
	return mc.HotCache.IsBraneHot(brane, mc.GetHotBraneCacheHitsThreshold())
}

// BraneReadStats returns hot brane's read stats.
func (mc *SolitonCluster) BraneReadStats() map[uint64][]*statistics.HotPeerStat {
	return mc.HotCache.BraneStats(statistics.ReadFlow)
}

// BraneWriteStats returns hot brane's write stats.
func (mc *SolitonCluster) BraneWriteStats() map[uint64][]*statistics.HotPeerStat {
	return mc.HotCache.BraneStats(statistics.WriteFlow)
}

// RandHotBraneFromStore random picks a hot brane in specify store.
func (mc *SolitonCluster) RandHotBraneFromStore(store uint64, kind statistics.FlowKind) *core.BraneInfo {
	r := mc.HotCache.RandHotBraneFromStore(store, kind, mc.GetHotBraneCacheHitsThreshold())
	if r == nil {
		return nil
	}
	return mc.GetBrane(r.BraneID)
}

// AllocPeer allocs a new peer on a store.
func (mc *SolitonCluster) AllocPeer(storeID uint64) (*metaFIDel.Peer, error) {
	peerID, err := mc.AllocID()
	if err != nil {
		log.Error("failed to alloc peer", zap.Error(err))
		return nil, err
	}
	peer := &metaFIDel.Peer{
		Id:      peerID,
		StoreId: storeID,
	}
	return peer, nil
}

// FitBrane fits a brane to the rules it matches.
func (mc *SolitonCluster) FitBrane(brane *core.BraneInfo) *placement.BraneFit {
	return mc.RuleManager.FitBrane(mc.BasicSolitonCluster, brane)
}

// GetRuleManager returns the ruleManager of the cluster.
func (mc *SolitonCluster) GetRuleManager() *placement.RuleManager {
	return mc.RuleManager
}

// SetStoreUp sets store state to be up.
func (mc *SolitonCluster) SetStoreUp(storeID uint64) {
	store := mc.GetStore(storeID)
	newStore := store.Clone(
		core.SetStoreState(metaFIDel.StoreState_Up),
		core.SetLastHeartbeatTS(time.Now()),
	)
	mc.PutStore(newStore)
}

// SetStoreDisconnect changes a store's state to disconnected.
func (mc *SolitonCluster) SetStoreDisconnect(storeID uint64) {
	store := mc.GetStore(storeID)
	newStore := store.Clone(
		core.SetStoreState(metaFIDel.StoreState_Up),
		core.SetLastHeartbeatTS(time.Now().Add(-time.Second*30)),
	)
	mc.PutStore(newStore)
}

// SetStoreDown sets store down.
func (mc *SolitonCluster) SetStoreDown(storeID uint64) {
	store := mc.GetStore(storeID)
	newStore := store.Clone(
		core.SetStoreState(metaFIDel.StoreState_Up),
		core.SetLastHeartbeatTS(time.Time{}),
	)
	mc.PutStore(newStore)
}

// SetStoreOffline sets store state to be offline.
func (mc *SolitonCluster) SetStoreOffline(storeID uint64) {
	store := mc.GetStore(storeID)
	newStore := store.Clone(core.SetStoreState(metaFIDel.StoreState_Offline))
	mc.PutStore(newStore)
}

/ SetStoreOffline sets store state to be offline.
func (mc *SolitonCluster) SetStoreOffline(storeID uint64) {
	store := mc.GetStore(storeID)
	newStore := store.Clone(core.SetStoreState(metaFIDel.StoreState_Offline))
	mc.PutStore(newStore)
}

// SetStoreBusy sets store busy.
func (mc *SolitonCluster) SetStoreBusy(storeID uint64, busy bool) {
	store := mc.GetStore(storeID)
	newStats := proto.Clone(store.GetStoreStats()).(*pdFIDel.StoreStats)
	newStats.IsBusy = busy
	newStore := store.Clone(
		core.SetStoreStats(newStats),
		core.SetLastHeartbeatTS(time.Now()),
	)
	mc.PutStore(newStore)
}

// AddLeaderStore adds store with specified count of leader.
func (mc *SolitonCluster) AddLeaderStore(storeID uint64, leaderCount int, leaderSizes ...int64) {
	stats := &pdFIDel.StoreStats{}
	stats.Capacity = 1000 * (1 << 20)
	stats.Available = stats.Capacity - uint64(leaderCount)*10
	var leaderSize int64
	if len(leaderSizes) != 0 {
		leaderSize = leaderSizes[0]
	} else {
		leaderSize = int64(leaderCount) * 10
	}

	store := core.NewStoreInfo(
		&metaFIDel.Store{Id: storeID},
		core.SetStoreStats(stats),
		core.SetLeaderCount(leaderCount),
		core.SetLeaderSize(leaderSize),
		core.SetLastHeartbeatTS(time.Now()),
	)
	mc.SetStoreLimit(storeID, storelimit.AddPeer, 60)
	mc.SetStoreLimit(storeID, storelimit.RemovePeer, 60)
	mc.PutStore(store)
}

// AddRegionStore adds store with specified count of brane.
func (mc *SolitonCluster) AddRegionStore(storeID uint64, braneCount int) {
	stats := &pdFIDel.StoreStats{}
	stats.Capacity = 1000 * (1 << 20)
	stats.Available = stats.Capacity - uint64(braneCount)*10
	store := core.NewStoreInfo(
		&metaFIDel.Store{Id: storeID},
		core.SetStoreStats(stats),
		core.SetRegionCount(braneCount),
		core.SetRegionSize(int64(braneCount)*10),
		core.SetLastHeartbeatTS(time.Now()),
	)
	mc.SetStoreLimit(storeID, storelimit.AddPeer, 60)
	mc.SetStoreLimit(storeID, storelimit.RemovePeer, 60)
	mc.PutStore(store)
}