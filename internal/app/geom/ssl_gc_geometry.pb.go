// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: geom/ssl_gc_geometry.proto

package geom

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

// A vector with two dimensions
type Vector2 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	X             *float32               `protobuf:"fixed32,1,req,name=x" json:"x,omitempty"`
	Y             *float32               `protobuf:"fixed32,2,req,name=y" json:"y,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Vector2) Reset() {
	*x = Vector2{}
	mi := &file_geom_ssl_gc_geometry_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Vector2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector2) ProtoMessage() {}

func (x *Vector2) ProtoReflect() protoreflect.Message {
	mi := &file_geom_ssl_gc_geometry_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector2.ProtoReflect.Descriptor instead.
func (*Vector2) Descriptor() ([]byte, []int) {
	return file_geom_ssl_gc_geometry_proto_rawDescGZIP(), []int{0}
}

func (x *Vector2) GetX() float32 {
	if x != nil && x.X != nil {
		return *x.X
	}
	return 0
}

func (x *Vector2) GetY() float32 {
	if x != nil && x.Y != nil {
		return *x.Y
	}
	return 0
}

// A vector with three dimensions
type Vector3 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	X             *float32               `protobuf:"fixed32,1,req,name=x" json:"x,omitempty"`
	Y             *float32               `protobuf:"fixed32,2,req,name=y" json:"y,omitempty"`
	Z             *float32               `protobuf:"fixed32,3,req,name=z" json:"z,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Vector3) Reset() {
	*x = Vector3{}
	mi := &file_geom_ssl_gc_geometry_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Vector3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector3) ProtoMessage() {}

func (x *Vector3) ProtoReflect() protoreflect.Message {
	mi := &file_geom_ssl_gc_geometry_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector3.ProtoReflect.Descriptor instead.
func (*Vector3) Descriptor() ([]byte, []int) {
	return file_geom_ssl_gc_geometry_proto_rawDescGZIP(), []int{1}
}

func (x *Vector3) GetX() float32 {
	if x != nil && x.X != nil {
		return *x.X
	}
	return 0
}

func (x *Vector3) GetY() float32 {
	if x != nil && x.Y != nil {
		return *x.Y
	}
	return 0
}

func (x *Vector3) GetZ() float32 {
	if x != nil && x.Z != nil {
		return *x.Z
	}
	return 0
}

var File_geom_ssl_gc_geometry_proto protoreflect.FileDescriptor

var file_geom_ssl_gc_geometry_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x65, 0x6f, 0x6d, 0x2f, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x67, 0x65,
	0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x07,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x02, 0x28, 0x02,
	0x52, 0x01, 0x79, 0x22, 0x33, 0x0a, 0x07, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x12, 0x0c,
	0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x02, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01,
	0x79, 0x18, 0x02, 0x20, 0x02, 0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18,
	0x03, 0x20, 0x02, 0x28, 0x02, 0x52, 0x01, 0x7a, 0x42, 0x54, 0x42, 0x12, 0x53, 0x73, 0x6c, 0x47,
	0x63, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x62,
	0x6f, 0x43, 0x75, 0x70, 0x2d, 0x53, 0x53, 0x4c, 0x2f, 0x73, 0x73, 0x6c, 0x2d, 0x67, 0x61, 0x6d,
	0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x65, 0x6f, 0x6d,
}

var (
	file_geom_ssl_gc_geometry_proto_rawDescOnce sync.Once
	file_geom_ssl_gc_geometry_proto_rawDescData = file_geom_ssl_gc_geometry_proto_rawDesc
)

func file_geom_ssl_gc_geometry_proto_rawDescGZIP() []byte {
	file_geom_ssl_gc_geometry_proto_rawDescOnce.Do(func() {
		file_geom_ssl_gc_geometry_proto_rawDescData = protoimpl.X.CompressGZIP(file_geom_ssl_gc_geometry_proto_rawDescData)
	})
	return file_geom_ssl_gc_geometry_proto_rawDescData
}

var file_geom_ssl_gc_geometry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_geom_ssl_gc_geometry_proto_goTypes = []any{
	(*Vector2)(nil), // 0: Vector2
	(*Vector3)(nil), // 1: Vector3
}
var file_geom_ssl_gc_geometry_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_geom_ssl_gc_geometry_proto_init() }
func file_geom_ssl_gc_geometry_proto_init() {
	if File_geom_ssl_gc_geometry_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_geom_ssl_gc_geometry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_geom_ssl_gc_geometry_proto_goTypes,
		DependencyIndexes: file_geom_ssl_gc_geometry_proto_depIdxs,
		MessageInfos:      file_geom_ssl_gc_geometry_proto_msgTypes,
	}.Build()
	File_geom_ssl_gc_geometry_proto = out.File
	file_geom_ssl_gc_geometry_proto_rawDesc = nil
	file_geom_ssl_gc_geometry_proto_goTypes = nil
	file_geom_ssl_gc_geometry_proto_depIdxs = nil
}
