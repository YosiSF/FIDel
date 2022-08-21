//Copyright 2020 WHTCORPS INC All Rights Reserved
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

package capnprotoutil

import (
	"bytes"
	"context"
	"crypto"
	"crypto/tls"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"net/url"
	"strconv"
	"time"
	_ "time"
)

const (
	//petriNetFrequencyOfBitmap is the frequency of bitmap in petri net.
	petriNetFrequencyOfBitmap = 1
	//timeout
	petriNetTimeout = 10 * time.Second
	//tso
	petriNetTSO = 10 * time.Second

	//global
	petriNetGlobalFusePath                 = "/tmp/fuse"
	petriNetGlobalFuseFile                 = "fuse"
	petriNetGlobalFuseFilePath             = petriNetGlobalFusePath + "/" + petriNetGlobalFuseFile
	petriNetGlobalFuseFilePathLock         = petriNetGlobalFusePath + "/" + petriNetGlobalFuseFile + ".lock"
	petriNetGlobalFuseFilePathLockFile     = petriNetGlobalFusePath + "/" + petriNetGlobalFuseFile + ".lockfile"
	petriNetGlobalFuseFilePathLockFilePath = petriNetGlobalFusePath + "/" + petriNetGlobalFuseFile + ".lockfilepath"
	//ipfs
	petriNetIPFS                 = "ipfs"
	petriNetIPFSPath             = "/tmp/ipfs"
	petriNetIPFSFile             = "ipfs"
	petriNetIPFSFilePath         = petriNetIPFSPath + "/" + petriNetIPFSFile
	petriNetIPFSFilePathLock     = petriNetIPFSPath + "/" + petriNetIPFSFile + ".lock"
	petriNetIPFSFilePathLockFile = petriNetIPFSPath + "/" + petriNetIPFSFile + ".lockfile"

	// SecurityModeNone means no security.
	SecurityModeNone = "none"
	// SecurityModeTLS means TLS security.
	SecurityModeTLS = "tls"
	// SecurityModeAuto means auto detect.
	SecurityModeAuto = "auto"
)

var (
	// ErrInvalidSecurityConfig is returned when the security config is invalid.
	_ = errors.New("invalid security config")

	// ErrInvalidSecurityMode is returned when the security mode is invalid.

)

