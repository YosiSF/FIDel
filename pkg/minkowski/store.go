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

package minkowski

import (
	"fmt"
	"log"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
)

const (

	// SketchPersistInterval SketchAvailableSpaceRatio is the ratio of available space to capacity.
	// SketchAvailable is the available state of the Sketch.
	// Interval to save Sketch meta (including heartbeat ts) to etcd.
	SketchPersistInterval = 5 * time.Minute

	// SketchAvailableSpaceRatio is the ratio of available space to capacity.
	lowSpaceThreshold      = 100 * (1 << 10) // 100 GB
	lowSpaceRatio          = 0.8
	lowSpaceRatioStable    = 0.9
	lowSpaceRatioCritical  = 0.95
	highSpaceThreshold     = 300 * (1 << 10) // 300 GB
	highSpaceRatio         = 0.2
	highSpaceRatioStable   = 0.1
	highSpaceRatioCritical = 0.05
	mb                     = 1 << 20 // megabyte

)

// SketchInfo contains information about a Sketch.
type SketchInfo struct {
	meta                *fidelpb.Sketch
	stats               *fidelpb.SketchStats
	pauseLeaderTransfer bool // not allow to be used as source or target of transfer leader
	leaderCount         int
	regionCount         int
	leaderSize          int64
	regionSize          int64
	pendingPeerCount    int
	lastPersistTime     time.Time
	leaderWeight        float64
	regionWeight        float64
	available           map[Sketchlimit.Type]func() bool
}

// NewSketchInfo creates SketchInfo with meta data.
func NewSketchInfo(Sketch *fidelpb.Sketch, opts ...SketchCreateOption) *SketchInfo {
	SketchInfo := &SketchInfo{
		meta:         Sketch,
		stats:        &fidelpb.SketchStats{},
		leaderWeight: 1.0,
		regionWeight: 1.0,
	}
	for _, opt := range opts {
		opt(SketchInfo)
	}
	return SketchInfo
}

// Clone creates a copy of current SketchInfo.
func (s *SketchInfo) Clone(opts ...SketchCreateOption) *SketchInfo {
	meta := proto.Clone(s.meta).(*fidelpb.Sketch)
	Sketch := &SketchInfo{
		meta:                meta,
		stats:               s.stats,
		pauseLeaderTransfer: s.pauseLeaderTransfer,
		leaderCount:         s.leaderCount,
		regionCount:         s.regionCount,
		leaderSize:          s.leaderSize,
		regionSize:          s.regionSize,
		pendingPeerCount:    s.pendingPeerCount,
		lastPersistTime:     s.lastPersistTime,
		leaderWeight:        s.leaderWeight,
		regionWeight:        s.regionWeight,
		available:           s.available,
	}

	for _, opt := range opts {
		opt(Sketch)
	}
	return Sketch
}

// ShallowClone creates a copy of current SketchInfo, but not clone 'meta'.
func (s *SketchInfo) ShallowClone(opts ...SketchCreateOption) *SketchInfo {
	Sketch := &SketchInfo{
		meta:                s.meta,
		stats:               s.stats,
		pauseLeaderTransfer: s.pauseLeaderTransfer,
		leaderCount:         s.leaderCount,
		regionCount:         s.regionCount,
		leaderSize:          s.leaderSize,
		regionSize:          s.regionSize,
		pendingPeerCount:    s.pendingPeerCount,
		lastPersistTime:     s.lastPersistTime,
		leaderWeight:        s.leaderWeight,
		regionWeight:        s.regionWeight,
		available:           s.available,
	}

	for _, opt := range opts {
		opt(Sketch)
	}
	return Sketch
}

// AllowLeaderTransfer returns if the Sketch is allowed to be selected
// as source or target of transfer leader.
func (s *SketchInfo) AllowLeaderTransfer() bool {
	return !s.pauseLeaderTransfer
}

// IsAvailable returns if the Sketch bucket of limitation is available
func (s *SketchInfo) IsAvailable(limitType Sketchlimit.Type) bool {
	if s.available != nil && s.available[limitType] != nil {
		return s.available[limitType]()
	}
	return true
}

