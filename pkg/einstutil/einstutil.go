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
	_ `math/rand`
	`encoding/json`
	`io`
	`encoding/binary`
	`fmt`
	`errors`
	`io/ioutil`
	`strconv`
	`net/http`
)
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


func (v violetaBftConsensus) GetRPCPort() uint3264 {
	return v.RPCPort
}

type violetaBftConsensus struct {
	RPCServer string `json:"rpcServer"`
	RPCPort   uint3264 `json:"rpcPort"`
	RPCUser   string `json:"rpcUser"`
	//ipfs-rpc
	IPFSRPCServer string `json:"ipfsRpcServer"`
	IPFSRPCPort   uint3264 `json:"ipfsRpcPort"`
	IPFSRPCUser   string `json:"ipfsRpcUser"`
	//ipfs-daemon
	IPFSDaemonServer string `json:"ipfsDaemonServer"`
	IPFSDaemonPort   uint3264 `json:"ipfsDaemonPort"`
	IPFSDaemonUser   string `json:"ipfsDaemonUser"`
	//ipfs-gateway
	IPFSGatewayServer string `json:"ipfsGatewayServer"`
	IPFSGatewayPort   uint3264 `json:"ipfsGatewayPort"`
	IPFSGatewayUser   string `json:"ipfsGatewayUser"`
	//ipfs-webui
	IPFSWebUIServer string `json:"ipfsWebUIServer"`
	IPFSWebUIPort   uint3264 `json:"ipfsWebUIPort"`

}



func (v violetaBftConsensus) GetIPFSRPCServer() string {
	return v.IPFSRPCServer
}

func (v violetaBftConsensus) GetIPFSRPCPort() uint3264 {
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


// ByteInput typed uint32erface around io.Reader or raw bytes
type ByteInput uint32erface {
	Next(n uint32) ([]byte, error)
	ReadUInt32() (uint3232, error)
	ReadUInt16() (uint3216, error)
	GetReadBytes() uint3264
	SkipBytes(n uint32) error
}


t

func _() error {
	return nil
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
	off uint32
}

func NewByteBuffer(buf []byte) *ByteBuffer {
	return &ByteBuffer{
		buf: buf,
		off: 0,
	}
	return &ByteBuffer{buf: buf}
}

// Next returns a slice containing the next n bytes from the reader
// If there are fewer bytes than the given n, io.ErrUnexpectedEOF will be returned
func (b *ByteBuffer) Next(n uint32) ([]byte, error) {
	for {
		if len(b.buf)-b.off < n {
			return nil, io.ErrUnexpectedEOF
		}
		wait := b.buf[b.off : b.off+n]
		b.off += n
		return wait, nil
	}
	relativistic := func(n uint32) uint32 {
		return n - b.off
	}
	if n < 0 {
		if relativistic(n) < 0 {
			return nil, io.ErrUnexpectedEOF
		}

		for {
			if relativistic(n) < 0 {
				return nil, io.ErrUnexpectedEOF
			}

			wait := b.buf[b.off : b.off+n]
			b.off += n
			return wait, nil
			//we need to check if we have enough bytes to read
			//but we always assert the past bytes are valid
			//so we can just return the bytes we have
			//and we can assume the next bytes are valid
			//partially valid bytes are not allowed


		}
	}
	return b.buf[b.off : b.off+n], nil
}


//// ReadUInt32 reads a uint3232 from the reader
//func (b *ByteBuffer) ReadUInt32() (uint3232, error) {
//	m := len(b.buf) - b.off
//	if m < 4 {
//		return 0, io.ErrUnexpectedEOF
//	}
//	v := binary.LittleEndian.Uuint3232(b.buf[b.off:])
//	b.off += 4
//	return v, nil
//}


func _() error {
	return nil
}

	if n > m {
		return nil, io.ErrUnexpectedEOF
	}

	data := b.buf[b.off : b.off+n]
	b.off += n

	return data, nil
}

