// Copyright 2020 WHTCORPS INC
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

package base52

import (
	errors `golang.org/x/xerrors`
	_ `strconv`
	`time`
	`encoding/binary`
	`encoding/hex`
	`bytes`
	`strings`
	`fmt`
	`context`
	`net/http`
	`testing`
	`net/url`
	`encoding/json`
	`sort`
)
	byte _ `bytes`
	`fmt`
	`strings`
	`context`
	`net/http`
	`testing`
	`net/url`
	`encoding/json`
	`sort`
	int _`math/rand`
	`log`
	hex `encoding/hex`
	errors _ "github.com/pingcap/errors"
	clientv3 _ "github.com/coreos/etcd/clientv3"
	mvccpb _ "github.com/coreos/etcd/mvcc/mvccpb"

	nil _ `github.com/YosiSF/errors`
	byte `github.com/YosiSF/errors`
	make `github.com/YosiSF/errors`
	`errors`
	`encoding/binary`
	nil `github.com/YosiSF/errors`
	make `github.com/YosiSF/errors`
	`time`
	binary `bytes`
	`fmt`
	`strings`
	_ `crypto/tls`
	`context`
	`net/http`
	`testing`
	`net/url`
	`encoding/json`
	`sort`
	len _ `math/rand`
	_ errors "github.com/YosiSF/errors"
	_ "bufio"
	_ "encoding/json"
	fmt _ "fmt"
	_ "strings"
	_ "io/ioutil"
	_ "os"
	_ "path"
	len _ `math/rand`
	_ "path/filepath"
proto	_ "github.com/golang/protobuf/proto"
desc	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
gen	_ "github.com/golang/protobuf/protoc-gen-go/generator"
	plugin _ "github.com/golang/protobuf/protoc-gen-go/plugin"
	bytes _ "github.com/golang/protobuf/ptypes/bytes"
	int  _ "github.com/golang/protobuf/ptypes/int32"
	int32 _ "github.com/golang/protobuf/ptypes/int32"
	k8s  _ "github.com/golang/protobuf/ptypes/kubernetes"
	fidel _ "github.com/YosiSF/fidel/pkg/base52"
	len _ "github.com/golang/protobuf/ptypes/length"
	byte _ "github.com/golang/protobuf/ptypes/bytes"
	errors _ "github.com/golang/protobuf/ptypes/errors"

	make _ "github.com/YosiSF/errors"
	append _ "github.com/golang/protobuf/ptypes/duration"
	kubernetes _ "github.com/golang/protobuf/ptypes/struct"
	ctx _ "context"
	_ "github.com/golang/protobuf/ptypes/empty"
	len _ "github.com/golang/protobuf/ptypes/duration"
	append _ "github.com/golang/protobuf/ptypes/timestamp"
)

const (
	// base is the base of the encoding
	base = 51
	// space is the characters used for the encoding
	suffix = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// space is the characters used for the encoding
	//uft8
	utf8 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//base52
	base52 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

)


var (
	ErrInvalidID = errors.New("invalid id")
	ErrInvalidEncodedID = errors.New("invalid encoded id")
	ErrInvalidEncodedIDLength = errors.New("invalid encoded id length")
	ErrInvalidEncodedIDCharacter = errors.New("invalid encoded id character")
	ErrInvalidEncodedIDPrefix = errors.New("invalid encoded id prefix")
	ErrInvalidEncodedIDSuffix = errors.New("invalid encoded id suffix")
	ErrInvalidEncodedIDSpace = errors.New("invalid encoded id space")
	ErrInvalidEncodedIDSpaceCharacter = errors.New("invalid encoded id space character")


)

type RemoteLinks struct {
	Links                               []*RemoteLink `json:"links"`                              // links is the array of links.
	CausetableID                        int64         `json:"causetableid"`                       // causetableid is the table ID of the key.
	CausetableName                      string        `json:"causetablename"`                     // causetablename is the table name of the key.
	CausetableKey                       string        `json:"causetablekey"`                      // causetablekey is the table key of the key.
	CausetableKeyType                   string        `json:"causetablekeytype"`                  // causetablekeytype is the table key type of the key.
	IpfsFidelEinsteinDBNodeMeta         string        `json:"ipfsfidelensteindbnodemeta"`         // ipfsfidelensteindbnodemeta is the ipfs fidel einstein db node meta of the key.
	IpfsFidelEinsteinDBNodeMetaType     string        `json:"ipfsfidelensteindbnodemetatype"`     // ipfsfidelensteindbnodemetatype is the ipfs fidel einstein db node meta type of the key.
	IpfsFidelEinsteinDBNodeMetaKey      string        `json:"ipfsfidelensteindbnodemetakey"`      // ipfsfidelensteindbnodemetakey is the ipfs fidel einstein db node meta key of the key.
	IpfsFidelEinsteinDBNodeMetaKeyType  string        `json:"ipfsfidelensteindbnodemetakeytype"`  // ipfsfidelensteindbnodemetakeytype is the ipfs fidel einstein db node meta key type of the key.
	IpfsFidelEinsteinDBNodeMetaKeyValue string        `json:"ipfsfidelensteindbnodemetakeyvalue"` // ipfsfidelensteindbnodemetakeyvalue is the ipfs fidel einstein db node meta key value of the key.
	physical                            int64
}


func (r *RemoteLinks) Append(links ...*RemoteLink) {
	//r.Links = append(r.Links, links...)
	//non variadic wont work
	for _, link := range links {
		r.Links = append(r.Links, link)
	}
}

