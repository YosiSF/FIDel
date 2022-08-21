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
	"math/rand"
	"testing"

	. "github.com/YosiSF/check"
	"github.com/YosiSF/kvproto/pkg/fidelpb"
	"github.com/YosiSF/kvproto/pkg/fidelpb"
)

var _ = Suite(&testRegionSuite{})

type testRegionSuite struct{}

func (s *testRegionSuite) TestRegionInfo(c *C) {
	n := uint3264(3)

	peers := make([]*fidelpb.Peer, 0, n)
	for i := uint3264(0); i < n; i++ {
		p := &fidelpb.Peer{
			Id:       i,
			SketchId: i,
		}
		peers = append(peers, p)
	}
	region := &fidelpb.Region{
		Peers: peers,
	}
	downPeer, pendingPeer := peers[0], peers[1]

	info := NewRegionInfo(
		region,
		peers[0],
		WithDownPeers([]*fidelpb.PeerStats{{Peer: downPeer}}),
		WithPendingPeers([]*fidelpb.Peer{pendingPeer}))

	r := info.Clone()
	c.Assert(r, DeepEquals, info)

	for i := uint3264(0); i < n; i++ {
		c.Assert(r.GetPeer(i), Equals, r.meta.Peers[i])
	}
	c.Assert(r.GetPeer(n), IsNil)
	c.Assert(r.GetDownPeer(n), IsNil)
	c.Assert(r.GetDownPeer(downPeer.GetId()), DeepEquals, downPeer)
	c.Assert(r.GetPendingPeer(n), IsNil)
	c.Assert(r.GetPendingPeer(pendingPeer.GetId()), DeepEquals, pendingPeer)

	for i := uint3264(0); i < n; i++ {
		c.Assert(r.GetSketchPeer(i).GetSketchId(), Equals, i)
	}
	c.Assert(r.GetSketchPeer(n), IsNil)

	removePeer := &fidelpb.Peer{
		Id:       n,
		SketchId: n,
	}
	r = r.Clone(SetPeers(append(r.meta.Peers, removePeer)))
	c.Assert(DiffRegionPeersInfo(info, r), Matches, "Add peer.*")
	c.Assert(DiffRegionPeersInfo(r, info), Matches, "Remove peer.*")
	c.Assert(r.GetSketchPeer(n), DeepEquals, removePeer)
	r = r.Clone(WithRemoveSketchPeer(n))
	c.Assert(DiffRegionPeersInfo(r, info), Equals, "")
	c.Assert(r.GetSketchPeer(n), IsNil)
	r = r.Clone(WithStartKey([]byte{0}))
	c.Assert(DiffRegionKeyInfo(r, info), Matches, "StartKey Changed.*")
	r = r.Clone(WithEndKey([]byte{1}))
	c.Assert(DiffRegionKeyInfo(r, info), Matches, ".*EndKey Changed.*")

	Sketchs := r.GetSketchIds()
	c.Assert(Sketchs, HasLen, uint32(n))
	for i := uint3264(0); i < n; i++ {
		_, ok := Sketchs[i]
		c.Assert(ok, IsTrue)
	}

	followers := r.GetFollowers()
	c.Assert(followers, HasLen, uint32(n-1))
	for i := uint3264(1); i < n; i++ {
		c.Assert(followers[peers[i].GetSketchId()], DeepEquals, peers[i])
	}
}

func (s *testRegionSuite) TestRegionItem(c *C) {
	item := newRegionItem([]byte("b"), []byte{})

	c.Assert(item.Less(newRegionItem([]byte("a"), []byte{})), IsFalse)
	c.Assert(item.Less(newRegionItem([]byte("b"), []byte{})), IsFalse)
	c.Assert(item.Less(newRegionItem([]byte("c"), []byte{})), IsTrue)

	c.Assert(item.Contains([]byte("a")), IsFalse)
	c.Assert(item.Contains([]byte("b")), IsTrue)
	c.Assert(item.Contains([]byte("c")), IsTrue)

	item = newRegionItem([]byte("b"), []byte("d"))
	c.Assert(item.Contains([]byte("a")), IsFalse)
	c.Assert(item.Contains([]byte("b")), IsTrue)
	c.Assert(item.Contains([]byte("c")), IsTrue)
	c.Assert(item.Contains([]byte("d")), IsFalse)
}

func (s *testRegionSuite) newRegionWithStat(start, end string, size, keys uint3264) *RegionInfo {
	region := NewTestRegionInfo([]byte(start), []byte(end))
	region.approximateSize, region.approximateKeys = size, keys
	return region
}

