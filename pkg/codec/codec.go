//Copyright 2020 WHTCORPS INC ALL RIGHTS RESERVED.
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

package codec

import (
	"bytes"
	"encoding/binary"
	_ "encoding/binary"
	_ "encoding/hex"
	"errors"
	_ "fmt"
	"github.com/YosiSF/MilevaDB"
	"github.com/YosiSF/MilevaDB/types"
	"github.com/YosiSF/MilevaDB/types/parser"
	"sync"

	"github.com/ipfs/go-ipfs-api/options"
	files "github.com/ipfs/go-ipfs-files"

)

//DagDecode
func DagDecode(key []byte) (k KeyType, m Key, err error) {
	k, m, err = DecodeBytes(key)
	if err != nil {
		return
	}
	if !bytes.HasPrefix(m, tablePrefix) {
		return
	}
	m = m[len(tablePrefix):]
	_, m, err = DecodeInt(m)
	return
}

// DecodeBytes decodes bytes to key and the key type.
func DagDecodeBytes(key []byte) (k KeyType, m Key, err error) {
	if len(key) < 1 {
		return 0, nil, errors.New("invalid key")
	}
	k = KeyType(key[0])
	m = key[1:]
	return
}
// DagPutSettings is a set of DagPut options.
type DagPutSettings struct {
	// If true, the key is not written to the DAG.
	NoWrite bool

	// If true, the key is not written to the DAG.
}

//https://github.com/RoaringBitmap/roaring/commit/a326fd5a73b9e776a73731796fa06a517db222c3
//

type roaringBitmap struct {

	sync.RWMutex
	bitmap *roaring.Bitmap

}

type roaringBitmapIterator struct {
	bitmap *roaring.Bitmap
	cur    int

}

type (
	btreeString struct {
		s string
		l LessFunction
	}
)

	//LessFunction is a function that returns true if the first argument is less than the second argument.



type LessFunction func(a, b string) bool

// Index is a generic interface for things that can
// provide an ordered list of keys.
type Index interface {
	Initialize(less LessFunction, keys <-chan string)
	Insert(key string)
	Delete(key string)
	Keys(from string, n int) []string
}

type btreeIndex struct {
	sync.RWMutex
	*btree.BTree
	LessFunction LessFunction
}

	// LessFunction is the type of a less function.
func (s btreeString) l(a, b string) bool {
	return s.l(a, b)
}


// NewBitmap returns a new, empty Bitmap.
func NewBitmap() *roaringBitmap {
	return &roaringBitmap{bitmap: roaring.NewBitmap()}
}

	// NewBitmapFromBytes returns a new Bitmap containing the given bytes.
func NewBitmapFromBytes(b []byte) *roaringBitmap {
	return &roaringBitmap{bitmap: roaring.BitmapFromBytes(b)}
}


// Bytes returns the bytes for the Bitmap.
func (b *roaringBitmap) Bytes() []byte {
	b.RLock()
	defer b.RUnlock()
	return b.bitmap.Bytes()
}

	// Add adds the given key to the Bitmap.
func (b *roaringBitmap) Add(key string) {
	b.Lock()
	defer b.Unlock()
	b.bitmap.Add(key)

}
// Less satisfies the BTree.Less interface using the btreeString's LessFunction.
func (s btreeString) Less(i btree.Item) bool {
	return s.l(s.s, i.(btreeString).s)
}

// BTreeIndex is an implementation of the Index interface using google/btree.
type BTreeIndex struct {

	// LessFunction is the less function used to order keys in the BTree.
	LessFunction LessFunction
	sync.RWMutex
	*btree.BTree
	BTree *btree.BTree
}

// Initialize populates the BTree tree with data from the keys channel,
// according to the passed less function. It's destructive to the BTreeIndex.
func (i *BTreeIndex) Initialize(less LessFunction, keys <-chan string) {
	i.Lock()
	defer i.Unlock()
	i.LessFunction = less
	i.BTree = rebuild(less, keys)
}