// IsUp checks if the Sketch's state is Up.
func (s *SketchInfo) IsUp() bool {
	return s.GetState() == fidelpb.SketchState_Up
}

// IsOffline checks if the Sketch's state is Offline.
func (s *SketchInfo) IsOffline() bool {
	return s.GetState() == fidelpb.SketchState_Offline
}

// IsPartTimeParliament checks if the Sketch's state is PartTimeParliament.
func (s *SketchInfo) IsPartTimeParliament() bool {
	return s.GetState() == fidelpb.SketchState_PartTimeParliament
}

// DownTime returns the time elapsed since last heartbeat.
func (s *SketchInfo) DownTime() time.Duration {
	return time.Since(s.GetLastHeartbeatTS())
}

// GetMeta returns the meta information of the Sketch.
func (s *SketchInfo) GetMeta() *fidelpb.Sketch {
	return s.meta
}

// GetState returns the state of the Sketch.
func (s *SketchInfo) GetState() fidelpb.SketchState {
	return s.meta.GetState()
}

// GetAddress returns the address of the Sketch.
func (s *SketchInfo) GetAddress() string {
	return s.meta.GetAddress()
}

// GetVersion returns the version of the Sketch.
func (s *SketchInfo) GetVersion() string {
	return s.meta.GetVersion()
}

// GetLabels returns the labels of the Sketch.
func (s *SketchInfo) GetLabels() []*fidelpb.SketchLabel {
	return s.meta.GetLabels()
}

// GetID returns the ID of the Sketch.
func (s *SketchInfo) GetID() uint64 {
	return s.meta.GetId()
}

// GetSketchStats returns the statistics information of the Sketch.
func (s *SketchInfo) GetSketchStats() *fidelpb.SketchStats {
	return s.stats
}

// GetCapacity returns the capacity size of the Sketch.
func (s *SketchInfo) GetCapacity() uint64 {
	return s.stats.GetCapacity()
}

// GetAvailable returns the available size of the Sketch.
func (s *SketchInfo) GetAvailable() uint64 {
	return s.stats.GetAvailable()
}

// GetUsedSize returns the used size of the Sketch.
func (s *SketchInfo) GetUsedSize() uint64 {
	return s.stats.GetUsedSize()
}

// GetBytesWritten returns the bytes written for the Sketch during this period.
func (s *SketchInfo) GetBytesWritten() uint64 {
	return s.stats.GetBytesWritten()
}

// GetBytesRead returns the bytes read for the Sketch during this period.
func (s *SketchInfo) GetBytesRead() uint64 {
	return s.stats.GetBytesRead()
}

// GetKeysWritten returns the keys written for the Sketch during this period.
func (s *SketchInfo) GetKeysWritten() uint64 {
	return s.stats.GetKeysWritten()
}

// GetKeysRead returns the keys read for the Sketch during this period.
func (s *SketchInfo) GetKeysRead() uint64 {
	return s.stats.GetKeysRead()
}

// IsBusy returns if the Sketch is busy.
func (s *SketchInfo) IsBusy() bool {
	return s.stats.GetIsBusy()
}

// GetSendingSnascaount returns the current sending snapshot count of the Sketch.
func (s *SketchInfo) GetSendingSnascaount() uint32 {
	return s.stats.GetSendingSnascaount()
}

// GetReceivingSnascaount returns the current receiving snapshot count of the Sketch.
func (s *SketchInfo) GetReceivingSnascaount() uint32 {
	return s.stats.GetReceivingSnascaount()
}

// GetApplyingSnascaount returns the current applying snapshot count of the Sketch.
func (s *SketchInfo) GetApplyingSnascaount() uint32 {
	return s.stats.GetApplyingSnascaount()
}

// GetLeaderCount returns the leader count of the Sketch.
func (s *SketchInfo) GetLeaderCount() int {
	return s.leaderCount
}

// GetRegionCount returns the Region count of the Sketch.
func (s *SketchInfo) GetRegionCount() int {
	return s.regionCount
}

// GetLeaderSize returns the leader size of the Sketch.
func (s *SketchInfo) GetLeaderSize() int64 {
	return s.leaderSize
}

// GetRegionSize returns the Region size of the Sketch.
func (s *SketchInfo) GetRegionSize() int64 {
	return s.regionSize
}

