// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: handshake.proto

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

type PeerType int32

const (
	PeerType_CLIENT         PeerType = 0
	PeerType_STORAGE_SERVER PeerType = 1
	PeerType_NAMING_SERVER  PeerType = 2
)

// Enum value maps for PeerType.
var (
	PeerType_name = map[int32]string{
		0: "CLIENT",
		1: "STORAGE_SERVER",
		2: "NAMING_SERVER",
	}
	PeerType_value = map[string]int32{
		"CLIENT":         0,
		"STORAGE_SERVER": 1,
		"NAMING_SERVER":  2,
	}
)

func (x PeerType) Enum() *PeerType {
	p := new(PeerType)
	*p = x
	return p
}

func (x PeerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PeerType) Descriptor() protoreflect.EnumDescriptor {
	return file_handshake_proto_enumTypes[0].Descriptor()
}

func (PeerType) Type() protoreflect.EnumType {
	return &file_handshake_proto_enumTypes[0]
}

func (x PeerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PeerType.Descriptor instead.
func (PeerType) EnumDescriptor() ([]byte, []int) {
	return file_handshake_proto_rawDescGZIP(), []int{0}
}

type PeerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerType PeerType `protobuf:"varint,1,opt,name=peerType,proto3,enum=pb.PeerType" json:"peerType,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PeerInfo) Reset() {
	*x = PeerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handshake_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerInfo) ProtoMessage() {}

func (x *PeerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_handshake_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerInfo.ProtoReflect.Descriptor instead.
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return file_handshake_proto_rawDescGZIP(), []int{0}
}

func (x *PeerInfo) GetPeerType() PeerType {
	if x != nil {
		return x.PeerType
	}
	return PeerType_CLIENT
}

func (x *PeerInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_handshake_proto protoreflect.FileDescriptor

var file_handshake_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x48, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x28, 0x0a, 0x08, 0x70, 0x65, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x08, 0x70, 0x65, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x2a,
	0x3d, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43,
	0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x4f, 0x52, 0x41,
	0x47, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4e,
	0x41, 0x4d, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x02, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_handshake_proto_rawDescOnce sync.Once
	file_handshake_proto_rawDescData = file_handshake_proto_rawDesc
)

func file_handshake_proto_rawDescGZIP() []byte {
	file_handshake_proto_rawDescOnce.Do(func() {
		file_handshake_proto_rawDescData = protoimpl.X.CompressGZIP(file_handshake_proto_rawDescData)
	})
	return file_handshake_proto_rawDescData
}

var file_handshake_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_handshake_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_handshake_proto_goTypes = []interface{}{
	(PeerType)(0),    // 0: pb.PeerType
	(*PeerInfo)(nil), // 1: pb.PeerInfo
}
var file_handshake_proto_depIdxs = []int32{
	0, // 0: pb.PeerInfo.peerType:type_name -> pb.PeerType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_handshake_proto_init() }
func file_handshake_proto_init() {
	if File_handshake_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_handshake_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerInfo); i {
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
			RawDescriptor: file_handshake_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_handshake_proto_goTypes,
		DependencyIndexes: file_handshake_proto_depIdxs,
		EnumInfos:         file_handshake_proto_enumTypes,
		MessageInfos:      file_handshake_proto_msgTypes,
	}.Build()
	File_handshake_proto = out.File
	file_handshake_proto_rawDesc = nil
	file_handshake_proto_goTypes = nil
	file_handshake_proto_depIdxs = nil
}