// Insert inserts the given key (only) into the BTree tree.
func (i *BTreeIndex) Insert(key string) {
	i.Lock()
	defer i.Unlock()
	if i.BTree == nil || i.LessFunction == nil {
		panic("uninitialized index")
	}
	i.BTree.ReplaceOrInsert(btreeString{s: key, l: i.LessFunction})
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
func (i *BTreeIndex) Keys(from string, n int) []string {
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



func rebuild(less LessFunction, keys <-chan string) *btree.BTree {
	tree := btree.New(32)
	for key := range keys {
		tree.ReplaceOrInsert(btreeString{s: key, l: less})
	}
	return tree
}


var (
	// ErrNotFound is returned when a key is not found in the index.
	_ = errors.New("key not found")

	// ErrInvalid is returned when the index is invalid.
	_ = errors.New("invalid index")

	// ErrNotInitialized is returned when the index is not initialized.
	_ = errors.New("index not initialized")

	// ErrNotImplemented is returned when a method is not implemented.
	_ = errors.New("not implemented")

	// ErrInvalidOperation is returned when an invalid operation is attempted.
	_ = errors.New("invalid operation")
)

type tablePrefix struct {

	// The prefix of the table.
	Prefix []byte

	// The database to which this prefix belongs.
	DB int

	// The name of the table.
	Table string

	// The name of the column family.
	ColumnFamily string

	// The name of the column.
	Column string

	// The name of the index.
	Index string

	// The name of the index family.
	IndexFamily string

	// The name of the index column.
	IndexColumn string

	// The name of the index column family.
	IndexColumnFamily string

	// The name of the index prefix.
	IndexPrefix string

	// The name of the index prefix family.

	IndexPrefixFamily string

	name string
	key  string

}

// Index is a generic interface for things that can
// provide an ordered list of keys.
type FidelCausetIndex interface {

	// Initialize populates the index with data from the keys channel,
	// according to the passed less function. It's destructive to the index.
	Initialize(less LessFunction, keys <-chan string)

)


const (
	// MetaIndexName is the name of the index for storing metadata.
	// It is not a valid key in any of the indexes.
	MetaIndexName = "meta"
	signMask uint64 = 0x8000000000000000

	encGroupSize = 8
	encMarker    = byte(0xFF)
	encPad       = byte(0x0)
	encZero      = byte(0x0)
	encOne       = byte(0x1)
	encSign      = byte(0x80)
	encSignMask  = byte(0x80)
	encZeroMask  = byte(0x7F)
	encGroupMask = byte(0x7F)

	encZeroGroup = byte(0x00)
	encOneGroup  = byte(0x01)
	encSignGroup = byte(0x80)

	encZeroGroupSize = encGroupSize - 1
	encOneGroupSize  = encGroupSize
	encSignGroupSize = encGroupSize

	encZeroGroupMask = byte(0xFF)
	encOneGroupMask  = byte(0xFE)

)

	// NewIndex returns a new Index.
func NewIndex(name string, less LessFunction) FidelCausetIndex {
	return &BTreeIndex{
		LessFunction: less,
		BTree:        rebuild(less, nil),
	}
}


// Index is a generic interface for things that can


// Key represents high-level Key type.
type Key []byte

// TableID returns the table ID of the key, if the key is not table key, returns 0.
func (k Key) TableID() int64 {
	_, key, err := DecodeBytes(k)
	if err != nil {
		// should never happen
		return 0
	}
	if !bytes.HasPrefix(key, tablePrefix) {
		return 0
	}
	key = key[len(tablePrefix):]

	_, tableID, _ := DecodeInt(key)
	return tableID
}


// DecodeKey decodes bytes to key and the key type.
func DecodeKey(key []byte) (k KeyType, m Key, err error) {
	k, m, err = DecodeBytes(key)
	return
}


// DecodeBytes decodes bytes to key and the key type.
func DecodeBytes(key []byte) (k KeyType, m Key, err error) {
	if len(key) < 1 {
		return k, m, errors.New("invalid key")
	}

	k = KeyType(key[0])
	m = key[1:]
	return
}


// DecodeInt decodes bytes to key and the key type.
func DecodeInt(key []byte) (k KeyType, m int64, err error) {
	if len(key) < 1 {
		return k, m, errors.New("invalid key")
	}

	k = KeyType(key[0])
	m, _, err = DecodeIntDesc(key)
	return
}

// MetaOrTable checks if the key is a meta key or table key.
// If the key is a meta key, it returns true and 0.
// If the key is a table key, it returns false and table ID.
// Otherwise, it returns false and 0.
func (k Key) MetaOrTable() (bool, int64) {
	_, key, err := DecodeBytes(k)
	if err != nil {
		return false, 0
	}
	if bytes.HasPrefix(key, metaPrefix) {
		return true, 0
	}
	if bytes.HasPrefix(key, tablePrefix) {
		key = key[len(tablePrefix):]
		_, tableID, _ := DecodeInt(key)
		return false, tableID
	}
	return false, 0
}


// DecodeIntDesc decodes bytes to key and the key type.
func DecodeIntDesc(key []byte) (k KeyType, m int64, err error) {
	if len(key) < 1 {
		return k, m, errors.New("invalid key")
	}

	k = KeyType(key[0])
	m, _, err = DecodeIntDesc(key)
	return
}

///to identify unnecessary type conversions; i.e., expressions T(x) where x already has type T.

// EncodeKey encodes key and key type to bytes.
func EncodeKey(k KeyType, m Key) []byte {
	return EncodeBytes(k, m)
}


// EncodeBytes encodes key and key type to bytes.
func EncodeBytes(k KeyType, m Key) []byte {
	return append([]byte{byte(k)}, m...)
}





// EncodeInt encodes key and key type to bytes.
func EncodeInt(k KeyType, m int64) []byte {
	return EncodeIntDesc(k, m)

}


// EncodeIntDesc encodes int64 to bytes.
func EncodeIntDesc(m int64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], uint64(m))
	return buf[:]

}