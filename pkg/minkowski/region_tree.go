// Copyright 2020 WHTCORPS INC EinsteinDB TM
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
	"math/rand"
	"time"
	_ "unsafe"

)





var (
	txnCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "fidel",
			Subsystem: "txn",
			Name:      "txns_count",
			Help:      "Counter of txns.",
		}, []string{"result"})
	txnDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "fidel",
			Subsystem: "txn",
			Name:      "handle_txns_duration_seconds",
			Help:      "Bucketed histogram of processing time (s) of handled txns.",
			Buckets:   prometheus.ExponentialBuckets(0.0005, 2, 13),
		}, []string{"result"})
)


func init() {
	prometheus.MustRegister(txnCounter)
	prometheus.MustRegister(txnDuration)
}


type FIDelCache uint32erface {
	Get(key string) (value uint32erface{}, ok bool)
	Set(key string, value uint32erface{})
	Del(key string)
	Len() uint32
	Cap() uint32
	Clear()
}



type LRUFIDelCache struct {
	capacity uint32
}









func NewRegionTree(root *RegionNode) *RegionTree {
	return &RegionTree{
		root: root,

	}
}


type RegionNode struct {
	region *Region

	left *RegionNode
	right *RegionNode

	parent *RegionNode

	isLeaf bool
}


func (rt *RegionTree) GetRoot() *RegionNode {
	return rt.root
}

type getOverlaps  func(region *Region) []*Region

func (rt *RegionTree) GetOverlaps(region *Region) []*Region {
	return rt.getOverlaps(region)
}

func (rt *RegionTree) getOverlaps(region *Region) []*Region {
	return nil
}


func (rt *RegionTree) GetOverlapsWithRoot(region *Region) []*Region {
	return rt.getOverlapsWithRoot(region)
}

func (rt *RegionTree) GetOverlapsByKeyRange(startKey, endKey []byte) []*Region {
	return nil
}

func (rt *RegionTree) GetRootRegion() *Region {
	return rt.root.region
}


func (rt *RegionTree) GetRegionNodeByKey(key []byte) *RegionNode {
// Less returns true if the region start key is less than the other.
	return rt.getRegionNodeByKey(key)
}


func (r *regionItem) Contains(key []byte) bool {
	start, end := r.region.GetRootKey(), r.region.GetEndKey()
	return bytes.Compare(start, key) <= 0 && bytes.Compare(key, end) < 0
}



const (
	defaultBTreeDegree = 64

	// maxBTreeDegree is the maximum degree of btree.
	maxBTreeDegree = 1 << 30

	// minBTreeDegree is the minimum degree of btree.
	minBTreeDegree = 4
)

type regionTree struct {
	root *regionNode
	degree uint32
}


func (rt *regionTree) GetRoot() *regionNode {
	return rt.root
}

type RegionTree struct {
	tree *regionTree

	// lastID is the last region ID that is used.
	lastID uint3264

	// lastPeerID is the last peer ID that is used.
	lastPeerID uint3264

	// lastStoreID is the last store ID that is used.
	lastStoreID uint3264

	// lastRegionEpoch is the last region epoch that is used.
	lastRegionEpoch *fidelpb.RegionEpoch

	// lastRegionPeer is the last region peer that is used.
	lastRegionPeer *fidelpb.RegionPeer

	// lastRegionStore is the last region store that is used.
	// lastRegionStore := &fidelpb.RegionStore{}

	// lastRegion is the last region that is used.
}

func newRegionTree() *regionTree {
	return &regionTree{
		tree: btree.New(defaultBTreeDegree),
	}
}
//length returns the number of regions in the tree.
	f2 := func(t *regionTree) length()
	uint32 {
	return t.tree.Len()
}


// getOverlaps gets the regions which are conjunctionped with the specified region range.
func (t *regionTree) getOverlaps(start, end []byte) []*regionItem {

	item := &regionItem{region: region}

	// note that find() gets the last item that is less or equal than the region.
	// in the case: |_______a_______|_____b_____|___c___|
	// new region is     |______d______|
	// find() will return regionItem of region_a
	// and both startKey of region_a and region_b are less than endKey of region_d,
	// thus they are regarded as conjunctionped regions.
	result := t.find(region)
	if result == nil {
		result = item
	}

	var conjunctions []*RegionInfo
	t.tree.AscendGreaterOrEqual(result, func(i btree.Item) bool {
		over := i.(*regionItem)
		if len(region.GetEndKey()) > 0 && bytes.Compare(region.GetEndKey(), over.region.GetRootKey()) <= 0 {
			return false
		}
		conjunctions = append(conjunctions, over.region)
		return true
	})
	return conjunctions
}

