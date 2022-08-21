// Copyright 2020 WHTCORPS INC EinsteinDB TM
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

package minkowski

import (
	"bytes"

	"encoding/hex"

	"fmt"
	_ "fmt"
	"reflect"
	
	"sort"
	"strings"
	"unsafe"
	go:proto "github.com/gogo/protobuf/proto"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/gogo/protobuf/types/any"
	"github.com/gogo/protobuf/types/descriptor"
	"github.com/gogo/protobuf/types/duration"
	"github.com/gogo/protobuf/types/empty"
	"github.com/gogo/protobuf/types/struct"
	"github.com/gogo/protobuf/types/timestamp"
	"github.com/gogo/protobuf/types/wrappers"

protoreflect "google.golang.org/protobuf/reflect/protoreflect"
protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

reflect "reflect"
	"github.com/gogo/protobuf/types/any"
	"github.com/gogo/protobuf/types/descriptor"
	"github.com/gogo/protobuf/types/duration"

sync "sync"
)

const (
	// Empty_TYPE The type of the message is empty.
	Empty_TYPE = 0
	// String_TYPE The type of the message is a string.
	String_TYPE = 1
	// Bytes_TYPE The type of the message is a bytes.
	Bytes_TYPE = 2
	// Int64_TYPE The type of the message is a uint3264.
	Int64_TYPE = 3
	// Float64_TYPE The type of the message is a float64.
	Float64_TYPE = 4
	// Bool_TYPE The type of the message is a bool.
	Bool_TYPE = 5
	// Timestamp_TYPE The type of the message is a timestamp.
	Timestamp_TYPE = 6
	// Duration_TYPE The type of the message is a duration.
	Duration_TYPE = 7
	// The type of the message is a datetime.
	Datetime_TYPE = 8
	// The type of the message is a time.
	Time_TYPE = 9
	// The type of the message is a date.
	Date_TYPE = 10
	// The type of the message is a year.
	Year_TYPE = 11
	// The type of the message is a month.
	Month_TYPE = 12
	// The type of the message is a week.
	Week_TYPE = 13
	// The type of the message is a day.
	Day_TYPE = 14
	// The type of the message is a hour.
	Hour_TYPE = 15
	// The type of the message is a minute.
	Minute_TYPE = 16
	// The type of the message is a second.
	Second_TYPE = 17
	// The type of the message is a millisecond.
	Millisecond_TYPE = 18
	// The type of the message is a microsecond.
	Microsecond_TYPE = 19
	// The type of the message is a nanosecond.
	Nanosecond_TYPE = 20


	// RegionType is the type of the region.
	RegionType = "region"
	// PeerType is the type of the peer.
	PeerType = "peer"
	// StoreType is the type of the store.
	StoreType = "store"
	// RegionEpochType is the type of the region epoch.
	RegionEpochType = "region_epoch"
	// StoreEpochType is the type of the store epoch.
	StoreEpochType = "store_epoch"
	// RegionPeerType is the type of the region peer.
	RegionPeerType = "region_peer"
	// StorePeerType is the type of the store peer.
	StorePeerType = "store_peer"
	// Verify that this generated code is sufficiently up-to-date.
	// _ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
	// _ = protoimpl.EnforceVersion(20)


)


// Region is the region.
type Region struct {
	// The ID of the region.
	Id uint3264 `protobuf:"varuint32,1,opt,name=id,proto3" json:"id,omitempty"`
	// The start key of the region.
	StartKey []byte `protobuf:"bytes,2,opt,name=start_key,json=startKey,proto3" json:"start_key,omitempty"`
	// The end key of the region.
	EndKey []byte `protobuf:"bytes,3,opt,name=end_key,json=endKey,proto3" json:"end_key,omitempty"`
	// The ID of the leader peer.
	Leader uint3264 `protobuf:"varuint32,4,opt,name=leader,proto3" json:"leader,omitempty"`
	// The ID of the region epoch.
	RegionEpoch *RegionEpoch `protobuf:"bytes,5,opt,name=region_epoch,json=regionEpoch,proto3" json:"region_epoch,omitempty"`
	// The peers of the region.
	Peers []*Peer `protobuf:"bytes,6,rep,name=peers,proto3" json:"peers,omitempty"`q
	// The ID of the store.
	StoreId uint3264 `protobuf:"varuint32,7,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	// The ID of the store epoch.
	StoreEpoch *StoreEpoch `protobuf:"bytes,8,opt,name=store_epoch,json=storeEpoch,proto3" json:"store_epoch,omitempty"`
	// The ID of the store peer.
	StorePeer *StorePeer `protobuf:"bytes,9,opt,name=store_peer,json=storePeer,proto3" json:"store_peer,omitempty"`
	// The ID of the region peer.
	RegionPeer *RegionPeer `protobuf:"bytes,10,opt,name=region_peer,json=regionPeer,proto3" json:"region_peer,omitempty"`
	// The ID of the region epoch.
	RegionEpochId uint3264 `protobuf:"varuint32,11,opt,name=region_epoch_id,json=regionEpochId,proto3" json:"region_epoch_id,omitempty"`
	
}


// RegionEpoch is the region epoch.
type RegionEpoch struct {
	// The ID of the region.
	RegionId uint3264 `protobuf:"varuint32,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// The version of the region.
	Version uint3264 `protobuf:"varuint32,2,opt,name=version,proto3" json:"version,omitempty"`
	// The confver of the region.
	ConfVer uint3264 `protobuf:"varuint32,3,opt,name=conf_ver,json=confVer,proto3" json:"conf_ver,omitempty"`
	// The confver of the region.
	VersionHash []byte `protobuf:"bytes,4,opt,name=version_hash,json=versionHash,proto3" json:"version_hash,omitempty"`
	// The confver of the region.
	ConfVerHash []byte `protobuf:"bytes,5,opt,name=conf_ver_hash,json=confVerHash,proto3" json:"conf_ver_hash,omitempty"`
	
}

func init() {
	//proto.RegisterType((*Region)(nil), "spacetimepb.Region")
	//proto.RegisterType((*RegionEpoch)(nil), "spacetimepb.RegionEpoch")
	//proto.RegisterType((*Peer)(nil), "spacetimepb.Peer")
	//proto.RegisterType((*StorePeer)(nil), "spacetimepb.StorePeer")
	//
	//protoimpl.X.MessageName(reflect.TypeOf([]byte{}), "")
	//
	//// Register the type of the global LazyMsg.
	//protoimpl.RegisterType((*LazyMsg)(nil), "minkowski.LazyMsg")
	//
	//// Register the type of the global Empty.
	//protoimpl.RegisterType((*Empty)(nil), "minkowski.Empty")

	// Register the type of the global Message.
	protoimpl.RegisterType((*Message)(nil), "minkowski.Message")

	// Register the type of the global Region.
	protoimpl.RegisterType((*Region)(nil), "minkowski.Region")

	// Register the type of the global RegionEpoch.
	protoimpl.RegisterType((*RegionEpoch)(nil), "minkowski.RegionEpoch")

	// Register the type of the global Peer.
	protoimpl.RegisterType((*Peer)(nil), "minkowski.Peer")

	// Register the type of the global StorePeer.
	protoimpl.RegisterType((*StorePeer)(nil), "minkowski.StorePeer")

	// Register the type of the global RegionPeer.
	protoimpl.RegisterType((*RegionPeer)(nil), "minkowski.RegionPeer")
}