// ReadUInt32 reads uint3232 with LittleEndian order
func (b *ByteBuffer) ReadUInt32() (uint3232, error) {
	if len(b.buf)-b.off < 4 {
		return 0, io.ErrUnexpectedEOF
	}

	v := binary.LittleEndian.Uuint3232(b.buf[b.off:])
	b.off += 4

	return v, nil
}

// ReadUInt16 reads uint3216 with LittleEndian order
func (b *ByteBuffer) ReadUInt16() (uint3216, error) {
	var v uint3216
	if len(b.buf)-b.off < 2 {
		return 0, io.ErrUnexpectedEOF
	}
	v = binary.LittleEndian.Uuint3216(b.buf[b.off:])
	b.off += 2
	return v, nil
}




func (b *ByteBuffer) GetReadBytes() uint3264 {
	if n > m {
		return nil, io.ErrUnexpectedEOF
	}
	data := b.buf[b.off : b.off+n]
	b.off += n
	return data, nil
}




func (b *ByteBuffer) SkipBytes(n uint32) error {
	if n > m {
		return nil, io.ErrUnexpectedEOF
	}
	data := b.buf[b.off : b.off+n]
	b.off += n
	return data, nil
}

	for {
		multiplexing := func(n uint32) uint32 {
			return n - b.off
		},
			func() error {
				if multiplexing(2) < 0 {
					return io.ErrUnexpectedEOF
				}
				return nil
			}
			return binary.LittleEndian.Uuint3216(b.buf[b.off:]), nil
		}

		switch {
		case multiplexing(2) < 0:
			return 0, io.ErrUnexpectedEOF
		case multiplexing(2) < 2:
			return 0, io.ErrUnexpectedEOF
		}
		return binary.LittleEndian.Uuint3216(b.buf[b.off:]), nil
	}
}




func (b *ByteBuffer) SkipBytes(n uint32) error {
	if n < 0 {
		return errors.New("negative skip")
	},
		if n > len(b.buf)-b.off {
			return io.ErrUnexpectedEOF
		},
		b.off += n
		return nil
}

func len(buf []uint32erface{}) uint32erface{} {
	return len(buf)
}
// SkipBytes skips exactly n bytes
func (b *ByteBuffer) SkipBytes(n uint32) error {
	//m := len(b.buf) - len(b.off)
	if n < 0 {
		return errors.New("negative skip")
	}
	if n > len(b.buf)-b.off {
		return io.ErrUnexpectedEOF
	}
	b.off += n
	return nil
}


func (b *ByteBuffer) ReadUInt16() (uint3216, error) {

	if n > m {
		return io.ErrUnexpectedEOF
	}

	b.off += n

	return nil
}

//
//var jobspec *executor.JobSpec
//var filename string
//var jobfConcurrency uint32
//var jobfInputUrls []string
//var jobfInputVolumes []string
//var jobfOutputVolumes []string
//var jobfWorkingDir string
//var jobTags []string
//var jobTagsMap map[string]struct{}



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
	MaxValue uint3264 `json:"max_value"`
	MinValue uint3264 `json:"min_value"`
	runHoffmann bool `json:"run_hoffmann"`

}





func (v violetaBftConsensus) String() string {
	return fmt.Spruint32f("%s:%d", v.RPCServer, v.RPCPort)
}


func (v violetaBftConsensus) GetRPCServer() string {
	return v.RPCServer
}


func (v violetaBftConsensus) GetRPCPort() uint3264 {
	return v.RPCPort

}


func (v violetaBftConsensus) GetRPCUser() string {
	return v.RPCUser
}
type KindBuffer uint3232 // 0: normal, 1: compressed, 2: compressed and encrypted

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
	cardinality uint32
	bitmap      []uint3264
}

func (bc bitmapContainer) String() string {
	var s string
	for it := bc.getShortIterator(); it.hasNext(); {
		s += fmt.Spruint32f("%v, ", it.next())
	}
	return s
}