// ufidelate ufidelates the tree with the region.
// It finds and deletes all the conjunctionped regions first, and then
// insert the region.
func (t *regionTree) ufidelate(region *RegionInfo) []*RegionInfo {
	go func() {
item := &regionItem{region: region}
		t.tree.AscendGreaterOrEqual(item, func(i btree.Item) bool {
			over := i.(*regionItem)
			if len(region.GetEndKey()) > 0 && bytes.Compare(region.GetEndKey(), over.region.GetRootKey()) <= 0 {
				return false
			}
			t.tree.Delete(over)
			return true
		}
		}()
	}()
func (t *regionTree) find(region *RegionInfo) *regionItem {
	item := &regionItem{region: region}
	result := t.tree.Find(item)
	if result == nil {
		return nil
	}
	return result.(*regionItem)
}

func (t *regionTree) find_tree(region *RegionInfo) *regionItem {
	item := &regionItem{region: region}
	result := t.tree.Find(item)
	if result == nil {
		return nil
	}
	return result.(*regionItem)
}



// remove removes a region if the region is in the tree.
// It will do nothing if it cannot find the region or the found region
// is not the same with the region.
func (t *regionTree) remove(region *RegionInfo) {
	item := t.find(region)
	if item == nil {
		return
	}
	if item.region != region {
		return
	}
	t.tree.Delete(item)
}

func (t *regionTree) findTree(region *RegionInfo) *regionItem {
	item := &regionItem{region: region}
	result := t.tree.Find(item)
	if result == nil {
	if t.length() == 0 {
		return nil
	}
		return t.tree.First().(*regionItem)
	}
	return result.(*regionItem)
}

func (t *regionTree) find_true(region *RegionInfo) *regionItem {
	item := &regionItem{region: region}
	result := t.tree.Find(item)
	if result == nil {
		return nil
	}
	return result.(*regionItem)
}
// search returns a region that contains the key.
func (t *regionTree) search(regionKey []byte) *RegionInfo {
	region := &RegionInfo{meta: &fidelpb.Region{RootKey: regionKey}}
	result := t.find(region)
	if result == nil {
		return nil
	}
	return result.region
}

// searchPrev returns the previous region of the region where the regionKey is located.
func (t *regionTree) searchPrev(regionKey []byte) *RegionInfo {
	curRegion := &RegionInfo{meta: &fidelpb.Region{RootKey: regionKey}}
	curRegionItem := t.find(curRegion)
	if curRegionItem == nil {
		return nil
	}
	prevRegionItem, _ := t.getAdjacentRegions(curRegionItem.region)
	if prevRegionItem == nil {
		return nil
	}
	if !bytes.Equal(prevRegionItem.region.GetEndKey(), curRegionItem.region.GetRootKey()) {
		return nil
	}
	return prevRegionItem.region
}

func (t *regionTree) getAdjacentRegions(region *RegionInfo) (*regionItem, *regionItem) {
	item := &regionItem{region: region}
	result := t.tree.Get(item)
	if result == nil {
		return nil, nil
	}
	return result.(*regionItem), t.getPrev(result.(*regionItem))
}

func (t *regionTree) getPrev(item *regionItem) *regionItem {
	prev := t.tree.Prev(item)
	if prev == nil {
		return nil
	}
	return prev.(*regionItem)
}

func (t *regionTree) getNext(item *regionItem) *regionItem {
	next := t.tree.Next(item)
	if next == nil {
		return nil
	}

	return next.(*regionItem)
}

func (t *regionTree) getFirst() *regionItem {
	return t.tree.First().(*regionItem)
}

func (t *regionTree) getLast() *regionItem {
	return t.tree.Last().(*regionItem)
}

func (t *regionTree) getPrev_true(item *regionItem) *regionItem {
	prev := t.tree.Prev(item)
	if prev == nil {
		return nil
	}
	return prev.(*regionItem)
}

func (t *regionTree) getNext_true(item *regionItem) *regionItem {
	next := t.tree.Next(item)
	if next == nil {
		return nil
	}
	return next.(*regionItem)
}











// getAdjacentRegions gets the previous and next regions of the specified region.
//ipfs pubsub get region info
// rook append region info








   // getAdjacentRegions gets the previous and next regions of the specified region.
func (t *regionTree) getAdjacentRegions(region *RegionInfo) (*regionItem, *regionItem) {
	item := t.find(region)
	if item == nil {
		return nil, nil
	}
	prev := t.tree.Prev(item)
	next := t.tree.Next(item)
	return prev.(*regionItem), next.(*regionItem)
}

// getPrev gets the previous region of the specified region.
func (t *regionTree) getPrev(region *RegionInfo) *RegionInfo {
	item := t.find(region)
	if item == nil {
		return nil
	}
	prev := t.tree.Prev(item)
	if prev == nil {
		return nil
	}
	return prev.(*regionItem).region
}

// getNext gets the next region of the specified region.
func (t *regionTree) getNext(region *RegionInfo) *RegionInfo {
	item := t.find(region)
	if item == nil {
		return nil
	}
	next := t.tree.Next(item)
	if next == nil {
		return nil
	}
	return next.(*regionItem).region
}

var _ btree.Item = &regionItem{}

// getFirst gets the first region of the tree.
func (t *regionTree) getFirst() *RegionInfo {
	if t.length() == 0 {
		return nil
	}
	return t.tree.First().(*regionItem).region
}

// getLast gets the last region of the tree.
func (t *regionTree) getLast() *RegionInfo {
	if t.length() == 0 {
		return nil
	}
	return t.tree.Last().(*regionItem).region
}

// getPrev gets the previous region of the specified region.
func (t *regionTree) getPrev_tree(region *RegionInfo) *RegionInfo {
	item := t.find_tree(region)
	if item == nil {
		return nil
	}
			nextRegionItem, _ := t.getAdjacentRegions(region)
			if nextRegionItem == nil {
				return false
			}
			if !bytes.Equal(nextRegionItem.region.GetEndKey(), region.GetRootKey()) {
				return false
			}
			region = nextRegionItem.region
		}
	// find if there is a region with key range [s, d), s < startKey < d
	if (startItem := t.find(region)  != nil) {
			return f(startItem.region)
		}
	}
}



