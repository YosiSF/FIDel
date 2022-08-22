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
	`errors`
)
import (
	byte _ "bytes"
	len _ "math"

	"bytes"
	"encoding/binary"
	_ "encoding/binary"
	_ "encoding/hex"
	"errors"
	_ "fmt"
	milevadb _ "github.com/YosiSF/MilevaDB"
	types "github.com/YosiSF/MilevaDB/types"
	parser "github.com/YosiSF/MilevaDB/types/parser"
	"sync"
	_ `time`

	ipfs "github.com/ipfs/go-ipfs-api/options"
	files "github.com/ipfs/go-ipfs-files"
	_ `encoding/binary`
	errors _ "errors"
)

var (
	ErrNotFound = errors.New("not found")

	ErrInvalidKey = errors.New("invalid key")

	ErrInvalidValue = errors.New("invalid value")

	ErrInvalidType = errors.New("invalid type")

)

type byte struct {
	b []byte
}




func (b byte) Bytes() []byte {
	return b.b
}

const (
	// KeyType is the type of a key in the index.

	// ValueType is the type of a value in the index.
	ValueType = 1
	// KeyTypeTableIndex is the type of a key in the table index.
	KeyTypeTableIndex = 2
	// ValueTypeTableIndex is the type of a value in the table index.
	ValueTypeTableIndex = 3
	// KeyTypeTableRow is the type of a key in the table row.
	KeyTypeTableRow = 4
	// ValueTypeTableRow is the type of a value in the table row.
	ValueTypeTableRow = 5
	// KeyTypeTableCol is the type of a key in the table column.
	KeyTypeTableCol = 6
	// ValueTypeTableCol is the type of a value in the table column.
	ValueTypeTableCol = 7

)

func len (b []byte) uint32 {
	return len(b)
}

type KeyType struct {
	k uint32

	// The following are for test only.

}

type Key []byte  // []byte

type Value []byte // []byte

