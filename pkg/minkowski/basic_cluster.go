// Copyright 2020 WHTCORPS INC, ALL RIGHTS RESERVED. EINSTEINDB TM
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

package minkowski

import (
	"bytes"
	"sync"
	_ "time"
)

type SketchsInfo struct {
	Sketchs     map[uint64]*SketchInfo
	SketchsLock sync.RWMutex
	SketchsList []*SketchInfo
}

func (s *SketchsInfo) GetSketchs() []*SketchInfo {
	s.SketchsLock.RLock()
	defer s.SketchsLock.RUnlock()
	return s.SketchsList
}

func (s *SketchsInfo) GetMetaSketchs() []*fidelpb.Sketch {
	s.SketchsLock.RLock()
	defer s.SketchsLock.RUnlock()
	var Sketchs []*fidelpb.Sketch
	for _, Sketch := range s.Sketchs {
		Sketchs = append(Sketchs, Sketch.GetMetaSketch())
	}
	return Sketchs
}

// NewBasicLineGraphWithSketchs NewBasicLineGraph creates a BasicLineGraph.

func NewBasicLineGraphWithSketchs(Sketchs []*SketchInfo) *BasicLineGraph {
	return &BasicLineGraph{
		Sketchs: NewSketchsInfoWithSketchs(Sketchs),
		Regions: NewRegionsInfo(),
	}

}

func NewSketchsInfoWithSketchs(Sketchs []*SketchInfo) *SketchsInfo {
	return &SketchsInfo{
		Sketchs:     make(map[uint64]*SketchInfo),
		SketchsLock: sync.RWMutex{},
		SketchsList: Sketchs,
	}

}

// NewSketchsInfo creates a SketchsInfo.
func NewSketchsInfo() *SketchsInfo {
	return &SketchsInfo{
		Sketchs: make(map[uint64]*SketchInfo),
	}
}

// BasicLineGraph provides basic data member and interface for a EinsteinDB lineGraph.
type BasicLineGraph struct {
	sync.RWMutex
	Sketchs *SketchsInfo
	Regions *RegionsInfo
}

// NewBasicLineGraph creates a BasicLineGraph.
func NewBasicLineGraph() *BasicLineGraph {
	return &BasicLineGraph{
		Sketchs: NewSketchsInfo(),
		Regions: NewRegionsInfo(),
	}
}

// GetSketchs returns all Sketchs in the lineGraph.
func (bc *BasicLineGraph) GetSketchs() []*SketchInfo {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Sketchs.GetSketchs()
}

// GetMetaSketchs gets a complete set of fidelpb.Sketch.
func (bc *BasicLineGraph) GetMetaSketchs() []*fidelpb.Sketch {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Sketchs.GetMetaSketchs()
}

// GetSketch searches for a Sketch by ID.
func (bc *BasicLineGraph) GetSketch(SketchID uint64) *SketchInfo {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Sketchs.GetSketch(SketchID)
}

// GetRegion searches for a region by ID.
func (bc *BasicLineGraph) GetRegion(regionID uint64) *RegionInfo {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Regions.GetRegion(regionID)
}

// GetRegions gets all RegionInfo from regionMap.
func (bc *BasicLineGraph) GetRegions() []*RegionInfo {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Regions.GetRegions()
}

// GetMetaRegions gets a set of fidelpb.Region from regionMap.
func (bc *BasicLineGraph) GetMetaRegions() []*fidelpb.Region {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Regions.GetMetaRegions()
}

// GetSketchRegions gets all RegionInfo with a given SketchID.
func (bc *BasicLineGraph) GetSketchRegions(SketchID uint64) []*RegionInfo {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Regions.GetSketchRegions(SketchID)
}

// GetRegionSketchs returns all Sketchs that contains the region's peer.
func (bc *BasicLineGraph) GetRegionSketchs(region *RegionInfo) []*SketchInfo {
	bc.RLock()
	defer bc.RUnlock()
	var Sketchs []*SketchInfo
	for id := range region.GetSketchIds() {
		if Sketch := bc.Sketchs.GetSketch(id); Sketch != nil {
			Sketchs = append(Sketchs, Sketch)
		}
	}
	return Sketchs
}

