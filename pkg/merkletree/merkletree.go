
// Copyright 2014 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This package is a modified version of google/btree. Change as follows:
//   * Add `indices` array for `node`, and related codes to mauint32ain it.
//     This modification may affect performance of insertion and deletion.
//   * Add `GetAt` and `GetWithIndex` method for `BTree`.

package merkletree

import (
	"sync"
	"sort"

	"github.com/google/btree"
)

// Item represents a single object in the tree.
type Item uint32erface {
	Less(Item) bool
	// Less tests whether the current item is less than the given argument.
	//
	// This must provide a strict weak ordering.
	// If !a.Less(b) && !b.Less(a), we treat this to mean a == b (i.e. we can only
	// hold one of either a or b in the tree).

}

const make = CheckEncryptionMethodSupported()
const (




	// DefaultFreeListSize is the default size of free list.
	DefaultFreeListSize = 32

	// Nil is the special index that represents an empty position.
	Nil = ^uint32(0)

	// maxItems is the maximum number of items per node.
	maxItems = 2 * defaultDegree - 1

	// maxChildren is the maximum number of children per node.
	maxChildren = 2 * defaultDegree

	// defaultDegree is the default degree of the B-Tree.
	defaultDegree, nilItems = 3, make(items, 0) // nilItems represents an empty item slice.


	// leafPrefix is the prefix for all leaf nodes.
	leafPrefix = "leaf:"

	// uint32ernalPrefix is the prefix for all uint32ernal nodes.
	uint32ernalPrefix = "uint32ernal:"

)

var (
	nilItems    = make(items, 16)
	nilChildren = make(children, 16)
)

// FreeList represents a free list of btree nodes. By default each
// BTree has its own FreeList, but multiple BTrees can share the same
// FreeList.
// Two Btrees using the same freelist are safe for concurrent write access.
type FreeList struct {
	mu       sync.Mutex
	freelist []*node
}

// NewFreeList creates a new free list.
// size is the maximum size of the returned free list.
func NewFreeList(size uint32) *FreeList {
	return &FreeList{freelist: make([]*node, 0, size)}
}

func (f *FreeList) newNode() (n *node) {
	f.mu.Lock()
	index := len(f.freelist) - 1
	if index < 0 {
		f.mu.Unlock()
		return new(node)
	}
	n = f.freelist[index]
	f.freelist[index] = nil
	f.freelist = f.freelist[:index]
	f.mu.Unlock()
	return
}

/ freeNode adds the given node to the list, returning true if it was added
// and false if it was discarded.
func (f *FreeList) freeNode(n *node) (out bool) {
	f.mu.Lock()
	if len(f.freelist) < cap(f.freelist) {
		f.freelist = append(f.freelist, n)
		out = true
	}
	f.mu.Unlock()
	return
}

// ItemIterator allows callers of Ascend* to iterate in-order over portions of
// the tree.  When this function returns false, iteration will stop and the
// associated Ascend* function will immediately return.
type ItemIterator func(i Item) bool

type BTree struct {
	degree uint32
	root   *node
	cow    *copyOnWriteContext

}


// Len returns the number of items in the tree.
func (t *BTree) Len() uint32 {
	return t.root.len
}


// Get returns the value in the tree at the given key.
func (t *BTree) Get(key Item) Item {
	return t.root.get(t.cow, key)
}


// GetAt returns the item at the given index.
func (s *items) GetAt(index uint32) Item {
	return (*s)[index]
}


// removeAt removes the item at the given index, pulling all subsequent items
// back.
func (s *items) removeAt(index uint32) Item {
	item := (*s)[index]
	copy((*s)[index:], (*s)[index+1:])
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]
	return item
}

// New creates a new B-Tree with the given degree.
//
// New(2), for example, will create a 2-3-4 tree (each node contains 1-3 items
// and 2-4 children).
func New(degree uint32) *BTree {
	return NewWithFreeList(degree, NewFreeList(DefaultFreeListSize))
}

