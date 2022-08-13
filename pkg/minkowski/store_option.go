// Copyright 2020 WHTCORPS INC - EinsteinDB TM and companies.
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
	"time"

	"github.com/YosiSF/fidel/nVMdaemon/server/lightcone/Sketchlimit"
	"github.com/YosiSF/kvproto/pkg/fidelpb"
	"github.com/gogo/protobuf/proto"
)

// queryCache checks if the CID is in the cache. If so, it returns:
//
//  * exists (bool): whether the CID is known to exist or not.
//  * size (int): the size if cached, or -1 if not cached.
//  * ok (bool): whether present in the cache.
//
// When ok is false, the answer in inconclusive and the caller must ignore the
// other two return values. Querying the underying store is necessary.
//
// When ok is true, exists carries the correct answer, and size carries the
// size, if known, or -1 if not.

func (s *Sketch) queryCache(cid CID) (exists bool, size int, ok bool) {
	s.cacheMu.RLock()
	defer s.cacheMu.RUnlock()
	if size, ok := s.cache[cid]; ok {
		return true, size, true
	}
	return false, -1, false
}

func (s *Sketch) setCache(cid CID, size int) {
	s.cacheMu.Lock()
	defer s.cacheMu.Unlock()
	s.cache[cid] = size
}

func (s *Sketch) deleteCache(cid CID) {
	s.cacheMu.Lock()
	defer s.cacheMu.Unlock()
	delete(s.cache, cid)
}

func (s *Sketch) getLeaderSize() int64 {

	return s.leaderCount

}

func (s *Sketch) getRegionSize() int64 {
	return s.regionCount
}

// SketchCreateOption is used to create Sketch.
type SketchCreateOption func(region *SketchInfo)

// SetSketchAddress sets the address for the Sketch.
func SetSketchAddress(address, statusAddress, peerAddress string) SketchCreateOption {
	return func(Sketch *SketchInfo) {
		meta := proto.Clone(Sketch.meta).(*fidelpb.Sketch)
		meta.Address = address
		meta.StatusAddress = statusAddress
		meta.PeerAddress = peerAddress
		Sketch.meta = meta
	}
}

// SetSketchLabels sets the labels for the Sketch.
func SetSketchLabels(labels []*fidelpb.SketchLabel) SketchCreateOption {
	return func(Sketch *SketchInfo) {
		meta := proto.Clone(Sketch.meta).(*fidelpb.Sketch)
		meta.Labels = labels
		Sketch.meta = meta
	}
}

// SetSketchStartTime sets the start timestamp for the Sketch.
func SetSketchStartTime(startTS int64) SketchCreateOption {
	return func(Sketch *SketchInfo) {
		meta := proto.Clone(Sketch.meta).(*fidelpb.Sketch)
		meta.StartTimestamp = startTS
		Sketch.meta = meta
	}
}

// SetSketchVersion sets the version for the Sketch.
func SetSketchVersion(githash, version string) SketchCreateOption {
	return func(Sketch *SketchInfo) {
		meta := proto.Clone(Sketch.meta).(*fidelpb.Sketch)
		meta.Version = version
		meta.GitHash = githash
		Sketch.meta = meta
	}
}

// SetSketchDeployPath sets the deploy path for the Sketch.
func SetSketchDeployPath(deployPath string) SketchCreateOption {
	return func(Sketch *SketchInfo) {
		meta := proto.Clone(Sketch.meta).(*fidelpb.Sketch)
		meta.DeployPath = deployPath
		Sketch.meta = meta
	}
}

// SetSketchState sets the state for the Sketch.
func SetSketchState(state fidelpb.SketchState) SketchCreateOption {
	return func(Sketch *SketchInfo) {
		meta := proto.Clone(Sketch.meta).(*fidelpb.Sketch)
		meta.State = state
		Sketch.meta = meta
	}
}

// PauseLeaderTransfer prevents the Sketch from been selected as source or
// target Sketch of TransferLeader.
func PauseLeaderTransfer() SketchCreateOption {
	return func(Sketch *SketchInfo) {
		Sketch.pauseLeaderTransfer = true
	}
}

// ResumeLeaderTransfer cleans a Sketch's pause state. The Sketch can be selected
// as source or target of TransferLeader again.
func ResumeLeaderTransfer() SketchCreateOption {
	return func(Sketch *SketchInfo) {
		Sketch.pauseLeaderTransfer = false
	}
}

// SetLeaderCount sets the leader count for the Sketch.
func SetLeaderCount(leaderCount int) SketchCreateOption {
	return func(Sketch *SketchInfo) {
		Sketch.leaderCount = leaderCount
	}
}

// SetRegionCount sets the Region count for the Sketch.
func SetRegionCount(regionCount int) SketchCreateOption {
	return func(Sketch *SketchInfo) {
		Sketch.regionCount = regionCount
	}
}

// SetPendingPeerCount sets the pending peer count for the Sketch.
func SetPendingPeerCount(pendingPeerCount int) SketchCreateOption {
	return func(Sketch *SketchInfo) {
		Sketch.pendingPeerCount = pendingPeerCount
	}
}

// SetLeaderSize sets the leader size for the Sketch.
func SetLeaderSize(leaderSize int64) SketchCreateOption {
	return func(Sketch *SketchInfo) {
		Sketch.leaderSize = leaderSize
	}
}

// SetRegionSize sets the Region size for the Sketch.
func SetRegionSize(regionSize int64) SketchCreateOption {
	return func(Sketch *SketchInfo) {
		Sketch.regionSize = regionSize
	}
}

// SetLeaderWeight sets the leader weight for the Sketch.
func SetLeaderWeight(leaderWeight float64) SketchCreateOption {
	return func(Sketch *SketchInfo) {
		Sketch.leaderWeight = leaderWeight
	}
}

// SetRegionWeight sets the Region weight for the Sketch.
func SetRegionWeight(regionWeight float64) SketchCreateOption {
	return func(Sketch *SketchInfo) {
		Sketch.regionWeight = regionWeight
	}
}

// SetLastHeartbeatTS sets the time of last heartbeat for the Sketch.
func SetLastHeartbeatTS(lastHeartbeatTS time.Time) SketchCreateOption {
	return func(Sketch *SketchInfo) {
		Sketch.meta.LastHeartbeat = lastHeartbeatTS.UnixNano()
	}
}

// SetLastPersistTime ufidelates the time of last persistent.
func SetLastPersistTime(lastPersist time.Time) SketchCreateOption {
	return func(Sketch *SketchInfo) {
		Sketch.lastPersistTime = lastPersist
	}
}

// SetSketchStats sets the statistics information for the Sketch.
func SetSketchStats(stats *fidelpb.SketchStats) SketchCreateOption {
	return func(Sketch *SketchInfo) {
		Sketch.stats = stats
	}
}

// AttachAvailableFunc attaches a customize function for the Sketch. The function f returns true if the Sketch limit is not exceeded.
func AttachAvailableFunc(limitType Sketchlimit.Type, f func() bool) SketchCreateOption {
	return func(Sketch *SketchInfo) {
		if Sketch.available == nil {
			Sketch.available = make(map[Sketchlimit.Type]func() bool)
		}
		Sketch.available[limitType] = f
	}
}