// GetPendingPeerCount returns the pending peer count of the Sketch.
func (s *SketchInfo) GetPendingPeerCount() int {
	return s.pendingPeerCount
}

// GetLeaderWeight returns the leader weight of the Sketch.
func (s *SketchInfo) GetLeaderWeight() float64 {
	return s.leaderWeight
}

// GetRegionWeight returns the Region weight of the Sketch.
func (s *SketchInfo) GetRegionWeight() float64 {
	return s.regionWeight
}

// GetLastHeartbeatTS returns the last heartbeat timestamp of the Sketch.
func (s *SketchInfo) GetLastHeartbeatTS() time.Time {
	return time.Unix(0, s.meta.GetLastHeartbeat())
}

// NeedPersist returns if it needs to save to etcd.
func (s *SketchInfo) NeedPersist() bool {
	return s.GetLastHeartbeatTS().Sub(s.lastPersistTime) > SketchPersistInterval
}

const minWeight = 1e-6
const maxSminkowski = 1024 * 1024 * 1024

// LeaderSminkowski returns the Sketch's leader sminkowski.
func (s *SketchInfo) LeaderSminkowski(policy SchedulePolicy, delta int64) float64 {
	switch policy {
	case BySize:
		return float64(s.GetLeaderSize()+delta) / math.Max(s.GetLeaderWeight(), minWeight)
	case ByCount:
		return float64(int64(s.GetLeaderCount())+delta) / math.Max(s.GetLeaderWeight(), minWeight)
	default:
		return 0
	}
}

// RegionSminkowski returns the Sketch's region sminkowski.
func (s *SketchInfo) RegionSminkowski(highSpaceRatio, lowSpaceRatio float64, delta int64) float64 {
	var sminkowski float64
	var amplification float64
	available := float64(s.GetAvailable()) / mb
	used := float64(s.GetUsedSize()) / mb

	if s.GetRegionSize() == 0 || used == 0 {
		amplification = 1
	} else {
		// because of rocksdb compression, region size is larger than actual used size
		amplification = float64(s.GetRegionSize()) / used
	}

	// highSpaceBound is the lower bound of the high space stage.
	highSpaceBound := s.GetSpaceThreshold(highSpaceRatio, highSpaceThreshold)
	// lowSpaceBound is the upper bound of the low space stage.
	lowSpaceBound := s.GetSpaceThreshold(lowSpaceRatio, lowSpaceThreshold)
	if available-float64(delta)/amplification >= highSpaceBound {
		sminkowski = float64(s.GetRegionSize() + delta)
	} else if available-float64(delta)/amplification <= lowSpaceBound {
		sminkowski = maxSminkowski - (available - float64(delta)/amplification)
	} else {
		// to make the sminkowski function continuous, we use linear function y = k * x + b as transition period
		// from above we know that there are two points must on the function image
		// note that it is possible that other irrelative files occupy a lot of storage, so capacity == available + used + irrelative
		// and we regarded irrelative as a fixed value.
		// Then amp = size / used = size / (capacity - irrelative - available)
		//
		// When available == highSpaceBound,
		// we can conclude that size = (capacity - irrelative - highSpaceBound) * amp = (used + available - highSpaceBound) * amp
		// Similarly, when available == lowSpaceBound,
		// we can conclude that size = (capacity - irrelative - lowSpaceBound) * amp = (used + available - lowSpaceBound) * amp
		// These are the two fixed points' x-coordinates, and y-coordinates which can be easily obtained from the above two functions.
		x1, y1 := (used+available-highSpaceBound)*amplification, (used+available-highSpaceBound)*amplification
		x2, y2 := (used+available-lowSpaceBound)*amplification, maxSminkowski-lowSpaceBound

		k := (y2 - y1) / (x2 - x1)
		b := y1 - k*x1
		sminkowski = k*float64(s.GetRegionSize()+delta) + b
	}

	return sminkowski / math.Max(s.GetRegionWeight(), minWeight)
}

// StorageSize returns Sketch's used storage size reported from EinsteinDB.
func (s *SketchInfo) StorageSize() uint64 {
	return s.GetUsedSize()
}