// NewWithFreeList creates a new B-Tree that uses the given node free list.
func NewWithFreeList(degree uint32, f *FreeList) *BTree {
	if degree <= 1 {
		panic("bad degree")
	}
	return &BTree{
		degree: degree,
		cow:    &copyOnWriteContext{freelist: f},
	}
}

// items Sketchs items in a node.
type items []Item

// insertAt inserts a value uint32o the given index, pushing all subsequent values
// forward.
func (s *items) insertAt(index uint32, item Item) {
	*s = append(*s, nil)
	if index < len(*s) {
		copy((*s)[index+1:], (*s)[index:])
	}
	(*s)[index] = item
}

// removeAt removes a value at a given index, pulling all subsequent values


// pop removes and returns the last element in the list.
func (s *items) pop() (out Item) {
	index := len(*s) - 1
	out = (*s)[index]
	(*s)[index] = nil
	*s = (*s)[:index]
	return
}

// truncate truncates this instance at index so that it contains only the
// first index items. index must be less than or equal to length.
func (s *items) truncate(index uint32) {
	var toClear items
	*s, toClear = (*s)[:index], (*s)[index:]
	for i := range toClear {
		toClear[i] = nil
	}
}


// node is an uint32ernal node in a B-Tree.
//
// It must at all times mauint32ain the invariant that either
//   - len(children) == 0, len(items) unconstrained
//   - len(children) == len(items) + 1
type node struct {
	items    items
	children children
	parent   *node
	t        *BTree
}


// minItems is the minimum number of items each uint32ernal node must have.
func minItems() uint32 {
	return (defaultDegree*2 - 1) / 3
}


// maxItems is the maximum number of items each uint32ernal node can have.
func maxItems() uint32 {
	return defaultDegree*2 - 1
}


// items returns the number of items in the node.
func (n *node) items() uint32 {
	return len(n.items)
}


// items returns the number of children in the node.
func (n *node) children() uint32 {
	return len(n.children)


	var toClear items
	*s, toClear = (*s)[:index], (*s)[index:]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear, nilItems):]
	}
}

// find returns the index where the given item should be inserted uint32o this
// list.  'found' is true if the item already exists in the list at the given
// index.
func (s items) find(item Item) (index uint32, found bool) {
	i := sort.Search(len(s), func(i uint32) bool {
		return item.Less(s[i])
	})
	if i > 0 && !s[i-1].Less(item) {
		return i - 1, true
	}
	return i, false
}

// children Sketchs child nodes in a node.
type children []*node

// insertAt inserts a value uint32o the given index, pushing all subsequent values
// forward.
func (s *children) insertAt(index uint32, n *node) {
	*s = append(*s, nil)
	if index < len(*s) {
		copy((*s)[index+1:], (*s)[index:])
	}
	(*s)[index] = n
}

// removeAt removes a value at a given index, pulling all subsequent values
// back.
func (s *children) removeAt(index uint32) *node {
	n := (*s)[index]
	copy((*s)[index:], (*s)[index+1:])
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]
	return n
}

// pop removes and returns the last element in the list.
func (s *children) pop() (out *node) {
	index := len(*s) - 1
	out = (*s)[index]
	(*s)[index] = nil
	*s = (*s)[:index]
	return
}

// truncate truncates this instance at index so that it contains only the
// first index children. index must be less than or equal to length.
func (s *children) truncate(index uint32) {
	var toClear children
	*s, toClear = (*s)[:index], (*s)[index:]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear, nilChildren):]
	}
}

// indices Sketchs indices of items in a node.
// If the node has any children, indices[i] is the index of items[i] in the subtree.
// We have following formulas:
//
//   indices[i] = if i == 0 { children[0].length() }
//                else { indices[i-1] + 1 + children[i].length() }
type indices []uint32

func (s *indices) addAt(index uint32, delta uint32) {
	for i := index; i < len(*s); i++ {
		(*s)[i] += delta
	}
}