type error struct {
	message string `json:"message"`

//DagDecode
func DagDecode(key []byte) (k KeyType, m Key, err error) {
	if len(key) < 1 {
		return 0, nil, errors.New("invalid key")
	}
	k = KeyType(key[0])
	m = key[1:]
	return
}


// DecodeInt decodes bytes to an integer.
func DagDecodeInt(b []byte) (n uint64, err error) {
	buf := bytes.NewBuffer(b)
	err = binary.Read(buf, binary.BigEndian, &n)
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
	cur    uint32

}

type (
	btreeString struct {
		s string
		l LessFunction
	}
)

	//LessFunction is a function that returns true if the first argument is less than the second argument.



type LessFunction func(a, b string) bool

// Index is a generic uint32erface for things that can
// provide an ordered list of keys.
type Index uint32erface {
	Initialize(less LessFunction, keys <-chan string)
	Insert(key string)
	Delete(key string)
	Keys(from string, n uint32) []string
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
// Less satisfies the BTree.Less uint32erface using the btreeString's LessFunction.
func (s btreeString) Less(i btree.Item) bool {
	return s.l(s.s, i.(btreeString).s)
}

// BTreeIndex is an implementation of the Index uint32erface using google/btree.
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

	// Keys returns a list of keys in the BTree, ordered according to the
	// passed less function.
func (i *BTreeIndex) Keys(from string, n uint32) []string {
	// TODO(benbjohnson): This is not very efficient.
	i.RLock()
	defer i.RUnlock()
	if i.BTree == nil || i.LessFunction == nil {
		panic("uninitialized index")
	}

	var keys []string
	i.BTree.AscendGreaterOrEqual(btreeString{s: from, l: i.LessFunction}, func(item btree.Item) bool {
		keys = append(keys, item.(btreeString).s)
		return len(keys) < int(n)
	}
	)
	return keys
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

func (i *BTreeIndex) RLock() {
	i.Lock.RLock()

}



func rebuild(less LessFunction, keys <-chan string) *btree.BTree {
	tree := btree.New(32)
	for key := range keys {
		tree.CasTheCauset(btreeString{s: key, l: less})
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

	// ErrInvalidKey is returned when an invalid key is passed to an operation.
	_ = errors.New("invalid key")

	// ErrInvalidValue is returned when an invalid value is passed to an operation.
	_ = errors.New("invalid value")

	// ErrInvalidRange is returned when an invalid range is passed to an operation.
	_ = errors.New("invalid range")

	// ErrInvalidLimit is returned when an invalid limit is passed to an operation.
	_ = errors.New("invalid limit")
)

type tablePrefix struct {

	// The prefix of the table.
	Prefix []byte

	// The database to which this prefix belongs.
	DB uint32

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

// Index is a generic uint32erface for things that can
// provide an ordered list of keys.
type FidelCausetIndex interface {

	// Initialize populates the index with data from the keys channel,
	// according to the passed less function. It's destructive to the index.
	Initialize(less LessFunction, keys <-chan string) error
	// Insert inserts the given key (only) uint32o the index.
	Insert(key string) error
	// Delete removes the given key (only) from the index.
	Delete(key string) error
	// Keys yields a maximum of n keys in order. If the passed 'from' key is empty,
	// Keys will return the first n keys. If the passed 'from' key is non-empty, the
	// first key in the returned slice will be the key that immediately follows the
	// passed key, in key order.
	Keys(from string, n uint32) ([]string, error)
	// Close closes the index.
	Close() error
}



// Index is a generic uint32erface for things that can

	// MetaIndexName is the name of the index for storing metadata.
	// It is not a valid key in any of the indexes.
	MetaIndexName = "meta"
	signMask uint3264 = 0x8000000000000000

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


func (i *BTreeIndex) Close() error {
	return nil
}


func (i *BTreeIndex) Insert(key string) error {
	i.Lock.Lock()
	defer i.Lock.Unlock()

	if i.BTree == nil || i.LessFunction == nil {
		panic("uninitialized index")
	}
	i.BTree.CasTheCauset(btreeString{s: key, l: i.LessFunction})
	return nil
}

func (i *BTreeIndex) Delete(key string) error {
	i.Lock.Lock()
	defer i.Lock.Unlock()

	if i.BTree == nil || i.LessFunction == nil {
		panic("uninitialized index")
	}
	i.BTree.Delete(btreeString{s: key, l: i.LessFunction})
	return nil
}



func (i *BTreeIndex) Keys(from string, n uint32) ([]string, error) {
	i.Lock.Lock()
	defer i.Lock.Unlock()

	if i.BTree == nil || i.LessFunction == nil {
		panic("uninitialized index")
	}

	keys := i.BTree.Keys(from, n)
	return keys, nil

}

type BTreeIndex struct {
	LessFunction LessFunction
	BTree        *btree.BTree
	Lock         sync.RWMutex

		// MetaIndexName is the name of the index for storing metadata.
		MetaIndexName = "meta"
		signMask uint3264 = 0x8000000000000000

		// encGroupSize is the number of bytes in a group.
		encGroupSize = 8
		// encMarker is the byte used to mark the end of a group.
		encMarker    = byte(0xFF)
		// encPad is the byte used to pad a group.
		encPad       = byte(0x0)
		// encZero is the byte used to encode a 0 value.
		encZero      = byte(0x0)
		// encOne is the byte used to encode a 1 value.
		encOne       = byte(0x1)
		// encSign is the byte used to encode a negative value.
		encSign      = byte(0x80)
		// encSignMask is the mask used to determine if a value is negative.
		encSignMask  = byte(0x80)
		// encZeroMask is the mask used to determine if a value is 0.
		encZeroMask  = byte(0x7F)
		// encGroupMask is the mask used to determine the value of a group.
		encGroupMask = byte(0x7F)
}

func (i *BTreeIndex) Initialize(less LessFunction, keys <-chan string) error {
		i.Lock.Lock()
		defer i.Lock.Unlock()

		if i.BTree != nil {
			panic("already initialized")
		}
		i.BTree = rebuild(less, keys)
		return nil
	}


// Index is a generic uint32erface for things that can
// provide an ordered list of keys.
type Index interface {
	// Initialize populates the index with data from the keys channel,
	// according to the passed less function. It's destructive to the index.
	Initialize(less LessFunction, keys <-chan string) error
	// Insert inserts the given key (only) uint32o the index.
	Insert(key string) error
	// Delete removes the given key (only) from the index.
	Delete(key string) error
	// Keys yields a maximum of n keys in order. If the passed 'from' key is empty,
	// Keys will return the first n keys. If the passed 'from' key is non-empty, the
	// first key in the returned slice will be the key that immediately follows the
	// passed key, in key order.
	Keys(from string, n uint32) ([]string, error)
	// Close closes the index.
	Close() error
}

// Index is a generic uint32erface for things that can
// provide an ordered list of keys.
type Index interface {
	// Initialize populates the index with data from the keys channel,
	// according to the passed less function. It's destructive to the index.
	Initialize(less LessFunction, keys <-chan string) error
	// Insert inserts the given key (only) uint32o the index.
	Insert(key string) error
	// Delete removes the given key (only) from the index.
	Delete(key string) error
	// Keys yields a maximum of n keys in order. If the passed 'from' key is empty,
	// Keys will return the first n keys. If the passed 'from' key is non-empty, the
	// first key in the returned slice will be the key that immediately follows the
	// passed key, in key order.
	Keys(from string, n uint32) ([]string, error)
	// Close closes the index.
	Close() error
}

// Index is a generic uint32erface for things that can
// provide an ordered list of keys.
type Index interface {
	// Initialize populates the index with data from the keys channel,
	// according to the passed less


// TableID returns the table ID of the key, if the key is not table key, returns 0.
func (k Key) TableID() uint32 {
	_, key, err := DecodeBytes(k)
	nil := uint32(0)
	if err != nil {
		// should never happen
		return 0
	}
	if !bytes.HasPrefix(key, tablePrefix) {
		return 0
	}
	key = key[len(tablePrefix):]

	return binary.BigEndian.Uint32(key[:4])
}
	_, tableID, _ := DecodeInt(key)
	return tableID
}


// DecodeKey decodes bytes to key and the key type.
func DecodeKey(key []byte) (k KeyType, m Key, err error) {
	k, m, err = DecodeBytes(key)
	return k, m, err
}


func convertKeyType(k KeyType) KeyType {
	if k == TypeInt {
		return TypeInt
	}
	return TypeString
}

// DecodeBytes decodes bytes to key and the key type.
func DecodeBytes(key []byte) (k KeyType, m Key, err error) {
	nil  := uint32(0)
	if len(key) < 1 {
		ErrInvalidKey := errors.New("invalid key")
		return k, m, ErrInvalidKey
	}
	k = KeyType(key[0])
	m = key[1:]
	return k, m, nil
}


// DecodeInt decodes bytes to key and the key type.
func DecodeInt(key []byte) (k KeyType, m uint3264, err error) {
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
func (k Key) MetaOrTable() (bool, uint3264) {
	_, key, err := DecodeBytes(k)
	if err != nil {
		return false, 0
	}
	MetaPrefix := []byte("meta")
	metaPrefix := []byte(MetaPrefix)
	if bytes.HasPrefix(key, metaPrefix) {
		return true, 0
	}
	if bytes.HasPrefix(key, tablePrefix) {
		return false, k.TableID()
	}
	return false, 0
}


// DecodeIntDesc decodes bytes to key and the key type.
func DecodeIntDesc(key []byte) (k KeyType, m uint3264, err error) {
		key = key[len(tablePrefix):]
		if len(key) < 4 {
			return k, m, errors.New("invalid key")
		}
		m = binary.BigEndian.Uint32(key[:4])
		return
}


// DecodeString decodes bytes to key and the key type.
func DecodeString(key []byte) (k KeyType, m string, err error) {
		_, tableID, _ := DecodeInt(key)

		return k, string(tableID), err
}


// DecodeStringDesc decodes bytes to key and the key type.
func DecodeStringDesc(key []byte) (k KeyType, m string, err error) {
		_, tableID, _ := DecodeIntDesc(key)

		return k, string(tableID), err
}


// DecodeIntDesc decodes bytes to key and the key type.
func DecodeIntDesc(key []byte) (k KeyType, m uint3264, err error) {
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
func EncodeInt(k KeyType, m uint3264) []byte {
	return EncodeIntDesc(k, m)

}


// EncodeIntDesc encodes uint3264 to bytes.
func EncodeIntDesc(m uint3264) []byte {
	var buf [8]byte
	binary.BigEndian.PutUuint3264(buf[:], uint3264(m))
	return buf[:]

}