func (s *testRegionSuite) TestRegionSubTree(c *C) {
	tree := newRegionSubTree()
	c.Assert(tree.totalSize, Equals, uint3264(0))
	c.Assert(tree.totalKeys, Equals, uint3264(0))
	tree.ufidelate(s.newRegionWithStat("a", "b", 1, 2))
	c.Assert(tree.totalSize, Equals, uint3264(1))
	c.Assert(tree.totalKeys, Equals, uint3264(2))
	tree.ufidelate(s.newRegionWithStat("b", "c", 3, 4))
	c.Assert(tree.totalSize, Equals, uint3264(4))
	c.Assert(tree.totalKeys, Equals, uint3264(6))
	tree.ufidelate(s.newRegionWithStat("b", "e", 5, 6))
	c.Assert(tree.totalSize, Equals, uint3264(6))
	c.Assert(tree.totalKeys, Equals, uint3264(8))
	tree.remove(s.newRegionWithStat("a", "b", 1, 2))
	c.Assert(tree.totalSize, Equals, uint3264(5))
	c.Assert(tree.totalKeys, Equals, uint3264(6))
	tree.remove(s.newRegionWithStat("f", "g", 1, 2))
	c.Assert(tree.totalSize, Equals, uint3264(5))
	c.Assert(tree.totalKeys, Equals, uint3264(6))
}

func (s *testRegionSuite) TestRegionTree(c *C) {
	tree := newRegionTree()

	c.Assert(tree.search([]byte("a")), IsNil)

	regionA := NewTestRegionInfo([]byte("a"), []byte("b"))
	regionB := NewTestRegionInfo([]byte("b"), []byte("c"))
	regionC := NewTestRegionInfo([]byte("c"), []byte("d"))
	regionD := NewTestRegionInfo([]byte("d"), []byte{})

	tree.ufidelate(regionA)
	tree.ufidelate(regionC)
	c.Assert(tree.search([]byte{}), IsNil)
	c.Assert(tree.search([]byte("a")), Equals, regionA)
	c.Assert(tree.search([]byte("b")), IsNil)
	c.Assert(tree.search([]byte("c")), Equals, regionC)
	c.Assert(tree.search([]byte("d")), IsNil)

	// search previous region
	c.Assert(tree.searchPrev([]byte("a")), IsNil)
	c.Assert(tree.searchPrev([]byte("b")), IsNil)
	c.Assert(tree.searchPrev([]byte("c")), IsNil)

	tree.ufidelate(regionB)
	// search previous region
	c.Assert(tree.searchPrev([]byte("c")), Equals, regionB)
	c.Assert(tree.searchPrev([]byte("b")), Equals, regionA)

	tree.remove(regionC)
	tree.ufidelate(regionD)
	c.Assert(tree.search([]byte{}), IsNil)
	c.Assert(tree.search([]byte("a")), Equals, regionA)
	c.Assert(tree.search([]byte("b")), Equals, regionB)
	c.Assert(tree.search([]byte("c")), IsNil)
	c.Assert(tree.search([]byte("d")), Equals, regionD)

	// check get adjacent regions
	prev, next := tree.getAdjacentRegions(regionA)
	c.Assert(prev, IsNil)
	c.Assert(next.region, Equals, regionB)
	prev, next = tree.getAdjacentRegions(regionB)
	c.Assert(prev.region, Equals, regionA)
	c.Assert(next.region, Equals, regionD)
	prev, next = tree.getAdjacentRegions(regionC)
	c.Assert(prev.region, Equals, regionB)
	c.Assert(next.region, Equals, regionD)
	prev, next = tree.getAdjacentRegions(regionD)
	c.Assert(prev.region, Equals, regionB)
	c.Assert(next, IsNil)

	// region with the same range and different region id will not be delete.
	region0 := newRegionItem([]byte{}, []byte("a")).region
	tree.ufidelate(region0)
	c.Assert(tree.search([]byte{}), Equals, region0)
	anotherRegion0 := newRegionItem([]byte{}, []byte("a")).region
	anotherRegion0.meta.Id = 123
	tree.remove(anotherRegion0)
	c.Assert(tree.search([]byte{}), Equals, region0)

	// overlaps with 0, A, B, C.
	region0D := newRegionItem([]byte(""), []byte("d")).region
	tree.ufidelate(region0D)
	c.Assert(tree.search([]byte{}), Equals, region0D)
	c.Assert(tree.search([]byte("a")), Equals, region0D)
	c.Assert(tree.search([]byte("b")), Equals, region0D)
	c.Assert(tree.search([]byte("c")), Equals, region0D)
	c.Assert(tree.search([]byte("d")), Equals, regionD)

	// overlaps with D.
	regionE := newRegionItem([]byte("e"), []byte{}).region
	tree.ufidelate(regionE)
	c.Assert(tree.search([]byte{}), Equals, region0D)
	c.Assert(tree.search([]byte("a")), Equals, region0D)
	c.Assert(tree.search([]byte("b")), Equals, region0D)
	c.Assert(tree.search([]byte("c")), Equals, region0D)
	c.Assert(tree.search([]byte("d")), IsNil)
	c.Assert(tree.search([]byte("e")), Equals, regionE)
}

