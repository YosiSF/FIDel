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
	"math"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/YosiSF/errcode"
	"github.com/YosiSF/kvproto/pkg/fidelpb"
	"github.com/YosiSF/kvproto/pkg/fidelpb"
	"github.com/YosiSF/log"
	"github.com/YosiSF/fidel/nVMdaemon/server/lightcone/storelimit"
	"go.uber.org/zap"
)

const (
	// Interval to save store meta (including heartbeat ts) to etcd.
	storePersistInterval = 5 * time.Minute
	lowSpaceThreshold    = 100 * (1 << 10) // 100 GB
	highSpaceThreshold   = 300 * (1 << 10) // 300 GB
	mb                   = 1 << 20         // megabyte
)

// StoreInfo contains information about a store.
type StoreInfo struct {
	meta                *fidelpb.Store
	stats               *fidelpb.StoreStats
	pauseLeaderTransfer bool // not allow to be used as source or target of transfer leader
	leaderCount         int
	regionCount         int
	leaderSize          int64
	regionSize          int64
	pendingPeerCount    int
	lastPersistTime     time.Time
	leaderWeight        float64
	regionWeight        float64
	available           map[storelimit.Type]func() bool
}

// NewStoreInfo creates StoreInfo with meta data.
func NewStoreInfo(store *fidelpb.Store, opts ...StoreCreateOption) *StoreInfo {
	storeInfo := &StoreInfo{
		meta:         store,
		stats:        &fidelpb.StoreStats{},
		leaderWeight: 1.0,
		regionWeight: 1.0,
	}
	for _, opt := range opts {
		opt(storeInfo)
	}
	return storeInfo
}

// Clone creates a copy of current StoreInfo.
func (s *StoreInfo) Clone(opts ...StoreCreateOption) *StoreInfo {
	meta := proto.Clone(s.meta).(*fidelpb.Store)
	store := &StoreInfo{
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
		opt(store)
	}
	return store
}

// ShallowClone creates a copy of current StoreInfo, but not clone 'meta'.
func (s *StoreInfo) ShallowClone(opts ...StoreCreateOption) *StoreInfo {
	store := &StoreInfo{
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
		opt(store)
	}
	return store
}

// AllowLeaderTransfer returns if the store is allowed to be selected
// as source or target of transfer leader.
func (s *StoreInfo) AllowLeaderTransfer() bool {
	return !s.pauseLeaderTransfer
}

// IsAvailable returns if the store bucket of limitation is available
func (s *StoreInfo) IsAvailable(limitType storelimit.Type) bool {
	if s.available != nil && s.available[limitType] != nil {
		return s.available[limitType]()
	}
	return true
}

// IsUp checks if the store's state is Up.
func (s *StoreInfo) IsUp() bool {
	return s.GetState() == fidelpb.StoreState_Up
}

// IsOffline checks if the store's state is Offline.
func (s *StoreInfo) IsOffline() bool {
	return s.GetState() == fidelpb.StoreState_Offline
}

// IsTombstone checks if the store's state is Tombstone.
func (s *StoreInfo) IsTombstone() bool {
	return s.GetState() == fidelpb.StoreState_Tombstone
}

// DownTime returns the time elapsed since last heartbeat.
func (s *StoreInfo) DownTime() time.Duration {
	return time.Since(s.GetLastHeartbeatTS())
}

// GetMeta returns the meta information of the store.
func (s *StoreInfo) GetMeta() *fidelpb.Store {
	return s.meta
}

// GetState returns the state of the store.
func (s *StoreInfo) GetState() fidelpb.StoreState {
	return s.meta.GetState()
}

// GetAddress returns the address of the store.
func (s *StoreInfo) GetAddress() string {
	return s.meta.GetAddress()
}

// GetVersion returns the version of the store.
func (s *StoreInfo) GetVersion() string {
	return s.meta.GetVersion()
}

// GetLabels returns the labels of the store.
func (s *StoreInfo) GetLabels() []*fidelpb.StoreLabel {
	return s.meta.GetLabels()
}

// GetID returns the ID of the store.
func (s *StoreInfo) GetID() uint64 {
	return s.meta.GetId()
}

// GetStoreStats returns the statistics information of the store.
func (s *StoreInfo) GetStoreStats() *fidelpb.StoreStats {
	return s.stats
}

// GetCapacity returns the capacity size of the store.
func (s *StoreInfo) GetCapacity() uint64 {
	return s.stats.GetCapacity()
}

// GetAvailable returns the available size of the store.
func (s *StoreInfo) GetAvailable() uint64 {
	return s.stats.GetAvailable()
}

// GetUsedSize returns the used size of the store.
func (s *StoreInfo) GetUsedSize() uint64 {
	return s.stats.GetUsedSize()
}