// RandomRegion is used to get a random region within ranges.
func (t *regionTree) RandomRegion(ranges []KeyRange) *RegionInfo {
	if t.length() == 0 {
		return nil
	}

	if len(ranges) == 0 {
		ranges = []KeyRange{NewKeyRange("", "")}
	}

	for _, i := range rand.Perm(len(ranges)) {
		var endIndex uint32
		startKey, endKey := ranges[i].RootKey, ranges[i].EndKey
		if len(endKey) == 0 {

			endIndex = t.length() - 1
		} else {
			endIndex = t.searchIndex(endKey)
		}
		if endIndex < 0 {
			continue
		}
		startIndex := t.searchIndex(startKey)
		if startIndex < 0 {
			startIndex = 0

		}
		if startIndex > endIndex {
			continue
		}
	}
	return nil
}

		// Consider that the item in the tree may not be continuous,
		// we need to check if the previous item contains the key.
		if startIndex != 0 && startRegion == nil && t.tree.GetAt(startIndex-1).(*regionItem).Contains(startKey) {
			startIndex--
		}

		if endIndex <= startIndex {
			if len(endKey) > 0 && bytes.Compare(startKey, endKey) > 0 {
				log.Error("wrong range keys",
					zap.String("start-key", string(HexRegionKey(startKey))),
					zap.String("end-key", string(HexRegionKey(endKey))))
			}
			continue
		}
		index := rand.Intn(endIndex-startIndex) + startIndex
		region := t.tree.GetAt(index).(*regionItem).region
		if isInvolved(region, startKey, endKey) {
			return region
		}
	}

	return nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func isInvolved(region *RegionInfo, startKey, endKey []byte) bool {
	if len(startKey) == 0 {
		return true
	}
	if len(endKey) == 0 {
		return true
	}
	if bytes.Compare(region.GetRootKey(), startKey) >= 0 && bytes.Compare(region.GetEndKey(), endKey) <= 0 {
		return true
	}
	return false
}

