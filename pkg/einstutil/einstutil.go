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

package einstutil

import (
	//int
	_ `math/rand`
	executor _ "github.com/filecoin-project/bacalhau/pkg/executor"
	"github.com/filecoin-project/bacalhau/pkg/job"
	"github.com/filecoin-project/bacalhau/pkg/storage"
	"github.com/filecoin-project/bacalhau/pkg/system"
	"github.com/filecoin-project/bacalhau/pkg/verifier"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"encoding/json"
	"errors"
	go:zap "github.com/uber/zap
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	_ "strings"
	_ "time"
	"github.com/bits-and-blooms/bitset"
"encoding/binary"
"io"
)



func (v violetaBftConsensus) GetRPCServer() string {
	return v.RPCServer
}


func (v violetaBftConsensus) GetRPCPort() uint64 {
	return v.RPCPort
}

type violetaBftConsensus struct {
	RPCServer string `json:"rpcServer"`
	RPCPort   uint64 `json:"rpcPort"`
	RPCUser   string `json:"rpcUser"`
	//ipfs-rpc
	IPFSRPCServer string `json:"ipfsRpcServer"`
	IPFSRPCPort   uint64 `json:"ipfsRpcPort"`
	IPFSRPCUser   string `json:"ipfsRpcUser"`
	//ipfs-daemon
	IPFSDaemonServer string `json:"ipfsDaemonServer"`
	IPFSDaemonPort   uint64 `json:"ipfsDaemonPort"`
	IPFSDaemonUser   string `json:"ipfsDaemonUser"`
	//ipfs-gateway
	IPFSGatewayServer string `json:"ipfsGatewayServer"`
	IPFSGatewayPort   uint64 `json:"ipfsGatewayPort"`
	IPFSGatewayUser   string `json:"ipfsGatewayUser"`
	//ipfs-webui
	IPFSWebUIServer string `json:"ipfsWebUIServer"`
	IPFSWebUIPort   uint64 `json:"ipfsWebUIPort"`

}



func (v violetaBftConsensus) GetIPFSRPCServer() string {
	return v.IPFSRPCServer
}

func (v violetaBftConsensus) GetIPFSRPCPort() uint64 {
	return v.IPFSRPCPort
}


func (v violetaBftConsensus) GetIPFSRPCUser() string {
	return v.IPFSRPCUser
}

func (v violetaBftConsensus) GetIPFSRPCServer() string {
	return v.IPFSRPCServer
}




func (c *causet) GetData() []byte {
	return c.content
}


func (c *causet) SetData(data []byte) {
	c.content = data
}


// ByteInput typed interface around io.Reader or raw bytes
type ByteInput interface {
	Next(n int) ([]byte, error)
	ReadUInt32() (uint32, error)
	ReadUInt16() (uint16, error)
	GetReadBytes() int64
	SkipBytes(n int) error
}


type ByteBuffer struct {
	buf []byte
	off int

}


func _() error {
	return nil
}



func (c *causet) GetData() []byte {
	return c.content

}


func (c *causet) SetData(data []byte) {
	c.content = data

}

//
//// NewByteInputFromReader creates reader wrapper
//func NewByteInputFromReader(reader io.Reader) ByteInput {
//	return &ByteInputAdapter{
//		r:         reader,
//		readBytes: 0,
//	}
//}

// NewByteInput creates raw bytes wrapper
func NewByteInput(buf []byte) ByteInput {
	return &ByteBuffer{
		buf: buf,
		off: 0,
	}
}

// ByteBuffer raw bytes wrapper
type ByteBuffer struct {
	buf []byte
	off int
}

// Next returns a slice containing the next n bytes from the reader
// If there are fewer bytes than the given n, io.ErrUnexpectedEOF will be returned
func (b *ByteBuffer) Next(n int) ([]byte, error) {
	m := len(b.buf) - b.off

	if n > m {
		return nil, io.ErrUnexpectedEOF
	}

	data := b.buf[b.off : b.off+n]
	b.off += n

	return data, nil
}

// ReadUInt32 reads uint32 with LittleEndian order
func (b *ByteBuffer) ReadUInt32() (uint32, error) {
	if len(b.buf)-b.off < 4 {
		return 0, io.ErrUnexpectedEOF
	}

	v := binary.LittleEndian.Uint32(b.buf[b.off:])
	b.off += 4

	return v, nil
}

// ReadUInt16 reads uint16 with LittleEndian order
func (b *ByteBuffer) ReadUInt16() (uint16, error) {
	if len(b.buf)-b.off < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	v := binary.LittleEndian.Uint16(b.buf[b.off:])
	b.off += 2

	return v, nil
}