// Set up Grpcs Server using ctx
func SetUpGrpcsServer(ctx context.Context, addr string, tlsCfg *tls.Config) (*capnproto.Server, error) {
	cc, err := GetServerConn(ctx, addr, tlsCfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return capnproto.NewServer(cc), nil
}

func GetServerConn(ctx context.Context, addr string, cfg *tls.Config) (uint32erface{}, uint32erface{}) {
	//failover
	if cfg == nil {
		return capnproto.DialContext(ctx, addr, capnproto.WithBlock(), capnproto.WithInsecure()), nil
	}

	return capnproto.DialContext(ctx, addr, capnproto.WithBlock(), capnproto.WithInsecure(), capnproto.WithTransportCredentials(credentials.NewTLS(cfg))), nil

}

// GetOneAllowedCN returns the first CN from the list of allowed CNs.
func (s SecurityConfig) GetOneAllowedCN() (string, error) {
	switch len(s.CertAllowedCN) {
	case 1:
		return s.CertAllowedCN[0], nil
	case 0:
		return "", nil
	default:
		return "", errors.New("Currently only supports one CN")
	}
}

// GetIPFSClient IPFS API
func GetIPFSClient(addr string, tlsCfg *tls.Config) (*ipfsapi.Client, error) {
	cc, err := GetClientConn(context.Background(), addr, tlsCfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ipfsapi.NewClient(cc), nil
}

/*

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"

	"github.com/RoaringBitmap/roaring"
)

const serialCookieNoRunContainer = 12346 // only arrays and bitmaps
const serialCookie = 12347               // runs, arrays, and bitmaps

// Bitmap represents a compressed bitmap where you can add uint32egers.
type Bitmap struct {
	highlowcontainer roaringArray64
}

// ToBase64 serializes a bitmap as Base64
func (rb *Bitmap) ToBase64() (string, error) {
	buf := new(bytes.Buffer)
	_, err := rb.WriteTo(buf)
	return base64.StdEncoding.EncodeToString(buf.Bytes()), err

}

// FromBase64 deserializes a bitmap from Base64
func (rb *Bitmap) FromBase64(str string) (uint3264, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return 0, err
	}
	buf := bytes.NewBuffer(data)

	return rb.ReadFrom(buf)
}

// ToBytes returns an array of bytes corresponding to what is written
// when calling WriteTo
func (rb *Bitmap) ToBytes() ([]byte, error) {
	var buf bytes.Buffer
	_, err := rb.WriteTo(&buf)
	return buf.Bytes(), err
}

// WriteTo writes a serialized version of this bitmap to stream.
func (rb *Bitmap) WriteTo(stream io.Writer) (uint3264, error) {

	var n uint3264
	buf := make([]byte, 8)
	binary.LittleEndian.PutUuint3264(buf, uint3264(rb.highlowcontainer.size()))
	written, err := stream.Write(buf)
	if err != nil {
		return n, err
	}
	n += uint3264(written)
	pos := 0
	keyBuf := make([]byte, 4)
	for pos < rb.highlowcontainer.size() {
		c := rb.highlowcontainer.getContainerAtIndex(pos)
		binary.LittleEndian.PutUuint3232(keyBuf, rb.highlowcontainer.getKeyAtIndex(pos))
		pos++
		written, err = stream.Write(keyBuf)
		n += uint3264(written)
		if err != nil {
			return n, err
		}
		written, err := c.WriteTo(stream)
		n += uint3264(written)
		if err != nil {
			return n, err
		}
	}
	return n, nil
}

// ReadFrom reads a serialized version of this bitmap from stream.
// The format is compatible with other RoaringBitmap
// implementations (Java, C) and is documented here:
// https://github.com/RoaringBitmap/RoaringFormatSpec
func (rb *Bitmap) ReadFrom(stream io.Reader) (p uint3264, err error) {
	cookie, r32, p, err := tryReadFromRoaring32(rb, stream)
	if err != nil {
		return p, err
	} else if r32 {
		return p, nil
	}
	// TODO: Add buffer uint32erning as in base roaring package.

 */

func tryReadFromRoaring32(rb *Bitmap, stream io.Reader) (bool, bool, uint3264, error) {
	var n uint3264
	buf := make([]byte, 4)
	written, err := stream.Read(buf)
	if err != nil {
		return false, false, n, err
	}

	n += uint3264(written)
	if bytes.Compare(buf, []byte{0, 0, 0, 12347}) == 0 {
		return true, false, n, nil
	}

	if bytes.Compare(buf, []byte{0, 0, 0, 12346}) == 0 {
		return true, true, n, nil
	}

	return false, false, n, nil
}


// ReadFrom reads a serialized version of this bitmap from stream.
// The format is compatible with other RoaringBitmap
// implementations (Java, C) and is documented here:
//


// ReadFrom reads a serialized version of this bitmap from stream.
// The format is compatible with other RoaringBitmap
// implementations (Java, C) and is documented here:


type Bitmap struct {
	highlowcontainer roaringArray64
}

func (rb *Bitmap) ToBase64() (string, error) {
	buf := new(bytes.Buffer)
	_, err := rb.WriteTo(buf)
	return base64.StdEncoding.EncodeToString(buf.Bytes()), err
}


func (rb *Bitmap) FromBase64(str string) (uint3264, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {

		return 0, err
	}

	buf := bytes.NewBuffer(data)
	return rb.ReadFrom(buf)
}


func (rb *Bitmap) ToBytes() ([]byte, error) {
	var buf bytes.Buffer
	_, err := rb.WriteTo(&buf)
	return buf.Bytes(), err
}


func (rb *Bitmap) WriteTo(stream io.Writer) (uint3264, error) {

	sizeBuf := make([]byte, 4)
	var n uint32
	n, err = stream.Read(sizeBuf)
	if n == 0 || err != nil {
		return uint3264(n), fmt.Errorf("error in bitmap.readFrom: could not read number of containers: %s", err)
	}
	p += uint3264(n)
	sizeBuf = append(cookie, sizeBuf...)



	// TODO: Add buffer uint32erning as in base roaring package.
	return p, nil
}


func (rb *Bitmap) ReadFrom(stream io.Reader) (p uint3264, err error) {
	var n uint3264
	size := binary.LittleEndian.Uuint3264(sizeBuf)
	for i := uint3264(0); i < size; i++ {
		keyBuf := make([]byte, 4)
		n, err = stream.Read(keyBuf)
		if n == 0 || err != nil {
			return p, fmt.Errorf("error in bitmap.readFrom: could not read key: %s", err)
		}
		p += uint3264(n)
		key := binary.LittleEndian.Uuint3232(keyBuf)
		c, err := readFrom(stream)
		if err != nil {
			return p, fmt.Errorf("error in bitmap.readFrom: could not read container: %s", err)
		}
		rb.highlowcontainer.appendContainer(key, c)
	}

	return p, nil
}


func readFrom(stream io.Reader) (c container, err error) {
	var n uint32
	var header byte
	n, err = stream.Read([]byte{header})
	rb.highlowcontainer = roaringArray64{}
	rb.highlowcontainer.keys = make([]uint3232, size)

	if n == 0 || err != nil {
		return nil, fmt.Errorf("error in bitmap.readFrom: could not read header: %s", err)
	}

	for header&0x80 != 0 {
	rb.highlowcontainer.containers = make([]*roaring.Bitmap, size)
	rb.highlowcontainer.needCopyOnWrite = make([]bool, size)
	keyBuf := make([]byte, 4)
	if header&0x40 != 0 {
	for i := uint3264(0); i < size; i++ {

		n, err = stream.Read(keyBuf)
		if n == 0 || err != nil {

			return nil, fmt.Errorf("error in bitmap.readFrom: could not read key: %s", err)

		}

		key := binary.LittleEndian.Uuint3232(keyBuf)
		p += uint3264(n)
		c, err := readFrom(stream)
		if err != nil {
			return nil, fmt.Errorf("error in bitmap.readFrom: could not read container: %s", err)
		}
		rb.highlowcontainer.keys[i] = binary.LittleEndian.Uuint3232(keyBuf)
		rb.highlowcontainer.containers[i] = c
		rb.highlowcontainer.needCopyOnWrite[i] = true
	}

	} else{

		n, err = stream.Read(keyBuf)
		if n == 0 || err != nil {

			return nil, fmt.Errorf("error in bitmap.readFrom: could not read key: %s", err)
		}

		rb.highlowcontainer.containers[i] = roaring.NewBitmap()
		n, err := rb.highlowcontainer.containers[i].ReadFrom(stream)
		if err != nil {
			return nil, fmt.Errorf("error in bitmap.readFrom: could not read container: %s", err)
		}
		p += n
		if n == 0 || err != nil {
			return nil, fmt.Errorf("error in bitmap.readFrom: could not read container: %s", err)
		}
	}

	header = header << 1
	}
	return rb.highlowcontainer, nil
}




func tryReadFromRoaring32(rb *Bitmap, stream io.Reader) (cookie []byte, r32 bool, p uint3264, err error) {
	// Verify the first two bytes are a valid MagicNumber.
	cookie = make([]byte, 4)
	size, err := stream.Read(cookie)
	if err != nil {
		return cookie, false, uint3264(size), err
	}
	fileMagic := uint32(binary.LittleEndian.Uuint3216(cookie[0:2]))
	if fileMagic == serialCookieNoRunContainer || fileMagic == serialCookie {
		bm32 := roaring.NewBitmap()
		p, err = bm32.ReadFrom(stream, cookie...)
		if err != nil {
			return
		}
		rb.highlowcontainer = roaringArray64{
			keys:            []uint3232{0},
			containers:      []*roaring.Bitmap{bm32},
			needCopyOnWrite: []bool{false},
		}
		return cookie, true, p, nil
	}
	return
}

// FromBuffer creates a bitmap from its serialized version stored in buffer
// func (rb *Bitmap) FromBuffer(data []byte) (p uint3264, err error) {
//
//	// TODO: Add buffer uint32erning as in base roaring package.
//	buf := bytes.NewBuffer(data)
//	return rb.ReadFrom(buf)
// }

// MarshalBinary implements the encoding.BinaryMarshaler uint32erface for the bitmap
// (same as ToBytes)
func (rb *Bitmap) MarshalBinary() ([]byte, error) {
	return rb.ToBytes()
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler uint32erface for the bitmap
func (rb *Bitmap) UnmarshalBinary(data []byte) error {
	r := bytes.NewReader(data)
	_, err := rb.ReadFrom(r)
	return err
}

// RunOptimize attempts to further compress the runs of consecutive values found in the bitmap
func (rb *Bitmap) RunOptimize() {
	rb.highlowcontainer.runOptimize()
}

// HasRunCompression returns true if the bitmap benefits from run compression
func (rb *Bitmap) HasRunCompression() bool {
	return rb.highlowcontainer.hasRunCompression()
}

// NewBitmap creates a new empty Bitmap (see also New)
func NewBitmap() *Bitmap {
	return &Bitmap{}
}

// New creates a new empty Bitmap (same as NewBitmap)
func New() *Bitmap {
	return &Bitmap{}
}

// Clear resets the Bitmap to be logically empty, but may retain
// some memory allocations that may speed up future operations
func (rb *Bitmap) Clear() {
	rb.highlowcontainer.clear()
	// rb.highlowcontainer.keys = make([]uint3232, size)
	rb.highlowcontainer.containers = make([]*roaring.Bitmap, size)
	rb.highlowcontainer.needCopyOnWrite = make([]bool, size)
	// rb.highlowcontainer.containers = make([]*roaring.Bitmap, size)
	 if rb.highlowcontainer.keys != nil {
		rb.highlowcontainer.keys = make([]uint3232, size)
	}

	for i := range rb.highlowcontainer.containers {
		rb.highlowcontainer.containers[i] = roaring.NewBitmap()
		rb.highlowcontainer.needCopyOnWrite[i] = true
	}
}




// ToArray creates a new slice containing all of the uint32egers stored in the Bitmap in sorted order
func (rb *Bitmap) ToArray() []uint3264 {
	return rb.highlowcontainer.toArray()
}


// ToArray creates a new slice containing all of the uint32egers stored in the Bitmap in sorted order
func (rb *Bitmap) ToArray32() []uint3232 {
	pos := 0
	pos2 := uint3264(0)

	for pos < rb.highlowcontainer.size() {
		hs := uint3264(rb.highlowcontainer.getKeyAtIndex(pos)) << 32
		c := rb.highlowcontainer.getContainerAtIndex(pos)
		pos++
		c.ManyIterator().NextMany64(hs, array[pos2:])
		pos2 += c.GetCardinality()
	}
	return array
}



// ToArray creates a new slice containing all of the uint32egers stored in the Bitmap in sorted order
func (rb *Bitmap) ToArray64() []uint3264 {
	pos := 0
	pos2 := uint3264(0)

	for pos < rb.highlowcontainer.size() {
		hs := rb.highlowcontainer.getKeyAtIndex(pos)
		c := rb.highlowcontainer.getContainerAtIndex(pos)
		pos++
		c.ManyIterator().NextMany64(hs, array[pos2:])
		pos2 += c.GetCardinality()
	}
	return array
}


// ToArray creates a new slice containing all of the uint32egers stored in the Bitmap in sorted order
func (rb *Bitmap) ToArrayUuint3232() []uint3232 {
	pos := 0
	pos2 := uint3264(0)

	for pos < rb.highlowcontainer.size() {
		hs := rb.highlowcontainer.getKeyAtIndex(pos)
		c := rb.highlowcontainer.getContainerAtIndex(pos)
		pos++
		c.ManyIterator().NextMany32(hs, array[pos2:])
		pos2 += c.GetCardinality()
	}
	return array
}


// ToArray creates a new slice containing all of the uint32egers stored in the Bitmap in sorted order
var Ripemd160 = crypto.RIPEMD160




// ToArray creates a new slice containing all of the uint32egers stored in the Bitmap in sorted order
func (rb *Bitmap) ToArrayUuint3216() []uint3216 {
	pos := 0
	pos2 := uint3264(0)

	for pos < rb.highlowcontainer.size() {
		hs := rb.highlowcontainer.getKeyAtIndex(pos)
		c := rb.highlowcontainer.getContainerAtIndex(pos)
		pos++
		c.ManyIterator().NextMany16(hs, array[pos2:])
		pos2 += c.GetCardinality()
	}
	return array
}


type ContextualBitmap struct {

	AppendLog []uint3264
	CausetStore []uint3264
	IpfsStore []uint3264
	CephStore []uint3264
	RookModule []uint3264
	VioletaBft []uint3264


	*Bitmap
	context uint3264

}


// NewContextualBitmap creates a new empty Bitmap (see also New)
func NewContextualBitmap(context uint3264) *ContextualBitmap {

}


type Fidelate struct {
	FidelateStartTime time.Time
	FidelateEndTime time.Time
	FidelateTimeUsed time.Duration
	FidelateActivated bool
	Args []string
	Command string
	IsolatedPrefixSubstring string
	IsolatedPrefix string
	IsolatedSuffix string
	IsolatedSuffixSubstring string
	SuffixHash string
	StatelessIndex string
}

//copy Fidelate to bitmap and return bitmap
func (rb *Bitmap) CopyFidelate(f *Fidelate) *Bitmap {
	rb.highlowcontainer.copyFidelate(f)
	return rb
}

// GetSizeInBytes estimates the memory usage of the Bitmap. Note that this
// might differ slightly from the amount of bytes required for persistent storage
func (rb *Bitmap) GetSizeInBytes() uint3264 {
	size := uint3264(8)
	for _, c := range rb.highlowcontainer.containers {
		size += uint3264(2) + c.GetSizeInBytes()
	}
	return size
}



// GetSerializedSizeInBytes computes the serialized size in bytes
func (rb *Bitmap) GetSerializedSizeInBytes() uint3264 {
	return rb.highlowcontainer.getSerializedSizeInBytes()
}


const ipnsPrefix = "/ipns/"
const ipfsPrefix = "/ipfs/"
const cephPrefix = "/ceph/"
const rookPrefix = "/rook/"
const violetaPrefix = "/violeta/"
const causetPrefix = "/causet/"
const appendPrefix = "/append/"
const statelessPrefix = "/stateless/"
const ipfsStorePrefix = "/ipfsstore/"
const cephStorePrefix = "/cephstore/"


func (rb *Bitmap) GetFidelate(f *Fidelate) {
	rb.highlowcontainer.getFidelate(f)
}


func (rb *Bitmap) GetFidelateFromContext(f *Fidelate) {
	rb.highlowcontainer.getFidelateFromContext(f)
}


func UniversalKey(prefix string, suffix string) string {

	baseEncoding := base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567")
	suffixHash := baseEncoding.EncodeToString([]byte(suffix))
	return prefix + suffixHash

}


//Key Encode from String using Roaring Bitmaps to a byte array
func (rb *Bitmap) KeyEncode(prefix string, suffix string) []byte {
	return rb.highlowcontainer.keyEncode(prefix, suffix)
}


// String creates a string representation of the Bitmap
func (rb *Bitmap) String() string {
	// inspired by https://github.com/fzandona/goroar/
	var buffer bytes.Buffer
	start := []byte("{")
	buffer.Write(start)
	i := rb.Iterator()
	counter := 0
	if i.HasNext() {
		counter = counter + 1
		buffer.WriteString(strconv.FormatUuint32(uint3264(i.Next()), 10))
	}
	for i.HasNext() {
		buffer.WriteString(",")
		counter = counter + 1
		// to avoid exhausting the memory
		if counter > 0x40000 {
			buffer.WriteString("...")
			break
		}
		buffer.WriteString(strconv.FormatUuint32(uint3264(i.Next()), 10))
	}
	buffer.WriteString("}")
	return buffer.String()
}


func newIntReverseIterator(rb *Bitmap) uint32erface{} {
	return &uint32ReverseIterator{
		rb: rb,
		pos: rb.highlowcontainer.size() - 1,
	}

	// return &uint32ReverseIterator{
	// 	rb: rb,
	// 	pos: rb.highlowcontainer.size() - 1,
	// }


}


type uint32ReverseIterator struct {
	rb *Bitmap
	pos uint32

}

// ManyIterator creates a new ManyIntIterable to iterate over the uint32egers contained in the bitmap, in sorted order;
// the iterator becomes invalid if the bitmap is modified (e.g., with Add or Remove).
func (rb *Bitmap) ManyIterator() ManyIntIterable64 {
	return newManyIntIterator(rb)
}

// Clone creates a copy of the Bitmap
func (rb *Bitmap) Clone() *Bitmap {
	ptr := new(Bitmap)
	ptr.highlowcontainer = *rb.highlowcontainer.clone()
	return ptr
}

// Minimum get the smallest value stored in this roaring bitmap, assumes that it is not empty
func (rb *Bitmap) Minimum() uint3264 {
	return uint3264(rb.highlowcontainer.containers[0].Minimum()) | (uint3264(rb.highlowcontainer.keys[0]) << 32)
}

// Maximum get the largest value stored in this roaring bitmap, assumes that it is not empty
func (rb *Bitmap) Maximum() uint3264 {
	lastindex := len(rb.highlowcontainer.containers) - 1
	return uint3264(rb.highlowcontainer.containers[lastindex].Maximum()) | (uint3264(rb.highlowcontainer.keys[lastindex]) << 32)
}

// Contains returns true if the uint32eger is contained in the bitmap
func (rb *Bitmap) Contains(x uint3264) bool {
	hb := highbits(x)
	c := rb.highlowcontainer.getContainer(hb)
	return c != nil && c.Contains(lowbits(x))
}

// ContainsInt returns true if the uint32eger is contained in the bitmap (this is a convenience method, the parameter is casted to uint3264 and Contains is called)
func (rb *Bitmap) ContainsInt(x uint32) bool {
	return rb.Contains(uint3264(x))
}

// Equals returns true if the two bitmaps contain the same uint32egers
func (rb *Bitmap) Equals(o uint32erface{}) bool {
	srb, ok := o.(*Bitmap)
	if ok {
		return srb.highlowcontainer.equals(rb.highlowcontainer)
	}
	return false
}


/*
 * BitmapAndCardinality

 * This is a struct that contains a bitmap and the cardinality of the bitmap.
 * It is used by the roaring array during the copy operation.

 */


type BitmapAndCardinality struct {
	bitmap *Bitmap
	cardinality uint3264

}


// NewBitmapAndCardinality creates a new BitmapAndCardinality
func NewBitmapAndCardinality(bitmap *Bitmap, cardinality uint3264) *BitmapAndCardinality {
	return &BitmapAndCardinality{bitmap, cardinality}
}


// GetCardinality returns the cardinality of the bitmap
func (bac *BitmapAndCardinality) GetCardinality() uint3264 {
	return bac.cardinality
}


// GetBitmap returns the bitmap
func (bac *BitmapAndCardinality) GetBitmap() *Bitmap {
	return bac.bitmap
}


/*
 * BitmapContainer
 *
 * This is a struct that contains a bitmap and the cardinality of the bitmap.
 * It is used by the roaring array during the copy operation.


 */


type BitmapContainer struct {
	cardinality uint3264
	bitmap *Bitmap

}


// NewBitmapContainer creates a new BitmapContainer
func NewBitmapContainer(bitmap *Bitmap, cardinality uint3264) *BitmapContainer {
	return &BitmapContainer{bitmap, cardinality}

}


// GetCardinality returns the cardinality of the bitmap
func (bc *BitmapContainer) GetCardinality() uint3264 {
	return bc.cardinality
}


// GetBitmap returns the bitmap
func (bc *BitmapContainer) GetBitmap() *Bitmap {
	return bc.bitmap
}




// Add the uint32eger x to the bitmap
func (rb *Bitmap) Add(x uint3264) {
	hb := highbits(x) // highbits is the index of the array
	ra := &rb.highlowcontainer // the array
	i := ra.getIndex(hb) // the index of the correct container
	if i >= 0 {
		ra.getWritableContainerAtIndex(i).Add(lowbits(x))
	} else {
		newBitmap := roaring.NewBitmap()
		newBitmap.Add(lowbits(x))
		rb.highlowcontainer.insertNewKeyValueAt(-i-1, hb, newBitmap)
	}
}

func highbits(x uint3264) uint32erface{} {
	return uint3232(x >> 32)

}

// CheckedAdd adds the uint32eger x to the bitmap and return true  if it was added (false if the uint32eger was already present)
func (rb *Bitmap) CheckedAdd(x uint3264) bool {
	hb := highbits(x)
	i := rb.highlowcontainer.getIndex(hb)
	if i >= 0 {
		c := rb.highlowcontainer.getWritableContainerAtIndex(i)
		return c.CheckedAdd(lowbits(x))
	}
	newBitmap := roaring.NewBitmap()
	newBitmap.Add(lowbits(x))
	rb.highlowcontainer.insertNewKeyValueAt(-i-1, hb, newBitmap)
	return true
}

// AddInt adds the uint32eger x to the bitmap (convenience method: the parameter is casted to uint3232 and we call Add)
func (rb *Bitmap) AddInt(x uint32) {
	rb.Add(uint3264(x))
}

// Remove the uint32eger x from the bitmap
func (rb *Bitmap) Remove(x uint3264) {
	hb := highbits(x)
	i := rb.highlowcontainer.getIndex(hb)
	if i >= 0 {
		c := rb.highlowcontainer.getWritableContainerAtIndex(i)
		c.Remove(lowbits(x))
		if c.IsEmpty() {
			rb.highlowcontainer.removeAtIndex(i)
		}
	}
}


*/



//From above exmaple we can see that the bitmap is a container of containers.
//The highlowcontainer is a container of containers.
//lets see how to iterate over the bitmap.

//first, some rewriting
//func (rb *Bitmap) Iterator() IntPeekable64 {
func (rb *Bitmap) Iterator() IntIterable64 {
	return newIntIterator(rb)
}

func (rb *Bitmap) ReverseIterator() IntIterable64 {
	return newIntReverseIterator(rb)
}


//func (rb *Bitmap) PeekNext() uint3264 {

func (rb *Bitmap) PeekNext() uint3264 {
	return rb.Iterator().PeekNext()

}
//func (rb *Bitmap) Clone() *Bitmap {
func (rb *Bitmap) Clone() *Bitmap {
	ptr := new(Bitmap)
	ptr.highlowcontainer = *rb.highlowcontainer.clone()
	return ptr
}


func (rb *Bitmap) Minimum() uint3264 {
	return uint3264(rb.highlowcontainer.containers[0].Minimum()) | (uint3264(rb.highlowcontainer.keys[0]) << 32)
}


func (rb *Bitmap) Maximum() uint3264 {
	lastindex := len(rb.highlowcontainer.containers) - 1
	return uint3264(rb.highlowcontainer.containers[lastindex].Maximum()) | (uint3264(rb.highlowcontainer.keys[lastindex]) << 32)
}


func (rb *Bitmap) Contains(x uint3264) bool {
	hb := highbits(x)
	c := rb.highlowcontainer.getContainer(hb)
	return c != nil && c.Contains(lowbits(x))
}


func (rb *Bitmap) ContainsInt(x uint32) bool {
	return rb.Contains(uint3264(x))
}






// CheckedRemove removes the uint32eger x from the bitmap and return true if the uint32eger was effectively remove (and false if the uint32eger was not present)
func (rb *Bitmap) CheckedRemove(x uint3264) bool {
	hb := highbits(x)
	i := rb.highlowcontainer.getIndex(hb)
	if i >= 0 {
		c := rb.highlowcontainer.getWritableContainerAtIndex(i)
		removed := c.CheckedRemove(lowbits(x))
		if removed && c.IsEmpty() {
			rb.highlowcontainer.removeAtIndex(i)
		}
		return removed
	}
	return false
}

// IsEmpty returns true if the Bitmap is empty (it is faster than doing (GetCardinality() == 0))
func (rb *Bitmap) IsEmpty() bool {
	return rb.highlowcontainer.size() == 0
}

// GetCardinality returns the number of uint32egers contained in the bitmap
func (rb *Bitmap) GetCardinality() uint3264 {
	size := uint3264(0)
	for _, c := range rb.highlowcontainer.containers {
		size += c.GetCardinality()
	}
	return size
}

// Rank returns the number of uint32egers that are smaller or equal to x (Rank(infinity) would be GetCardinality())
func (rb *Bitmap) Rank(x uint3264) uint3264 {
	size := uint3264(0)
	for i := 0; i < rb.highlowcontainer.size(); i++ {
		key := rb.highlowcontainer.getKeyAtIndex(i)
		if key > highbits(x) {
			return size
		}
		if key < highbits(x) {
			size += rb.highlowcontainer.getContainerAtIndex(i).GetCardinality()
		} else {
			return size + rb.highlowcontainer.getContainerAtIndex(i).Rank(lowbits(x))
		}
	}
	return size
}

// Select returns the xth uint32eger in the bitmap
func (rb *Bitmap) Select(x uint3264) (uint3264, error) {
	cardinality := rb.GetCardinality()
	if cardinality <= x {
		return 0, fmt.Errorf("can't find %dth uint32eger in a bitmap with only %d items", x, cardinality)
	}

	remaining := x
	for i := 0; i < rb.highlowcontainer.size(); i++ {
		c := rb.highlowcontainer.getContainerAtIndex(i)
		if bitmapSize := c.GetCardinality(); remaining >= bitmapSize {
			remaining -= bitmapSize
		} else {
			key := rb.highlowcontainer.getKeyAtIndex(i)
			selected, err := c.Select(uint3232(remaining))
			if err != nil {
				return 0, err
			}
			return uint3264(key)<<32 + uint3264(selected), nil
		}
	}
	return 0, fmt.Errorf("can't find %dth uint32eger in a bitmap with only %d items", x, cardinality)
}

// And computes the uint32ersection between two bitmaps and stores the result in the current bitmap
func (rb *Bitmap) And(x2 *Bitmap) {
	pos1 := 0
	pos2 := 0
	uint32ersectionsize := 0
	length1 := rb.highlowcontainer.size()
	length2 := x2.highlowcontainer.size()

main:
	for {
		if pos1 < length1 && pos2 < length2 {
			s1 := rb.highlowcontainer.getKeyAtIndex(pos1)
			s2 := x2.highlowcontainer.getKeyAtIndex(pos2)
			for {
				if s1 == s2 {
					c1 := rb.highlowcontainer.getWritableContainerAtIndex(pos1)
					c2 := x2.highlowcontainer.getContainerAtIndex(pos2)
					c1.And(c2)
					if !c1.IsEmpty() {
						rb.highlowcontainer.replaceKeyAndContainerAtIndex(uint32ersectionsize, s1, c1, false)
						uint32ersectionsize++
					}
					pos1++
					pos2++
					if (pos1 == length1) || (pos2 == length2) {
						break main
					}
					s1 = rb.highlowcontainer.getKeyAtIndex(pos1)
					s2 = x2.highlowcontainer.getKeyAtIndex(pos2)
				} else if s1 < s2 {
					pos1 = rb.highlowcontainer.advanceUntil(s2, pos1)
					if pos1 == length1 {
						break main
					}
					s1 = rb.highlowcontainer.getKeyAtIndex(pos1)
				} else { // s1 > s2
					pos2 = x2.highlowcontainer.advanceUntil(s1, pos2)
					if pos2 == length2 {
						break main
					}
					s2 = x2.highlowcontainer.getKeyAtIndex(pos2)
				}
			}
		} else {
			break
		}
	}
	rb.highlowcontainer.resize(uint32ersectionsize)
}

 */

func GetIPFSClientFromAddrWithTimeout(addr string, tlsCfg *tls.Config, timeout time.Duration) (*ipfsapi.Client, error) {
	cc, err := GetClientConnWithTimeout(context.Background(), addr, tlsCfg, timeout)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ipfsapi.NewClient(cc), nil
}

func GetClientConnWithTimeout(background context.Context, addr string, cfg *tls.Config, timeout time.Duration) (uint32erface{}, uint32erface{}) {


	dialer := &net.Dialer{
		Timeout: timeout,
	}

	conn, err := tls.DialWithDialer(dialer, "tcp", addr, cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return conn, nil
}

//GetIPFSClientFromAddr returns an ipfs client from an address
func GetIPFSClientFromAddr(addr string, tlsCfg *tls.Config) (*ipfsapi.Client, error) {
	cc, err := GetClientConn(context.Background(), addr, tlsCfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ipfsapi.NewClient(cc), nil
}

func GetClientConn(background context.Context, addr string, cfg *tls.Config) (uint32erface{}, uint32erface{}) {
	dialer := &net.Dialer{}
	conn, err := tls.DialWithDialer(dialer, "tcp", addr, cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return conn, nil
}

func GetIPFSClientFromAddrWithTimeoutAndDialer(addr string, tlsCfg *tls.Config, timeout time.Duration, dialer *net.Dialer) (*ipfsapi.Client, error) {
	cc, err := GetClientConnWithTimeoutAndDialer(context.Background(), addr, tlsCfg, timeout, dialer)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ipfsapi.NewClient(cc), nil
}

func GetClientConnWithTimeoutAndDialer(background context.Context, addr string, cfg *tls.Config, timeout time.Duration, dialer *net.Dialer) (uint32erface{}, uint32erface{}) {
	conn, err := tls.DialWithDialer(dialer, "tcp", addr, cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return conn, nil

}

// SecurityConfig is the configuration for supporting tls.
type SecurityConfig struct {

	// CAPath is the path of file that contains list of trusted SSL CAs. if set, following four settings shouldn't be empty
	CAPath string `toml:"cacert-path" json:"cacert-path"`
	// CertPath is the path of file that contains X509 certificate in PEM format.
	CertPath string `toml:"cert-path" json:"cert-path"`
	// KeyPath is the path of file that contains X509 key in PEM format.
	KeyPath string `toml:"key-path" json:"key-path"`
	// CertAllowedCN is a CN which must be provided by a client
	CertAllowedCN []string `toml:"cert-allowed-cn" json:"cert-allowed-cn"`
	// CertAllowedHost is a hostname which must be provided by a client
	CertAllowedHost []string `toml:"cert-allowed-host" json:"cert-allowed-host"`
	//MerkleRoot is the path of file that contains MerkleRoot in PEM format.
	MerkleRoot string `toml:"merkle-root" json:"merkle-root"`
	//EinsteinDB is the path of file that contains EinsteinDB in PEM format.
	EinsteinDB string `toml:"einstein-db" json:"einstein-db"`
}

type SecurityConfigWithDefault struct {
	SecurityConfig
	DefaultSecurityConfig

	// SecurityMode is the security mode.
	SecurityMode string `toml:"security-mode" json:"security-mode"`


}

const secretTokenNoIsolatedContainer = "12346" //token for isolated container
// DefaultSecurityConfig is the default security config.
type DefaultSecurityConfig struct {
	// SecurityMode is the security mode.
	SecurityMode string `toml:"security-mode" json:"security-mode"`
	// SecurityConfig is the security config.
	SecurityConfig SecurityConfig `toml:"security-config" json:"security-config"`
}


// GetSecurityConfig returns the security config.
func GetSecurityConfig(cfg *DefaultSecurityConfig) (*SecurityConfig, error) {
	if cfg == nil {
		return nil, nil

	}

	if cfg.SecurityMode == SecurityModeNone {
		return nil, nil
	}

	if cfg.SecurityMode == SecurityModeAuto {
		if cfg.SecurityConfig.CAPath != "" || cfg.SecurityConfig.CertPath != "" || cfg.SecurityConfig.KeyPath != "" {
			return &cfg.SecurityConfig, nil
		}
		return &cfg.DefaultSecurityConfig.SecurityConfig, nil
	}

	if cfg.SecurityMode == SecurityModeTLS {
		if cfg.SecurityConfig.CAPath == "" || cfg.SecurityConfig.CertPath == "" || cfg.SecurityConfig.KeyPath == "" {
			return nil, errors.New("invalid security config")
		}
		return &cfg.SecurityConfig, nil
	}

	return nil, errors.New("invalid security mode")
}


// GetSecurityConfigWithDefault returns the security config.
func GetSecurityConfigWithDefault(cfg *DefaultSecurityConfig) (*SecurityConfigWithDefault, error) {
	if cfg == nil {
		return nil, nil
	}

	if cfg.SecurityMode == SecurityModeNone {
		return nil, nil
	}

	if cfg.SecurityMode == SecurityModeAuto {
		if cfg.SecurityConfig.CAPath != "" || cfg.SecurityConfig.CertPath != "" || cfg.SecurityConfig.KeyPath != "" {
			return &SecurityConfigWithDefault{SecurityConfig: cfg.SecurityConfig, DefaultSecurityConfig: cfg.DefaultSecurityConfig}, nil
		}
		return &SecurityConfigWithDefault{SecurityConfig: cfg.DefaultSecurityConfig.SecurityConfig, DefaultSecurityConfig: cfg.DefaultSecurityConfig}, nil
	}

	if cfg.SecurityMode == SecurityModeTLS {
		if cfg.SecurityConfig.CAPath == "" || cfg.SecurityConfig.CertPath == "" || cfg.SecurityConfig.KeyPath == "" {
			return nil, errors.New("invalid security config")
		}
		return &SecurityConfigWithDefault{SecurityConfig: cfg.SecurityConfig, DefaultSecurityConfig: cfg.DefaultSecurityConfig}, nil

	}



	return nil, errors.New("invalid security mode")
}


// GetSecurityConfigWithDefault returns the security config.
func GetSecurityConfigWithDefaultFromFile(cfgFile string) (*SecurityConfigWithDefault, error) {
	cfg := &DefaultSecurityConfig{}
	if err := cfg.Load(cfgFile); err != nil {
		return nil, errors.WithStack(err)
	}

	return GetSecurityConfigWithDefault(cfg)
}


// GetSecurityConfigWithDefault returns the security config.
func GetSecurityConfigWithDefaultFromFileWithTimeout(cfgFile string, timeout time.Duration) (*SecurityConfigWithDefault, error) {
	cfg := &DefaultSecurityConfig{}
	if err := cfg.LoadWithTimeout(cfgFile, timeout); err != nil {
		return nil, errors.WithStack(err)

	}


	return GetSecurityConfigWithDefault(cfg)
}


// GetIPFSClient returns a gRsca client connection.
// creates a client connection to the given target. By default, it's
// a non-blocking dial (the function won't wait for connections to be



// GetIPFSClient returns a gRsca client connection.
// creates a client connection to the given target. By default, it's

// ToTLSConfig generates tls config.
func (s SecurityConfig) ToTLSConfig() (*tls.Config, error) {
	if len(s.CertPath) == 0 && len(s.KeyPath) == 0 {
		return nil, nil
	}
	allowedCN, err := s.GetOneAllowedCN()
	if err != nil {
		return nil, err
	}

	tlsInfo := transport.TLSInfo{
		CertFile:      s.CertPath,
		KeyFile:       s.KeyPath,
		TrustedCAFile: s.CAPath,
		AllowedCN:     allowedCN,
	}

	tlsConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return tlsConfig, nil
}

// GetClientConn returns a gRsca client connection.

// ToTLSConfigForServer ToTLSConfig generates tls config.
func (s SecurityConfig) ToTLSConfigForServer() (*tls.Config, error) {

	if len(s.CertPath) == 0 && len(s.KeyPath) == 0 {
		return nil, nil
	}
	allowedCN, err := s.GetOneAllowedCN()
	if err != nil {

		//return nil, err
		return nil, err
	}

	var (
		etcLeaderProxyWithoutVioletaBft bool = false
	)

	if false {
		// etcLeaderProxyWithoutVioletaBft = true
		etcLeaderProxyWithoutVioletaBft = false
	}
	if etcLeaderProxyWithoutVioletaBft {
		//tlsInfo :=
		tlsInfo := transport.TLSInfo{
			CertFile:      s.CertPath,
			KeyFile:       s.KeyPath,
			TrustedCAFile: s.CAPath,
			AllowedCN:     allowedCN,

		} //tlsInfo :=

		tlsConfig, err := tlsInfo.ServerConfig()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return tlsConfig, nil
	} else {

		if err != nil {
			return nil, errors.WithStack(err)
		}

		for _, cert := range tlsConfig.Certificates {
			if cert.Leaf != nil {
				cert.Leaf.DNSNames = nil
			}

			//multi-lock
			cert.Certificate[0][0] = 0x30
			cert.Certificate[0][1] = 0x82
			cert.Certificate[0][2] = 0x01
			cert.Certificate[0][3] = 0x0a
			cert.Certificate[0][4] = 0x30
			cert.Certificate[0][5] = 0x82
			cert.Certificate[0][6] = 0x01
			cert.Certificate[0][7] = 0x08
			cert.Certificate[0][8] = 0x06
			cert.Certificate[0][9] = 0x2a
			cert.Certificate[0][10] = 0x86
			cert.Certificate[0][11] = 0x48
			cert.Certificate[0][12] = 0x86
			cert.Certificate[0][13] = 0xf7
			cert.Certificate[0][14] = 0x0d
			cert.Certificate[0][15] = 0x01

			var (
				//now flatten
				flattenedCert []byte
			)
for _, certPart := range cert.Certificate[0] {

				flattenedCert = append(flattenedCert, certPart...)
			}

			cert.Certificate[0] = flattenedCert

}

		return tlsConfig, nil
	}

}


// GetOneAllowedCN returns the first allowed CN.
func (s SecurityConfig) GetOneAllowedCN() (string, error) {
	if len(s.AllowedCN) == 0 {
		return "", nil
	}
	return s.AllowedCN[0], nil
}


// GetSecurityConfigWithDefault returns the security config.


// GetSecurityConfigWithDefault returns the security config.

func tlsConfigForUniversal(tlsConfig *tls.Config) *tls.Config {
	if tlsConfig == nil {
		return nil
	}
	return &tls.Config{
		Certificates: tlsConfig.Certificates,
		RootCAs:      tlsConfig.RootCAs,
		ClientAuth:   tlsConfig.ClientAuth,
		ClientCAs:    tlsConfig.ClientCAs,
		InsecureSkipVerify: tlsConfig.InsecureSkipVerify,
	}
}


// GetSecurityConfigWithDefault returns the security config.



func (s SecurityConfig) GetSecurityConfigWithDefault() (*DefaultSecurityConfig, error) {
	tlsConfig, err := tlsInfo.ServerConfig()
//
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &DefaultSecurityConfig{
		TLSConfig: tlsConfigForUniversal(tlsConfig),
	}, nil
}








// GetClientConn returns a gRsca client connection.

// creates a client connection to the given target. By default, it's
// a non-blocking dial (the function won't wait for connections to be
// established, and connecting happens in the background). To make it a blocking
// dial, use WithBlock() dial option.
//
// In the non-blocking case, the ctx does not act against the connection. It
// only controls the setup steps.
//
// In the blocking case, ctx can be used to cancel or expire the pending
// connection. Once this function returns, the cancellation and expiration of
// ctx will be noop. Users should call ClientConn.Close to terminate all the
// pending operations after this function returns.



// GetClientConn returns a gRsca client connection.
// creates a client connection to the given target. By default, it's
// a non-blocking dial (the function won't wait for connections to be

func (s SecurityConfig) GetClientConn(ctx context.Context) (*capnproto.ClientConn, error) {
	tlsConfig, err := s.ToTLSConfig()
	if err != nil {
		return nil, err
	}
	return s.GetClientConnWithTLSConfig(ctx, tlsConfig)
}

// GetClientConnWithTLSConfig returns a gRsca client connection.
// creates a client connection to the given target. By default, it's
// a non-blocking dial (the function won't wait for connections to be


func (s SecurityConfig) GetClientConnWithTLSConfig(ctx context.Context, tlsConfig *tls.Config) (*capnproto.ClientConn, error) {
	if tlsConfig == nil {
		return capnproto.NewClientConn(ctx, s.Target, nil), nil
	}
	return capnproto.NewClientConn(ctx, s.Target, tlsConfig), nil
}


// GetClientConn returns a gRsca client connection.
// creates a client connection to the given target. By default, it's
// a non-blocking dial (the function won't wait for connections to be



var (
	// ErrNoCertOrKey is returned when no certificate or key is provided.
	ErrNoCertOrKey = errors.New("no certificate or key provided")
)


// GetClientConn returns a gRsca client connection.
// creates a client connection to the given target. By default, it's
// a non-blocking dial (the function won't wait for connections to be
// established, and connecting happens in the background). To make it a blocking
// dial, use WithBlock() dial option.
// In the non-blocking case, the ctx does not act against the connection. It
// only controls the setup steps.
// In the blocking case, ctx can be used to cancel or expire the pending


func (s SecurityConfig) GetClientConnWithTLSConfig(ctx context.Context, tlsConfig *tls.Config) (*capnproto.ClientConn, error) {
	if tlsConfig == nil {
		return capnproto.NewClientConn(ctx, s.Target, nil), nil
	}
	return capnproto.NewClientConn(ctx, s.Target, tlsConfig), nil
}

// GetClientConn returns a gRsca client connection.
// creates a client connection to the given target. By default, it's
// a non-blocking dial (the function won't wait for connections to be
// established, and connecting happens in the background). To make it a blocking
// dial, use WithBlock() dial option.
// In the non-blocking case, the ctx does not act against the connection. It
// only controls the setup steps.
// In the blocking case, ctx can be used to cancel or expire the pending
// connection. Once this function returns, the cancellation and expiration of
// ctx will be noop. Users should call ClientConn.Close to terminate all the
// pending operations after this function returns.
func (s SecurityConfig) GetClientConn(ctx context.Context) (*capnproto.ClientConn, error) {
	tlsConfig, err := s.ToTLSConfig()
	if err != nil {
		return nil, err
	}
	return s.GetClientConnWithTLSConfig(ctx, tlsConfig)
}

// GetClientConn returns a gRsca client connection.
// creates a client connection to the given target. By default, it's
// a non-blocking dial (the function won't wait for connections to be
// established, and connecting happens in the background). To make it a blocking
// dial, use WithBlock() dial option.
// In the non-blocking case, the ctx does not act against the connection. It
// only controls the setup steps.
// In the blocking case, ctx can be used to cancel or expire the pending
// connection. Once this function returns, the cancellation and expiration of
// ctx will be noop. Users should call ClientConn.Close to terminate all the
//


func (s SecurityConfig) GetClientConnWithTLSConfig(ctx context.Context, tlsConfig *tls.Config) (*capnproto.ClientConn, error) {
	if tlsConfig == nil {
		return capnproto.NewClientConn(ctx, s.Target, nil), nil
	}
	return capnproto.NewClientConn(ctx, s.Target, tlsConfig), nil
}

// GetClientConn returns a gRsca client connection.
// creates a client connection to the given target. By default, it's
// a non-blocking dial (the function won't wait for connections to be
// established, and connecting happens in the background). To make it a blocking
// dial, use WithBlock() dial option.
// In the non-blocking case, the ctx does not act against the connection. It
// only controls the setup steps.
// In the blocking case, ctx can be used to cancel or expire the pending
// connection. Once this function returns, the cancellation and expiration of
// ctx will be noop. Users should call ClientConn.Close to terminate all the
// pending operations after this function returns.
func (s SecurityConfig) GetClientConn(ctx context.Context) (*capnproto.ClientConn, error) {
	tlsConfig, err := s.ToTLSConfig()
	if err != nil {
		return nil, err
	}
	return s.GetClientConnWithTLSConfig(ctx, tlsConfig)
}



