// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: pkg/proto/camera/camera.proto

package camera

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_camera_camera_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_camera_camera_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_pkg_proto_camera_camera_proto_rawDescGZIP(), []int{0}
}

type Photo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Photo) Reset() {
	*x = Photo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_camera_camera_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Photo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Photo) ProtoMessage() {}

func (x *Photo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_camera_camera_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Photo.ProtoReflect.Descriptor instead.
func (*Photo) Descriptor() ([]byte, []int) {
	return file_pkg_proto_camera_camera_proto_rawDescGZIP(), []int{1}
}

func (x *Photo) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_pkg_proto_camera_camera_proto protoreflect.FileDescriptor

var file_pkg_proto_camera_camera_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x6d, 0x65,
	0x72, 0x61, 0x2f, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x1d, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x32,
	0x3c, 0x0a, 0x0d, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2b, 0x0a, 0x09, 0x54, 0x61, 0x6b, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x2e,
	0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x63,
	0x61, 0x6d, 0x65, 0x72, 0x61, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x22, 0x00, 0x42, 0x14, 0x5a,
	0x12, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x6d,
	0x65, 0x72, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_camera_camera_proto_rawDescOnce sync.Once
	file_pkg_proto_camera_camera_proto_rawDescData = file_pkg_proto_camera_camera_proto_rawDesc
)

func file_pkg_proto_camera_camera_proto_rawDescGZIP() []byte {
	file_pkg_proto_camera_camera_proto_rawDescOnce.Do(func() {
		file_pkg_proto_camera_camera_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_camera_camera_proto_rawDescData)
	})
	return file_pkg_proto_camera_camera_proto_rawDescData
}

var file_pkg_proto_camera_camera_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_camera_camera_proto_goTypes = []interface{}{
	(*Empty)(nil), // 0: camera.Empty
	(*Photo)(nil), // 1: camera.Photo
}
var file_pkg_proto_camera_camera_proto_depIdxs = []int32{
	0, // 0: camera.CameraService.TakePhoto:input_type -> camera.Empty
	1, // 1: camera.CameraService.TakePhoto:output_type -> camera.Photo
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_camera_camera_proto_init() }
func file_pkg_proto_camera_camera_proto_init() {
	if File_pkg_proto_camera_camera_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_camera_camera_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_proto_camera_camera_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Photo); i {
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
			RawDescriptor: file_pkg_proto_camera_camera_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_camera_camera_proto_goTypes,
		DependencyIndexes: file_pkg_proto_camera_camera_proto_depIdxs,
		MessageInfos:      file_pkg_proto_camera_camera_proto_msgTypes,
	}.Build()
	File_pkg_proto_camera_camera_proto = out.File
	file_pkg_proto_camera_camera_proto_rawDesc = nil
	file_pkg_proto_camera_camera_proto_goTypes = nil
	file_pkg_proto_camera_camera_proto_depIdxs = nil
}