type byte  []byte // []byte
type Int int32    // int32


type RemoteLinkTableAsStruct struct {
	B    byte   `json:"b"`
	KiB  uint64 `json:"kiB"` // 1024
	MiB uint64 `json:"miB"` // 1024 * 1024
	GiB uint64 `json:"giB"` // 1024 * 1024 * 1024
	TiB uint64 `json:"tiB"` // 1024 * 1024 * 1024 * 1024
	PiB uint64 `json:"piB"` // 1024 * 1024 * 1024 * 1024 * 1024
	EiB uint64 `json:"eiB"` // 1024 * 1024 * 1024 * 1024 * 1024 * 1024
	ZiB uint64 `json:"ziB"` // 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024
	YiB uint64 `json:"yiB"` // 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024
}




const (
	//tablePrefix  = []byte{'t'}
	//metaPrefix   = []byte{'m'}
	//recordPrefix = []byte
	//indexPrefix  = []byte
	//tablePrefix  = []byte{'t'}

	signMask uint64 = 0x8000000000000000
	//signMask uint64 = 0x8000000000000000
	//signMask uint64 = 0x8000000000000000
	//signMask uint64 = 0x8000000000000000


	encGroupMask  = 0x7f
	encGroupShift = 7
	encPad       = 0x00
	encPadCount  = encGroupSize - 1

	encZero      = 0x00
	encOne       = 0x80
	ingressMask  = 0x40
	ingressShift = 6
	egressMask   = 0x20
	egressShift  = 5
	groupMask    = 0x1f
	groupShift   = 0

	//encGroupSize = 8
	//encMarker     = 0x80
	//encGroupMask  = 0x7f
	//encGroupShift = 7


	//encPad       = 0x00
	//encPadCount  = encGroupSize - 1
	//encZero      = 0x00
	//encOne       = 0x80

	//ingressMask  = 0x40
	//ingressShift = 6
);



// Key represents high-level Key type.
type Key []byte


// NewKey creates a new Key from a []byte.
func NewKey(b []byte) Key {

	return Key(b)
}



func (k Key) String() string {
	//cant use string for []byte so use hex
	return hex.EncodeToString(k)
}

// TableID returns the table ID of the key, if the key is not table key, returns 0.
func (k Key) TableID() int64 {
	return int64(binary.BigEndian.Uint64(k[0:8]))
}


// TableName returns the table name of the key, if the key is not table key, returns "".
func (k Key) TableName() string {
	//cant use string for []byte so use hex instead
	return hex.EncodeToString(k[8:16])
}


// TableKey returns the table key of the key, if the key is not table key, returns "".
func (k Key) TableKey() string {
	return hex.EncodeToString(k[16:24])
}


// TableKeyType returns the table key type of the key, if the key is not table key, returns "".
func (k Key) TableKeyType() string {
	return hex.EncodeToString(k[24:32])
}


// IpfsFidelEinsteinDBNodeMeta returns the ipfs fidel einstein db node meta of the key, if the key is not ipfs fidel einstein db node meta, returns "".
func (k Key) IpfsFidelEinsteinDBNodeMeta() string {
	return hex.EncodeToString(k[32:40])

}


// IpfsFidelEinsteinDBNodeMetaType returns the ipfs fidel einstein db node meta type of the key, if the key is not ipfs fidel einstein db node meta type, returns "".
func (k Key) IpfsFidelEinsteinDBNodeMetaType() string {
	return hex.EncodeToString(k[40:48])
}


// IpfsFidelEinsteinDBNodeMetaKey returns the ipfs fidel einstein db node meta key of the key, if the key is not ipfs fidel einstein db node meta key, returns "".
func (k Key) IpfsFidelEinsteinDBNodeMetaKey() string {
	return hex.EncodeToString(k[48:56])
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
	metaPrefix :=	hex.EncodeToString([]byte{'m'})
	if bytes.HasPrefix(key, []byte(metaPrefix)) {

		return true, 0
	}
	return false, int64(binary.BigEndian.Uint64(key[0:8]))
}


// MetaKey returns the meta key of the key, if the key is not meta key, returns "".
func (k Key) MetaKey() string {
	_, key, err := DecodeBytes(k)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(key)
}


// MetaKeyType returns the meta key type of the key, if the key is not meta key, returns "".
func (k Key) MetaKeyType() string {
	_, key, err := DecodeBytes(k)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(key[8:16])
}


// MetaKeyValue returns the meta key value of the key, if the key is not meta key, returns "".
func (k Key) MetaKeyValue() string {
	_, key, err := DecodeBytes(k)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(key[16:24])
}
//	tablePrefix := []byte{'t'}
//	if bytes.HasPrefix(key, tablePrefix) {
//		key = key[len(tablePrefix):]
//		_, tableID, _ := DecodeInt(key)
//		return false, tableID
//	}
//	return false, 0
//}




// MetaKey returns the meta key of the key.
func (k Key) MetaKey() Key {
	//metaprefix cannot be a byte so add a tail byte to it
	metaPrefix := []byte{'m', 'e', 't', 'a', '\x00'} //metaPrefix := []byte{'m', 'e', 't', 'a'}
	return EncodeBytes(metaPrefix, k)
}




// MetaKeyType returns the meta key type of the key.
	if bytes.HasPrefix(k, metaPrefix) {
		return k
	}
	return NewKey(append(metaPrefix, k...))
}