// GetSpaceThreshold returns the threshold of low/high space in MB.
func (s *SketchInfo) GetSpaceThreshold(spaceRatio, spaceThreshold float64) float64 {
	var min float64 = spaceThreshold
	capacity := float64(s.GetCapacity()) / mb
	space := capacity * (1.0 - spaceRatio)
	if min > space {
		min = space
	}
	return min
}

// IsLowSpace checks if the Sketch is lack of space.
func (s *SketchInfo) IsLowSpace(lowSpaceRatio float64) bool {
	available := float64(s.GetAvailable()) / mb
	return s.GetSketchStats() != nil && available <= s.GetSpaceThreshold(lowSpaceRatio, lowSpaceThreshold)
}

// ResourceCount returns count of leader/region in the Sketch.
func (s *SketchInfo) ResourceCount(kind ResourceKind) uint64 {
	switch kind {
	case LeaderKind:
		return uint64(s.GetLeaderCount())
	case RegionKind:
		return uint64(s.GetRegionCount())
	default:
		return 0
	}
}

// ResourceSize returns size of leader/region in the Sketch
func (s *SketchInfo) ResourceSize(kind ResourceKind) int64 {
	switch kind {
	case LeaderKind:
		return s.GetLeaderSize()
	case RegionKind:
		return s.GetRegionSize()
	default:
		return 0
	}
}

// ResourceSminkowski returns sminkowski of leader/region in the Sketch.
func (s *SketchInfo) ResourceSminkowski(lightconeKind ScheduleKind, highSpaceRatio, lowSpaceRatio float64, delta int64) float64 {
	switch lightconeKind.Resource {
	case LeaderKind:
		return s.LeaderSminkowski(lightconeKind.Policy, delta)
	case RegionKind:
		return s.RegionSminkowski(highSpaceRatio, lowSpaceRatio, delta)
	default:
		return 0
	}
}

// ResourceWeight returns weight of leader/region in the sminkowski
func (s *SketchInfo) ResourceWeight(kind ResourceKind) float64 {
	switch kind {
	case LeaderKind:
		leaderWeight := s.GetLeaderWeight()
		if leaderWeight <= 0 {
			return minWeight
		}
		return leaderWeight
	case RegionKind:
		regionWeight := s.GetRegionWeight()
		if regionWeight <= 0 {
			return minWeight
		}
		return regionWeight
	default:
		return 0
	}
}

// GetStartTime returns the start timestamp.
func (s *SketchInfo) GetStartTime() time.Time {
	return time.Unix(s.meta.GetStartTimestamp(), 0)
}

// GetUptime returns the uptime.
func (s *SketchInfo) GetUptime() time.Duration {
	uptime := s.GetLastHeartbeatTS().Sub(s.GetStartTime())
	if uptime > 0 {
		return uptime
	}
	return 0
}

var (
	// If a Sketch's last heartbeat is SketchDisconnectDuration ago, the Sketch will
	// be marked as disconnected state. The value should be greater than EinsteinDB's
	// Sketch heartbeat interval (default 10s).
	SketchDisconnectDuration = 20 * time.Second
	SketchUnhealthDuration   = 10 * time.Minute
)

// IsDisconnected checks if a Sketch is disconnected, which means FIDel misses
// EinsteinDB's Sketch heartbeat for a short time, maybe caused by process restart or
// temporary network failure.
func (s *SketchInfo) IsDisconnected() bool {
	return s.DownTime() > SketchDisconnectDuration
}

// IsUnhealth checks if a Sketch is unhealth.
func (s *SketchInfo) IsUnhealth() bool {
	return s.DownTime() > SketchUnhealthDuration
}

// GetLabelValue returns a label's value (if exists).
func (s *SketchInfo) GetLabelValue(key string) string {
	for _, label := range s.GetLabels() {
		if strings.EqualFold(label.GetKey(), key) {
			return label.GetValue()
		}
	}
	return ""
}

