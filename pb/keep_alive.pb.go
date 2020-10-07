// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: keep_alive.proto

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

type KeepAliveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *KeepAliveRequest) Reset() {
	*x = KeepAliveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keep_alive_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeepAliveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeepAliveRequest) ProtoMessage() {}

func (x *KeepAliveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keep_alive_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeepAliveRequest.ProtoReflect.Descriptor instead.
func (*KeepAliveRequest) Descriptor() ([]byte, []int) {
	return file_keep_alive_proto_rawDescGZIP(), []int{0}
}

func (x *KeepAliveRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type KeepAliveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *KeepAliveResponse) Reset() {
	*x = KeepAliveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keep_alive_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeepAliveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeepAliveResponse) ProtoMessage() {}

func (x *KeepAliveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keep_alive_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeepAliveResponse.ProtoReflect.Descriptor instead.
func (*KeepAliveResponse) Descriptor() ([]byte, []int) {
	return file_keep_alive_proto_rawDescGZIP(), []int{1}
}

func (x *KeepAliveResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_keep_alive_proto protoreflect.FileDescriptor

var file_keep_alive_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x2c, 0x0a, 0x10, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c,
	0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x43, 0x0a, 0x09, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65,
	0x12, 0x36, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4b,
	0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keep_alive_proto_rawDescOnce sync.Once
	file_keep_alive_proto_rawDescData = file_keep_alive_proto_rawDesc
)

func file_keep_alive_proto_rawDescGZIP() []byte {
	file_keep_alive_proto_rawDescOnce.Do(func() {
		file_keep_alive_proto_rawDescData = protoimpl.X.CompressGZIP(file_keep_alive_proto_rawDescData)
	})
	return file_keep_alive_proto_rawDescData
}

var file_keep_alive_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_keep_alive_proto_goTypes = []interface{}{
	(*KeepAliveRequest)(nil),  // 0: pb.KeepAliveRequest
	(*KeepAliveResponse)(nil), // 1: pb.KeepAliveResponse
}
var file_keep_alive_proto_depIdxs = []int32{
	0, // 0: pb.KeepAlive.Check:input_type -> pb.KeepAliveRequest
	1, // 1: pb.KeepAlive.Check:output_type -> pb.KeepAliveResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_keep_alive_proto_init() }
func file_keep_alive_proto_init() {
	if File_keep_alive_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_keep_alive_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeepAliveRequest); i {
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
		file_keep_alive_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeepAliveResponse); i {
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
			RawDescriptor: file_keep_alive_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_keep_alive_proto_goTypes,
		DependencyIndexes: file_keep_alive_proto_depIdxs,
		MessageInfos:      file_keep_alive_proto_msgTypes,
	}.Build()
	File_keep_alive_proto = out.File
	file_keep_alive_proto_rawDesc = nil
	file_keep_alive_proto_goTypes = nil
	file_keep_alive_proto_depIdxs = nil
}