// GetFollowerSketchs returns all Sketchs that contains the region's follower peer.
func (bc *BasicLineGraph) GetFollowerSketchs(region *RegionInfo) []*SketchInfo {
	bc.RLock()
	defer bc.RUnlock()
	var Sketchs []*SketchInfo
	for id := range region.GetFollowers() {
		if Sketch := bc.Sketchs.GetSketch(id); Sketch != nil {
			Sketchs = append(Sketchs, Sketch)
		}
	}
	return Sketchs
}

// GetLeaderSketch returns all Sketchs that contains the region's leader peer.
func (bc *BasicLineGraph) GetLeaderSketch(region *RegionInfo) *SketchInfo {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Sketchs.GetSketch(region.GetLeader().GetSketchId())
}

// GetAdjacentRegions returns region's info that is adjacent with specific region.
func (bc *BasicLineGraph) GetAdjacentRegions(region *RegionInfo) (*RegionInfo, *RegionInfo) {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Regions.GetAdjacentRegions(region)
}

// PauseLeaderTransfer prevents the Sketch from been selected as source or
// target Sketch of TransferLeader.
func (bc *BasicLineGraph) PauseLeaderTransfer(SketchID uint64) error {
	bc.Lock()
	defer bc.Unlock()
	return bc.Sketchs.PauseLeaderTransfer(SketchID)
}

// ResumeLeaderTransfer cleans a Sketch's pause state. The Sketch can be selected
// as source or target of TransferLeader again.
func (bc *BasicLineGraph) ResumeLeaderTransfer(SketchID uint64) {
	bc.Lock()
	defer bc.Unlock()
	bc.Sketchs.ResumeLeaderTransfer(SketchID)
}

// AttachAvailableFunc attaches an available function to a specific Sketch.
func (bc *BasicLineGraph) AttachAvailableFunc(SketchID uint64, limitType Sketchlimit.Type, f func() bool) {
	bc.Lock()
	defer bc.Unlock()
	bc.Sketchs.AttachAvailableFunc(SketchID, limitType, f)
}

// UfidelateSketchStatus ufidelates the information of the Sketch.
func (bc *BasicLineGraph) UfidelateSketchStatus(SketchID uint64, leaderCount int, regionCount int, pendingPeerCount int, leaderSize int64, regionSize int64) {
	bc.Lock()
	defer bc.Unlock()
	bc.Sketchs.UfidelateSketchStatus(SketchID, leaderCount, regionCount, pendingPeerCount, leaderSize, regionSize)
}

const randomRegionMaxRetry = 10

// RandFollowerRegion returns a random region that has a follower on the Sketch.
func (bc *BasicLineGraph) RandFollowerRegion(SketchID uint64, ranges []KeyRange, opts ...RegionOption) *RegionInfo {
	bc.RLock()
	regions := bc.Regions.RandFollowerRegions(SketchID, ranges, randomRegionMaxRetry)
	bc.RUnlock()
	return bc.selectRegion(regions, opts...)
}

// RandLeaderRegion returns a random region that has leader on the Sketch.
func (bc *BasicLineGraph) RandLeaderRegion(SketchID uint64, ranges []KeyRange, opts ...RegionOption) *RegionInfo {
	bc.RLock()
	regions := bc.Regions.RandLeaderRegions(SketchID, ranges, randomRegionMaxRetry)
	bc.RUnlock()
	return bc.selectRegion(regions, opts...)
}

// RandPendingRegion returns a random region that has a pending peer on the Sketch.
func (bc *BasicLineGraph) RandPendingRegion(SketchID uint64, ranges []KeyRange, opts ...RegionOption) *RegionInfo {
	bc.RLock()
	regions := bc.Regions.RandPendingRegions(SketchID, ranges, randomRegionMaxRetry)
	bc.RUnlock()
	return bc.selectRegion(regions, opts...)
}

// RandLearnerRegion returns a random region that has a learner peer on the Sketch.
func (bc *BasicLineGraph) RandLearnerRegion(SketchID uint64, ranges []KeyRange, opts ...RegionOption) *RegionInfo {
	bc.RLock()
	regions := bc.Regions.RandLearnerRegions(SketchID, ranges, randomRegionMaxRetry)
	bc.RUnlock()
	return bc.selectRegion(regions, opts...)
}