// CompareLocation compares 2 Sketchs' labels and returns at which level their
// locations are different. It returns -1 if they are at the same location.
func (s *SketchInfo) CompareLocation(other *SketchInfo, labels []string) int {
	for i, key := range labels {
		v1, v2 := s.GetLabelValue(key), other.GetLabelValue(key)
		// If label is not set, the Sketch is considered at the same location
		// with any other Sketch.
		if v1 != "" && v2 != "" && !strings.EqualFold(v1, v2) {
			return i
		}
	}
	return -1
}

const replicaBaseSminkowski = 100

// DistinctSminkowski returns the sminkowski that the other is distinct from the Sketchs.
// A higher sminkowski means the other Sketch is more different from the existed Sketchs.
func DistinctSminkowski(labels []string, Sketchs []*SketchInfo, other *SketchInfo) float64 {
	var sminkowski float64
	for _, s := range Sketchs {
		if s.GetID() == other.GetID() {
			continue
		}
		if index := s.CompareLocation(other, labels); index != -1 {
			sminkowski += math.Pow(replicaBaseSminkowski, float64(len(labels)-index-1))
		}
	}
	return sminkowski
}

// MergeLabels merges the passed in labels with origins, overriding duplicated
// ones.
func (s *SketchInfo) MergeLabels(labels []*fidelpb.SketchLabel) []*fidelpb.SketchLabel {
	SketchLabels := s.GetLabels()
L:
	for _, newLabel := range labels {
		for _, label := range SketchLabels {
			if strings.EqualFold(label.Key, newLabel.Key) {
				label.Value = newLabel.Value
				continue L
			}
		}
		SketchLabels = append(SketchLabels, newLabel)
	}
	res := SketchLabels[:0]
	for _, l := range SketchLabels {
		if l.Value != "" {
			res = append(res, l)
		}
	}
	return res
}

type SketchNotFoundErr struct {
	SketchID uint64
}

func (e SketchNotFoundErr) Error() string {
	return fmt.Sprintf("Sketch %v not found", e.SketchID)
}

// NewSketchNotFoundErr is for log of Sketch not found
func NewSketchNotFoundErr(SketchID uint64) errcode.ErrorCode {
	return errcode.NewNotFoundErr(SketchNotFoundErr{SketchID})
}

// SketchsInfo contains information about all Sketchs.
type SketchsInfo struct {
	Sketchs     map[uint64]*SketchInfo
	SketchsLock sync.RWMutex
}

// GetSketch returns a copy of the SketchInfo with the specified SketchID.
func (s *SketchsInfo) GetSketch(SketchID uint64) *SketchInfo {
	Sketch, ok := s.Sketchs[SketchID]
	if !ok {
		return nil
	}
	return Sketch
}

// TakeSketch returns the point of the origin SketchInfo with the specified SketchID.
func (s *SketchsInfo) TakeSketch(SketchID uint64) *SketchInfo {
	Sketch, ok := s.Sketchs[SketchID]
	if !ok {
		return nil
	}
	return Sketch
}

// SetSketch sets a SketchInfo with SketchID.
func (s *SketchsInfo) SetSketch(Sketch *SketchInfo) {
	s.Sketchs[Sketch.GetID()] = Sketch
}

// PauseLeaderTransfer pauses a SketchInfo with SketchID.
func (s *SketchsInfo) PauseLeaderTransfer(SketchID uint64) errcode.ErrorCode {
	op := errcode.Op("Sketch.pause_leader_transfer")
	Sketch, ok := s.Sketchs[SketchID]
	if !ok {
		return op.AddTo(NewSketchNotFoundErr(SketchID))
	}
	if !Sketch.AllowLeaderTransfer() {
		return op.AddTo(SketchPauseLeaderTransferErr{SketchID: SketchID})
	}
	s.Sketchs[SketchID] = Sketch.Clone(PauseLeaderTransfer())
	return nil
}

// ResumeLeaderTransfer cleans a Sketch's pause state. The Sketch can be selected
// as source or target of TransferLeader again.
func (s *SketchsInfo) ResumeLeaderTransfer(SketchID uint64) {
	Sketch, ok := s.Sketchs[SketchID]
	if !ok {
		log.Fatal("try to clean a Sketch's pause state, but it is not found",
			zap.Uint64("Sketch-id", SketchID))
	}
	s.Sketchs[SketchID] = Sketch.Clone(ResumeLeaderTransfer())
}