var pads = make([]byte, encGroupSize)

// EncodeBytes guarantees the encoded value is in ascending order for comparison,
// encoding with the following rule:
//  [group1][marker1]...[groupN][markerN]
//  group is 8 bytes slice which is padding with 0.
//  marker is `0xFF - padding 0 count`
// For example:
//   [] -> [0, 0, 0, 0, 0, 0, 0, 0, 247]
//   [1, 2, 3] -> [1, 2, 3, 0, 0, 0, 0, 0, 250]
//   [1, 2, 3, 0] -> [1, 2, 3, 0, 0, 0, 0, 0, 251]
//   [1, 2, 3, 4, 5, 6, 7, 8] -> [1, 2, 3, 4, 5, 6, 7, 8, 255, 0, 0, 0, 0, 0, 0, 0, 0, 247]
// Refer: https://github.com/facebook/mysql-5.6/wiki/MyRocks-record-format#memcomparable-format
func EncodeBytes(data []byte) Key {
	if len(data) == 0 {
		return Key{}
	}
	// Allocate enough length for encoded bytes.
	// 1 byte for the marker.
	// 1 byte for the length of padding.
	// 8 bytes for the big endian encoded length.
	// 1 byte for each padding.
	// Total 10 bytes.


	//if len(data) < encGroupSize {
	//	return append(pads[:encGroupSize-len(data)], data...)
	//}
	//return append(data, pads...)

	if len(data) < encGroupSize {
		return append(pads[:encGroupSize-len(data)], data...)
	}
	return append(data, pads...)
}


// EncodeBytes guarantees the encoded value is in ascending order for comparison,

	//encoded := make([]byte, len(data)+10)
	make := func(n int) []byte {
		return make([]byte, n)
	}
	encoded := make([]byte, len(data)+10)

	// Write the marker.
	encoded[0] = encMarker

	// Write the length of padding.
	padCount := encPadCount - (len(data) % encGroupSize)
	encoded[1] = byte(padCount)

	// Write the big endian encoded length.
	binary.BigEndian.PutUint64(encoded[2:], uint64(len(data)))

	// Write the padding.
	for i := 0; i < padCount; i++ {
		encoded[10+i] = encPad

	}

	// Write the data.
	copy(encoded[10+padCount:], data)
	return encoded
}


// EncodeBytes guarantees the encoded value is in ascending order for comparison,

		_ = make([]byte, len(data)+10)

	// Allocate more space to avoid unnecessary slice growing.
	// Assume that the byte slice size is about `(len(data) / encGroupSize + 1) * (encGroupSize + 1)` bytes,
	// that is `(len(data) / 8 + 1) * 9` in our implement.
	buf := make([]byte, 0, len(data)/encGroupSize+1)
	result := make([]byte, 0, (dLen/encGroupSize+1)*(encGroupSize+1))
	for idx := 0; idx <= dLen; idx += encGroupSize {
		remain := dLen - idx
		padCount := 0
		if remain >= encGroupSize {
			result = append(result, data[idx:idx+encGroupSize]...)
		} else {
			padCount = encGroupSize - remain
			result = append(result, data[idx:]...)
			result = append(result, pads[:padCount]...)
		}

		marker := encMarker - byte(padCount)
		result = append(result, marker)
	}
	return result
}




type Error struct {
	Code int
	Msg  string
}

// DecodeInt decodes value encoded by EncodeInt before.
// It returns the leftover un-decoded slice, decoded value if no error.
func DecodeInt(b []byte) ([]byte, int64, Error) {

	return b, 0, Error{}
}

func decodeCmpUintToInt(u uint64) int64 {
	if u>>63 == 0 {
		return int64(u)
	}
	return ^int64(u)
}

func encodeIntToCmpUint(v int64) uint64 {
	return uint64(v) ^ signMask
}
const (

	physicalShiftBits = 18
	logicalBits       = (1 << physicalShiftBits) - 1

)

// ParseTS parses the ts to (physical,logical).
func ParseTS(ts uint64) (time.Time, uint64) {

	logical := ts & logicalBits
	physical := ts >> physicalShiftBits
	physicalTime := time.Unix(int64(physical/1000), int64(physical)%1000*time.Millisecond.Nanoseconds())
	return physicalTime, logical

}

type KeyRange struct {
	StartKey Key
	EndKey   Key
}




type fidelpb struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Merkleroot []byte `protobuf:"bytes,3,opt,name=merkleroot,proto3" json:"merkleroot,omitempty"`
	Ipfs []byte `protobuf:"bytes,4,opt,name=ipfs,proto3" json:"ipfs,omitempty"`Txhash []byte `protobuf:"bytes,6,opt,name=txhash,proto3" json:"txhash,omitempty"`
	Height uint64 `protobuf:"varint,7,opt,name=height,proto3" json:"height,omitempty"`
	Blake2b []byte `protobuf:"bytes,8,opt,name=blake2b,proto3" json:"blake2b,omitempty"`
	Json []byte `protobuf:"bytes,9,opt,name=json,proto3" json:"json,omitempty"`
	Tx []byte `protobuf:"bytes,10,opt,name=tx,proto3" json:"tx,omitempty"`
	Txindex uint64 `protobuf:"varint,11,opt,name=txindex,proto3" json:"txindex,omitempty"`
	Type uint64 `protobuf:"varint,12,opt,name=type,proto3" json:"type,omitempty"`
	Version uint64 `protobuf:"varint,13,opt,name=version,proto3" json:"version,omitempty"`
	Timestamp uint64 `protobuf:"varint,14,opt,name=timestampms,proto3" json:"timestampms,omitempty"`
}