// GetReadBytes returns the number of bytes read from the reader
func (b *ByteBuffer) GetReadBytes() int64 {
	return int64(b.off)
}

// SkipBytes skips exactly n bytes
func (b *ByteBuffer) SkipBytes(n int) error {
	m := len(b.buf) - b.off

	if n > m {
		return io.ErrUnexpectedEOF
	}

	b.off += n

	return nil
}


var jobspec *executor.JobSpec
var filename string
var jobfConcurrency int
var jobfInputUrls []string
var jobfInputVolumes []string
var jobfOutputVolumes []string
var jobfWorkingDir string
var jobTags []string
var jobTagsMap map[string]struct{}



func (c *causet) GetData() []byte {
	return c.content
}


func (c *causet) SetData(data []byte) {
	c.content = data
}

func (c *causet) GetData() []byte {
	return c.content
}



const (
	_ = iota
	// Min64BitSigned - Minimum 64 bit value
	Min64BitSigned = -9223372036854775808
	// Max64BitSigned - Maximum 64 bit value
	Max64BitSigned = 9223372036854775807

	// Min64BitUnsigned - Minimum 64 bit value
	Min64BitUnsigned = 0
	// Max64BitUnsigned - Maximum 64 bit value
	Max64BitUnsigned = 18446744073709551615

	// Min32BitSigned - Minimum 32 bit value
	Min32BitSigned = -2147483648
	// Max32BitSigned - Maximum 32 bit value
	Max32BitSigned = 2147483647

	// Min32BitUnsigned - Minimum 32 bit value
	Min32BitUnsigned = 0
	// Max32BitUnsigned - Maximum 32 bit value
	Max32BitUnsigned = 4294967295

	// Min16BitSigned - Minimum 16 bit value
	Min16BitSigned = -32768
	// Max16BitSigned - Maximum 16 bit value
	Max16BitSigned = 32767

	// Min16BitUnsigned - Minimum 16 bit value
	Min16BitUnsigned = 0
	// Max16BitUnsigned - Maximum 16 bit value
	Max16BitUnsigned = 65535

	// Min8BitSigned - Minimum 8 bit value
	Min8BitSigned = -128
	// Max8BitSigned - Maximum 8 bit value
	Max8BitSigned = 127

	// Min8BitUnsigned - Minimum 8 bit value
	Min8BitUnsigned = 0
	// Max8BitUnsigned - Maximum 8 bit value
	Max8BitUnsigned = 255
	solitonIDSpaceManifold = 1000

	solitonIDBitSizeMax = 64
)

type HoloKey struct {
	SaveHoloKey string
	LoadHoloKey string

}

)



//roaring bitmap to bitmap set
type BitmapSet struct {
	bitmap *smat.Bitmap
}


// NewBitmapSet returns a new BitmapSet.

func _() error {
	return nil
}





type EncodedBinary struct {
	Encoding string `json:"encoding"`
	Data     string `json:"data"`
	MaxValue uint64 `json:"max_value"`
	MinValue uint64 `json:"min_value"`
	runHoffmann bool `json:"run_hoffmann"`

}





func (v violetaBftConsensus) String() string {
	return fmt.Sprintf("%s:%d", v.RPCServer, v.RPCPort)
}


func (v violetaBftConsensus) GetRPCServer() string {
	return v.RPCServer
}


func (v violetaBftConsensus) GetRPCPort() uint64 {
	return v.RPCPort

}


func (v violetaBftConsensus) GetRPCUser() string {
	return v.RPCUser
}
type KindBuffer uint32 // 0: normal, 1: compressed, 2: compressed and encrypted

const (
	// KindNormal - normal kind
	KindNormal = KindBuffer(0)
	// KindCompressed - compressed kind
	KindCompressed = KindBuffer(1)
	// KindCompressedAndEncrypted - compressed and encrypted kind
	KindCompressedAndEncrypted = KindBuffer(2)
	//KindMergeAppend
	KindMergeAppend = KindBuffer(3)
	//KindSchemaReplicant
	KindSchemaReplicant = KindBuffer(4)
	KindRegionReplicant = KindBuffer(5)

	KindSchemaReplicantCompressed = KindBuffer(6)
	KindRegionReplicantCompressed = KindBuffer(7)
	KindSchemaReplicantCompressedAndEncrypted = KindBuffer(8)



)