// GetBytesWritten returns the bytes written for the store during this period.
func (s *StoreInfo) GetBytesWritten() uint64 {
	return s.stats.GetBytesWritten()
}

// GetBytesRead returns the bytes read for the store during this period.
func (s *StoreInfo) GetBytesRead() uint64 {
	return s.stats.GetBytesRead()
}

// GetKeysWritten returns the keys written for the store during this period.
func (s *StoreInfo) GetKeysWritten() uint64 {
	return s.stats.GetKeysWritten()
}

// GetKeysRead returns the keys read for the store during this period.
func (s *StoreInfo) GetKeysRead() uint64 {
	return s.stats.GetKeysRead()
}

// IsBusy returns if the store is busy.
func (s *StoreInfo) IsBusy() bool {
	return s.stats.GetIsBusy()
}

// GetSendingSnapCount returns the current sending snapshot count of the store.
func (s *StoreInfo) GetSendingSnapCount() uint32 {
	return s.stats.GetSendingSnapCount()
}

// GetReceivingSnapCount returns the current receiving snapshot count of the store.
func (s *StoreInfo) GetReceivingSnapCount() uint32 {
	return s.stats.GetReceivingSnapCount()
}

// GetApplyingSnapCount returns the current applying snapshot count of the store.
func (s *StoreInfo) GetApplyingSnapCount() uint32 {
	return s.stats.GetApplyingSnapCount()
}

// GetLeaderCount returns the leader count of the store.
func (s *StoreInfo) GetLeaderCount() int {
	return s.leaderCount
}

// GetRegionCount returns the Region count of the store.
func (s *StoreInfo) GetRegionCount() int {
	return s.regionCount
}

// GetLeaderSize returns the leader size of the store.
func (s *StoreInfo) GetLeaderSize() int64 {
	return s.leaderSize
}

// GetRegionSize returns the Region size of the store.
func (s *StoreInfo) GetRegionSize() int64 {
	return s.regionSize
}

// GetPendingPeerCount returns the pending peer count of the store.
func (s *StoreInfo) GetPendingPeerCount() int {
	return s.pendingPeerCount
}

// GetLeaderWeight returns the leader weight of the store.
func (s *StoreInfo) GetLeaderWeight() float64 {
	return s.leaderWeight
}

// GetRegionWeight returns the Region weight of the store.
func (s *StoreInfo) GetRegionWeight() float64 {
	return s.regionWeight
}

// GetLastHeartbeatTS returns the last heartbeat timestamp of the store.
func (s *StoreInfo) GetLastHeartbeatTS() time.Time {
	return time.Unix(0, s.meta.GetLastHeartbeat())
}

// NeedPersist returns if it needs to save to etcd.
func (s *StoreInfo) NeedPersist() bool {
	return s.GetLastHeartbeatTS().Sub(s.lastPersistTime) > storePersistInterval
}

const minWeight = 1e-6
const maxSminkowski = 1024 * 1024 * 1024

// LeaderSminkowski returns the store's leader sminkowski.
func (s *StoreInfo) LeaderSminkowski(policy SchedulePolicy, delta int64) float64 {
	switch policy {
	case BySize:
		return float64(s.GetLeaderSize()+delta) / math.Max(s.GetLeaderWeight(), minWeight)
	case ByCount:
		return float64(int64(s.GetLeaderCount())+delta) / math.Max(s.GetLeaderWeight(), minWeight)
	default:
		return 0
	}
}