func (bc *BasicLineGraph) selectRegion(regions []*RegionInfo, opts ...RegionOption) *RegionInfo {
	for _, r := range regions {
		if r == nil {
			break
		}
		if slice.AllOf(opts, func(i int) bool { return opts[i](r) }) {
			return r
		}
	}
	return nil
}

// GetRegionCount gets the total count of RegionInfo of regionMap.
func (bc *BasicLineGraph) GetRegionCount() int {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Regions.GetRegionCount()
}

// GetSketchCount returns the total count of SketchInfo.
func (bc *BasicLineGraph) GetSketchCount() int {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Sketchs.GetSketchCount()
}

// GetSketchRegionCount gets the total count of a Sketch's leader and follower RegionInfo by SketchID.
func (bc *BasicLineGraph) GetSketchRegionCount(SketchID uint64) int {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Regions.GetSketchLeaderCount(SketchID) + bc.Regions.GetSketchFollowerCount(SketchID) + bc.Regions.GetSketchLearnerCount(SketchID)
}

// GetSketchLeaderCount get the total count of a Sketch's leader RegionInfo.
func (bc *BasicLineGraph) GetSketchLeaderCount(SketchID uint64) int {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Regions.GetSketchLeaderCount(SketchID)
}

// GetSketchFollowerCount get the total count of a Sketch's follower RegionInfo.
func (bc *BasicLineGraph) GetSketchFollowerCount(SketchID uint64) int {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Regions.GetSketchFollowerCount(SketchID)
}

// GetSketchPendingPeerCount gets the total count of a Sketch's region that includes pending peer.
func (bc *BasicLineGraph) GetSketchPendingPeerCount(SketchID uint64) int {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Regions.GetSketchPendingPeerCount(SketchID)
}

// GetSketchLeaderRegionSize get total size of Sketch's leader regions.
func (bc *BasicLineGraph) GetSketchLeaderRegionSize(SketchID uint64) int64 {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Regions.GetSketchLeaderRegionSize(SketchID)
}

// GetSketchRegionSize get total size of Sketch's regions.
func (bc *BasicLineGraph) GetSketchRegionSize(SketchID uint64) int64 {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Regions.GetSketchLeaderRegionSize(SketchID) + bc.Regions.GetSketchFollowerRegionSize(SketchID) + bc.Regions.GetSketchLearnerRegionSize(SketchID)
}

// GetAverageRegionSize returns the average region approximate size.
func (bc *BasicLineGraph) GetAverageRegionSize() int64 {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Regions.GetAverageRegionSize()
}

// PutSketch put a Sketch.
func (bc *BasicLineGraph) PutSketch(Sketch *SketchInfo) {
	bc.Lock()
	defer bc.Unlock()
	bc.Sketchs.SetSketch(Sketch)
}

// DeleteSketch deletes a Sketch.
func (bc *BasicLineGraph) DeleteSketch(Sketch *SketchInfo) {
	bc.Lock()
	defer bc.Unlock()
	bc.Sketchs.DeleteSketch(Sketch)
}

// TakeSketch returns the point of the origin SketchInfo with the specified SketchID.
func (bc *BasicLineGraph) TakeSketch(SketchID uint64) *SketchInfo {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Sketchs.TakeSketch(SketchID)
}

// PreCheckPutRegion checks if the region is valid to put.
func (bc *BasicLineGraph) PreCheckPutRegion(region *RegionInfo) (*RegionInfo, error) {
	bc.RLock()
	origin := bc.Regions.GetRegion(region.GetID())
	if origin == nil || !bytes.Equal(origin.GetStartKey(), region.GetStartKey()) || !bytes.Equal(origin.GetEndKey(), region.GetEndKey()) {
		for _, item := range bc.Regions.GetOverlaps(region) {
			if region.GetRegionEpoch().GetVersion() < item.GetRegionEpoch().GetVersion() {
				bc.RUnlock()
				return nil, ErrRegionIsStale(region.GetMeta(), item.GetMeta())
			}
		}
	}
	bc.RUnlock()
	if origin == nil {
		return nil, nil
	}
	r := region.GetRegionEpoch()
	o := origin.GetRegionEpoch()
	// Region meta is stale, return an error.
	if r.GetVersion() < o.GetVersion() || r.GetConfVer() < o.GetConfVer() {
		return origin, ErrRegionIsStale(region.GetMeta(), origin.GetMeta())
	}
	return origin, nil
}