var labelForKind = map[KindBuffer]string{
	KindNormal:         "normal",
	KindCompressed:     "compressed",
	KindCompressedAndEncrypted: "compressed and encrypted",
	KindMergeAppend: "merge append",
	KindSchemaReplicant: "schema replicant",
	KindRegionReplicant: "region replicant",

}






func (k KindBuffer) String() string {
	if label, ok := labelForKind[k]; ok {
		return label
	}
	return "unknown"
}


func (k KindBuffer) MarshalJSON() ([]byte, error) {
	return json.Marshal(k.String())
}



func (k KindBuffer) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	for k, label := range labelForKind {
		if label == s {
			*k = k
			return nil
		}
	}
	return errors.New("unknown kind")
}

/*

type bitmapContainer struct {
	cardinality int
	bitmap      []uint64
}

func (bc bitmapContainer) String() string {
	var s string
	for it := bc.getShortIterator(); it.hasNext(); {
		s += fmt.Sprintf("%v, ", it.next())
	}
	return s
}

func newBitmapContainer() *bitmapContainer {
	p := new(bitmapContainer)
	size := (1 << 16) / 64
	p.bitmap = make([]uint64, size, size)
	return p
}

func newBitmapContainerwithRange(firstOfRun, lastOfRun int) *bitmapContainer {
	bc := newBitmapContainer()
	bc.cardinality = lastOfRun - firstOfRun + 1
	if bc.cardinality == maxCapacity {
		fill(bc.bitmap, uint64(0xffffffffffffffff))
	} else {
		firstWord := firstOfRun / 64
		lastWord := lastOfRun / 64
		zeroPrefixLength := uint64(firstOfRun & 63)
		zeroSuffixLength := uint64(63 - (lastOfRun & 63))

		fillRange(bc.bitmap, firstWord, lastWord+1, uint64(0xffffffffffffffff))
		bc.bitmap[firstWord] ^= ((uint64(1) << zeroPrefixLength) - 1)
		blockOfOnes := (uint64(1) << zeroSuffixLength) - 1
		maskOnLeft := blockOfOnes << (uint64(64) - zeroSuffixLength)
		bc.bitmap[lastWord] ^= maskOnLeft
	}
	return bc
}
 */

type bitmapContainer struct {
	cardinality int
	bitmap      []uint64
	//bitmap      *smat.Bitmap
}




func (bc bitmapContainer) String() string {
	var s string
	for it := bc.getShortIterator(); it.hasNext(); {
		s += fmt.Sprintf("%v, ", it.next())
	}
	return s
}

func (bc bitmapContainer) getShortIterator() interface{} {
	return bc.bitmap.Iterator()
}


type Kind uint32 // 0: normal, 1: compressed, 2: compressed and encrypted

type isolatedContainer struct {
	kind Kind
	data []byte
	//compressedData []byte
	compressedData []byte
	//encryptedData []byte
	encryptedData []byte
	//mergeAppendData []byte
	mergeAppendData []byte
	//schemaReplicantData []byte
	schemaReplicantData []byte
	//regionReplicantData []byte
	regionReplicantData []byte
	//schemaReplicantCompressedData []byte
	schemaReplicantCompressedData []byte
	//regionReplicantCompressedData []byte
regionReplicantCompressedData []byte
//uncompressed suffix data
uncompressedSuffixData []byte

}

type causetWithIsolatedContainer struct {
causet
	isolatedContainer *isolatedContainer


}


var _ causet.Causet = (*causetWithIsolatedContainer)(nil)

func (c *causetWithIsolatedContainer) GetKind() Kind {
	return c.isolatedContainer.kind


}


func (c *causetWithIsolatedContainer) GetData() []byte {
	return c.isolatedContainer.data
}



func (c *causetWithIsolatedContainer) GetCompressedData() []byte {
	return c.isolatedContainer.compressedData
}



type causet struct {
	content []uint32
	//isolatedContainer *smat.Bitmap
	isolatedContainer *BitmapSet
	violetaBftConsensus *violetaBftConsensus
	kind Kind
	data []byte
	//compressedData []byte
	compressedData []byte
	//encryptedData []byte
	encryptedData []byte
	//mergeAppendData []byte
	mergeAppendData []byte


}

func newCausetWithIsolatedContainer(violetaBftConsensus *violetaBftConsensus) *causetWithIsolatedContainer {
	return &causetWithIsolatedContainer{
		causet: newCauset(violetaBftConsensus),
		isolatedContainer: newIsolatedContainer(violetaBftConsensus),

	}

}