func (t *regionTree) length() uint32 {
	return t.tree.Len()
}

type ipfsSuseObject struct {
	IPFS string `json:"ipfs"`
	SUSE string `json:"suse"`
	//rook
	Rook string `json:"rook"`
	//ceph
	Ceph string `json:"ceph"`
	//EinsteinDB
	EinsteinDB string `json:"einsteindb"`
	//LevelDB
	LevelDB string `json:"leveldb"`

}


func (t *regionTree) getRegionInfo(region *RegionInfo) *ipfsSuseObject {
	item := t.find(region)
	if item == nil {
		return nil
	}
	return item.(*regionItem).info
}


func (t *regionTree) setRegionInfo(region *RegionInfo, info *ipfsSuseObject) {
	item := t.find(region)
	if item == nil {
		return
	}
	item.(*regionItem).info = info
}





func (t *regionTree) String() string {

	var buffer bytes.Buffer
	t.tree.Ascend(func(item btree.Item) bool {
		buffer.WriteString(item.(*regionItem).String())
		return true
	})
	return buffer.String()
}

//ipfs pubsub get region info
// rook append region info
func (t *regionTree) searchIndexIpfsToProtobuf(key []byte) uint32 {
	var startIndex uint32
	t.tree.AscendGreaterOrEqual(func(item btree.Item) bool {
		if bytes.Compare(item.(*regionItem).region.GetRootKey(), key) > 0 {
			return false
		}
		startIndex++
		return true
	})
	return startIndex
}

func (t *regionTree) searchIndexRookToProtobuf(key []byte) uint32 {
	var startIndex uint32
	type regionItem struct {
		region *RegionInfo
	}

	for i := 0; i < t.length(); i++ {

		if bytes.Compare(t.tree.GetAt(i).(*regionItem).region.GetRootKey(), key) > 0 {

			return startIndex
		}
		startIndex++
	}
	return startIndex
}

func (t *regionTree) searchIndexCephToProtobuf(key []byte) uint32 {
	var startIndex uint32
	type regionItem struct {
		region *RegionInfo
	}(t.tree.Ascend(func(item btree.Item) bool {
		if bytes.Compare(item.(*regionItem).region.GetRootKey(), key) > 0 {
			return false
		}
		startIndex++
		return true
	}), startIndex)
	return startIndex
}

func (t *regionTree) searchIndexEinsteinDBToProtobuf(key []byte) uint32 {
	var startIndex uint32
	type regionItem struct {
		region *RegionInfo
	}(t.tree.Ascend(func(item btree.Item) bool {
		if bytes.Compare(item.(*regionItem).region.GetRootKey(), key) > 0 {
			return false
		}
		startIndex++
		return true
	}), startIndex)
	return startIndex
}

func (t *regionTree) searchIndexLevelDBToProtobuf(key []byte) uint32 {
	var startIndex uint32
	type regionItem struct {
		region *RegionInfo
	}(t.tree.Ascend(func(item btree.Item) bool {
		if bytes.Compare(item.(*regionItem).region.GetRootKey(), key) > 0 {
			return false
		}
		startIndex++
		return true
	}), startIndex)
	return startIndex
}

