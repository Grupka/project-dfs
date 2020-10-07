// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: ss_discovery.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DiscoverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DiscoverRequest) Reset() {
	*x = DiscoverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ss_discovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverRequest) ProtoMessage() {}

func (x *DiscoverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ss_discovery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverRequest.ProtoReflect.Descriptor instead.
func (*DiscoverRequest) Descriptor() ([]byte, []int) {
	return file_ss_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *DiscoverRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DiscoveredStorage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alias string `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`
}

func (x *DiscoveredStorage) Reset() {
	*x = DiscoveredStorage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ss_discovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoveredStorage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveredStorage) ProtoMessage() {}

func (x *DiscoveredStorage) ProtoReflect() protoreflect.Message {
	mi := &file_ss_discovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoveredStorage.ProtoReflect.Descriptor instead.
func (*DiscoveredStorage) Descriptor() ([]byte, []int) {
	return file_ss_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *DiscoveredStorage) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

type DiscoverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageInfo []*DiscoveredStorage `protobuf:"bytes,1,rep,name=storageInfo,proto3" json:"storageInfo,omitempty"`
}

func (x *DiscoverResponse) Reset() {
	*x = DiscoverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ss_discovery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverResponse) ProtoMessage() {}

func (x *DiscoverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ss_discovery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverResponse.ProtoReflect.Descriptor instead.
func (*DiscoverResponse) Descriptor() ([]byte, []int) {
	return file_ss_discovery_proto_rawDescGZIP(), []int{2}
}

func (x *DiscoverResponse) GetStorageInfo() []*DiscoveredStorage {
	if x != nil {
		return x.StorageInfo
	}
	return nil
}

var File_ss_discovery_proto protoreflect.FileDescriptor

var file_ss_discovery_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x73, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x25, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22,
	0x29, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x22, 0x4b, 0x0a, 0x10, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x4b, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x37, 0x0a, 0x08, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70,
	0x62, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ss_discovery_proto_rawDescOnce sync.Once
	file_ss_discovery_proto_rawDescData = file_ss_discovery_proto_rawDesc
)

func file_ss_discovery_proto_rawDescGZIP() []byte {
	file_ss_discovery_proto_rawDescOnce.Do(func() {
		file_ss_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_ss_discovery_proto_rawDescData)
	})
	return file_ss_discovery_proto_rawDescData
}

var file_ss_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ss_discovery_proto_goTypes = []interface{}{
	(*DiscoverRequest)(nil),   // 0: pb.DiscoverRequest
	(*DiscoveredStorage)(nil), // 1: pb.DiscoveredStorage
	(*DiscoverResponse)(nil),  // 2: pb.DiscoverResponse
}
var file_ss_discovery_proto_depIdxs = []int32{
	1, // 0: pb.DiscoverResponse.storageInfo:type_name -> pb.DiscoveredStorage
	0, // 1: pb.StorageDiscovery.Discover:input_type -> pb.DiscoverRequest
	2, // 2: pb.StorageDiscovery.Discover:output_type -> pb.DiscoverResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ss_discovery_proto_init() }
func file_ss_discovery_proto_init() {
	if File_ss_discovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ss_discovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverRequest); i {
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
		file_ss_discovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoveredStorage); i {
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
		file_ss_discovery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverResponse); i {
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
			RawDescriptor: file_ss_discovery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ss_discovery_proto_goTypes,
		DependencyIndexes: file_ss_discovery_proto_depIdxs,
		MessageInfos:      file_ss_discovery_proto_msgTypes,
	}.Build()
	File_ss_discovery_proto = out.File
	file_ss_discovery_proto_rawDesc = nil
	file_ss_discovery_proto_goTypes = nil
	file_ss_discovery_proto_depIdxs = nil
}