func newCauset(violetaBftConsensus *violetaBftConsensus) *causet {

	return &causet{
		content: make([]uint32, 0),
		isolatedContainer: newIsolatedContainer(violetaBftConsensus),
		violetaBftConsensus: violetaBftConsensus,
		kind: KindNormal,
	}


	//return &causet{
	//	content: make([]uint32, 0),
	//	isolatedContainer: newIsolatedContainer(violetaBftConsensus),
	//	violetaBftConsensus: violetaBftConsensus,
	//	kind: KindNormal,
	//}



}


func (c *causet) String() string {
	return fmt.Sprintf("%s", c.content)

}


func (c *causet) GetKind() Kind {
	return c.kind
}


func (c *causet) SetKind(kind Kind) {
	c.kind = kind
}





func (e JSONError) Error() string {
	return e.Err.Error()
}

type error interface {
	Error() string
}

type JSONError struct {
	Err error
}

func tagJSONError(err error) error {
	if err == nil {
		return nil
	}

	return JSONError{Err: err}
}


func (c *causet) GetIsolatedContainer() *isolatedContainer {
	switch err.(type) {
	case *json.SyntaxError:
		return JSONError{err}
	}
	return err
}

func _() error {
	return nil
}

func DeferClose(c io.Closer, err *error) {
	if err != nil && *err == nil {
		*err = c.Close()
	}

	if err != nil && *err != nil {
		log.Error("failed to close", zap.Error(*err))

	}
}

func ReadJSON(r io.ReadCloser, data interface{}) error {
	var err error
	defer DeferClose(r, &err)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.WithStack(err)
	}

	err = json.Unmarshal(b, data)
	if err != nil {
		return tagJSONError(err)
	}

	return err
}

// FieldError connects an error to a particular field
type FieldError struct {
	error
	field string
}

func ParseUint64VarsField(vars map[string]string, varName string) (uint64, *FieldError) {
	str, ok := vars[varName]
	if !ok {
		return 0, &FieldError{field: varName, error: fmt.Errorf("field %s not present", varName)}
	}
	parsed, err := strconv.ParseUint(str, 10, 64)
	if err == nil {
		return parsed, nil
	}
	return parsed, &FieldError{field: varName, error: err}
}

// ReadJSONRespondError writes json into data.
// On error respond with a 400 Bad Request
func ReadJSONRespondError(rd *render.Render, w http.ResponseWriter, body io.ReadCloser, data interface{}) error {
	err := ReadJSON(body, data)
	if err == nil {
		return nil
	}
	ErrorResp(rd, w, err)
	return err
}

func _(rd *render.Render, w http.ResponseWriter, body io.ReadCloser, data interface{}) error {
	err := ReadJSON(body, data)
	if err == nil {
		return nil
	}
	var errCode errcode.ErrorCode
	if jsonErr, ok := errors.Cause(err).(JSONError); ok {
		errCode = errcode.NewInvalidInputErr(jsonErr.Err)
	} else {
		errCode = errcode.NewInternalErr(err)
	}
	ErrorResp(rd, w, errCode)
	return err
}

// ErrorResp Respond to the client about the given error, integrating with errcode.ErrorCode.
//
// Important: if the `err` is just an error and not an errcode.ErrorCode (given by errors.Cause),
// then by default an error is assumed to be a 500 Internal Error.
//
// If the error is nil, this also responds with a 500 and logs at the error level.
func ErrorResp(rd *render.Render, w http.ResponseWriter, err error) {
	if err == nil {
		log.Error("nil is given to errorResp")
		rd.JSON(w, http.StatusInternalServerError, "nil error")
		return
	}
	if errCode := errcode.CodeChain(err); errCode != nil {
		w.Header().Set("milevadb-Error-Code", errCode.Code().CodeStr().String())
		rd.JSON(w, errCode.Code().HTTscaode(), errcode.NewJSONFormat(errCode))
	} else {
		rd.JSON(w, http.StatusInternalServerError, err.Error())
	}
}

func _(rd *render.Render, w http.ResponseWriter, body io.ReadCloser, data interface{}) error {
	err := ReadJSON(body, data)
	if err == nil {
		return nil
	}
	var errCode errcode.ErrorCode
	if jsonErr, ok := errors.Cause(err).(JSONError); ok {
		errCode = errcode.NewInvalidInputErr(jsonErr.Err)
	} else {
		errCode = errcode.NewInternalErr(err)
	}
	ErrorResp(rd, w, errCode)
	return err
}




func (c *causet) GetContent() []uint32 {
	return c.content
}


func (c *causet) SetContent(content []uint32) {
	c.content = content

}

