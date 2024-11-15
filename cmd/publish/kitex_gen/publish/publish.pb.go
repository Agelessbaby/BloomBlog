// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.3
// source: publish.proto

package publish

import (
	context "context"
	feed "github.com/Agelessbaby/BloomBlog/cmd/publish/kitex_gen/feed"
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

// 请求发布帖子
type BloomblogPublishActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`                                // 用户鉴权token
	Images      [][]byte `protobuf:"bytes,2,rep,name=images,proto3" json:"images,omitempty"`                              // 帖子图片数据（支持多张图片）
	TextContent string   `protobuf:"bytes,3,opt,name=text_content,json=textContent,proto3" json:"text_content,omitempty"` // 帖子文字内容
	Title       string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`                                // 帖子标题
}

func (x *BloomblogPublishActionRequest) Reset() {
	*x = BloomblogPublishActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BloomblogPublishActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BloomblogPublishActionRequest) ProtoMessage() {}

func (x *BloomblogPublishActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BloomblogPublishActionRequest.ProtoReflect.Descriptor instead.
func (*BloomblogPublishActionRequest) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{0}
}

func (x *BloomblogPublishActionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *BloomblogPublishActionRequest) GetImages() [][]byte {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *BloomblogPublishActionRequest) GetTextContent() string {
	if x != nil {
		return x.TextContent
	}
	return ""
}

func (x *BloomblogPublishActionRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

// 发布帖子操作的响应
type BloomblogPublishActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  *string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
}

func (x *BloomblogPublishActionResponse) Reset() {
	*x = BloomblogPublishActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BloomblogPublishActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BloomblogPublishActionResponse) ProtoMessage() {}

func (x *BloomblogPublishActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BloomblogPublishActionResponse.ProtoReflect.Descriptor instead.
func (*BloomblogPublishActionResponse) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{1}
}

func (x *BloomblogPublishActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BloomblogPublishActionResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

// 请求获取用户发布的帖子列表
type BloomblogPublishListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`                  // 用户鉴权token
}

func (x *BloomblogPublishListRequest) Reset() {
	*x = BloomblogPublishListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BloomblogPublishListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BloomblogPublishListRequest) ProtoMessage() {}

func (x *BloomblogPublishListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BloomblogPublishListRequest.ProtoReflect.Descriptor instead.
func (*BloomblogPublishListRequest) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{2}
}

func (x *BloomblogPublishListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BloomblogPublishListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// 获取用户已发布帖子列表的响应
type BloomblogPublishListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32        `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  *string      `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	PostList   []*feed.Post `protobuf:"bytes,3,rep,name=post_list,json=postList,proto3" json:"post_list,omitempty"`          // 用户发布的帖子列表
}

func (x *BloomblogPublishListResponse) Reset() {
	*x = BloomblogPublishListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BloomblogPublishListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BloomblogPublishListResponse) ProtoMessage() {}

func (x *BloomblogPublishListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BloomblogPublishListResponse.ProtoReflect.Descriptor instead.
func (*BloomblogPublishListResponse) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{3}
}

func (x *BloomblogPublishListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BloomblogPublishListResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

func (x *BloomblogPublishListResponse) GetPostList() []*feed.Post {
	if x != nil {
		return x.PostList
	}
	return nil
}

var File_publish_proto protoreflect.FileDescriptor

var file_publish_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x1a, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x20, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x62, 0x6c,
	0x6f, 0x67, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x65, 0x78, 0x74, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74,
	0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x22, 0x77, 0x0a, 0x21, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x22, 0x4f, 0x0a, 0x1e, 0x62, 0x6c, 0x6f,
	0x6f, 0x6d, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9e, 0x01, 0x0a, 0x1f, 0x62,
	0x6c, 0x6f, 0x6f, 0x6d, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x22, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67,
	0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x32, 0xd6, 0x01, 0x0a, 0x0a,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x72, 0x76, 0x12, 0x66, 0x0a, 0x0d, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x62, 0x6c, 0x6f, 0x67, 0x5f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x60, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x27, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x62, 0x6c, 0x6f, 0x6f,
	0x6d, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x41, 0x67, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x62, 0x61, 0x62, 0x79, 0x2f, 0x42,
	0x6c, 0x6f, 0x6f, 0x6d, 0x42, 0x6c, 0x6f, 0x67, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_publish_proto_rawDescOnce sync.Once
	file_publish_proto_rawDescData = file_publish_proto_rawDesc
)

func file_publish_proto_rawDescGZIP() []byte {
	file_publish_proto_rawDescOnce.Do(func() {
		file_publish_proto_rawDescData = protoimpl.X.CompressGZIP(file_publish_proto_rawDescData)
	})
	return file_publish_proto_rawDescData
}

var file_publish_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_publish_proto_goTypes = []interface{}{
	(*BloomblogPublishActionRequest)(nil),  // 0: publish.bloomblog_publish_action_request
	(*BloomblogPublishActionResponse)(nil), // 1: publish.bloomblog_publish_action_response
	(*BloomblogPublishListRequest)(nil),    // 2: publish.bloomblog_publish_list_request
	(*BloomblogPublishListResponse)(nil),   // 3: publish.bloomblog_publish_list_response
	(*feed.Post)(nil),                      // 4: feed.Post
}
var file_publish_proto_depIdxs = []int32{
	4, // 0: publish.bloomblog_publish_list_response.post_list:type_name -> feed.Post
	0, // 1: publish.PublishSrv.PublishAction:input_type -> publish.bloomblog_publish_action_request
	2, // 2: publish.PublishSrv.PublishList:input_type -> publish.bloomblog_publish_list_request
	1, // 3: publish.PublishSrv.PublishAction:output_type -> publish.bloomblog_publish_action_response
	3, // 4: publish.PublishSrv.PublishList:output_type -> publish.bloomblog_publish_list_response
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_publish_proto_init() }
func file_publish_proto_init() {
	if File_publish_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_publish_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BloomblogPublishActionRequest); i {
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
		file_publish_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BloomblogPublishActionResponse); i {
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
		file_publish_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BloomblogPublishListRequest); i {
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
		file_publish_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BloomblogPublishListResponse); i {
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
	file_publish_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_publish_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_publish_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_publish_proto_goTypes,
		DependencyIndexes: file_publish_proto_depIdxs,
		MessageInfos:      file_publish_proto_msgTypes,
	}.Build()
	File_publish_proto = out.File
	file_publish_proto_rawDesc = nil
	file_publish_proto_goTypes = nil
	file_publish_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.11.3. DO NOT EDIT.

type PublishSrv interface {
	PublishAction(ctx context.Context, req *BloomblogPublishActionRequest) (res *BloomblogPublishActionResponse, err error)
	PublishList(ctx context.Context, req *BloomblogPublishListRequest) (res *BloomblogPublishListResponse, err error)
}