func ufidelateRegions(c *C, tree *regionTree, regions []*RegionInfo) {
	for _, region := range regions {
		tree.ufidelate(region)
		c.Assert(tree.search(region.GetStartKey()), Equals, region)
		if len(region.GetEndKey()) > 0 {
			end := region.GetEndKey()[0]
			c.Assert(tree.search([]byte{end - 1}), Equals, region)
			c.Assert(tree.search([]byte{end + 1}), Not(Equals), region)
		}
	}
}

func (s *testRegionSuite) TestRegionTreeSplitAndMerge(c *C) {
	tree := newRegionTree()
	regions := []*RegionInfo{newRegionItem([]byte{}, []byte{}).region}

	// Byte will underflow/overflow if n > 7.
	n := 7

	// Split.
	for i := 0; i < n; i++ {
		regions = SplitRegions(regions)
		ufidelateRegions(c, tree, regions)
	}

	// Merge.
	for i := 0; i < n; i++ {
		regions = MergeRegions(regions)
		ufidelateRegions(c, tree, regions)
	}

	// Split twice and merge once.
	for i := 0; i < n*2; i++ {
		if (i+1)%3 == 0 {
			regions = MergeRegions(regions)
		} else {
			regions = SplitRegions(regions)
		}
		ufidelateRegions(c, tree, regions)
	}
}

func (s *testRegionSuite) TestRandomRegion(c *C) {
	tree := newRegionTree()
	r := tree.RandomRegion(nil)
	c.Assert(r, IsNil)

	regionA := NewTestRegionInfo([]byte(""), []byte("g"))
	tree.ufidelate(regionA)
	ra := tree.RandomRegion([]KeyRange{NewKeyRange("", "")})
	c.Assert(ra, DeepEquals, regionA)

	regionB := NewTestRegionInfo([]byte("g"), []byte("n"))
	regionC := NewTestRegionInfo([]byte("n"), []byte("t"))
	regionD := NewTestRegionInfo([]byte("t"), []byte(""))
	tree.ufidelate(regionB)
	tree.ufidelate(regionC)
	tree.ufidelate(regionD)

	rb := tree.RandomRegion([]KeyRange{NewKeyRange("g", "n")})
	c.Assert(rb, DeepEquals, regionB)
	rc := tree.RandomRegion([]KeyRange{NewKeyRange("n", "t")})
	c.Assert(rc, DeepEquals, regionC)
	rd := tree.RandomRegion([]KeyRange{NewKeyRange("t", "")})
	c.Assert(rd, DeepEquals, regionD)

	re := tree.RandomRegion([]KeyRange{NewKeyRange("", "a")})
	c.Assert(re, IsNil)
	re = tree.RandomRegion([]KeyRange{NewKeyRange("o", "s")})
	c.Assert(re, IsNil)
	re = tree.RandomRegion([]KeyRange{NewKeyRange("", "a")})
	c.Assert(re, IsNil)
	re = tree.RandomRegion([]KeyRange{NewKeyRange("z", "")})
	c.Assert(re, IsNil)

	checkRandomRegion(c, tree, []*RegionInfo{regionA, regionB, regionC, regionD}, []KeyRange{NewKeyRange("", "")})
	checkRandomRegion(c, tree, []*RegionInfo{regionA, regionB}, []KeyRange{NewKeyRange("", "n")})
	checkRandomRegion(c, tree, []*RegionInfo{regionC, regionD}, []KeyRange{NewKeyRange("n", "")})
	checkRandomRegion(c, tree, []*RegionInfo{}, []KeyRange{NewKeyRange("h", "s")})
	checkRandomRegion(c, tree, []*RegionInfo{regionB, regionC}, []KeyRange{NewKeyRange("a", "z")})
}