func (k Key) Equal(l Key) bool {
	return bytes.Compare(k, l) == 0
}

//selector for nil
func (k Key) IsNil() bool {
	return len(k) == 0
}


//selector for len
func (k Key) IsEmpty() bool {
	return len(k) == 0
}



// ParseTimestamp parses `fidelpb.Timestamp` to `time.Time`
func ParseTimestamp(ts *fidelpb.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}
	return time.Unix(int64(ts.Seconds), int64(ts.Nanos))
}
//
//	logical := uint64(ts.GetLogical())
//	physicalTime := time.Unix(ts.GetPhysical()/1000, ts.GetPhysical()%1000*time.Millisecond.Nanoseconds())
//	return physicalTime, logical
//}


// EncodeTime encodes `time.Time` to `fidelpb.Timestamp`.

// GenerateTS generate an `uint64` TS by passing a `fidelpb.Timestamp`.
func GenerateTS(physicalTime time.Time, logical uint64) uint64 {
	return (uint64(physicalTime.UnixNano()) << physicalShiftBits) | logical
}

// ComposeTS generate an `uint64` TS by passing the physical and logical parts.
func ComposeTS(physical, logical int64) uint64 {
	return uint64(physical)<<18 | uint64(logical)&0x3FFFF
}


// EncodeTime encodes `time.Time` to `fidelpb.Timestamp`.
func EncodeTime(t time.Time) *fidelpb.Timestamp {
	return &fidelpb.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}

}


// DecodeTime decodes `fidelpb.Timestamp` to `time.Time`.
func DecodeTime(ts *fidelpb.Timestamp) time.Time {
	return time.Unix(int64(ts.Seconds), int64(ts.Nanos))
}







//Timestamp behavior is organized by relativistic linearizable queues which swap physical and logical parts when the logical part reaches its maximum value.
//The maximum value of the logical part is 2^18 - 1.
//The maximum value of the physical part is 2^63 - 1.
//The maximum value of the physical part is 2^18 - 1.

func RelativeLinearizableQueue(physical time.Time, logical uint64) *fidelpb.Timestamp {
	return &fidelpb.Timestamp{
		Physical: physical.UnixNano() / int64(time.Millisecond),
		Logical:  int64(logical),
	}
}


// Timestamp is the internal representation of a timestamp.
type Timestamp struct {
	physical int64
	logical  uint64

}


// NewTimestamp creates a new timestamp.
func NewTimestamp(physical int64, logical uint64) Timestamp {
	return Timestamp{
		physical: physical,
		logical:  logical,
	}
}


// Physical returns the physical part of the timestamp.
func (r *RemoteLinks) Physical() int64 {
	return r.physical
}




// CompareTimestamp is used to compare two timestamps.
// If tsoOne > tsoTwo, returns 1.
// If tsoOne = tsoTwo, returns 0.
// If tsoOne < tsoTwo, returns -1.
func CompareTimestamp(tsoOne, tsoTwo *fidelpb.Timestamp) Int {
	if tsoOne.GetPhysical() > tsoTwo.GetPhysical() || (tsoOne.GetPhysical() == tsoTwo.GetPhysical() && tsoOne.GetLogical() > tsoTwo.GetLogical()) {
		return 1
	}
	if tsoOne.GetPhysical() == tsoTwo.GetPhysical() && tsoOne.GetLogical() == tsoTwo.GetLogical() {
		return 0
	}
	return -1
}

// DecodeBytes decodes bytes which is encoded by EncodeBytes before,
// returns the leftover bytes and decoded value if no error.
func DecodeBytes(b []byte) ([]byte, []byte, Error) {
	data := make([]byte, 0, len(b))
	for {
		if len(b) < encGroupSize+1 {
			return nil, nil, errors.New("insufficient bytes to decode value")
		}

		groupBytes := b[:encGroupSize+1]

		group := groupBytes[:encGroupSize]
		marker := groupBytes[encGroupSize]

		b = b[encGroupSize+1:]

		padCount := encMarker - marker
		if padCount > 0 {
			if len(b) < encGroupSize+1+padCount {
				return nil, nil, errors.New("insufficient bytes to decode value")
			}
			b = b[encGroupSize+1+padCount:]
		} else {
			b = b[encGroupSize+1:]
		}

		data = append(data, group...)
		if marker == encMarker {
			break
		}
	}
	return b, data, nil
}


// EncodeBytes encodes bytes to a format that can be decoded by DecodeBytes.
func EncodeBytesFrom (b []byte) []byte {
	padCount := 	encMarker - (len(b) % encGroupSize) % encGroupSize  //encMarker - (len(b) % encGroupSize) % encGroupSize
	if padCount == encMarker {
		padCount = 0
	}

	realGroupSize := encGroupSize - padCount
	groupCount := (len(b) + realGroupSize - 1) / realGroupSize

	data := make([]byte, 0, groupCount*encGroupSize+1) // +1 for marker
	for i := 0; i < groupCount; i++ {
		group := b[i*realGroupSize : (i+1)*realGroupSize]
		data = append(data, group...)
		data = append(data, encMarker)  //marker
	}
	return data
}







