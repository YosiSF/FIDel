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

import "bytes"

var (
	tablePrefix  = []byte{'t'}
	metaPrefix   = []byte{'m'}
	recordPrefix = []byte{'r'}
)

const (
	signMask uint64 = 0x8000000000000000

	encGroupSize = 8
	encMarker    = byte(0xFF)
	encPad       = byte(0x0)
)

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

var pads = make([]byte, encGroupSize)