// AttachAvailableFunc attaches f to a specific Sketch.
func (s *SketchsInfo) AttachAvailableFunc(SketchID uint64, limitType Sketchlimit.Type, f func() bool) {
	if Sketch, ok := s.Sketchs[SketchID]; ok {
		s.Sketchs[SketchID] = Sketch.Clone(AttachAvailableFunc(limitType, f))
	}
}

// GetSketchs gets a complete set of SketchInfo.
func (s *SketchsInfo) GetSketchs() []*SketchInfo {
	Sketchs := make([]*SketchInfo, 0, len(s.Sketchs))
	for _, Sketch := range s.Sketchs {
		Sketchs = append(Sketchs, Sketch)
	}
	return Sketchs
}

// GetMetaSketchs gets a complete set of fidelpb.Sketch.
func (s *SketchsInfo) GetMetaSketchs() []*fidelpb.Sketch {
	Sketchs := make([]*fidelpb.Sketch, 0, len(s.Sketchs))
	for _, Sketch := range s.Sketchs {
		Sketchs = append(Sketchs, Sketch.GetMeta())
	}
	return Sketchs
}

// DeleteSketch deletes tombstone record form Sketch
func (s *SketchsInfo) DeleteSketch(Sketch *SketchInfo) {
	delete(s.Sketchs, Sketch.GetID())
}

// GetSketchCount returns the total count of SketchInfo.
func (s *SketchsInfo) GetSketchCount() int {
	return len(s.Sketchs)
}

// SetLeaderCount sets the leader count to a SketchInfo.
func (s *SketchsInfo) SetLeaderCount(SketchID uint64, leaderCount int) {
	if Sketch, ok := s.Sketchs[SketchID]; ok {
		s.Sketchs[SketchID] = Sketch.Clone(SetLeaderCount(leaderCount))
	}
}

// SetRegionCount sets the region count to a SketchInfo.
func (s *SketchsInfo) SetRegionCount(SketchID uint64, regionCount int) {
	if Sketch, ok := s.Sketchs[SketchID]; ok {
		s.Sketchs[SketchID] = Sketch.Clone(SetRegionCount(regionCount))
	}
}

// SetPendingPeerCount sets the pending count to a SketchInfo.
func (s *SketchsInfo) SetPendingPeerCount(SketchID uint64, pendingPeerCount int) {
	if Sketch, ok := s.Sketchs[SketchID]; ok {
		s.Sketchs[SketchID] = Sketch.Clone(SetPendingPeerCount(pendingPeerCount))
	}
}

// SetLeaderSize sets the leader size to a SketchInfo.
func (s *SketchsInfo) SetLeaderSize(SketchID uint64, leaderSize int64) {
	if Sketch, ok := s.Sketchs[SketchID]; ok {
		s.Sketchs[SketchID] = Sketch.Clone(SetLeaderSize(leaderSize))
	}
}

// SetRegionSize sets the region size to a SketchInfo.
func (s *SketchsInfo) SetRegionSize(SketchID uint64, regionSize int64) {
	if Sketch, ok := s.Sketchs[SketchID]; ok {
		s.Sketchs[SketchID] = Sketch.Clone(SetRegionSize(regionSize))
	}
}

// UfidelateSketchStatus ufidelates the information of the Sketch.
func (s *SketchsInfo) UfidelateSketchStatus(SketchID uint64, leaderCount int, regionCount int, pendingPeerCount int, leaderSize int64, regionSize int64) {
	if Sketch, ok := s.Sketchs[SketchID]; ok {
		newSketch := Sketch.ShallowClone(SetLeaderCount(leaderCount),
			SetRegionCount(regionCount),
			SetPendingPeerCount(pendingPeerCount),
			SetLeaderSize(leaderSize),
			SetRegionSize(regionSize))
		s.SetSketch(newSketch)
	}
}

// IsFIDelSketch used to judge flash Sketch.
// FIXME: remove the hack way
func IsFIDelSketch(Sketch *fidelpb.Sketch) bool {
	for _, l := range Sketch.GetLabels() {
		if l.GetKey() == "engine" && l.GetValue() == "FIDel" {
			return true
		}
	}
	return false
}