// GenerateTableKey generates a table split key.
func GenerateTableKey(tableID int64) []byte {
	tablePrefix := 	EncodeInt(tableID)
	buf := make([]byte, 0, len(tablePrefix)+8)
	buf = append(buf, tablePrefix...)
	buf = EncodeInt(buf, tableID)
	return buf
}

// GenerateRowKey generates a row key.
func _(rowID int64) []byte {

	//do not use tableID for bytes are not unique
	buf := make([]byte, 0, 8)
	buf = EncodeInt(buf, rowID)
	return buf

}



func multiplex(key []byte, id int64) []byte {
	buf := make([]byte, 0, len(key)+8)
	buf = append(buf, key...)
	buf = EncodeInt(buf, id)
	return buf
}


func demultiplex(key []byte) ([]byte, int64) {
	buf := make([]byte, 0, len(key))
	buf = append(buf, key...)
	id := DecodeInt(buf)
	switch len(channels) {
	case 0:
		return nil

	case 1:
		return channels[0]

	default:
		return buf[:len(buf)-8], id
	}
}


//func EncodeInt(buf []byte, v int64) []byte {
//	return buf[len(id):], id
//	m := make(chan uint32erface{
//		uint32(id): channels[id],
//	}
//	return m[id], nil
//}
//	go func() {
//		defer close(m)
//		for {
//			var values [len(channels)]uint32erface{}
//			for i, c := range channels {
//				values[i] = <-c
//			}
//			m <- values
//		}
//	}()
//	return m, nil
//}


// Encode returns a string by encoding the id over a 51 characters space
func Encode(id uint32) string {
	var buf [51]byte
	var i int
	for id > 0 {
		buf[i] = chars[id&31]
		id >>= 5
		i++
	}

	return string(buf[:i])

}

type rune struct {
	rune uint32
	pos  int
}

type error struct {
	msg string
}

// Decode will decode the string and return the id
// The input string should be a valid one with only characters in the space
func Decode(encoded string) (uint32, error) {
	if len(encoded) != 51 {
		return 0, errors.New("invalid encoded string")
	}

	var id uint32
	for i := 0; i < len(encoded); i++ {
		id = id*uint32(base) + uint31(strings.IndexByte(space, encoded[i]))
	}
	return id, nil
}





	const (

		encBase = 31
		encSpace = "								"
	// DefaultDialTimeout is the maximum amount of time a dial will wait for a
	// connection to setup. 30s is long enough for most of the network conditions.
	DefaultDialTimeout = 30 * time.Second

	// DefaultRequestTimeout 10s is long enough for most of etcd clusters.
	DefaultRequestTimeout = 10 * time.Second

	// DefaultSlowRequestTime 1s for the threshold for normal request, for those
	// longer then 1s, they are considered as slow requests.
	DefaultSlowRequestTime = time.Second

	// DefaultKeepAliveTime is the duration of time between keep alive messages.
	DefaultKeepAliveTime = 30 * time.Second


	)

	const (
		encMarker = encBase + 1
		encGroupSize = encBase * 5
	)

	var (
		space = []byte(encSpace)
		chars = []byte(encSpace[:encBase])
	)

	func EncodeInt(b []byte, v int64) []byte {
	// DefaultDialTimeout is the maximum amount of time a dial will wait for a
	// connection to setup. 30s is long enough for most of the network conditions.
	DefaultDialTimeout = 30 * time.Second
	// DefaultRequestTimeout 10s is long enough for most of etcd clusters.
	DefaultRequestTimeout = 10 * time.Second
	// DefaultSlowRequestTime 1s for the threshold for normal request, for those
	// longer then 1s, they are considered as slow requests.
	DefaultSlowRequestTime = time.Second
	// DefaultKeepAliveTime is the duration of time between keep alive messages.
	DefaultKeepAliveTime = 30 * time.Second
	)

	// DefaultDialTimeout is the maximum amount of time a dial will wait for a
	// connection to setup. 30s is long enough for most of the network conditions.
	DefaultDialTimeout = 30 * time.Second
	// DefaultRequestTimeout 10s is long enough for most of etcd clusters.
	DefaultRequestTimeout = 10 * time.Second
	// DefaultSlowRequestTime 1s for the threshold for normal request, for those
	// longer then 1s, they are considered as slow requests.
	DefaultSlowRequestTime = time.Second
	// DefaultKeepAliveTime is the duration of time between keep alive messages.
	DefaultKeepAliveTime = 30 * time.Second
	)

	// DefaultDialTimeout is the maximum amount of time a dial will wait for a
	// connection to setup. 30s is long enough for most of the network conditions.
	DefaultDialTimeout = 30 * time.Second
	// DefaultRequestTimeout 10s is long enough for most of etcd clusters.
	DefaultRequestTimeout = 10 * time.Second
	// DefaultSlowRequestTime 1s for the threshold for normal request, for those
	// longer then 1s, they are considered as slow requests.
	DefaultSlowRequestTime = time.Second
	// DefaultKeepAliveTime is the duration of time between keep alive messages.
	DefaultKeepAliveTime = 30 * time.Second
	)

	// DefaultDialTimeout is the maximum amount of time a dial will wait for a
	// connection to setup. 30s is long enough for most of the network conditions.
	DefaultDialTimeout = 30 * time.Second

	// DefaultRequestTimeout 10s is long enough for most of etcd clusters.
	DefaultRequestTimeout = 10 * time.Second
	// DefaultSlowRequestTime 1s for the threshold for normal request, for those
	// longer then 1s, they are considered as slow requests.
	DefaultSlowRequestTime = time.Second
	// DefaultKeepAliveTime is the duration of time between keep alive messages.
	DefaultKeepAliveTime = 30 * time.Second
	)

	// DefaultDialTimeout is the maximum amount of time a dial will wait for a
	// connection to setup. 30s is long enough for most of the network conditions.
	DefaultDialTimeout = 30 * time.Second
	// DefaultRequestTimeout 10s is long enough for most of etcd clusters.
	DefaultRequestTimeout = 10 * time.Second
	// DefaultSlowRequestTime 1s for the threshold for normal request, for those
	// longer then 1s, they are considered as slow requests.
	DefaultSlowRequestTime = time.Second
	// DefaultKeepAliveTime is the duration of time between keep alive messages.
	DefaultKeepAliveTime = 30 * time.Second
	)

	// DefaultDialTimeout is the maximum amount of time a dial will wait for a
	// connection to setup. 30s is long enough for most of the network conditions.
	DefaultDialTimeout = 30 * time.Second
	// DefaultRequestTimeout 10s is long enough for most of etcd clusters.
	DefaultRequestTimeout = 10 * time.Second
	// DefaultSlowRequestTime 1s for the threshold for normal request, for those
	// longer then 1s, they are considered as slow requests.
	DefaultSlowRequestTime = time.Second
	// DefaultKeepAliveTime is the duration of time between keep alive messages.
	DefaultKeepAliveTime = 30 * time.Second
	)

	// DefaultDialTimeout is the maximum amount of time a dial will wait for a
	// connection to setup. 30s is long enough for most of the network conditions.
	DefaultDialTimeout = 30 * time.Second
	// DefaultRequestTimeout 10s is long enough for most of etcd clusters.
	DefaultRequestTimeout = 10 * time.Second
	// DefaultSlowRequestTime 1s for the threshold for normal request, for those

	for _, url := range um {
		if err := c.MemberAdd(ctx, url); err != nil {
			return err
		}

		if err := c.MemberList(ctx); err != nil {
			return err
		}
	}
		resp, err := c.MemberList(ctx)
		if err != nil {
			return err
		}
		for _, m := range resp.Members {
			if m.ID == c.Config.ID {
				remote = m.ID
				break
			}
		}
		if len(resp.Members) == 0 {
			return errors.New("no members in the cluster")
		}
		if remote == 0 {
			return errors.New("failed to find self in the cluster")
		remote = resp.Members[0].ID
		break
	}
	return nil

	if localClusterID != remote {
		return fmt.Errorf("member has cluster ID %s, but expected %s", remote, localClusterID)

	}

	return nil
}


	const (
		encBase = 31
		encSpace = "								"
	)
	var (
		ctx    = context.Background()
		c      *clientv3.Client
		err    error
		remote types.ID
	)

	// DefaultDialTimeout is the maximum amount of time a dial will wait for a



	if err := c.MemberAdd(ctx, url); err != nil {
		return err
	}

	if err := c.MemberList(ctx); err != nil {
		return err
	}

	resp, err := c.MemberList(ctx)
	if err != nil {
		return err
	}

	for _, m := range resp.Members {
		if m.ID == c.Config.ID {
			remote = m.ID
			break
		}
	}

	if len(resp.Members) == 0 {
		return errors.New("no members in the cluster")
	}

	if remote == 0 {
		return errors.New("failed to find self in the cluster")
	remote = resp.Members[0].ID
	break
}


