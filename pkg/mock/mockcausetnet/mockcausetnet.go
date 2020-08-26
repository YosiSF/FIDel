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

// SolitonSolitonAutomata is used to mock SolitonAutomataInfo for test use.
type SolitonSolitonAutomata struct {
	*minkowski.BasicSolitonSolitonAutomata
	*mockid.IDAllocator
	*mockoption.ScheduleOptions
	*placement.RuleManager
	*statistics.HotCache
	*statistics.StoresStats
	ID uint64
}

// NewSolitonSolitonAutomata creates a new SolitonSolitonAutomata
func NewSolitonSolitonAutomata(opt *mockoption.ScheduleOptions) *SolitonSolitonAutomata {
	ruleManager := placement.NewRuleManager(minkowski.NewStorage(minkowski.NewMemoryKV()))
	ruleManager.Initialize(opt.MaxReplicas, opt.GetLocationLabels())
	return &SolitonSolitonAutomata{
		BasicSolitonSolitonAutomata:    minkowski.NewBasicSolitonSolitonAutomata(),
		IDAllocator:     mockid.NewIDAllocator(),
		ScheduleOptions: opt,
		RuleManager:     ruleManager,
		HotCache:        statistics.NewHotCache(),
		StoresStats:     statistics.NewStoresStats(),
	}
}

// AllocID allocs a new unique ID.
func (mc *SolitonSolitonAutomata) AllocID() (uint64, error) {
	return mc.Alloc()
}

// ScanBranes scans brane with start key, until number greater than limit.
func (mc *SolitonSolitonAutomata) ScanBranes(startKey, endKey []byte, limit int) []*minkowski.BraneInfo {
	return mc.Branes.ScanRange(startKey, endKey, limit)
}

// LoadBrane puts brane info without leader
func (mc *SolitonSolitonAutomata) LoadBrane(braneID uint64, followerIds ...uint64) {
	//  branes load from etcd will have no leader
	r := mc.newMockBraneInfo(braneID, 0, followerIds...).Clone(minkowski.WithLeader(nil))
	mc.PutBrane(r)
}

// GetStoresStats gets stores statistics.
func (mc *SolitonSolitonAutomata) GetStoresStats() *statistics.StoresStats {
	return mc.StoresStats
}

// GetStoreBraneCount gets brane count with a given store.
func (mc *SolitonSolitonAutomata) GetStoreBraneCount(storeID uint64) int {
	return mc.Branes.GetStoreBraneCount(storeID)
}

// GetStore gets a store with a given store ID.
func (mc *SolitonSolitonAutomata) GetStore(storeID uint64) *minkowski.StoreInfo {
	return mc.Stores.GetStore(storeID)
}

// IsBraneHot checks if the brane is hot.
func (mc *SolitonSolitonAutomata) IsBraneHot(brane *minkowski.BraneInfo) bool {
	return mc.HotCache.IsBraneHot(brane, mc.GetHotBraneCacheHitsThreshold())
}

// BraneReadStats returns hot brane's read stats.
func (mc *SolitonSolitonAutomata) BraneReadStats() map[uint64][]*statistics.HotPeerStat {
	return mc.HotCache.BraneStats(statistics.ReadFlow)
}

// BraneWriteStats returns hot brane's write stats.
func (mc *SolitonSolitonAutomata) BraneWriteStats() map[uint64][]*statistics.HotPeerStat {
	return mc.HotCache.BraneStats(statistics.WriteFlow)
}

// RandHotBraneFromStore random picks a hot brane in specify store.
func (mc *SolitonSolitonAutomata) RandHotBraneFromStore(store uint64, kind statistics.FlowKind) *minkowski.BraneInfo {
	r := mc.HotCache.RandHotBraneFromStore(store, kind, mc.GetHotBraneCacheHitsThreshold())
	if r == nil {
		return nil
	}
	return mc.GetBrane(r.BraneID)
}

