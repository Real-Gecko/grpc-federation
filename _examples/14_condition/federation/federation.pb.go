// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: federation/federation.proto

package federation

import (
	_ "example/post"
	_ "github.com/mercari/grpc-federation/grpc/federation"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{0}
}

func (x *GetPostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *GetPostResponse) Reset() {
	*x = GetPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostResponse) ProtoMessage() {}

func (x *GetPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostResponse.ProtoReflect.Descriptor instead.
func (*GetPostResponse) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{1}
}

func (x *GetPostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	User  *User  `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{2}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_federation_federation_proto protoreflect.FileDescriptor

var file_federation_federation_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6f,
	0x72, 0x67, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x70, 0x6f, 0x73, 0x74,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x67, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x33, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x42, 0x09, 0x9a, 0x4a, 0x06, 0x12, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x52,
	0x04, 0x70, 0x6f, 0x73, 0x74, 0x3a, 0x1f, 0x9a, 0x4a, 0x1c, 0x0a, 0x1a, 0x0a, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x6a, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0a, 0x0a, 0x02, 0x69, 0x64,
	0x12, 0x04, 0x24, 0x2e, 0x69, 0x64, 0x22, 0xa3, 0x03, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0x9a, 0x4a, 0x09,
	0x12, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x9a, 0x4a,
	0x0c, 0x12, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x0d, 0x9a, 0x4a, 0x0a, 0x12, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x5b, 0x30, 0x5d, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x9c, 0x02,
	0x9a, 0x4a, 0x98, 0x02, 0x0a, 0x39, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x12, 0x0a, 0x24, 0x2e, 0x69,
	0x64, 0x20, 0x21, 0x3d, 0x20, 0x27, 0x27, 0x72, 0x26, 0x0a, 0x18, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x0a, 0x0a, 0x02, 0x69, 0x64, 0x12, 0x04, 0x24, 0x2e, 0x69, 0x64, 0x0a,
	0x1d, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x0b, 0x72, 0x65, 0x73, 0x20, 0x21, 0x3d, 0x20,
	0x6e, 0x75, 0x6c, 0x6c, 0x5a, 0x08, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x0a, 0x35,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x0c, 0x70, 0x6f, 0x73, 0x74, 0x20, 0x21, 0x3d, 0x20,
	0x6e, 0x75, 0x6c, 0x6c, 0x6a, 0x1f, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x0c, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x0a, 0x0f, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5a, 0x06,
	0x5b, 0x70, 0x6f, 0x73, 0x74, 0x5d, 0x0a, 0x47, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x20, 0x21, 0x3d, 0x20, 0x6e, 0x75, 0x6c, 0x6c, 0x62, 0x30, 0x0a,
	0x0d, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x72, 0x12, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x62, 0x1f,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x12, 0x0c, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x0a,
	0x2b, 0x12, 0x10, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x29, 0x20,
	0x3e, 0x20, 0x30, 0x7a, 0x17, 0x12, 0x15, 0x12, 0x11, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5b, 0x30,
	0x5d, 0x2e, 0x69, 0x64, 0x20, 0x3d, 0x3d, 0x20, 0x27, 0x27, 0x18, 0x03, 0x22, 0x26, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0e, 0x9a, 0x4a, 0x0b, 0x12, 0x09, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x52, 0x02, 0x69, 0x64, 0x32, 0x66, 0x0a, 0x11, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x1a, 0x03, 0x9a, 0x4a, 0x00, 0x42, 0x9d, 0x01, 0x0a,
	0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x0f, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x66, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x4f, 0x46, 0x58, 0xaa, 0x02, 0x0e, 0x4f, 0x72,
	0x67, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xca, 0x02, 0x0e, 0x4f,
	0x72, 0x67, 0x5c, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xe2, 0x02, 0x1a,
	0x4f, 0x72, 0x67, 0x5c, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x4f, 0x72, 0x67,
	0x3a, 0x3a, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_federation_federation_proto_rawDescOnce sync.Once
	file_federation_federation_proto_rawDescData = file_federation_federation_proto_rawDesc
)

func file_federation_federation_proto_rawDescGZIP() []byte {
	file_federation_federation_proto_rawDescOnce.Do(func() {
		file_federation_federation_proto_rawDescData = protoimpl.X.CompressGZIP(file_federation_federation_proto_rawDescData)
	})
	return file_federation_federation_proto_rawDescData
}

var file_federation_federation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_federation_federation_proto_goTypes = []interface{}{
	(*GetPostRequest)(nil),  // 0: org.federation.GetPostRequest
	(*GetPostResponse)(nil), // 1: org.federation.GetPostResponse
	(*Post)(nil),            // 2: org.federation.Post
	(*User)(nil),            // 3: org.federation.User
}
var file_federation_federation_proto_depIdxs = []int32{
	2, // 0: org.federation.GetPostResponse.post:type_name -> org.federation.Post
	3, // 1: org.federation.Post.user:type_name -> org.federation.User
	0, // 2: org.federation.FederationService.GetPost:input_type -> org.federation.GetPostRequest
	1, // 3: org.federation.FederationService.GetPost:output_type -> org.federation.GetPostResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_federation_federation_proto_init() }
func file_federation_federation_proto_init() {
	if File_federation_federation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_federation_federation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostRequest); i {
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
		file_federation_federation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostResponse); i {
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
		file_federation_federation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_federation_federation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_federation_federation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_federation_federation_proto_goTypes,
		DependencyIndexes: file_federation_federation_proto_depIdxs,
		MessageInfos:      file_federation_federation_proto_msgTypes,
	}.Build()
	File_federation_federation_proto = out.File
	file_federation_federation_proto_rawDesc = nil
	file_federation_federation_proto_goTypes = nil
	file_federation_federation_proto_depIdxs = nil
}