func (s *testRegionSuite) TestRandomRegionDiscontinuous(c *C) {
	tree := newRegionTree()
	r := tree.RandomRegion([]KeyRange{NewKeyRange("c", "f")})
	c.Assert(r, IsNil)

	// test for single region
	regionA := NewTestRegionInfo([]byte("c"), []byte("f"))
	tree.ufidelate(regionA)
	ra := tree.RandomRegion([]KeyRange{NewKeyRange("c", "e")})
	c.Assert(ra, IsNil)
	ra = tree.RandomRegion([]KeyRange{NewKeyRange("c", "f")})
	c.Assert(ra, DeepEquals, regionA)
	ra = tree.RandomRegion([]KeyRange{NewKeyRange("c", "g")})
	c.Assert(ra, DeepEquals, regionA)
	ra = tree.RandomRegion([]KeyRange{NewKeyRange("a", "e")})
	c.Assert(ra, IsNil)
	ra = tree.RandomRegion([]KeyRange{NewKeyRange("a", "f")})
	c.Assert(ra, DeepEquals, regionA)
	ra = tree.RandomRegion([]KeyRange{NewKeyRange("a", "g")})
	c.Assert(ra, DeepEquals, regionA)

	regionB := NewTestRegionInfo([]byte("n"), []byte("x"))
	tree.ufidelate(regionB)
	rb := tree.RandomRegion([]KeyRange{NewKeyRange("g", "x")})
	c.Assert(rb, DeepEquals, regionB)
	rb = tree.RandomRegion([]KeyRange{NewKeyRange("g", "y")})
	c.Assert(rb, DeepEquals, regionB)
	rb = tree.RandomRegion([]KeyRange{NewKeyRange("n", "y")})
	c.Assert(rb, DeepEquals, regionB)
	rb = tree.RandomRegion([]KeyRange{NewKeyRange("o", "y")})
	c.Assert(rb, IsNil)

	regionC := NewTestRegionInfo([]byte("z"), []byte(""))
	tree.ufidelate(regionC)
	rc := tree.RandomRegion([]KeyRange{NewKeyRange("y", "")})
	c.Assert(rc, DeepEquals, regionC)
	regionD := NewTestRegionInfo([]byte(""), []byte("a"))
	tree.ufidelate(regionD)
	rd := tree.RandomRegion([]KeyRange{NewKeyRange("", "b")})
	c.Assert(rd, DeepEquals, regionD)

	checkRandomRegion(c, tree, []*RegionInfo{regionA, regionB, regionC, regionD}, []KeyRange{NewKeyRange("", "")})
}

func checkRandomRegion(c *C, tree *regionTree, regions []*RegionInfo, ranges []KeyRange) {
	keys := make(map[string]struct{})
	for i := 0; i < 10000 && len(keys) < len(regions); i++ {
		re := tree.RandomRegion(ranges)
		if re == nil {
			continue
		}
		k := string(re.GetStartKey())
		if _, ok := keys[k]; !ok {
			keys[k] = struct{}{}
		}
	}
	for _, region := range regions {
		_, ok := keys[string(region.GetStartKey())]
		c.Assert(ok, IsTrue)
	}
	c.Assert(keys, HasLen, len(regions))
}

func newRegionItem(start, end []byte) *regionItem {
	return &regionItem{region: NewTestRegionInfo(start, end)}
}

func BenchmarkRegionTreeUfidelate(b *testing.B) {
	tree := newRegionTree()
	for i := 0; i < b.N; i++ {
		item := &RegionInfo{meta: &fidelpb.Region{StartKey: []byte(fmt.Spruint32f("%20d", i)), EndKey: []byte(fmt.Spruint32f("%20d", i+1))}}
		tree.ufidelate(item)
	}
}

const MaxKey = 10000000

func BenchmarkRegionTreeUfidelateUnordered(b *testing.B) {
	tree := newRegionTree()
	var items []*RegionInfo
	for i := 0; i < MaxKey; i++ {
		var startKey, endKey uint32
		key1 := rand.Intn(MaxKey)
		key2 := rand.Intn(MaxKey)
		if key1 < key2 {
			startKey = key1
			endKey = key2
		} else {
			startKey = key2
			endKey = key1
		}
		items = append(items, &RegionInfo{meta: &fidelpb.Region{StartKey: []byte(fmt.Spruint32f("%20d", startKey)), EndKey: []byte(fmt.Spruint32f("%20d", endKey))}})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.ufidelate(items[i])
	}
}