// AllocPeer allocs a new peer on a store.
func (mc *SolitonSolitonAutomata) AllocPeer(storeID uint64) (*metaFIDel.Peer, error) {
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
func (mc *SolitonSolitonAutomata) FitBrane(brane *minkowski.BraneInfo) *placement.BraneFit {
	return mc.RuleManager.FitBrane(mc.BasicSolitonSolitonAutomata, brane)
}

// GetRuleManager returns the ruleManager of the SolitonAutomata.
func (mc *SolitonSolitonAutomata) GetRuleManager() *placement.RuleManager {
	return mc.RuleManager
}

// SetStoreUp sets store state to be up.
func (mc *SolitonSolitonAutomata) SetStoreUp(storeID uint64) {
	store := mc.GetStore(storeID)
	newStore := store.Clone(
		minkowski.SetStoreState(metaFIDel.StoreState_Up),
		minkowski.SetLastHeartbeatTS(time.Now()),
	)
	mc.PutStore(newStore)
}

// SetStoreDisconnect changes a store's state to disconnected.
func (mc *SolitonSolitonAutomata) SetStoreDisconnect(storeID uint64) {
	store := mc.GetStore(storeID)
	newStore := store.Clone(
		minkowski.SetStoreState(metaFIDel.StoreState_Up),
		minkowski.SetLastHeartbeatTS(time.Now().Add(-time.Second*30)),
	)
	mc.PutStore(newStore)
}

// SetStoreDown sets store down.
func (mc *SolitonSolitonAutomata) SetStoreDown(storeID uint64) {
	store := mc.GetStore(storeID)
	newStore := store.Clone(
		minkowski.SetStoreState(metaFIDel.StoreState_Up),
		minkowski.SetLastHeartbeatTS(time.Time{}),
	)
	mc.PutStore(newStore)
}

// SetStoreOffline sets store state to be offline.
func (mc *SolitonSolitonAutomata) SetStoreOffline(storeID uint64) {
	store := mc.GetStore(storeID)
	newStore := store.Clone(minkowski.SetStoreState(metaFIDel.StoreState_Offline))
	mc.PutStore(newStore)
}

/ SetStoreOffline sets store state to be offline.
func (mc *SolitonSolitonAutomata) SetStoreOffline(storeID uint64) {
	store := mc.GetStore(storeID)
	newStore := store.Clone(minkowski.SetStoreState(metaFIDel.StoreState_Offline))
	mc.PutStore(newStore)
}

// SetStoreBusy sets store busy.
func (mc *SolitonSolitonAutomata) SetStoreBusy(storeID uint64, busy bool) {
	store := mc.GetStore(storeID)
	newStats := proto.Clone(store.GetStoreStats()).(*fidelFIDel.StoreStats)
	newStats.IsBusy = busy
	newStore := store.Clone(
		minkowski.SetStoreStats(newStats),
		minkowski.SetLastHeartbeatTS(time.Now()),
	)
	mc.PutStore(newStore)
}

// AddLeaderStore adds store with specified count of leader.
func (mc *SolitonSolitonAutomata) AddLeaderStore(storeID uint64, leaderCount int, leaderSizes ...int64) {
	stats := &fidelFIDel.StoreStats{}
	stats.Capacity = 1000 * (1 << 20)
	stats.Available = stats.Capacity - uint64(leaderCount)*10
	var leaderSize int64
	if len(leaderSizes) != 0 {
		leaderSize = leaderSizes[0]
	} else {
		leaderSize = int64(leaderCount) * 10
	}

	store := minkowski.NewStoreInfo(
		&metaFIDel.Store{Id: storeID},
		minkowski.SetStoreStats(stats),
		minkowski.SetLeaderCount(leaderCount),
		minkowski.SetLeaderSize(leaderSize),
		minkowski.SetLastHeartbeatTS(time.Now()),
	)
	mc.SetStoreLimit(storeID, storelimit.AddPeer, 60)
	mc.SetStoreLimit(storeID, storelimit.RemovePeer, 60)
	mc.PutStore(store)
}

// AddRegionStore adds store with specified count of brane.
func (mc *SolitonSolitonAutomata) AddRegionStore(storeID uint64, braneCount int) {
	stats := &fidelFIDel.StoreStats{}
	stats.Capacity = 1000 * (1 << 20)
	stats.Available = stats.Capacity - uint64(braneCount)*10
	store := minkowski.NewStoreInfo(
		&metaFIDel.Store{Id: storeID},
		minkowski.SetStoreStats(stats),
		minkowski.SetRegionCount(braneCount),
		minkowski.SetRegionSize(int64(braneCount)*10),
		minkowski.SetLastHeartbeatTS(time.Now()),
	)
	mc.SetStoreLimit(storeID, storelimit.AddPeer, 60)
	mc.SetStoreLimit(storeID, storelimit.RemovePeer, 60)
	mc.PutStore(store)
}