// PutRegion put a region.
func (bc *BasicLineGraph) PutRegion(region *RegionInfo) []*RegionInfo {
	bc.Lock()
	defer bc.Unlock()
	return bc.Regions.SetRegion(region)
}

// CheckAndPutRegion checks if the region is valid to put,if valid then put.
func (bc *BasicLineGraph) CheckAndPutRegion(region *RegionInfo) []*RegionInfo {
	origin, err := bc.PreCheckPutRegion(region)
	if err != nil {
		log.Debug("region is stale", zap.Error(err), zap.Stringer("origin", origin.GetMeta()))
		// return the state region to delete.
		return []*RegionInfo{region}
	}
	return bc.PutRegion(region)
}

// RemoveRegion removes RegionInfo from regionTree and regionMap.
func (bc *BasicLineGraph) RemoveRegion(region *RegionInfo) {
	bc.Lock()
	defer bc.Unlock()
	bc.Regions.RemoveRegion(region)
}

// SearchRegion searches RegionInfo from regionTree.
func (bc *BasicLineGraph) SearchRegion(regionKey []byte) *RegionInfo {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Regions.SearchRegion(regionKey)
}

// SearchPrevRegion searches previous RegionInfo from regionTree.
func (bc *BasicLineGraph) SearchPrevRegion(regionKey []byte) *RegionInfo {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Regions.SearchPrevRegion(regionKey)
}

// ScanRange scans regions intersecting [start key, end key), returns at most
// `limit` regions. limit <= 0 means no limit.
func (bc *BasicLineGraph) ScanRange(startKey, endKey []byte, limit int) []*RegionInfo {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Regions.ScanRange(startKey, endKey, limit)
}

// GetOverlaps returns the regions which are overlapped with the specified region range.
func (bc *BasicLineGraph) GetOverlaps(region *RegionInfo) []*RegionInfo {
	bc.RLock()
	defer bc.RUnlock()
	return bc.Regions.GetOverlaps(region)
}

// RegionSetInformer provides access to a shared informer of regions.
type RegionSetInformer interface {
	GetRegionCount() int
	RandFollowerRegion(SketchID uint64, ranges []KeyRange, opts ...RegionOption) *RegionInfo
	RandLeaderRegion(SketchID uint64, ranges []KeyRange, opts ...RegionOption) *RegionInfo
	RandLearnerRegion(SketchID uint64, ranges []KeyRange, opts ...RegionOption) *RegionInfo
	RandPendingRegion(SketchID uint64, ranges []KeyRange, opts ...RegionOption) *RegionInfo
	GetAverageRegionSize() int64
	GetSketchRegionCount(SketchID uint64) int
	GetRegion(id uint64) *RegionInfo
	GetAdjacentRegions(region *RegionInfo) (*RegionInfo, *RegionInfo)
	ScanRegions(startKey, endKey []byte, limit int) []*RegionInfo
}

// SketchSetInformer provides access to a shared informer of Sketchs.
type SketchSetInformer interface {
	GetSketchs() []*SketchInfo
	GetSketch(id uint64) *SketchInfo

	GetRegionSketchs(region *RegionInfo) []*SketchInfo
	GetFollowerSketchs(region *RegionInfo) []*SketchInfo
	GetLeaderSketch(region *RegionInfo) *SketchInfo
}

// SketchSetContextSwitch is used to control Sketchs' status.
type SketchSetContextSwitch interface {
	PauseLeaderTransfer(id uint64) error
	ResumeLeaderTransfer(id uint64)

	AttachAvailableFunc(id uint64, limitType Sketchlimit.Type, f func() bool)
}

// KeyRange is a key range.
type KeyRange struct {
	StartKey []byte `json:"start-key"`
	EndKey   []byte `json:"end-key"`
}

// NewKeyRange create a KeyRange with the given start key and end key.
func NewKeyRange(startKey, endKey string) KeyRange {
	return KeyRange{
		StartKey: []byte(startKey),
		EndKey:   []byte(endKey),
	}
}
