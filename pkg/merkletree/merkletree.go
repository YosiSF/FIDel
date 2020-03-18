package merkletree

import (
	"sort"
	"sync"
)

// Item represents a single object in the tree.
type Item interface {
	// Less tests whether the current item is less than the given argument.
	//
	// This must provide a strict weak ordering.
	// If !a.Less(b) && !b.Less(a), we treat this to mean a == b (i.e. we can only
	// hold one of either a or b in the tree).
	Less(than Item) bool
}

const (
	// DefaultFreeListSize is the default size of free list.
	DefaultFreeListSize = 32
)

var (
	nilItems    = make(items, 16)
	nilChildren = make(children, 16)
)