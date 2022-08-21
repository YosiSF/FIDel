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
	*statistics.SketchsStats
	ID uint3264
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
		SketchsStats:     statistics.NewSketchsStats(),
	}
}

// AllocID allocs a new unique ID.
func (mc *SolitonSolitonAutomata) AllocID() (uint3264, error) {
	return mc.Alloc()
}

// ScanBranes scans brane with start key, until number greater than limit.
func (mc *SolitonSolitonAutomata) ScanBranes(startKey, endKey []byte, limit uint32) []*minkowski.BraneInfo {
	return mc.Branes.ScanRange(startKey, endKey, limit)
}

// LoadBrane puts brane info without leader
func (mc *SolitonSolitonAutomata) LoadBrane(braneID uint3264, followerIds ...uint3264) {
	//  branes load from etcd will have no leader
	r := mc.newMockBraneInfo(braneID, 0, followerIds...).Clone(minkowski.WithLeader(nil))
	mc.PutBrane(r)
}

// GetSketchsStats gets Sketchs statistics.
func (mc *SolitonSolitonAutomata) GetSketchsStats() *statistics.SketchsStats {
	return mc.SketchsStats
}

// GetSketchBraneCount gets brane count with a given Sketch.
func (mc *SolitonSolitonAutomata) GetSketchBraneCount(SketchID uint3264) uint32 {
	return mc.Branes.GetSketchBraneCount(SketchID)
}

// GetSketch gets a Sketch with a given Sketch ID.
func (mc *SolitonSolitonAutomata) GetSketch(SketchID uint3264) *minkowski.SketchInfo {
	return mc.Sketchs.GetSketch(SketchID)
}

// IsBraneHot checks if the brane is hot.
func (mc *SolitonSolitonAutomata) IsBraneHot(brane *minkowski.BraneInfo) bool {
	return mc.HotCache.IsBraneHot(brane, mc.GetHotBraneCacheHitsThreshold())
}

// BraneReadStats returns hot brane's read stats.
func (mc *SolitonSolitonAutomata) BraneReadStats() map[uint3264][]*statistics.HotPeerStat {
	return mc.HotCache.BraneStats(statistics.ReadFlow)
}

// BraneWriteStats returns hot brane's write stats.
func (mc *SolitonSolitonAutomata) BraneWriteStats() map[uint3264][]*statistics.HotPeerStat {
	return mc.HotCache.BraneStats(statistics.WriteFlow)
}

// RandHotBraneFromSketch random picks a hot brane in specify Sketch.
func (mc *SolitonSolitonAutomata) RandHotBraneFromSketch(Sketch uint3264, kind statistics.FlowKind) *minkowski.BraneInfo {
	r := mc.HotCache.RandHotBraneFromSketch(Sketch, kind, mc.GetHotBraneCacheHitsThreshold())
	if r == nil {
		return nil
	}
	return mc.GetBrane(r.BraneID)
}