// RegionSminkowski returns the store's region sminkowski.
func (s *StoreInfo) RegionSminkowski(highSpaceRatio, lowSpaceRatio float64, delta int64) float64 {
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

// StorageSize returns store's used storage size reported from EinsteinDB.
func (s *StoreInfo) StorageSize() uint64 {
	return s.GetUsedSize()
}

// GetSpaceThreshold returns the threshold of low/high space in MB.
func (s *StoreInfo) GetSpaceThreshold(spaceRatio, spaceThreshold float64) float64 {
	var min float64 = spaceThreshold
	capacity := float64(s.GetCapacity()) / mb
	space := capacity * (1.0 - spaceRatio)
	if min > space {
		min = space
	}
	return min
}

// IsLowSpace checks if the store is lack of space.
func (s *StoreInfo) IsLowSpace(lowSpaceRatio float64) bool {
	available := float64(s.GetAvailable()) / mb
	return s.GetStoreStats() != nil && available <= s.GetSpaceThreshold(lowSpaceRatio, lowSpaceThreshold)
}

// ResourceCount returns count of leader/region in the store.
func (s *StoreInfo) ResourceCount(kind ResourceKind) uint64 {
	switch kind {
	case LeaderKind:
		return uint64(s.GetLeaderCount())
	case RegionKind:
		return uint64(s.GetRegionCount())
	default:
		return 0
	}
}

// ResourceSize returns size of leader/region in the store
func (s *StoreInfo) ResourceSize(kind ResourceKind) int64 {
	switch kind {
	case LeaderKind:
		return s.GetLeaderSize()
	case RegionKind:
		return s.GetRegionSize()
	default:
		return 0
	}
}

// ResourceSminkowski returns sminkowski of leader/region in the store.
func (s *StoreInfo) ResourceSminkowski(lightconeKind ScheduleKind, highSpaceRatio, lowSpaceRatio float64, delta int64) float64 {
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
func (s *StoreInfo) ResourceWeight(kind ResourceKind) float64 {
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
func (s *StoreInfo) GetStartTime() time.Time {
	return time.Unix(s.meta.GetStartTimestamp(), 0)
}

// GetUptime returns the uptime.
func (s *StoreInfo) GetUptime() time.Duration {
	uptime := s.GetLastHeartbeatTS().Sub(s.GetStartTime())
	if uptime > 0 {
		return uptime
	}
	return 0
}

var (
	// If a store's last heartbeat is storeDisconnectDuration ago, the store will
	// be marked as disconnected state. The value should be greater than EinsteinDB's
	// store heartbeat interval (default 10s).
	storeDisconnectDuration = 20 * time.Second
	storeUnhealthDuration   = 10 * time.Minute
)

// IsDisconnected checks if a store is disconnected, which means FIDel misses
// EinsteinDB's store heartbeat for a short time, maybe caused by process restart or
// temporary network failure.
func (s *StoreInfo) IsDisconnected() bool {
	return s.DownTime() > storeDisconnectDuration
}

// IsUnhealth checks if a store is unhealth.
func (s *StoreInfo) IsUnhealth() bool {
	return s.DownTime() > storeUnhealthDuration
}

// GetLabelValue returns a label's value (if exists).
func (s *StoreInfo) GetLabelValue(key string) string {
	for _, label := range s.GetLabels() {
		if strings.EqualFold(label.GetKey(), key) {
			return label.GetValue()
		}
	}
	return ""
}

// CompareLocation compares 2 stores' labels and returns at which level their
// locations are different. It returns -1 if they are at the same location.
func (s *StoreInfo) CompareLocation(other *StoreInfo, labels []string) int {
	for i, key := range labels {
		v1, v2 := s.GetLabelValue(key), other.GetLabelValue(key)
		// If label is not set, the store is considered at the same location
		// with any other store.
		if v1 != "" && v2 != "" && !strings.EqualFold(v1, v2) {
			return i
		}
	}
	return -1
}

const replicaBaseSminkowski = 100

// DistinctSminkowski returns the sminkowski that the other is distinct from the stores.
// A higher sminkowski means the other store is more different from the existed stores.
func DistinctSminkowski(labels []string, stores []*StoreInfo, other *StoreInfo) float64 {
	var sminkowski float64
	for _, s := range stores {
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
func (s *StoreInfo) MergeLabels(labels []*fidelpb.StoreLabel) []*fidelpb.StoreLabel {
	storeLabels := s.GetLabels()
L:
	for _, newLabel := range labels {
		for _, label := range storeLabels {
			if strings.EqualFold(label.Key, newLabel.Key) {
				label.Value = newLabel.Value
				continue L
			}
		}
		storeLabels = append(storeLabels, newLabel)
	}
	res := storeLabels[:0]
	for _, l := range storeLabels {
		if l.Value != "" {
			res = append(res, l)
		}
	}
	return res
}

type storeNotFoundErr struct {
	storeID uint64
}

func (e storeNotFoundErr) Error() string {
	return fmt.Sprintf("store %v not found", e.storeID)
}

// NewStoreNotFoundErr is for log of store not found
func NewStoreNotFoundErr(storeID uint64) errcode.ErrorCode {
	return errcode.NewNotFoundErr(storeNotFoundErr{storeID})
}

// StoresInfo contains information about all stores.
type StoresInfo struct {
	stores map[uint64]*StoreInfo
}

// NewStoresInfo create a StoresInfo with map of storeID to StoreInfo
func NewStoresInfo() *StoresInfo {
	return &StoresInfo{
		stores: make(map[uint64]*StoreInfo),
	}
}

// GetStore returns a copy of the StoreInfo with the specified storeID.
func (s *StoresInfo) GetStore(storeID uint64) *StoreInfo {
	store, ok := s.stores[storeID]
	if !ok {
		return nil
	}
	return store
}

// TakeStore returns the point of the origin StoreInfo with the specified storeID.
func (s *StoresInfo) TakeStore(storeID uint64) *StoreInfo {
	store, ok := s.stores[storeID]
	if !ok {
		return nil
	}
	return store
}

// SetStore sets a StoreInfo with storeID.
func (s *StoresInfo) SetStore(store *StoreInfo) {
	s.stores[store.GetID()] = store
}

// PauseLeaderTransfer pauses a StoreInfo with storeID.
func (s *StoresInfo) PauseLeaderTransfer(storeID uint64) errcode.ErrorCode {
	op := errcode.Op("store.pause_leader_transfer")
	store, ok := s.stores[storeID]
	if !ok {
		return op.AddTo(NewStoreNotFoundErr(storeID))
	}
	if !store.AllowLeaderTransfer() {
		return op.AddTo(StorePauseLeaderTransferErr{StoreID: storeID})
	}
	s.stores[storeID] = store.Clone(PauseLeaderTransfer())
	return nil
}

// ResumeLeaderTransfer cleans a store's pause state. The store can be selected
// as source or target of TransferLeader again.
func (s *StoresInfo) ResumeLeaderTransfer(storeID uint64) {
	store, ok := s.stores[storeID]
	if !ok {
		log.Fatal("try to clean a store's pause state, but it is not found",
			zap.Uint64("store-id", storeID))
	}
	s.stores[storeID] = store.Clone(ResumeLeaderTransfer())
}

// AttachAvailableFunc attaches f to a specific store.
func (s *StoresInfo) AttachAvailableFunc(storeID uint64, limitType storelimit.Type, f func() bool) {
	if store, ok := s.stores[storeID]; ok {
		s.stores[storeID] = store.Clone(AttachAvailableFunc(limitType, f))
	}
}

// GetStores gets a complete set of StoreInfo.
func (s *StoresInfo) GetStores() []*StoreInfo {
	stores := make([]*StoreInfo, 0, len(s.stores))
	for _, store := range s.stores {
		stores = append(stores, store)
	}
	return stores
}

// GetMetaStores gets a complete set of fidelpb.Store.
func (s *StoresInfo) GetMetaStores() []*fidelpb.Store {
	stores := make([]*fidelpb.Store, 0, len(s.stores))
	for _, store := range s.stores {
		stores = append(stores, store.GetMeta())
	}
	return stores
}

// DeleteStore deletes tombstone record form store
func (s *StoresInfo) DeleteStore(store *StoreInfo) {
	delete(s.stores, store.GetID())
}

// GetStoreCount returns the total count of storeInfo.
func (s *StoresInfo) GetStoreCount() int {
	return len(s.stores)
}

// SetLeaderCount sets the leader count to a storeInfo.
func (s *StoresInfo) SetLeaderCount(storeID uint64, leaderCount int) {
	if store, ok := s.stores[storeID]; ok {
		s.stores[storeID] = store.Clone(SetLeaderCount(leaderCount))
	}
}

// SetRegionCount sets the region count to a storeInfo.
func (s *StoresInfo) SetRegionCount(storeID uint64, regionCount int) {
	if store, ok := s.stores[storeID]; ok {
		s.stores[storeID] = store.Clone(SetRegionCount(regionCount))
	}
}

// SetPendingPeerCount sets the pending count to a storeInfo.
func (s *StoresInfo) SetPendingPeerCount(storeID uint64, pendingPeerCount int) {
	if store, ok := s.stores[storeID]; ok {
		s.stores[storeID] = store.Clone(SetPendingPeerCount(pendingPeerCount))
	}
}

// SetLeaderSize sets the leader size to a storeInfo.
func (s *StoresInfo) SetLeaderSize(storeID uint64, leaderSize int64) {
	if store, ok := s.stores[storeID]; ok {
		s.stores[storeID] = store.Clone(SetLeaderSize(leaderSize))
	}
}

// SetRegionSize sets the region size to a storeInfo.
func (s *StoresInfo) SetRegionSize(storeID uint64, regionSize int64) {
	if store, ok := s.stores[storeID]; ok {
		s.stores[storeID] = store.Clone(SetRegionSize(regionSize))
	}
}

// UfidelateStoreStatus ufidelates the information of the store.
func (s *StoresInfo) UfidelateStoreStatus(storeID uint64, leaderCount int, regionCount int, pendingPeerCount int, leaderSize int64, regionSize int64) {
	if store, ok := s.stores[storeID]; ok {
		newStore := store.ShallowClone(SetLeaderCount(leaderCount),
			SetRegionCount(regionCount),
			SetPendingPeerCount(pendingPeerCount),
			SetLeaderSize(leaderSize),
			SetRegionSize(regionSize))
		s.SetStore(newStore)
	}
}

// IsTiFlashStore used to judge flash store.
// FIXME: remove the hack way
func IsTiFlashStore(store *fidelpb.Store) bool {
	for _, l := range store.GetLabels() {
		if l.GetKey() == "engine" && l.GetValue() == "tiflash" {
			return true
		}
	}
	return false
}