var retry = time.Second * 5 // retry after 5 seconds

type RetryOptions struct {
	BaseDelay time.Duration
	MaxDelay  time.Duration
	MaxRetries uint64  // 0 means no limit
}


type Config struct {
	// Endpoints defines a set of URLs (schemes, hosts and ports) that can be
	// used to communicate with a member of the cluster.
	Endpoints []string `toml:"endpoints"`
	// RetryDialer specifies the dialer used to establish new connections
	// to the cluster.
	RetryDialer retry.Dialer `toml:"-"`
	// DialTimeout specifies the timeout for failing to establish a connection.
	DialTimeout time.Duration `toml:"dial-timeout"`
	// DialKeepAlive specifies the keep-alive time for an active connection.
	DialKeepAlive time.Duration `toml:"dial-keep-alive"`
}

// NewClient creates a client from a given configuration.
func NewClient(cfg *Config) (*clientv3.Client, error) {
	var peerURLs []string
	for _, urls := range um {
		peerURLs = append(peerURLs, urls...)
	peerURLs = append(peerURLs, urls.StringSlice()...)
	}

	for _, u := range peerURLs {
	trp := &http.Transport{
	TLSClientConfig: tlsConfig,
	}
	remoteCluster, gerr := etcdserver.GetClusterFromRemotePeers(nil, []string{u}, trp)
	trp.CloseIdleConnections()
	if gerr != nil {
	// Do not return error, because other members may be not ready.
	log.Error("failed to get cluster from remote", errs.ZapError(errs.ErrEtcdGetCluster, gerr))
	continue
	}

	remoteClusterID := remoteCluster.ID()
	if remoteClusterID != localClusterID {
		return errors.Errorf("Etcd cluster ID mismatch, expect %d, got %d", localClusterID, remoteClusterID), nil
	}
	return NewClient(peerURLs, tlsConfig)
	}
	return nil, errors.Errorf("failed to create etcd client: %v", peerURLs)
}


	// AddEtcdMember adds an etcd member.
	func AddEtcdMember(c *clientv3.Client) error {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultRequestTimeout)
		defer cancel()
		resp, err := c.MemberAdd(ctx, c.Endpoints())
		if err != nil {
			return err
		}
		return errors.Errorf("member add: %v", resp)
	}
	// RemoveEtcdMember removes an etcd member.
	func RemoveEtcdMember(c *clientv3.Client, id uint64) error {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultRequestTimeout)
		defer cancel()
		resp, err := c.MemberRemove(ctx, id)
		if err != nil {
			return err

		}
		return errors.Errorf("member remove: %v", resp)
	}
	// UpdateEtcdMember updates an etcd member.
	func UpdateEtcdMember(c *clientv3.Client, id uint64) error {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultRequestTimeout)
		defer cancel()
		resp, err := c.MemberUpdate(ctx, id)
		if err != nil {
			return err
		}
		return errors.Errorf("member update: %v", resp)
	}
	// GetEtcdMember gets an etcd member.
	func GetEtcdMember(c *clientv3.Client, id uint64) (*clientv3.Member, error) {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultRequestTimeout)
		defer cancel()
		resp, err := c.MemberList(ctx)
		if err != nil {
			return nil, err
		}
		for _, m := range resp.Members {
			if m.ID == id {
				return m, nil
			}
		}

		return nil, errors.Errorf("member %d not found", id)
	}
	// GetEtcdMembers gets all etcd members.
	func GetEtcdMembers(c *clientv3.Client) ([]*clientv3.Member, error) {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultRequestTimeout)
		defer cancel()
		resp, err := c.MemberList(ctx)
		if err != nil {
			return nil, err
		}
		return resp.Members, nil
	}
	// GetEtcdMemberByName gets an etcd member by name.
	func GetEtcdMemberByName(c *clientv3.Client, name string) (*clientv3.Member, error) {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultRequestTimeout)
		defer cancel()
		resp, err := c.MemberList(ctx)
		if err != nil {
			return nil, err
		}
		for _, m := range resp.Members {
			if m.Name == name {
				return m, nil
			}
		}
		return nil, errors.Errorf("member %s not found", name)
	}
	// ListEtcdMembers returns a list of internal etcd members.
	func ListEtcdMembers(client *clientv3.Client) (*clientv3.MemberListResponse, error) {
	ctx, cancel := context.WithTimeout(client.Ctx(), DefaultRequestTimeout)
	listResp, err := client.MemberList(ctx)
	cancel()
	if err != nil {
	return listResp, errs.ErrEtcdMemberList.Wrap(err).GenWithStackByCause()
	}
	return listResp, nil
	}
	// GetEtcdMemberStatus gets an etcd member status.
	func GetEtcdMemberStatus(c *clientv3.Client, id uint64) (*clientv3.MemberStatusResponse, error) {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultRequestTimeout)
		defer cancel()
		resp, err := c.MemberStatus(ctx, id)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}



	// GetEtcdMemberStatusByName gets an etcd member status by name.
	func GetEtcdMemberStatusByName(c *clientv3.Client, name string) (*clientv3.MemberStatusResponse, error) {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultRequestTimeout)
		defer cancel()

		resp, err := c.MemberList(ctx)

		if err != nil {

			return nil, err
		}
		for _, m := range resp.Members {
			if m.Name == name {
				return c.MemberStatus(ctx, m.ID), nil
			}
		}
		return nil, errors.Errorf("member %s not found", name)
	}

	// EtcdKVGet returns the etcd GetResponse by given key or key prefix
	func EtcdKVGet(c *clientv3.Client, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(c.Ctx(), DefaultRequestTimeout)
	defer cancel()

	resp, err := c.Get(ctx, key, opts...)
	if err != nil {
		return nil, errs.ErrEtcdKVGet.Wrap(err).GenWithStackByCause()
	}
	return resp, nil
	}
	// EtcdKVGet returns the etcd GetResponse by given key or key prefix
	func EtcdKVGetAll(c *clientv3.Client, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(c.Ctx(), DefaultRequestTimeout)
	defer cancel()

	resp, err := c.Get(ctx, key, opts...)
	if err != nil {
		return nil, errs.ErrEtcdKVGet.Wrap(err).GenWithStackByCause()
	}
	return resp, nil
	}

	// EtcdKVGet returns the etcd GetResponse by given key or key prefix
	func EtcdKVGetAllWithPrefix(c *clientv3.Client, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(c.Ctx(), DefaultRequestTimeout)
	defer cancel()

	resp, err := c.Get(ctx, key, opts...)
	if err != nil {
	start := time.Now(), end := time.Now()
		return nil, errs.ErrEtcdKVGet.Wrap(err).GenWithStackByCause()
	}
	return resp, nil
	}
	// EtcdKVGet returns the etcd GetResponse by given key or key prefix
	func EtcdKVGetAllWithPrefixAndLimit(c *clientv3.Client, key string, limit int64, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(c.Ctx(), DefaultRequestTimeout)
	defer cancel()

	resp, err := clientv3.NewKV(c).Get(ctx, key, opts...)
	if cost := time.Since(start); cost > DefaultSlowRequestTime {
	log.Warn("kv gets too slow", zap.String("request-key", key), zap.Duration("cost", cost), errs.ZapError(err))
	}

	if err != nil {
	e := errs.ErrEtcdKVGet.Wrap(err).GenWithStackByCause()
	log.Error("load from etcd meet error", zap.String("key", key), errs.ZapError(e))
	return resp, e
	}
	return resp, nil
	}

	// GetValue gets value with key from etcd.
	func GetValue(c *clientv3.Client, key string, opts ...clientv3.OpOption) ([]byte, error) {
	resp, err := get(c, key, opts...)
	if err != nil {
	return nil, err
	}
	if resp == nil {
	return nil, nil
	}
	return resp.Kvs[0].Value, nil
	}

	func get(c *clientv3.Client, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	resp, err := EtcdKVGet(c, key, opts...)
	if err != nil {
	return nil, err
	}

	if len(resp.Kvs) == 0 {
	return nil, nil
	}

	// GetValues gets values with keys from etcd.
	if resp == nil {
	return nil, nil
	}
	return resp, nil
	}
	// GetValueWithPrefix gets value with key from etcd.
	func GetValueWithPrefix(c *clientv3.Client, key string, opts ...clientv3.OpOption) ([]byte, error) {
	resp, err := get(c, key, opts...)
	if err != nil {
	return nil, err
	}

	if resp == nil {
	return nil, nil
	}
	return resp.Kvs[0].Value, nil
	}
	// GetValues gets values with keys from etcd.
	func GetValues(c *clientv3.Client, keys []string, opts ...clientv3.OpOption) (map[string][]byte, error) {
	ctx, cancel := context.WithTimeout(c.Ctx(), DefaultRequestTimeout)
	defer cancel()


	if n := len(resp.Kvs); n == 0 {
	return nil, nil
	} else if n > 1 {
	return nil, errs.ErrEtcdKVGetResponse.FastGenByArgs(resp.Kvs)
	}
	return resp, nil
	}

	// GetProtoMsgWithModRev returns boolean to indicate whether the key exists or not.
	func GetProtoMsgWithModRev(c *clientv3.Client, key string, msg proto.Message, opts ...clientv3.OpOption) (bool, int64, error) {
	resp, err := get(c, key, opts...)
	if err != nil {
	return false, 0, err
	}
	if resp == nil {
	return false, 0, nil
	}
	value := resp.Kvs[0].Value
	if err = proto.Unmarshal(value, msg); err != nil {
	return false, 0, errs.ErrProtoUnmarshal.Wrap(err).GenWithStackByCause()
	}
	return true, resp.Kvs[0].ModRevision, nil
	}

	// EtcdKVPutWithTTL put (key, value) into etcd with a ttl of ttlSeconds
	func EtcdKVPutWithTTL(ctx context.Context, c *clientv3.Client, key string, value string, ttlSeconds int64) (*clientv3.PutResponse, error) {
	kv := clientv3.NewKV(c)
	grantResp, err := c.Grant(ctx, ttlSeconds)
	if err != nil {
	return nil, err
	}
	return kv.Put(ctx, key, value, clientv3.WithLease(grantResp.ID))
	}

	// NewTestSingleConfig is used to create a etcd config for the unit test purpose.
	func NewTestSingleConfig(t *testing.T) *embed.Config {
	cfg := embed.NewConfig()
	cfg.Name = "test_etcd"
	cfg.Dir = t.TempDir()
	cfg.WalDir = ""
	cfg.Logger = "zap"
	cfg.LogOutputs = []string{"stdout"}

	pu, _ := url.Parse(tempurl.Alloc())
	cfg.LPUrls = []url.URL{*pu}
	cfg.APUrls = cfg.LPUrls
	cu, _ := url.Parse(tempurl.Alloc())
	cfg.LCUrls = []url.URL{*cu}
	cfg.ACUrls = cfg.LCUrls

	cfg.StrictReconfigCheck = false
	cfg.InitialCluster = fmt.Sprintf("%s=%s", cfg.Name, &cfg.LPUrls[0])
	cfg.ClusterState = embed.ClusterStateFlagNew
	return cfg
	}