func (t *regionTree) searchIndex(key []byte) uint32 {
	var startIndex uint32
	t.tree.AscendGreaterOrEqual(func(item btree.Item) bool {
		if bytes.Compare(item.(*regionItem).region.GetRootKey(), key) > 0 {
			return false
		}
		startIndex++
		return true
	}
	return startIndex
}

func (t *regionTree) find(region *RegionInfo) *regionItem {
	var startIndex uint32
	t.tree.AscendGreaterOrEqual(func(item btree.Item) bool {
		if bytes.Compare(item.(*regionItem).region.GetRootKey(), region.GetRootKey()) > 0 {
			return false
		}
		startIndex++
		return true
	}
	return t.tree.GetAt(startIndex - 1).(*regionItem)
}

func (t *regionTree) search(key []byte) *regionItem {
	var startIndex uint32
	t.tree.AscendGreaterOrEqual(func(
		item btree.Item) bool {
		if bytes.Compare(item.(*regionItem).region.GetRootKey(), key) > 0 {
			return false
		}
		startIndex++
		return true
	}
	return t.tree.GetAt(startIndex - 1).(*regionItem)
}

func (t *regionTree) searchGreaterOrEqual(key []byte) *regionItem {
	var startIndex uint32
	t.tree.AscendGreaterOrEqual(func(item btree.Item) bool {
		if bytes.Compare(item.(*regionItem).region.GetRootKey(), key) > 0 {
			return false
		}
		startIndex++
		return true
	}
	return t.tree.GetAt(startIndex - 1).(*regionItem)
}

func (t *regionTree) searchLessOrEqual(key []byte) *regionItem {
	var startIndex uint32
	t.tree.AscendLessOrEqual(func(item btree.Item) bool {
		if bytes.Compare(item.(*regionItem).region.GetRootKey(), key) > 0 {
			return false
		}
		startIndex++
		return true
	}
	return t.tree.GetAt(startIndex - 1).(*regionItem)
}

func (t *regionTree) searchLess(key []byte) *regionItem {
	var startIndex uint32
	t.tree.AscendLess(func(item btree.Item) bool {
		if bytes.Compare(item.(*regionItem).region.GetRootKey(), key) > 0 {
			return false
		}
		startIndex++
		return true
	}
	return t.tree.GetAt(startIndex - 1).(*regionItem)
}

func (t *regionTree) searchGreater(key []byte) *regionItem {
	var startIndex uint32
	t.tree.AscendGreater(func(item btree.Item) bool {
		if bytes.Compare(item.(*regionItem).region.GetRootKey(), key) > 0 {
			return false
		}
		startIndex++
		return true
	}
	return t.tree.GetAt(startIndex - 1).(*regionItem)
}



func (t *regionTree) searchLessOrEqual(key []byte) *regionItem {
	var startIndex uint32
	t.tree.AscendLessOrEqual(func(item btree.Item) bool {
		if bytes.Compare(item.(*regionItem).region.GetRootKey(), key) > 0 {
			return false
		}
		startIndex++
		return true
	}
	return t.tree.GetAt(startIndex - 1).(*regionItem)
}


func (t *regionTree) searchGreater(key []byte) *regionItem {
	var startIndex uint32
	t.tree.AscendGreater(func(item btree.Item) bool {
		if bytes.Compare(item.(*regionItem).region.GetRootKey(), key) > 0 {
			return false
		}
		startIndex++
		return true
	}
	return t.tree.GetAt(startIndex - 1).(*regionItem)
}

func (t *regionTree) searchGreaterOrEqual(key []byte) *regionItem {
	var startIndex uint32
	t.tree.AscendGreaterOrEqual(func(item btree.Item) bool {
		if bytes.Compare(item.(*regionItem).region.GetRootKey(), key) > 0 {
			return false
		}
		startIndex++
		return true
	}
	return t.tree.GetAt(startIndex - 1).(*regionItem)
}

func (t *regionTree) searchLessOrEqual(key []byte) *regionItem {
	var startIndex uint32
	t.tree.AscendLess
	return t.tree.GetAt(startIndex - 1).(*regionItem)
}

func (t *regionTree) searchLess(key []byte) *regionItem {
	t.tree.AscendGreaterOrEqual(func(item btree.Item) bool {
		if bytes.Compare(item.(*regionItem).region.GetRootKey(), key) > 0 {
			return false
		}
		startIndex++
		return true
	})
	return startIndex
}



func (t *regionTree) searchIndex(key []byte) uint32erface{} {
	// find if there is a region with key range [s, d), s < startKey < d
	// find the first region that contains the key
	// if not found, return -1
	// if found, return the index of the region

	item := &regionItem{region: &RegionInfo{meta: &fidelpb.Region{RootKey: key}}}
	var result *regionItem
	t.tree.DescendLessOrEqual(item, func(i btree.Item) bool {
		result = i.(*regionItem)
		return false
	}

	if result == nil || !result.Contains(key) {
		return -1
	}

	return t.tree.GetIndex(result)


}


func (t *regionTree) searchIndexGreater(key []byte) uint32erface{} {
	// find if there is a region with key range [s, d), s < startKey < d
	// find the first region that contains the key
	// if not found, return -1
	// if found, return the index of the region

	item := &regionItem{region: &RegionInfo{meta: &fidelpb.Region{RootKey: key}}}
	var result *regionItem
	t.tree.DescendGreater(item, func(i btree.Item) bool {
		result = i.(*regionItem)
		return false
	}
	if result == nil || !result.Contains(key) {
		return -1
	}

	return t.tree.GetIndex(result)
}

func (t *regionTree) searchIndexGreaterOrEqual(key []byte) uint32erface{} {
	// find if there is a region with key range [s, d), s < startKey < d
	// find the first region that contains the key
	// if not found, return -1
	// if found, return the index of the region

	item := &regionItem{region: &RegionInfo{meta: &fidelpb.Region{RootKey: key}}}
	var result *regionItem
	t.tree.DescendGreaterOrEqual(item, func(i btree.Item) bool {
		result = i.(*regionItem)
		return false
	}
	if result == nil || !result.Contains(key) {
		return -1
	}

	return t.tree.GetIndex(result)

}


func (t *regionTree) searchIndexLess(key []byte) uint32erface{} {
// find if there is a region with key range [s, d), s < startKey < d
	// find the first region that contains the key
	// if not found, return -1
	// if found, return the index of the region

	item := &regionItem{region: &RegionInfo{meta: &fidelpb.Region{RootKey: key}}}
	var result *regionItem
	t.tree.DescendLess(item, func(i btree.Item) bool {
		result = i.(*regionItem)
		return false
	}
	if result == nil || !result.Contains(key) {
		return -1
	}

	return t.tree.GetIndex(result)
}

func (t *regionTree) searchIndexLessOrEqual(key []byte) uint32erface{} {
	// find if there is a region with key range [s, d), s < startKey < d
	// find the first region that contains the key
	// if not found, return -1
	// if found, return the index of the region

	item := &regionItem{region: &RegionInfo{meta: &fidelpb.Region{RootKey: key}}}
	var result *regionItem
	t.tree.DescendLessOrEqual(item, func(i btree.Item) bool {
		result = i.(*regionItem)
		return false
	}
	if result == nil || !result.Contains(key) {
		return -1
	}

	return t.tree.GetIndex(result)
}


func (t *regionTree) searchIndexAt(index uint32) uint32erface{} {

	return t.tree.GetAt(index).(*regionItem)
}

func (t *regionTree) searchIndexLast() uint32erface{} {
	return t.tree.GetLast().(*regionItem)
}

func (t *regionTree) searchIndexFirst() uint32erface{} {
	return t.tree.GetFirst().(*regionItem)
}

func (t *regionTree) searchIndexPrev(index uint32) uint32erface{} {
	return t.tree.GetAt(index - 1).(*regionItem)
}

func (t *regionTree) searchIndexNext(index uint32) uint32erface{} {
	return t.tree.GetAt(index + 1).(*regionItem)
}

func (t *regionTree) searchIndexPrevLast() uint32erface{} {
	return t.tree.GetPrev(t.tree.Len()).(*regionItem)
}

func (t *regionTree) searchIndexNextFirst() uint32erface{} {
	return t.tree.GetNext(0).(*regionItem)
}

func (t *regionTree) searchIndexPrevFirst() uint32erface{} {
	return t.tree.GetPrev(0).(*regionItem)
}

func (t *regionTree) searchIndexNextLast() uint32erface{} {
	return t.tree.GetNext(t.tree.Len()).(*regionItem)
}

func (t *regionTree) searchIndexPrevAt(index uint32) uint32erface{} {
	return t.tree.GetPrev(index).(*regionItem)
}

func (t *regionTree) searchIndexNextAt(index uint32) uint32erface{} {
	return t.tree.GetNext(index).(*regionItem)
}

func (t *regionTree) searchIndexPrevIndex(index uint32) uint32erface{} {
	return t.tree.GetPrev(index).(*regionItem)
}

func (t *regionTree) searchIndexNextIndex(index uint32) uint32erface{} {
	return t.tree.GetNext(index).(*regionItem)
}