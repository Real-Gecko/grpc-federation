// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: comment/comment.proto

package comment

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

type CommentType int32

const (
	CommentType_UNKNOWN CommentType = 0
	CommentType_TYPE1   CommentType = 1
	CommentType_TYPE2   CommentType = 2
)

// Enum value maps for CommentType.
var (
	CommentType_name = map[int32]string{
		0: "UNKNOWN",
		1: "TYPE1",
		2: "TYPE2",
	}
	CommentType_value = map[string]int32{
		"UNKNOWN": 0,
		"TYPE1":   1,
		"TYPE2":   2,
	}
)

func (x CommentType) Enum() *CommentType {
	p := new(CommentType)
	*p = x
	return p
}

func (x CommentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommentType) Descriptor() protoreflect.EnumDescriptor {
	return file_comment_comment_proto_enumTypes[0].Descriptor()
}

func (CommentType) Type() protoreflect.EnumType {
	return &file_comment_comment_proto_enumTypes[0]
}

func (x CommentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommentType.Descriptor instead.
func (CommentType) EnumDescriptor() ([]byte, []int) {
	return file_comment_comment_proto_rawDescGZIP(), []int{0}
}

var File_comment_comment_proto protoreflect.FileDescriptor

var file_comment_comment_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2a, 0x30, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x54, 0x59, 0x50, 0x45, 0x31, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x59, 0x50, 0x45, 0x32,
	0x10, 0x02, 0x42, 0x70, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x42, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x17, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58,
	0xaa, 0x02, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0xca, 0x02, 0x07, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0xe2, 0x02, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comment_comment_proto_rawDescOnce sync.Once
	file_comment_comment_proto_rawDescData = file_comment_comment_proto_rawDesc
)

func file_comment_comment_proto_rawDescGZIP() []byte {
	file_comment_comment_proto_rawDescOnce.Do(func() {
		file_comment_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_comment_proto_rawDescData)
	})
	return file_comment_comment_proto_rawDescData
}

var file_comment_comment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_comment_comment_proto_goTypes = []interface{}{
	(CommentType)(0), // 0: comment.CommentType
}
var file_comment_comment_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_comment_comment_proto_init() }
func file_comment_comment_proto_init() {
	if File_comment_comment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_comment_comment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_comment_comment_proto_goTypes,
		DependencyIndexes: file_comment_comment_proto_depIdxs,
		EnumInfos:         file_comment_comment_proto_enumTypes,
	}.Build()
	File_comment_comment_proto = out.File
	file_comment_comment_proto_rawDesc = nil
	file_comment_comment_proto_goTypes = nil
	file_comment_comment_proto_depIdxs = nil
}