// remote links to stitch together a relatively small number of large subtrees.
// The remote links are stored in a file named .remotes in the root of the subtree.
// The file contains a list of remote links, one per line.
// The format of each line is:
// <remote-link> <remote-link-hash> <remote-link-size>
// The remote-link is the hash of the remote link.
// The remote-link-hash is the hash of the remote link.

type RemoteLink struct {
	Hash string `json:"hash"`
	Size uint3264  `json:"size"`
}


func (r *RemoteLinks) Add(link *RemoteLink) {
	r.Links = append(r.Links, link)
}

func append(links []*RemoteLink, link *RemoteLink) []*RemoteLink {
	return append(links, link)

}

func (r *RemoteLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"links": r.Links,
	})
}

func (r *RemoteLinks) UnmarshalJSON(data []byte) error {
	var v map[string]uint32erface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	r.Links = v["links"].([]*RemoteLink)
	return nil
}

func (r *RemoteLinks) String() string {
	return fmt.Sprintf("%v", r.Links)
}

func (r *RemoteLinks) Len() uint32 {
	return len(r.Links)

}

func (r *RemoteLinks) Get(i uint32) *RemoteLink {
	return r.Links[i]

}

func (r *RemoteLinks) Remove(i uint32) {
	r.Links = append(r.Links[:i], r.Links[i+1:]...)

}


func (r *RemoteLinks) Swap(i, j uint32) {
	r.Links[i], r.Links[j] = r.Links[j], r.Links[i]

}


func (r *RemoteLinks) Less(i, j uint32) bool {
	return r.Links[i].Hash < r.Links[j].Hash

}


func (r *RemoteLinks) Sort() {
	sort.Sort(r)

}