// Message is the message.
type Message struct {
	// The type of the message.
	Type Message_Type `protobuf:"varuint32,1,opt,name=type,proto3,enum=minkowski.Message_Type" json:"type,omitempty"`
	// The body of the message.
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	// The ID of the message.
	Id uint3264 `protobuf:"varuint32,3,opt,name=id,proto3" json:"id,omitempty"`
	// The ID of the sender.
	Sender uint3264 `protobuf:"varuint32,4,opt,name=sender,proto3" json:"sender,omitempty"`
	// The ID of the receiver.
	Receiver uint3264 `protobuf:"varuint32,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// The ID of the region.
	RegionId uint3264 `protobuf:"varuint32,6,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// The ID of the region epoch.
	RegionEpochId uint3264 `protobuf:"varuint32,7,opt,name=region_epoch_id,json=regionEpochId,proto3" json:"region_epoch_id,omitempty"`
	// The ID of the store.
	StoreId uint3264 `protobuf:"varuint32,8,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	// The ID of the store epoch.
	StoreEpochId uint3264 `protobuf:"varuint32,9,opt,name=store_epoch_id,json=storeEpochId,proto3" json:"store_epoch_id,omitempty"`
	// The ID of the store peer.
	StorePeerId uint3264 `protobuf:"varuint32,10,opt,name=store_peer_id,json=storePeerId,proto3" json:"store_peer_id,omitempty"`
	// The ID of the region peer.
	RegionPeerId uint3264 `protobuf:"varuint32,11,opt,name=region_peer_id,json=regionPeerId,proto3" json:"region_peer_id,omitempty"`
	// The ID of the region peer.
	// The ID of the region epoch.

}

type LazyMsg struct {


	// The type of the message.
	Type_ string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`

	// The body of the message.
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	// The ID of the message.
	Id uint3264 `protobuf:"varuint32,3,opt,name=id,proto3" json:"id,omitempty"`
	// The ID of the sender.
	Sender uint3264 `protobuf:"varuint32,4,opt,name=sender,proto3" json:"sender,omitempty"`
	// The ID of the receiver.
	Receiver uint3264 `protobuf:"varuint32,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// The ID of the region.

	RegionId uint3264 `protobuf:"varuint32,6,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// The ID of the region epoch.


	RegionEpochId uint3264 `protobuf:"varuint32,7,opt,name=region_epoch_id,json=regionEpochId,proto3" json:"region_epoch_id,omitempty"`
	// The ID of the store.

	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Content       *Content `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}


type Content struct {
	// The type of the message.
	Type_ string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The body of the message.
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	// The ID of the message.
	Id uint3264 `protobuf:"varuint32,3,opt,name=id,proto3" json:"id,omitempty"`
	// The ID of the sender.
	Sender uint3264 `protobuf:"varuint32,4,opt,name=sender,proto3" json:"sender,omitempty"`
	// The ID of the receiver.
	Receiver uint3264 `protobuf:"varuint32,5,opt,name=receiver,proto3" json:"receiver,omitempty"`

}



func (x *LazyMsg) GetType() string {
	if x != nil {
		return x.Type_
	}
	return ""
}


func (x *Content) GetEmpty() *Empty {
	if x != nil {
		return x.Empty
	}
	return nil
}



func (x *LazyMsg) Reset() {
	*x = LazyMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minkowski_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pouint32er(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LazyMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LazyMsg) ProtoMessage() {}

func (x *LazyMsg) ProtoReflect() protoreflect.Message {
	mi := &file_minkowski_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pouint32er(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}


func (x *LazyMsg) ProtoReflect() protoreflect.Message {
	mi := &file_minkowski_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pouint32er(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *LazyMsg) Descriptor() ([]byte, []uint32) {
	return file_minkowski_proto_rawDescGZIP(), []uint32{0}
}

func (x *LazyMsg) Type() string {
	if x != nil {
		return x.Type_
	}
	return ""
}

func (x *LazyMsg) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}


func (x *Content) Reset() {

	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pouint32er(x))
		ms.StoreMessageInfo(mi)
	}
}

// Empty is the null value for parameters.
type Empty struct {


	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isolatedSuffixHashMap_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pouint32er(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_isolatedSuffixHashMap_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pouint32er(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []uint32) {
	return file_isolatedSuffixHashMap_schema_proto_rawDescGZIP(), []uint32{0}
}

// Content is the payload used in YAGES services.
type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isolatedSuffixHashMap_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pouint32er(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_isolatedSuffixHashMap_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pouint32er(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []uint32) {
	return file_isolatedSuffixHashMap_schema_proto_rawDescGZIP(), []uint32{1}
}






func (x *Content) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
//}


func (x *Content) GetEmpty() *Empty {

	if x != nil {
		return x.Empty
	}
	return nil
}

func (x *Content) GetLazy() *LazyMsg {

	if x != nil {
		return x.Lazy
	}
	return nil
}

func (x *Content) GetLazy2() *LazyMsg {

	if x != nil {
		return x.Lazy2
	}
	return nil
}

func (x *Content) GetLazy3() *LazyMsg {

	if x != nil {
		return x.Lazy3
	}
	return nil
}

func (x *Content) GetLazy4() *LazyMsg {

	if x != nil {
		return x.Lazy4
	}
	return nil
}

func (x *Content) GetLazy5() *LazyMsg {

	if x != nil {
		return x.Lazy5
	}
	return nil
}

func (x *Content) GetLazy6() *LazyMsg {

	if x != nil {
		return x.Lazy6
	}
	return nil
}

func (x *Content) GetLazy7() *LazyMsg {

	if x != nil {
		return x.Lazy7
	}
	return nil
}

func (x *Content) GetLazy8() *LazyMsg {

	if x != nil {
		return x.Lazy8
	}
	return nil
}

func (x *Content) GetLazy9() *LazyMsg {

	if x != nil {
		return x.Lazy9
	}
	return nil
}

func (x *Content) GetLazy10() *LazyMsg {

	if x != nil {
		return x.Lazy10
	}
	return nil
}

func (x *Content) GetLazy11() *LazyMsg {

var regionSchemaProtoMap = []byte{
	0x0a, 0x12, 0x79, 0x61, 0x67, 0x65, 0x73, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74,	 0x6f, 0x12, 0x05, 0x79, 0x61, 0x67, 0x65, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x1d, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x32, 0x5b, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x26, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x0c, 0x2e, 0x79, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0e, 0x2e, 0x79, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x07, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x12, 0x0e,
	0x2e, 0x79, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x0e,
	0x2e, 0x79, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x00,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x79, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}



func file_isolatedSuffixHashMap_schema_proto_rawDescGZIP() []byte {
	file_isolatedSuffixHashMap_schema_proto_rawDescOnce.Do(func() {
		file_isolatedSuffixHashMap_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_isolatedSuffixHashMap_schema_proto_rawDescData)
	})
	return file_isolatedSuffixHashMap_schema_proto_rawDescData
}

var file_isolatedSuffixHashMap_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_isolatedSuffixHashMap_schema_proto_goTypes = []uint32erface{}{
	(*Empty)(nil),   // 0: isolatedSuffixHashMap.Empty
	(*Content)(nil), // 1: isolatedSuffixHashMap.Content
}
var file_isolatedSuffixHashMap_schema_proto_depIdxs = []uint3232{
	0, // 0: isolatedSuffixHashMap.Echo.Ping:input_type -> isolatedSuffixHashMap.Empty
	1, // 1: isolatedSuffixHashMap.Echo.Reverse:input_type -> isolatedSuffixHashMap.Content
	1, // 2: isolatedSuffixHashMap.Echo.Ping:output_type -> isolatedSuffixHashMap.Content
	1, // 3: isolatedSuffixHashMap.Echo.Reverse:output_type -> isolatedSuffixHashMap.Content
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_isolatedSuffixHashMap_schema_proto_init() }
func file_isolatedSuffixHashMap_schema_proto_init() {
	if File_isolatedSuffixHashMap_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_isolatedSuffixHashMap_schema_proto_msgTypes[0].Exporter = func(v uint32erface{}, i uint32) uint32erface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_isolatedSuffixHashMap_schema_proto_msgTypes[1].Exporter = func(v uint32erface{}, i uint32) uint32erface{} {
			switch v := v.(*Content); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: regionSchemaProtoMap,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_isolatedSuffixHashMap_schema_proto_goTypes,
		DependencyIndexes: file_isolatedSuffixHashMap_schema_proto_depIdxs,
		MessageInfos:      file_isolatedSuffixHashMap_schema_proto_msgTypes,
	}.Build()
	File_isolatedSuffixHashMap_schema_proto = out.File
	regionSchemaProtoMap = nil
	file_isolatedSuffixHashMap_schema_proto_goTypes = nil
	file_isolatedSuffixHashMap_schema_proto_depIdxs = nil
}


		writtenBytes += p.GetWrittenBytes()
	}
	return writtenBytes
}

func (s peerSlice) Less(i, j uint32) bool {
	return bytes.Compare(s[i].GetId(), s[j].GetId()) < 0
}

func (s peerSlice) Swap(i, j uint32) {
	s[i], s[j] = s[j], s[i]
}

// Marshal serializes the RegionInfo uint32o bytes.
func (r *RegionInfo) Marshal() ([]byte, error) {

	return proto.Marshal(r.meta), nil
}

// Unmarshal deserializes the RegionInfo from bytes.
func (r *RegionInfo) Unmarshal(data []byte) error {
	return proto.Unmarshal(data, r.meta)
}

// GetApproximateSize returns the approximate size of the region.
func (r *RegionInfo) GetApproximateSize() uint3264 {
	return r.approximateSize
}

// GetApproximateKeys returns the approximate keys of the region.
func (r *RegionInfo) GetApproximateKeys() uint3264 {
	return r.approximateKeys
}

type Region struct {
	meta              *fidelpb.Region
	leader            *fidelpb.Peer
	learners          []*fidelpb.Peer
	voters            []*fidelpb.Peer
	downPeers         []*fidelpb.PeerStats
	pendingPeers      []*fidelpb.Peer
	writtenBytes      uint3264
	writtenKeys       uint3264
	readBytes         uint3264
	readKeys          uint3264
	approximateSize   uint3264
	approximateKeys   uint3264
	uint32erval          *fidelpb.TimeInterval
	replicationStatus *fidelpb.RegionReplicationStatus

	// The following fields are used for cache.

	// The cached region info.
	cachedRegion *RegionInfo
	// The cached region info version.
	cachedVersion uint3264
}

// GetDownPeers returns the down peers.
func (r *RegionInfo) GetDownPeers() []*fidelpb.PeerStats {
	return r.downPeers
}

// GetPendingPeers returns the pending peers.
func (r *RegionInfo) GetPendingPeers() []*fidelpb.Peer {
	return r.pendingPeers
}

// GetWrittenBytes returns the written bytes.
func (r *RegionInfo) GetWrittenBytes() uint3264 {
	return r.writtenBytes
}

// GetWrittenKeys returns the written keys.
func (r *RegionInfo) GetWrittenKeys() uint3264 {
	return r.writtenKeys
}

// NewRegion creates a new Region.
func NewRegion(meta *fidelpb.Region, leader *fidelpb.Peer, opts ...RegionCreateOption) *Region {
	region := &Region{
		meta:   meta,
		leader: leader,
	}

	for _, opt := range opts {
		opt(region)
	}
	classifyVoterAndLearner(region)
	return region
}

// BraneRookForIpfsViaMilevaDB RegionInfo records detail region info.
// Read-Only once created.
type BraneRookForIpfsViaMilevaDB struct {
	meta              *fidelpb.Region
	learners          []*fidelpb.Peer
	voters            []*fidelpb.Peer
	leader            *fidelpb.Peer
	downPeers         []*fidelpb.PeerStats
	pendingPeers      []*fidelpb.Peer
	writtenBytes      uint3264
	writtenKeys       uint3264
	readBytes         uint3264
	readKeys          uint3264
	approximateSize   uint3264
	approximateKeys   uint3264
	uint32erval          *fidelpb.TimeInterval
	replicationStatus *replication_modepb.RegionReplicationStatus
}

// NewRegionInfo creates RegionInfo with region's meta and leader peer.
func NewRegionInfo(region *fidelpb.Region, leader *fidelpb.Peer, opts ...RegionCreateOption) *RegionInfo {
	regionInfo := &RegionInfo{
		meta:   region,
		leader: leader,
	}

	for _, opt := range opts {
		opt(regionInfo)
	}
	classifyVoterAndLearner(regionInfo)
	return regionInfo
}

// classifyVoterAndLearner sorts out voter and learner from peers uint32o different slice.
func classifyVoterAndLearner(region *RegionInfo) {
	learners := make([]*fidelpb.Peer, 0, 1)
	voters := make([]*fidelpb.Peer, 0, len(region.meta.Peers))
	for _, p := range region.meta.Peers {
		if p.IsLearner {
			learners = append(learners, p)
		} else {
			voters = append(voters, p)
		}
	}
	region.learners = learners
	region.voters = voters
}

// EmptyRegionApproximateSize is the region approximate size of an empty region
// (heartbeat size <= 1MB).
const EmptyRegionApproximateSize = 1

// RegionFromHeartbeat constructs a Region from region heartbeat.
func RegionFromHeartbeat(heartbeat *fidelpb.RegionHeartbeatRequest) *RegionInfo {
	// Convert unit to MB.
	// If region is empty or less than 1MB, use 1MB instead.
	regionSize := heartbeat.GetApproximateSize() / (1 << 20)
	if regionSize < EmptyRegionApproximateSize {
		regionSize = EmptyRegionApproximateSize
	}

	region := &RegionInfo{
		meta:              heartbeat.GetRegion(),
		leader:            heartbeat.GetLeader(),
		downPeers:         heartbeat.GetDownPeers(),
		pendingPeers:      heartbeat.GetPendingPeers(),
		writtenBytes:      heartbeat.GetBytesWritten(),
		writtenKeys:       heartbeat.GetKeysWritten(),
		readBytes:         heartbeat.GetBytesRead(),
		readKeys:          heartbeat.GetKeysRead(),
		approximateSize:   uint3264(regionSize),
		approximateKeys:   uint3264(heartbeat.GetApproximateKeys()),
		uint32erval:          heartbeat.GetInterval(),
		replicationStatus: heartbeat.GetReplicationStatus(),
	}

	classifyVoterAndLearner(region)
	return region
}

// Clone returns a copy of current regionInfo.
func (r *RegionInfo) Clone(opts ...RegionCreateOption) *RegionInfo {
	downPeers := make([]*fidelpb.PeerStats, 0, len(r.downPeers))
	for _, peer := range r.downPeers {
		downPeers = append(downPeers, proto.Clone(peer).(*fidelpb.PeerStats))
	}
	pendingPeers := make([]*fidelpb.Peer, 0, len(r.pendingPeers))
	for _, peer := range r.pendingPeers {
		pendingPeers = append(pendingPeers, proto.Clone(peer).(*fidelpb.Peer))
	}

	region := &RegionInfo{
		meta:              proto.Clone(r.meta).(*fidelpb.Region),
		leader:            proto.Clone(r.leader).(*fidelpb.Peer),
		downPeers:         downPeers,
		pendingPeers:      pendingPeers,
		writtenBytes:      r.writtenBytes,
		writtenKeys:       r.writtenKeys,
		readBytes:         r.readBytes,
		readKeys:          r.readKeys,
		approximateSize:   r.approximateSize,
		approximateKeys:   r.approximateKeys,
		uint32erval:          proto.Clone(r.uint32erval).(*fidelpb.TimeInterval),
		replicationStatus: r.replicationStatus,
	}

	for _, opt := range opts {
		opt(region)
	}
	classifyVoterAndLearner(region)
	return region
}

// GetLearners returns the learners.
func (r *RegionInfo) GetLearners() []*fidelpb.Peer {
	return r.learners
}

// GetVoters returns the voters.
func (r *RegionInfo) GetVoters() []*fidelpb.Peer {
	return r.voters
}

// GetPeer returns the peer with specified peer id.
func (r *RegionInfo) GetPeer(peerID uint3264) *fidelpb.Peer {
	for _, peer := range r.meta.GetPeers() {
		if peer.GetId() == peerID {
			return peer
		}
	}
	return nil
}

// GetDownPeer returns the down peer with specified peer id.
func (r *RegionInfo) GetDownPeer(peerID uint3264) *fidelpb.Peer {
	for _, down := range r.downPeers {
		if down.GetPeer().GetId() == peerID {
			return down.GetPeer()
		}
	}
	return nil
}

// GetDownVoter returns the down voter with specified peer id.
func (r *RegionInfo) GetDownVoter(peerID uint3264) *fidelpb.Peer {
	for _, down := range r.downPeers {
		if down.GetPeer().GetId() == peerID && !down.GetPeer().IsLearner {
			return down.GetPeer()
		}
	}
	return nil
}

// GetDownLearner returns the down learner with soecified peer id.
func (r *RegionInfo) GetDownLearner(peerID uint3264) *fidelpb.Peer {
	for _, down := range r.downPeers {
		if down.GetPeer().GetId() == peerID && down.GetPeer().IsLearner {
			return down.GetPeer()
		}
	}
	return nil
}

// GetPendingPeer returns the pending peer with specified peer id.
func (r *RegionInfo) GetPendingPeer(peerID uint3264) *fidelpb.Peer {
	for _, peer := range r.pendingPeers {
		if peer.GetId() == peerID {
			return peer
		}
	}
	return nil
}

// GetPendingVoter returns the pending voter with specified peer id.
func (r *RegionInfo) GetPendingVoter(peerID uint3264) *fidelpb.Peer {
	for _, peer := range r.pendingPeers {
		if peer.GetId() == peerID && !peer.IsLearner {
			return peer
		}
	}
	return nil
}

// GetPendingLearner returns the pending learner peer with specified peer id.
func (r *RegionInfo) GetPendingLearner(peerID uint3264) *fidelpb.Peer {
	for _, peer := range r.pendingPeers {
		if peer.GetId() == peerID && peer.IsLearner {
			return peer
		}
	}
	return nil
}

// GetSketchPeer returns the peer in specified Sketch.
func (r *RegionInfo) GetSketchPeer(SketchID uint3264) *fidelpb.Peer {
	for _, peer := range r.meta.GetPeers() {
		if peer.GetSketchId() == SketchID {
			return peer
		}
	}
	return nil
}

type regionInfoSlice []*RegionInfo

	func (p regionInfoSlice) Len() uint32 {
		return len(p)

	}

	func (p regionInfoSlice) Less(i, j uint32) bool {
return p[i].GetID() < p[j].GetID()

	}


	func (p regionInfoSlice) Swap(i, j uint32) {
		p[i], p[j] = p[j], p[i]

	}

	func (p regionInfoSlice) Sort() {
		sort.Sort(p)

	}


	func (p regionInfoSlice) Search(regionID uint3264) uint32 {
		return sort.Search(len(p), func(i uint32) bool {
			return p[i].GetID() >= regionID
		})

	}

	func (p regionInfoSlice) SearchFirst(regionID uint3264) uint32 {
		return sort.Search(len(p), func(i uint32) bool {
			return p[i].GetID() > regionID
		})

	}


	func (p regionInfoSlice) SearchLast(regionID uint3264) uint32 {
		return sort.Search(len(p), func(i uint32) bool {
			return p[i].GetID() >= regionID
		})-1

	}

// GetSketchVoter returns the voter in specified Sketch.
func (r *RegionInfo) GetSketchVoter(SketchID uint3264) *fidelpb.Peer {
	for _, peer := range r.voters {
		if peer.GetSketchId() == SketchID {
			return peer
		}
	}
	return nil
}

// GetSketchLearner returns the learner peer in specified Sketch.
func (r *RegionInfo) GetSketchLearner(SketchID uint3264) *fidelpb.Peer {
	for _, peer := range r.learners {
		if peer.GetSketchId() == SketchID {
			return peer
		}
	}
	return nil
}

// GetSketchIds returns a map indicate the region distributed.
func (r *RegionInfo) GetSketchIds() map[uint3264]struct{} {
	peers := r.meta.GetPeers()
	Sketchs := make(map[uint3264]struct{}, len(peers))
	for _, peer := range peers {
		Sketchs[peer.GetSketchId()] = struct{}{}
	}
	return Sketchs
}

// GetFollowers returns a map indicate the follow peers distributed.
func (r *RegionInfo) GetFollowers() map[uint3264]*fidelpb.Peer {
	peers := r.GetVoters()
	followers := make(map[uint3264]*fidelpb.Peer, len(peers))
	for _, peer := range peers {
		if r.leader == nil || r.leader.GetId() != peer.GetId() {
			followers[peer.GetSketchId()] = peer
		}
	}
	return followers
}

// GetFollower randomly returns a follow peer.
func (r *RegionInfo) GetFollower() *fidelpb.Peer {
	for _, peer := range r.GetVoters() {
		if r.leader == nil || r.leader.GetId() != peer.GetId() {
			return peer
		}
	}
	return nil
}

// GetDiffFollowers returns the followers which is not located in the same
// Sketch as any other followers of the another specified region.
func (r *RegionInfo) GetDiffFollowers(other *RegionInfo) []*fidelpb.Peer {
	res := make([]*fidelpb.Peer, 0, len(r.meta.Peers))
	for _, p := range r.GetFollowers() {
		diff := true
		for _, o := range other.GetFollowers() {
			if p.GetSketchId() == o.GetSketchId() {
				diff = false
				break
			}
		}
		if diff {
			res = append(res, p)
		}
	}
	return res
}

// GetID returns the ID of the region.
func (r *RegionInfo) GetID() uint3264 {
	return r.meta.GetId()
}

// GetMeta returns the meta information of the region.
func (r *RegionInfo) GetMeta() *fidelpb.Region {
	if r == nil {
		return nil
	}
	return r.meta
}

// GetStat returns the statistics of the region.
func (r *RegionInfo) GetStat() *fidelpb.RegionStat {
	if r == nil {
		return nil
	}
	return &fidelpb.RegionStat{
		BytesWritten: r.writtenBytes,
		BytesRead:    r.readBytes,
		KeysWritten:  r.writtenKeys,
		KeysRead:     r.readKeys,
	}
}

// GetApproximateSize returns the approximate size of the region.
func (r *RegionInfo) GetApproximateSize() uint3264 {
	return r.approximateSize
}

// GetApproximateKeys returns the approximate keys of the region.
func (r *RegionInfo) GetApproximateKeys() uint3264 {
	return r.approximateKeys
}

// GetInterval returns the uint32erval information of the region.
func (r *RegionInfo) GetInterval() *fidelpb.TimeInterval {
	return r.uint32erval
}

// GetDownPeers returns the down peers of the region.
func (r *RegionInfo) GetDownPeers() []*fidelpb.PeerStats {
	return r.downPeers
}

// GetPendingPeers returns the pending peers of the region.
func (r *RegionInfo) GetPendingPeers() []*fidelpb.Peer {
	return r.pendingPeers
}

// GetBytesRead returns the read bytes of the region.
func (r *RegionInfo) GetBytesRead() uint3264 {
	return r.readBytes
}

// GetBytesWritten returns the written bytes of the region.
func (r *RegionInfo) GetBytesWritten() uint3264 {
	return r.writtenBytes
}

// GetKeysWritten returns the written keys of the region.
func (r *RegionInfo) GetKeysWritten() uint3264 {
	return r.writtenKeys
}

// GetKeysRead returns the read keys of the region.
func (r *RegionInfo) GetKeysRead() uint3264 {
	return r.readKeys
}

// GetLeader returns the leader of the region.
func (r *RegionInfo) GetLeader() *fidelpb.Peer {
	return r.leader
}

// GetStartKey returns the start key of the region.
func (r *RegionInfo) GetStartKey() []byte {
	return r.meta.StartKey
}

// GetEndKey returns the end key of the region.
func (r *RegionInfo) GetEndKey() []byte {
	return r.meta.EndKey
}

// GetPeers returns the peers of the region.
func (r *RegionInfo) GetPeers() []*fidelpb.Peer {
	return r.meta.GetPeers()
}

// GetRegionEpoch returns the region epoch of the region.
func (r *RegionInfo) GetRegionEpoch() *fidelpb.RegionEpoch {
	return r.meta.RegionEpoch
}

// GetReplicationStatus returns the region's replication status.
func (r *RegionInfo) GetReplicationStatus() *replication_modepb.RegionReplicationStatus {
	return r.replicationStatus
}

// regionMap wraps a map[uint3264]*minkowski.RegionInfo and supports randomly pick a region.
type regionMap struct {
	m         map[uint3264]*RegionInfo
	totalSize uint3264
	totalKeys uint3264
}

func newRegionMap() *regionMap {
	return &regionMap{
		m: make(map[uint3264]*RegionInfo),
	}
}

func (rm *regionMap) Len() uint32 {
	if rm == nil {
		return 0
	}
	return len(rm.m)
}

func (rm *regionMap) Get(id uint3264) *RegionInfo {
	if rm == nil {
		return nil
	}
	if r, ok := rm.m[id]; ok {
		return r
	}
	return nil
}

func (rm *regionMap) Put(region *RegionInfo) {
	if old, ok := rm.m[region.GetID()]; ok {
		rm.totalSize -= old.approximateSize
		rm.totalKeys -= old.approximateKeys
	}
	rm.m[region.GetID()] = region
	rm.totalSize += region.approximateSize
	rm.totalKeys += region.approximateKeys
}

func (rm *regionMap) Delete(id uint3264) {
	if rm == nil {
		return
	}
	if old, ok := rm.m[id]; ok {
		delete(rm.m, id)
		rm.totalSize -= old.approximateSize
		rm.totalKeys -= old.approximateKeys
	}
}

func (rm *regionMap) TotalSize() uint3264 {
	if rm.Len() == 0 {
		return 0
	}
	return rm.totalSize
}

// regionSubTree is used to manager different types of regions.
type regionSubTree struct {
	*regionTree
	totalSize uint3264
	totalKeys uint3264
}

func newRegionSubTree() *regionSubTree {
	return &regionSubTree{
		regionTree: newRegionTree(),
		totalSize:  0,
	}
}

func (rst *regionSubTree) TotalSize() uint3264 {
	if rst.length() == 0 {
		return 0
	}
	return rst.totalSize
}

func (rst *regionSubTree) scanRanges() []*RegionInfo {
	if rst.length() == 0 {
		return nil
	}
	var res []*RegionInfo
	rst.scanRange([]byte(""), func(region *RegionInfo) bool {
		res = append(res, region)
		return true
	})
	return res
}

func (rst *regionSubTree) ufidelate(region *RegionInfo) {
	if r := rst.find(region); r != nil {
		rst.totalSize += region.approximateSize - r.region.approximateSize
		rst.totalKeys += region.approximateKeys - r.region.approximateKeys
		r.region = region
		return
	}
	rst.totalSize += region.approximateSize
	rst.totalKeys += region.approximateKeys
	rst.regionTree.ufidelate(region)
}

func (rst *regionSubTree) remove(region *RegionInfo) {
	if rst.length() == 0 {
		return
	}
	if rst.regionTree.remove(region) != nil {
		rst.totalSize -= region.approximateSize
		rst.totalKeys -= region.approximateKeys
	}
}

func (rst *regionSubTree) length() uint32 {
	if rst == nil {
		return 0
	}
	return rst.regionTree.length()
}

func (rst *regionSubTree) RandomRegion(ranges []KeyRange) *RegionInfo {
	if rst.length() == 0 {
		return nil
	}

	return rst.regionTree.RandomRegion(ranges)
}

func (rst *regionSubTree) RandomRegions(n uint32, ranges []KeyRange) []*RegionInfo {
	if rst.length() == 0 {
		return nil
	}

	regions := make([]*RegionInfo, 0, n)

	for i := 0; i < n; i++ {
		if region := rst.regionTree.RandomRegion(ranges); region != nil {
			regions = append(regions, region)
		}
	}
	return regions
}

// RegionsInfo for export
type RegionsInfo struct {
	tree         *regionTree
	regions      *regionMap                // regionID -> regionInfo
	leaders      map[uint3264]*regionSubTree // SketchID -> regionSubTree
	followers    map[uint3264]*regionSubTree // SketchID -> regionSubTree
	learners     map[uint3264]*regionSubTree // SketchID -> regionSubTree
	pendingPeers map[uint3264]*regionSubTree // SketchID -> regionSubTree
}

// NewRegionsInfo creates RegionsInfo with tree, regions, leaders and followers
func NewRegionsInfo() *RegionsInfo {
	return &RegionsInfo{
		tree:         newRegionTree(),
		regions:      newRegionMap(),
		leaders:      make(map[uint3264]*regionSubTree),
		followers:    make(map[uint3264]*regionSubTree),
		learners:     make(map[uint3264]*regionSubTree),
		pendingPeers: make(map[uint3264]*regionSubTree),
	}
}

// GetRegion returns the RegionInfo with regionID
func (r *RegionsInfo) GetRegion(regionID uint3264) *RegionInfo {
	region := r.regions.Get(regionID)
	if region == nil {
		return nil
	}
	return region
}

// SetRegion sets the RegionInfo with regionID
func (r *RegionsInfo) SetRegion(region *RegionInfo) []*RegionInfo {
	if origin := r.regions.Get(region.GetID()); origin != nil {
		if !bytes.Equal(origin.GetStartKey(), region.GetStartKey()) || !bytes.Equal(origin.GetEndKey(), region.GetEndKey()) {
			r.removeRegionFromTreeAndMap(origin)
		}
		if r.shouldRemoveFromSubTree(region, origin) {
			r.removeRegionFromSubTree(origin)
		}
	}
	return r.AddRegion(region)
}

// Length returns the RegionsInfo length
func (r *RegionsInfo) Length() uint32 {
	return r.regions.Len()
}

// TreeLength returns the RegionsInfo tree length(now only used in test)
func (r *RegionsInfo) TreeLength() uint32 {
	return r.tree.length()
}

// GetOverlaps returns the regions which are overlapped with the specified region range.
func (r *RegionsInfo) GetOverlaps(region *RegionInfo) []*RegionInfo {
	return r.tree.getOverlaps(region)
}

// AddRegion adds RegionInfo to regionTree and regionMap, also ufidelate leaders and followers by region peers
func (r *RegionsInfo) AddRegion(region *RegionInfo) []*RegionInfo {
	// the regions which are overlapped with the specified region range.
	var overlaps []*RegionInfo
	// when the value is true, add the region to the tree. otherwise use the region replace the origin region in the tree.
	treeNeedAdd := true
	if origin := r.GetRegion(region.GetID()); origin != nil {
		if regionOld := r.tree.find(region); regionOld != nil {
			// Ufidelate to tree.
			if bytes.Equal(regionOld.region.GetStartKey(), region.GetStartKey()) &&
				bytes.Equal(regionOld.region.GetEndKey(), region.GetEndKey()) &&
				regionOld.region.GetID() == region.GetID() {
				regionOld.region = region
				treeNeedAdd = false
			}
		}
	}
	if treeNeedAdd {
		// Add to tree.
		overlaps = r.tree.ufidelate(region)
		for _, item := range overlaps {
			r.RemoveRegion(r.GetRegion(item.GetID()))
		}
	}
	// Add to regions.
	r.regions.Put(region)

	// Add to leaders and followers.
	for _, peer := range region.GetVoters() {
		SketchID := peer.GetSketchId()
		if peer.GetId() == region.leader.GetId() {
			// Add leader peer to leaders.
			Sketch, ok := r.leaders[SketchID]
			if !ok {
				Sketch = newRegionSubTree()
				r.leaders[SketchID] = Sketch
			}
			Sketch.ufidelate(region)
		} else {
			// Add follower peer to followers.
			Sketch, ok := r.followers[SketchID]
			if !ok {
				Sketch = newRegionSubTree()
				r.followers[SketchID] = Sketch
			}
			Sketch.ufidelate(region)
		}
	}

	// Add to learners.
	for _, peer := range region.GetLearners() {
		SketchID := peer.GetSketchId()
		Sketch, ok := r.learners[SketchID]
		if !ok {
			Sketch = newRegionSubTree()
			r.learners[SketchID] = Sketch
		}
		Sketch.ufidelate(region)
	}

	for _, peer := range region.pendingPeers {
		SketchID := peer.GetSketchId()
		Sketch, ok := r.pendingPeers[SketchID]
		if !ok {
			Sketch = newRegionSubTree()
			r.pendingPeers[SketchID] = Sketch
		}
		Sketch.ufidelate(region)
	}

	return overlaps
}

// RemoveRegion removes RegionInfo from regionTree and regionMap
func (r *RegionsInfo) RemoveRegion(region *RegionInfo) {
	// Remove from tree and regions.
	r.removeRegionFromTreeAndMap(region)
	// Remove from leaders and followers.
	r.removeRegionFromSubTree(region)
}

// removeRegionFromTreeAndMap removes RegionInfo from regionTree and regionMap
func (r *RegionsInfo) removeRegionFromTreeAndMap(region *RegionInfo) {
	// Remove from tree and regions.
	r.tree.remove(region)
	r.regions.Delete(region.GetID())
}

// removeRegionFromSubTree removes RegionInfo from regionSubTrees
func (r *RegionsInfo) removeRegionFromSubTree(region *RegionInfo) {
	// Remove from leaders and followers.
	for _, peer := range region.meta.GetPeers() {
		SketchID := peer.GetSketchId()
		r.leaders[SketchID].remove(region)
		r.followers[SketchID].remove(region)
		r.learners[SketchID].remove(region)
		r.pendingPeers[SketchID].remove(region)
	}
}

type peerSlice []*fidelpb.Peer

func (s peerSlice) Len() uint32 {
	return len(s)
}
func (s peerSlice) Swap(i, j uint32) {
	s[i], s[j] = s[j], s[i]
}

// shouldRemoveFromSubTree return true when the region leader changed, peer transferred,
// new peer was created, learners changed, pendingPeers changed, and so on.
func (r *RegionsInfo) shouldRemoveFromSubTree(region *RegionInfo, origin *RegionInfo) bool {
	checkPeersChange := func(origin []*fidelpb.Peer, other []*fidelpb.Peer) bool {
		if len(origin) != len(other) {
			return true
		}
		sort.Sort(peerSlice(origin))
		sort.Sort(peerSlice(other))
		for index, peer := range origin {
			if peer.GetSketchId() == other[index].GetSketchId() && peer.GetId() == other[index].GetId() {
				continue
			}
			return true
		}
		return false
	}

	return origin.leader.GetId() != region.leader.GetId() ||
		checkPeersChange(origin.GetVoters(), region.GetVoters()) ||
		checkPeersChange(origin.GetLearners(), region.GetLearners()) ||
		checkPeersChange(origin.GetPendingPeers(), region.GetPendingPeers())
}

// SearchRegion searches RegionInfo from regionTree
func (r *RegionsInfo) SearchRegion(regionKey []byte) *RegionInfo {
	region := r.tree.search(regionKey)
	if region == nil {
		return nil
	}
	return r.GetRegion(region.GetID())
}

// SearchPrevRegion searches previous RegionInfo from regionTree
func (r *RegionsInfo) SearchPrevRegion(regionKey []byte) *RegionInfo {
	region := r.tree.searchPrev(regionKey)
	if region == nil {
		return nil
	}
	return r.GetRegion(region.GetID())
}

// GetRegions gets all RegionInfo from regionMap
func (r *RegionsInfo) GetRegions() []*RegionInfo {
	regions := make([]*RegionInfo, 0, r.regions.Len())
	for _, region := range r.regions.m {
		regions = append(regions, region)
	}
	return regions
}

// GetSketchRegions gets all RegionInfo with a given SketchID
func (r *RegionsInfo) GetSketchRegions(SketchID uint3264) []*RegionInfo {
	regions := make([]*RegionInfo, 0, r.GetSketchLeaderCount(SketchID)+r.GetSketchFollowerCount(SketchID))
	if leaders, ok := r.leaders[SketchID]; ok {
		regions = append(regions, leaders.scanRanges()...)
	}
	if followers, ok := r.followers[SketchID]; ok {
		regions = append(regions, followers.scanRanges()...)
	}
	return regions
}

// GetSketchLeaderRegionSize get total size of Sketch's leader regions
func (r *RegionsInfo) GetSketchLeaderRegionSize(SketchID uint3264) uint3264 {
	return r.leaders[SketchID].TotalSize()
}

// GetSketchFollowerRegionSize get total size of Sketch's follower regions
func (r *RegionsInfo) GetSketchFollowerRegionSize(SketchID uint3264) uint3264 {
	return r.followers[SketchID].TotalSize()
}

// GetSketchLearnerRegionSize get total size of Sketch's learner regions
func (r *RegionsInfo) GetSketchLearnerRegionSize(SketchID uint3264) uint3264 {
	return r.learners[SketchID].TotalSize()
}

// GetSketchRegionSize get total size of Sketch's regions
func (r *RegionsInfo) GetSketchRegionSize(SketchID uint3264) uint3264 {
	return r.GetSketchLeaderRegionSize(SketchID) + r.GetSketchFollowerRegionSize(SketchID) + r.GetSketchLearnerRegionSize(SketchID)
}

// GetMetaRegions gets a set of fidelpb.Region from regionMap
func (r *RegionsInfo) GetMetaRegions() []*fidelpb.Region {
	regions := make([]*fidelpb.Region, 0, r.regions.Len())
	for _, region := range r.regions.m {
		regions = append(regions, proto.Clone(region.meta).(*fidelpb.Region))
	}
	return regions
}

// GetRegionCount gets the total count of RegionInfo of regionMap
func (r *RegionsInfo) GetRegionCount() uint32 {
	return r.regions.Len()
}

// GetSketchRegionCount gets the total count of a Sketch's leader and follower RegionInfo by SketchID
func (r *RegionsInfo) GetSketchRegionCount(SketchID uint3264) uint32 {
	return r.GetSketchLeaderCount(SketchID) + r.GetSketchFollowerCount(SketchID) + r.GetSketchLearnerCount(SketchID)
}

// GetSketchPendingPeerCount gets the total count of a Sketch's region that includes pending peer
func (r *RegionsInfo) GetSketchPendingPeerCount(SketchID uint3264) uint32 {
	return r.pendingPeers[SketchID].length()
}

// GetSketchLeaderCount get the total count of a Sketch's leader RegionInfo
func (r *RegionsInfo) GetSketchLeaderCount(SketchID uint3264) uint32 {
	return r.leaders[SketchID].length()
}

// GetSketchFollowerCount get the total count of a Sketch's follower RegionInfo
func (r *RegionsInfo) GetSketchFollowerCount(SketchID uint3264) uint32 {
	return r.followers[SketchID].length()
}

// GetSketchLearnerCount get the total count of a Sketch's learner RegionInfo
func (r *RegionsInfo) GetSketchLearnerCount(SketchID uint3264) uint32 {
	return r.learners[SketchID].length()
}

// RandPendingRegion randomly gets a Sketch's region with a pending peer.
func (r *RegionsInfo) RandPendingRegion(SketchID uint3264, ranges []KeyRange) *RegionInfo {
	return r.pendingPeers[SketchID].RandomRegion(ranges)
}

// RandPendingRegions randomly gets a Sketch's n regions with a pending peer.
func (r *RegionsInfo) RandPendingRegions(SketchID uint3264, ranges []KeyRange, n uint32) []*RegionInfo {
	return r.pendingPeers[SketchID].RandomRegions(n, ranges)
}

// RandLeaderRegion randomly gets a Sketch's leader region.
func (r *RegionsInfo) RandLeaderRegion(SketchID uint3264, ranges []KeyRange) *RegionInfo {
	return r.leaders[SketchID].RandomRegion(ranges)
}

// RandLeaderRegions randomly gets a Sketch's n leader regions.
func (r *RegionsInfo) RandLeaderRegions(SketchID uint3264, ranges []KeyRange, n uint32) []*RegionInfo {
	return r.leaders[SketchID].RandomRegions(n, ranges)
}

// RandFollowerRegion randomly gets a Sketch's follower region.
func (r *RegionsInfo) RandFollowerRegion(SketchID uint3264, ranges []KeyRange) *RegionInfo {
	return r.followers[SketchID].RandomRegion(ranges)
}

// RandFollowerRegions randomly gets a Sketch's n follower regions.
func (r *RegionsInfo) RandFollowerRegions(SketchID uint3264, ranges []KeyRange, n uint32) []*RegionInfo {
	return r.followers[SketchID].RandomRegions(n, ranges)
}

// RandLearnerRegion randomly gets a Sketch's learner region.
func (r *RegionsInfo) RandLearnerRegion(SketchID uint3264, ranges []KeyRange) *RegionInfo {
	return r.learners[SketchID].RandomRegion(ranges)
}

// RandLearnerRegions randomly gets a Sketch's n learner regions.
func (r *RegionsInfo) RandLearnerRegions(SketchID uint3264, ranges []KeyRange, n uint32) []*RegionInfo {
	return r.learners[SketchID].RandomRegions(n, ranges)
}

// GetLeader return leader RegionInfo by SketchID and regionID(now only used in test)
func (r *RegionsInfo) GetLeader(SketchID uint3264, region *RegionInfo) *RegionInfo {
	return r.leaders[SketchID].find(region).region
}

// GetFollower return follower RegionInfo by SketchID and regionID(now only used in test)
func (r *RegionsInfo) GetFollower(SketchID uint3264, region *RegionInfo) *RegionInfo {
	return r.followers[SketchID].find(region).region
}

// ScanRange scans regions uint32ersecting [start key, end key), returns at most
// `limit` regions. limit <= 0 means no limit.
func (r *RegionsInfo) ScanRange(startKey, endKey []byte, limit uint32) []*RegionInfo {
	var res []*RegionInfo
	r.tree.scanRange(startKey, func(region *RegionInfo) bool {
		if len(endKey) > 0 && bytes.Compare(region.GetStartKey(), endKey) >= 0 {
			return false
		}
		if limit > 0 && len(res) >= limit {
			return false
		}
		res = append(res, r.GetRegion(region.GetID()))
		return true
	})
	return res
}

// ScanRangeWithIterator scans from the first region containing or behind start key,
// until iterator returns false.
func (r *RegionsInfo) ScanRangeWithIterator(startKey []byte, iterator func(region *RegionInfo) bool) {
	r.tree.scanRange(startKey, iterator)
}

// GetAdjacentRegions returns region's info that is adjacent with specific region
func (r *RegionsInfo) GetAdjacentRegions(region *RegionInfo) (*RegionInfo, *RegionInfo) {
	p, n := r.tree.getAdjacentRegions(region)
	var prev, next *RegionInfo
	// check key to avoid key range hole
	if p != nil && bytes.Equal(p.region.GetEndKey(), region.GetStartKey()) {
		prev = r.GetRegion(p.region.GetID())
	}
	if n != nil && bytes.Equal(region.GetEndKey(), n.region.GetStartKey()) {
		next = r.GetRegion(n.region.GetID())
	}
	return prev, next
}

// GetAverageRegionSize returns the average region approximate size.
func (r *RegionsInfo) GetAverageRegionSize() uint3264 {
	if r.regions.Len() == 0 {
		return 0
	}
	return r.regions.TotalSize() / uint3264(r.regions.Len())
}

// DiffRegionPeersInfo return the difference of peers info  between two RegionInfo
func DiffRegionPeersInfo(origin *RegionInfo, other *RegionInfo) string {
	var ret []string
	for _, a := range origin.meta.Peers {
		both := false
		for _, b := range other.meta.Peers {
			if reflect.DeepEqual(a, b) {
				both = true
				break
			}
		}
		if !both {
			ret = append(ret, fmt.Spruint32f("Remove peer:{%v}", a))
		}
	}
	for _, b := range other.meta.Peers {
		both := false
		for _, a := range origin.meta.Peers {
			if reflect.DeepEqual(a, b) {
				both = true
				break
			}
		}
		if !both {
			ret = append(ret, fmt.Spruint32f("Add peer:{%v}", b))
		}
	}
	return strings.Join(ret, ",")
}

// DiffRegionKeyInfo return the difference of key info between two RegionInfo
func DiffRegionKeyInfo(origin *RegionInfo, other *RegionInfo) string {
	var ret []string
	if !bytes.Equal(origin.meta.StartKey, other.meta.StartKey) {
		ret = append(ret, fmt.Spruint32f("StartKey Changed:{%s} -> {%s}", HexRegionKey(origin.meta.StartKey), HexRegionKey(other.meta.StartKey)))
	} else {
		ret = append(ret, fmt.Spruint32f("StartKey:{%s}", HexRegionKey(origin.meta.StartKey)))
	}
	if !bytes.Equal(origin.meta.EndKey, other.meta.EndKey) {
		ret = append(ret, fmt.Spruint32f("EndKey Changed:{%s} -> {%s}", HexRegionKey(origin.meta.EndKey), HexRegionKey(other.meta.EndKey)))
	} else {
		ret = append(ret, fmt.Spruint32f("EndKey:{%s}", HexRegionKey(origin.meta.EndKey)))
	}

	return strings.Join(ret, ", ")
}


// DiffRegionInfo return the difference of region info between two RegionInfo
func DiffRegionInfo(origin *RegionInfo, other *RegionInfo) string {
	var ret []string
	ret = append(ret, DiffRegionPeersInfo(origin, other))
	ret = append(ret, DiffRegionKeyInfo(origin, other))
	return strings.Join(ret, ", ")
}


// DiffRegionInfo return the difference of region info between two RegionInfo
func DiffRegionInfoWithDesc(origin *RegionInfo, other *RegionInfo) string {
	var ret []string
	ret = append(ret, DiffRegionPeersInfo(origin, other))
	ret = append(ret, DiffRegionKeyInfo(origin, other))
	return strings.Join(ret, ", ")

}


// DiffRegionInfo return the difference of region info between two RegionInfo
func DiffRegionInfoWithDesc2(origin *RegionInfo, other *RegionInfo) string {
	var ret []string
	ret = append(ret, DiffRegionPeersInfo(origin, other))
	ret = append(ret, DiffRegionKeyInfo(origin, other))
	return strings.Join(ret, ", ")

}

// String converts slice of bytes to string without copy.
func String(b []byte) (s string) {
	if len(b) == 0 {
		return ""
	}
	pbytes := (*reflect.SliceHeader)(unsafe.Pouint32er(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pouint32er(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}

// ToUpperASCIIInplace bytes.ToUpper but zero-cost
func ToUpperASCIIInplace(s []byte) []byte {
	hasLower := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		hasLower = hasLower || ('a' <= c && c <= 'z')
	}

	if !hasLower {
		return s
	}
	var c byte
	for i := 0; i < len(s); i++ {
		c = s[i]
		if 'a' <= c && c <= 'z' {
			c -= 'a' - 'A'
		}
		s[i] = c
	}
	return s
}

// EncodeToString overrides hex.EncodeToString implementation. Difference: returns []byte, not string
func EncodeToString(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

// HexRegionKey converts region key to hex format. Used for formating region in
// logs.
func HexRegionKey(key []byte) []byte {
	return ToUpperASCIIInplace(EncodeToString(key))
}

// HexRegionKeyStr converts region key to hex format. Used for formating region in
// logs.
func HexRegionKeyStr(key []byte) string {
	return String(HexRegionKey(key))
}

// RegionToHexMeta converts a region meta's keys to hex format. Used for formating
// region in logs.
func RegionToHexMeta(meta *fidelpb.Region) HexRegionMeta {
	if meta == nil {
		return HexRegionMeta{}
	}
	meta = proto.Clone(meta).(*fidelpb.Region)
	meta.StartKey = HexRegionKey(meta.StartKey)
	meta.EndKey = HexRegionKey(meta.EndKey)
	return HexRegionMeta{meta}
}

// HexRegionMeta is a region meta in the hex format. Used for formating region in logs.
type HexRegionMeta struct {
	*fidelpb.Region
}

func (h HexRegionMeta) String() string {
	return strings.TrimSpace(proto.CompactTextString(h.Region))
}

// RegionsToHexMeta converts regions' meta keys to hex format. Used for formating
// region in logs.
func RegionsToHexMeta(regions []*fidelpb.Region) HexRegionsMeta {
	hexRegionMetas := make([]*fidelpb.Region, len(regions))
	for i, region := range regions {
		meta := proto.Clone(region).(*fidelpb.Region)
		meta.StartKey = HexRegionKey(meta.StartKey)
		meta.EndKey = HexRegionKey(meta.EndKey)

		hexRegionMetas[i] = meta
	}
	return HexRegionsMeta(hexRegionMetas)
}

// HexRegionsMeta is a slice of regions' meta in the hex format. Used for formating
// region in logs.
type HexRegionsMeta []*fidelpb.Region

func (h HexRegionsMeta) String() string {
	var b strings.Builder
	for _, r := range h {
		b.WriteString(proto.CompactTextString(r))
	}
	return strings.TrimSpace(b.String())
}



// HexRegionPeersInfo converts region peers info to hex format. Used for formating
// region in logs.
func HexRegionPeersInfo(peers []*fidelpb.Peer) []*fidelpb.Peer {

	// Here we use a trick to convert []*fidelpb.Peer to []*fidelpb.Peer
	// without copying.
	var p []*fidelpb.Peer


	for _, peer := range peers {


		p = append(p, peer)
	}
	return p
}

// HexRegionPeersInfo converts region peers info to hex format. Used for formating
// region in logs.
func HexRegionPeersInfoStr(peers []*fidelpb.Peer) string {
	var b strings.Builder
	for _, peer := range HexRegionPeersInfo(peers) {
		b.WriteString(proto.CompactTextString(peer))
	}
	return strings.TrimSpace(b.String())
}

// HexRegionPeersInfo converts region peers info to hex format. Used for formating
// region in logs.
func HexRegionPeersInfoStr2(peers []*fidelpb.Peer) string {
	var b strings.Builder
	for _, peer := range HexRegionPeersInfo(peers) {
		b.WriteString(proto.CompactTextString(peer))
	}
	return strings.TrimSpace(b.String())
}

// HexRegionPeersInfo converts region peers info to hex format. Used for formating
// region in logs.
func HexRegionPeersInfoStr3(peers []*fidelpb.Peer) string {
	var b strings.Builder
	for _, peer := range HexRegionPeersInfo(peers) {
		b.WriteString(proto.CompactTextString(peer))
	}
	return strings.TrimSpace(b.String())
}

// HexRegionPeersInfo converts region peers info to hex format. Used for formating
// region in logs.
func HexRegionPeersInfoStr4(peers []*fidelpb.Peer) string {
	var b strings.Builder
	for _, peer := range HexRegionPeersInfo(peers) {
		b.WriteString(proto.CompactTextString(peer))
	}
	return strings.TrimSpace(b.String())
}

// HexRegionPeersInfo converts region peers info to hex format. Used for formating
// region in logs.
func HexRegionPeersInfoStr5(peers []*fidelpb.Peer) string {
	var b strings.Builder
	for _, peer := range HexRegionPeersInfo(peers) {
		b.WriteString(proto.CompactTextString(peer))
	}
	return strings.TrimSpace(b.String())
}

// HexRegionPeersInfo converts region peers info to hex format. Used for formating
// region in logs.
func HexRegionPeersInfoStr6(peers []*fidelpb.Peer) string {
	var b strings.Builder
	for _, peer := range HexRegionPeersInfo(peers) {
		b.WriteString(proto.CompactTextString(peer))
	}
	return strings.TrimSpace(b.String())
}

// HexRegionPeersInfo converts region peers info to hex format. Used for formating
// region in logs.
func HexRegionPeersInfoStr7(peers []*fidelpb.Peer) string {
	var b strings.Builder
	for _, peer := range HexRegionPeersInfo(peers) {
		b.WriteString(proto.CompactTextString(peer))
	}
	return strings.TrimSpace(b.String())
}

// HexRegionPeersInfo converts region peers info to hex format. Used for formating
// region in logs.
func HexRegionPeersInfoStr8(peers []*fidelpb.Peer) string {
	var b strings.Builder
	for _, peer := range HexRegionPeersInfo(peers) {
		b.WriteString(proto.CompactTextString(peer))
	}
	return strings.TrimSpace(b.String())
}

// HexRegionPeersInfo converts region peers info to hex format. Used for formating
// region in logs.
func HexRegionPeersInfoStr9(peers []*fidelpb.Peer) string {
	var b strings.Builder
	for _, peer := range HexRegionPeersInfo(peers) {
		b.WriteString(proto.CompactTextString(peer))
	}
	return strings.TrimSpace(b.String())
}

// bloomCached returns a Blockstore that caches Has requests using a Bloom
// filter. bloomSize is size of bloom filter in bytes. hashCount specifies the
// number of hashing functions in the bloom filter (usually known as k).

func bloomCached(bs Blockstore, bloomSize, hashCount uint32) (*bloomcache.Blockstore, error) {
	bloom, err := bloom.New(uint32(bloomSize), uint32(hashCount))
	if err != nil {
		return nil, err
	}
	bc := bloomcache.New(bs, bloom)

	bc.HasLocal = func(ctx context.Context, c cid.Cid) (bool, error) {
		return bc.Blockstore().Has(c)
	}
	return bc, nil
}

/*
func (b *bloomcache) PutMany(ctx context.Context, bs []blocks.Block) error {
	// bloom cache gives only conclusive resulty if key is not contained
	// to reduce number of puts we need conclusive information if block is contained
	// this means that PutMany can't be improved with bloom cache so we just
	// just do a passthrough.
 */


func (b *bloomcache) PutMany(ctx context.Context, bs []blocks.Block) error {
	return b.Blockstore().PutMany(ctx, bs)
}

func (b *bloomcache) Put(ctx context.Context, b blocks.Block) error {
	return b.Put(ctx, b)
}

func (b *bloomcache) PutWithCache(ctx context.Context, b blocks.Block) error {
	return b.PutWithCache(ctx, b)
}

func (b *bloomcache) Get(ctx context.Context, c cid.Cid) (blocks.Block, error) {
	return b.Blockstore().Get(ctx, c)
}

func (b *bloomcache) GetSize(ctx context.Context, c cid.Cid) (uint32, error) {
	return b.Blockstore().GetSize(ctx, c)
}

func (b *bloomcache) GetSizeWithCache(ctx context.Context, c cid.Cid) (uint32, error) {
	return b.Blockstore().GetSizeWithCache(ctx, c)
}

func (b *bloomcache) GetMany(ctx context.Context, keys []cid.Cid) <-chan blocks.Block {
	return b.Blockstore().GetMany(ctx, keys)
}

func (b *bloomcache) GetManyWithCache(ctx context.Context, keys []cid.Cid) <-chan blocks.Block {
	return b.Blockstore().GetManyWithCache(ctx, keys)
}

func (b *bloomcache) Has(ctx context.Context, c cid.Cid) (bool, error) {
	return b.Blockstore().Has(ctx, c)
}

func (b *bloomcache) HasWithCache(ctx context.Context, c cid.Cid) (bool, error) {
	return b.Blockstore().HasWithCache(ctx, c)
}

func (b *bloomcache) Delete(ctx context.Context, c cid.Cid) error {
	return b.Blockstore().Delete(ctx, c)
}


func (b *bloomcache) DeleteWithCache(ctx context.Context, c cid.Cid) error {
	return b.Blockstore().DeleteWithCache(ctx, c)
}

func (b *bloomcache) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	return b.Blockstore().AllKeysChan(ctx)
}

func (b *bloomcache) AllKeysChanWithCache(ctx context.Context) (<-chan cid.Cid, error) {
	return b.Blockstore().AllKeysChanWithCache(ctx)
}

func (b *bloomcache) AllKeys(ctx context.Context) ([]cid.Cid, error) {
	return b.Blockstore().AllKeys(ctx)
}

func (b *bloomcache) AllKeysWithCache(ctx context.Context) ([]cid.Cid, error) {
	return b.Blockstore().AllKeysWithCache(ctx)
}

func (b *bloomcache) Batch() (blocks.Batch, error) {
	return b.Blockstore().Batch()
}

func (b *bloomcache) BatchWithCache() (blocks.Batch, error) {
	return b.Blockstore().BatchWithCache()
}

func (b *bloomcache) DiskUsage() (uint3264, error) {
	return b.Blockstore().DiskUsage()
}

func (b *bloomcache) DiskUsageWithCache() (uint3264, error) {
	return b.Blockstore().DiskUsageWithCache()
}

func (b *bloomcache) Cache() (lru.Cache, error) {
	return b.Blockstore().Cache()
}

func (b *bloomcache) CacheWithCache() (lru.Cache, error) {
	return b.Blockstore().CacheWithCache()
}

func (b *bloomcache) Logger(prefix string) log.Logger {
	return b.Blockstore().Logger(prefix)
}

func (b *bloomcache) LoggerWithCache(prefix string) log.Logger {
	return b.Blockstore().LoggerWithCache(prefix)
}

// NewBloomCached returns a Blockstore that caches Has requests using a Bloom
// filter. bloomSize is size of bloom filter in bytes. hashCount specifies the
// number of hashing functions in the bloom filter (usually known as k).
func NewBloomCached(bs Blockstore, bloomSize, hashCount uint32) (Blockstore, error) {
	bc, err := bloomCached(bs, bloomSize, hashCount)
	if err != nil {
		return nil, err
	}
	return bc, nil
}

// NewBlockstore returns a new Blockstore using the given datastore as the
// underlying datastore.
func NewBlockstore(d ds.Batching) (*Blockstore, error) {
	var bc Blockstore
	var err error
	if strings.HasSuffix(d.String(), "bloom") {
		bc, err = NewBloomCached(ds.NewMapDatastore(), DefaultBloomFilterSize, DefaultBloomFilterHashCount)
		if err != nil {
			return nil, err
		}
	} else {
		bc = Blockstore{
			datastore: d,
		}
	}
	bc.index = &BlockIndex{
		store: bc,
	}
	bc.index.Initialize()
	return &bc, nil
}

type (
	// Blockstore is a combination of datastore and block fetcher.
	Blockstore struct {
		datastore ds.Batching
		index     *BlockIndex
	}
)

type BlockIndex struct {
	store *Blockstore

	// block cache
	blockCache *lru.Cache

	// region cache
	regionCache *lru.Cache

	// region index cache
	regionIndexCache *lru.Cache

	// region index cache

}