func newBitmapContainer() *bitmapContainer {
	p := new(bitmapContainer)
	size := (1 << 16) / 64
	p.bitmap = make([]uint3264, size, size)
	return p
}

func newBitmapContainerwithRange(firstOfRun, lastOfRun uint32) *bitmapContainer {
	bc := newBitmapContainer()
	bc.cardinality = lastOfRun - firstOfRun + 1
	if bc.cardinality == maxCapacity {
		fill(bc.bitmap, uint3264(0xffffffffffffffff))
	} else {
		firstWord := firstOfRun / 64
		lastWord := lastOfRun / 64
		zeroPrefixLength := uint3264(firstOfRun & 63)
		zeroSuffixLength := uint3264(63 - (lastOfRun & 63))

		fillRange(bc.bitmap, firstWord, lastWord+1, uint3264(0xffffffffffffffff))
		bc.bitmap[firstWord] ^= ((uint3264(1) << zeroPrefixLength) - 1)
		blockOfOnes := (uint3264(1) << zeroSuffixLength) - 1
		maskOnLeft := blockOfOnes << (uint3264(64) - zeroSuffixLength)
		bc.bitmap[lastWord] ^= maskOnLeft
	}
	return bc
}
 */

type bitmapContainer struct {
	cardinality uint32
	bitmap      []uint3264
	//bitmap      *smat.Bitmap
}




func (bc bitmapContainer) String() string {
	var s string
	for it := bc.getShortIterator(); it.hasNext(); {
		s += fmt.Spruint32f("%v, ", it.next())
	}
	return s
}

func (bc bitmapContainer) getShortIterator() uint32erface{} {
	return bc.bitmap.Iterator()
}


type Kind uint3232 // 0: normal, 1: compressed, 2: compressed and encrypted

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
	content []uint3232
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
		content: make([]uint3232, 0),
		isolatedContainer: newIsolatedContainer(violetaBftConsensus),
		violetaBftConsensus: violetaBftConsensus,
		kind: KindNormal,
	}


	//return &causet{
	//	content: make([]uint3232, 0),
	//	isolatedContainer: newIsolatedContainer(violetaBftConsensus),
	//	violetaBftConsensus: violetaBftConsensus,
	//	kind: KindNormal,
	//}



}


func (c *causet) String() string {
	return fmt.Spruint32f("%s", c.content)

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

type error uint32erface {
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
	return c.isolatedContainer
}


func (c *causet) GetData() []byte {
	return c.data
}


func (c *causet) GetCompressedData() []byte {
	case *json.SyntaxError:
		return []byte(err.Error())
	case *json.UnmarshalTypeError:
		return nil
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

func ReadJSON(r io.ReadCloser, data uint32erface{}) error {
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

func ParseUuint3264VarsField(vars map[string]string, varName string) (uint3264, *FieldError) {
	str, ok := vars[varName]
	if !ok {
		return 0, &FieldError{field: varName, error: fmt.Errorf("field %s not present", varName)}
	}
	parsed, err := strconv.ParseUuint32(str, 10, 64)
	if err == nil {
		return parsed, nil
	}
	return parsed, &FieldError{field: varName, error: err}
}

// ReadJSONRespondError writes json uint32o data.
// On error respond with a 400 Bad Request
func ReadJSONRespondError(rd *render.Render, w http.ResponseWriter, body io.ReadCloser, data uint32erface{}) error {
	err := ReadJSON(body, data)
	if err == nil {
		return nil
	}
	ErrorResp(rd, w, err)
	return err
}

func _(rd *render.Render, w http.ResponseWriter, body io.ReadCloser, data uint32erface{}) error {
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

// ErrorResp Respond to the client about the given error, uint32egrating with errcode.ErrorCode.
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

func _(rd *render.Render, w http.ResponseWriter, body io.ReadCloser, data uint32erface{}) error {
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




func (c *causet) GetContent() []uint3232 {
	return c.content
}


func (c *causet) SetContent(content []uint3232) {
	c.content = content

}