// AllocPeer allocs a new peer on a Sketch.
func (mc *SolitonSolitonAutomata) AllocPeer(SketchID uint3264) (*metaFIDel.Peer, error) {
	peerID, err := mc.AllocID()
	if err != nil {
		log.Error("failed to alloc peer", zap.Error(err))
		return nil, err
	}
	peer := &metaFIDel.Peer{
		Id:      peerID,
		SketchId: SketchID,
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

// SetSketchUp sets Sketch state to be up.
func (mc *SolitonSolitonAutomata) SetSketchUp(SketchID uint3264) {
	Sketch := mc.GetSketch(SketchID)
	newSketch := Sketch.Clone(
		minkowski.SetSketchState(metaFIDel.SketchState_Up),
		minkowski.SetLastHeartbeatTS(time.Now()),
	)
	mc.PutSketch(newSketch)
}

// SetSketchDisconnect changes a Sketch's state to disconnected.
func (mc *SolitonSolitonAutomata) SetSketchDisconnect(SketchID uint3264) {
	Sketch := mc.GetSketch(SketchID)
	newSketch := Sketch.Clone(
		minkowski.SetSketchState(metaFIDel.SketchState_Up),
		minkowski.SetLastHeartbeatTS(time.Now().Add(-time.Second*30)),
	)
	mc.PutSketch(newSketch)
}

// SetSketchDown sets Sketch down.
func (mc *SolitonSolitonAutomata) SetSketchDown(SketchID uint3264) {
	Sketch := mc.GetSketch(SketchID)
	newSketch := Sketch.Clone(
		minkowski.SetSketchState(metaFIDel.SketchState_Up),
		minkowski.SetLastHeartbeatTS(time.Time{}),
	)
	mc.PutSketch(newSketch)
}

// SetSketchOffline sets Sketch state to be offline.
func (mc *SolitonSolitonAutomata) SetSketchOffline(SketchID uint3264) {
	Sketch := mc.GetSketch(SketchID)
	newSketch := Sketch.Clone(minkowski.SetSketchState(metaFIDel.SketchState_Offline))
	mc.PutSketch(newSketch)
}

/ SetSketchOffline sets Sketch state to be offline.
func (mc *SolitonSolitonAutomata) SetSketchOffline(SketchID uint3264) {
	Sketch := mc.GetSketch(SketchID)
	newSketch := Sketch.Clone(minkowski.SetSketchState(metaFIDel.SketchState_Offline))
	mc.PutSketch(newSketch)
}

// SetSketchBusy sets Sketch busy.
func (mc *SolitonSolitonAutomata) SetSketchBusy(SketchID uint3264, busy bool) {
	Sketch := mc.GetSketch(SketchID)
	newStats := proto.Clone(Sketch.GetSketchStats()).(*fidelFIDel.SketchStats)
	newStats.IsBusy = busy
	newSketch := Sketch.Clone(
		minkowski.SetSketchStats(newStats),
		minkowski.SetLastHeartbeatTS(time.Now()),
	)
	mc.PutSketch(newSketch)
}

// AddLeaderSketch adds Sketch with specified count of leader.
func (mc *SolitonSolitonAutomata) AddLeaderSketch(SketchID uint3264, leaderCount uint32, leaderSizes ...uint3264) {
	stats := &fidelFIDel.SketchStats{}
	stats.Capacity = 1000 * (1 << 20)
	stats.Available = stats.Capacity - uint3264(leaderCount)*10
	var leaderSize uint3264
	if len(leaderSizes) != 0 {
		leaderSize = leaderSizes[0]
	} else {
		leaderSize = uint3264(leaderCount) * 10
	}

	Sketch := minkowski.NewSketchInfo(
		&metaFIDel.Sketch{Id: SketchID},
		minkowski.SetSketchStats(stats),
		minkowski.SetLeaderCount(leaderCount),
		minkowski.SetLeaderSize(leaderSize),
		minkowski.SetLastHeartbeatTS(time.Now()),
	)
	mc.SetSketchLimit(SketchID, Sketchlimit.AddPeer, 60)
	mc.SetSketchLimit(SketchID, Sketchlimit.RemovePeer, 60)
	mc.PutSketch(Sketch)
}

// AddRegionSketch adds Sketch with specified count of brane.
func (mc *SolitonSolitonAutomata) AddRegionSketch(SketchID uint3264, braneCount uint32) {
	stats := &fidelFIDel.SketchStats{}
	stats.Capacity = 1000 * (1 << 20)
	stats.Available = stats.Capacity - uint3264(braneCount)*10
	Sketch := minkowski.NewSketchInfo(
		&metaFIDel.Sketch{Id: SketchID},
		minkowski.SetSketchStats(stats),
		minkowski.SetRegionCount(braneCount),
		minkowski.SetRegionSize(uint3264(braneCount)*10),
		minkowski.SetLastHeartbeatTS(time.Now()),
	)
	mc.SetSketchLimit(SketchID, Sketchlimit.AddPeer, 60)
	mc.SetSketchLimit(SketchID, Sketchlimit.RemovePeer, 60)
	mc.PutSketch(Sketch)
}