// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.4
// source: sf/substreams/v1/test.proto

package pbsubstreams

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number uint64 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Step   int32  `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_v1_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_v1_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_sf_substreams_v1_test_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Block) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Block) GetStep() int32 {
	if x != nil {
		return x.Step
	}
	return 0
}

type MapResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber uint64 `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	BlockHash   string `protobuf:"bytes,2,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
}

func (x *MapResult) Reset() {
	*x = MapResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_v1_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapResult) ProtoMessage() {}

func (x *MapResult) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_v1_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapResult.ProtoReflect.Descriptor instead.
func (*MapResult) Descriptor() ([]byte, []int) {
	return file_sf_substreams_v1_test_proto_rawDescGZIP(), []int{1}
}

func (x *MapResult) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *MapResult) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

var File_sf_substreams_v1_test_proto protoreflect.FileDescriptor

var file_sf_substreams_v1_test_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73,
	0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x22, 0x43, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x22, 0x4d, 0x0a, 0x09, 0x4d, 0x61, 0x70,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x66, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f,
	0x70, 0x62, 0x2f, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_substreams_v1_test_proto_rawDescOnce sync.Once
	file_sf_substreams_v1_test_proto_rawDescData = file_sf_substreams_v1_test_proto_rawDesc
)

func file_sf_substreams_v1_test_proto_rawDescGZIP() []byte {
	file_sf_substreams_v1_test_proto_rawDescOnce.Do(func() {
		file_sf_substreams_v1_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_substreams_v1_test_proto_rawDescData)
	})
	return file_sf_substreams_v1_test_proto_rawDescData
}

var file_sf_substreams_v1_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sf_substreams_v1_test_proto_goTypes = []interface{}{
	(*Block)(nil),     // 0: sf.substreams.v1.test.Block
	(*MapResult)(nil), // 1: sf.substreams.v1.test.MapResult
}
var file_sf_substreams_v1_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sf_substreams_v1_test_proto_init() }
func file_sf_substreams_v1_test_proto_init() {
	if File_sf_substreams_v1_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_substreams_v1_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_sf_substreams_v1_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapResult); i {
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
			RawDescriptor: file_sf_substreams_v1_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sf_substreams_v1_test_proto_goTypes,
		DependencyIndexes: file_sf_substreams_v1_test_proto_depIdxs,
		MessageInfos:      file_sf_substreams_v1_test_proto_msgTypes,
	}.Build()
	File_sf_substreams_v1_test_proto = out.File
	file_sf_substreams_v1_test_proto_rawDesc = nil
	file_sf_substreams_v1_test_proto_goTypes = nil
	file_sf_substreams_v1_test_proto_depIdxs = nil
}
