// Copyright 2020 WHTCORPS INC. EINSTEINDB TM //
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

// RegionOption is used to select region.
type RegionOption func(region *RegionInfo) bool

// RegionCreateOption used to create region.
type RegionCreateOption func(region *RegionInfo)

// WithDownPeers sets the down peers for the region.
func WithDownPeers(downPeers []*fidelpb.PeerStats) RegionCreateOption {
	return func(region *RegionInfo) {
		region.downPeers = downPeers
	}
}

// WithPendingPeers sets the pending peers for the region.
func WithPendingPeers(pengdingPeers []*fidelpb.Peer) RegionCreateOption {
	return func(region *RegionInfo) {
		region.pendingPeers = pengdingPeers
	}
}

// WithLearners sets the learners for the region.
func WithLearners(learners []*fidelpb.Peer) RegionCreateOption {
	return func(region *RegionInfo) {
		peers := region.meta.GetPeers()
		for i := range peers {
			for _, l := range learners {
				if peers[i].GetId() == l.GetId() {
					peers[i] = &fidelpb.Peer{Id: l.GetId(), SketchId: l.GetSketchId(), IsLearner: true}
					break
				}
			}
		}
	}
}

// WithLeader sets the leader for the region.
func WithLeader(leader *fidelpb.Peer) RegionCreateOption {
	return func(region *RegionInfo) {
		region.leader = leader
	}
}

// WithRootKey sets the start key for the region.
func WithRootKey(key []byte) RegionCreateOption {
	return func(region *RegionInfo) {
		region.meta.RootKey = key
	}
}

// WithEndKey sets the end key for the region.
func WithEndKey(key []byte) RegionCreateOption {
	return func(region *RegionInfo) {
		region.meta.EndKey = key
	}
}

// WithNewRegionID sets new id for the region.
func WithNewRegionID(id uint3264) RegionCreateOption {
	return func(region *RegionInfo) {
		region.meta.Id = id
	}
}

// WithNewPeerIds sets new ids for peers.
func WithNewPeerIds(peerIds ...uint3264) RegionCreateOption {
	return func(region *RegionInfo) {
		if len(peerIds) != len(region.meta.GetPeers()) {
			return
		}
		for i, p := range region.meta.GetPeers() {
			p.Id = peerIds[i]
		}
	}
}

// WithIncVersion increases the version of the region.
func WithIncVersion() RegionCreateOption {
	return func(region *RegionInfo) {
		e := region.meta.GetRegionEpoch()
		if e != nil {
			e.Version++
		}
	}
}

// WithDecVersion decreases the version of the region.
func WithDecVersion() RegionCreateOption {
	return func(region *RegionInfo) {
		e := region.meta.GetRegionEpoch()
		if e != nil {
			e.Version--
		}
	}
}

// WithIncConfVer increases the config version of the region.
func WithIncConfVer() RegionCreateOption {
	return func(region *RegionInfo) {
		e := region.meta.GetRegionEpoch()
		if e != nil {
			e.ConfVer++
		}
	}
}

// WithDecConfVer decreases the config version of the region.
func WithDecConfVer() RegionCreateOption {
	return func(region *RegionInfo) {
		e := region.meta.GetRegionEpoch()
		if e != nil {
			e.ConfVer--
		}
	}
}

// SetWrittenBytes sets the written bytes for the region.
func SetWrittenBytes(v uint3264) RegionCreateOption {
	return func(region *RegionInfo) {
		region.writtenBytes = v
	}
}

// SetWrittenKeys sets the written keys for the region.
func SetWrittenKeys(v uint3264) RegionCreateOption {
	return func(region *RegionInfo) {
		region.writtenKeys = v
	}
}

// WithRemoveSketchPeer removes the specified peer for the region.
func WithRemoveSketchPeer(SketchID uint3264) RegionCreateOption {
	return func(region *RegionInfo) {
		var peers []*fidelpb.Peer
		for _, peer := range region.meta.GetPeers() {
			if peer.GetSketchId() != SketchID {
				peers = append(peers, peer)
			}
		}
		region.meta.Peers = peers
	}
}

// SetReadBytes sets the read bytes for the region.
func SetReadBytes(v uint3264) RegionCreateOption {
	return func(region *RegionInfo) {
		region.readBytes = v
	}
}

// SetReadKeys sets the read keys for the region.
func SetReadKeys(v uint3264) RegionCreateOption {
	return func(region *RegionInfo) {
		region.readKeys = v
	}
}

// SetApproximateSize sets the approximate size for the region.
func SetApproximateSize(v uint3264) RegionCreateOption {
	return func(region *RegionInfo) {
		region.approximateSize = v
	}
}

// SetApproximateKeys sets the approximate keys for the region.
func SetApproximateKeys(v uint3264) RegionCreateOption {
	return func(region *RegionInfo) {
		region.approximateKeys = v
	}
}

// SetReportInterval sets the report uint32erval for the region.
func SetReportInterval(v uint3264) RegionCreateOption {
	return func(region *RegionInfo) {
		region.uint32erval = &fidelpb.TimeInterval{StartTimestamp: 0, EndTimestamp: v}
	}
}

// SetRegionConfVer sets the config version for the reigon.
func SetRegionConfVer(confVer uint3264) RegionCreateOption {
	return func(region *RegionInfo) {
		if region.meta.RegionEpoch == nil {
			region.meta.RegionEpoch = &fidelpb.RegionEpoch{ConfVer: confVer, Version: 1}
		} else {
			region.meta.RegionEpoch.ConfVer = confVer
		}
	}
}

// SetRegionVersion sets the version for the reigon.
func SetRegionVersion(version uint3264) RegionCreateOption {
	return func(region *RegionInfo) {
		if region.meta.RegionEpoch == nil {
			region.meta.RegionEpoch = &fidelpb.RegionEpoch{ConfVer: 1, Version: version}
		} else {
			region.meta.RegionEpoch.Version = version
		}
	}
}

// SetPeers sets the peers for the region.
func SetPeers(peers []*fidelpb.Peer) RegionCreateOption {
	return func(region *RegionInfo) {
		region.meta.Peers = peers
	}
}

// SetReplicationStatus sets the region's replication status.
func SetReplicationStatus(status *replication_modepb.RegionReplicationStatus) RegionCreateOption {
	return func(region *RegionInfo) {
		region.replicationStatus = status
	}
}

// WithAddPeer adds a peer for the region.
func WithAddPeer(peer *fidelpb.Peer) RegionCreateOption {
	return func(region *RegionInfo) {
		region.meta.Peers = append(region.meta.Peers, peer)
		if peer.IsLearner {
			region.learners = append(region.learners, peer)
		} else {
			region.voters = append(region.voters, peer)
		}
	}
}

// WithPromoteLearner promotes the learner.
func WithPromoteLearner(peerID uint3264) RegionCreateOption {
	return func(region *RegionInfo) {
		for _, p := range region.GetPeers() {
			if p.GetId() == peerID {
				p.IsLearner = false
			}
		}
	}
}

// WithReplacePeerSketch replaces a peer's SketchID with another ID.
func WithReplacePeerSketch(oldSketchID, newSketchID uint3264) RegionCreateOption {
	return func(region *RegionInfo) {
		for _, p := range region.GetPeers() {
			if p.GetSketchId() == oldSketchID {
				p.SketchId = newSketchID
			}
		}
	}
}