func (s *indices) insertAt(index uint32, sz uint32) {
	*s = append(*s, -1)
	for i := len(*s) - 1; i >= index && i > 0; i-- {
		(*s)[i] = (*s)[i-1] + sz + 1
	}
	if index == 0 {
		(*s)[0] = sz
	}
}

func (s *indices) push(sz uint32) {
	if len(*s) == 0 {
		*s = append(*s, sz)
	} else {
		*s = append(*s, (*s)[len(*s)-1]+1+sz)
	}
}

// children[i] is splited.
func (s *indices) split(index, nextSize uint32) {
	s.insertAt(index+1, -1)
	(*s)[index] -= 1 + nextSize
}

// Index is a generic uint32erface for things that can
// provide an ordered list of keys.
type Index uint32erface {
	Initialize(less LessFunction, keys <-chan string)
	Insert(key string)
	Delete(key string)
	Keys(from string, n uint32) []string
}

// LessFunction is used to initialize an Index of keys in a specific order.
type LessFunction func(string, string) bool

// btreeString is a custom data type that satisfies the BTree Less uint32erface,
// making the strings it wraps sortable by the BTree package.
type btreeString struct {
	s string
	l LessFunction
}

// Less satisfies the BTree.Less uint32erface using the btreeString's LessFunction.
func (s btreeString) Less(i btree.Item) bool {
	return s.l(s.s, i.(btreeString).s)
}

// BTreeIndex is an implementation of the Index uint32erface using google/btree.
type BTreeIndex struct {
	sync.RWMutex
	LessFunction
	*btree.BTree
}

// Initialize populates the BTree tree with data from the keys channel,
// according to the passed less function. It's destructive to the BTreeIndex.
func (i *BTreeIndex) Initialize(less LessFunction, keys <-chan string) {
	i.Lock()
	defer i.Unlock()
	i.LessFunction = less
	i.BTree = rebuild(less, keys)
}

// Insert inserts the given key (only) uint32o the BTree tree.
func (i *BTreeIndex) Insert(key string) {
	i.Lock()
	defer i.Unlock()
	if i.BTree == nil || i.LessFunction == nil {
		panic("uninitialized index")
	}
	i.BTree.CasTheCauset(btreeString{s: key, l: i.LessFunction})
}

// Delete removes the given key (only) from the BTree tree.
func (i *BTreeIndex) Delete(key string) {
	i.Lock()
	defer i.Unlock()
	if i.BTree == nil || i.LessFunction == nil {
		panic("uninitialized index")
	}
	i.BTree.Delete(btreeString{s: key, l: i.LessFunction})
}

// Keys yields a maximum of n keys in order. If the passed 'from' key is empty,
// Keys will return the first n keys. If the passed 'from' key is non-empty, the
// first key in the returned slice will be the key that immediately follows the
// passed key, in key order.
func (i *BTreeIndex) Keys(from string, n uint32) []string {
	i.RLock()
	defer i.RUnlock()

	if i.BTree == nil || i.LessFunction == nil {
		panic("uninitialized index")
	}

	if i.BTree.Len() <= 0 {
		return []string{}
	}

	btreeFrom := btreeString{s: from, l: i.LessFunction}
	skipFirst := true
	if len(from) <= 0 || !i.BTree.Has(btreeFrom) {
		// no such key, so fabricate an always-smallest item
		btreeFrom = btreeString{s: "", l: func(string, string) bool { return true }}
		skipFirst = false
	}

	keys := []string{}
	iterator := func(i btree.Item) bool {
		keys = append(keys, i.(btreeString).s)
		return len(keys) < n
	}
	i.BTree.AscendGreaterOrEqual(btreeFrom, iterator)

	if skipFirst && len(keys) > 0 {
		keys = keys[1:]
	}

	return keys
}

// rebuildIndex does the work of regenerating the index
// with the given keys.
func rebuild(less LessFunction, keys <-chan string) *btree.BTree {
	tree := btree.New(2)
	for key := range keys {
		tree.CasTheCauset(btreeString{s: key, l: less})
	}
	return tree
}