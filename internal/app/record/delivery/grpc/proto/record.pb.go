// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: internal/app/record/delivery/grpc/proto/record.proto

package proto

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

type GetAllUrlsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetAllUrlsRequest) Reset() {
	*x = GetAllUrlsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUrlsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUrlsRequest) ProtoMessage() {}

func (x *GetAllUrlsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUrlsRequest.ProtoReflect.Descriptor instead.
func (*GetAllUrlsRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_record_delivery_grpc_proto_record_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllUrlsRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type ShortenURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	UniqueId    uint64 `protobuf:"varint,2,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
}

func (x *ShortenURL) Reset() {
	*x = ShortenURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURL) ProtoMessage() {}

func (x *ShortenURL) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURL.ProtoReflect.Descriptor instead.
func (*ShortenURL) Descriptor() ([]byte, []int) {
	return file_internal_app_record_delivery_grpc_proto_record_proto_rawDescGZIP(), []int{1}
}

func (x *ShortenURL) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

func (x *ShortenURL) GetUniqueId() uint64 {
	if x != nil {
		return x.UniqueId
	}
	return 0
}

type GetAllUrlsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []*ShortenURL `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *GetAllUrlsResponse) Reset() {
	*x = GetAllUrlsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUrlsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUrlsResponse) ProtoMessage() {}

func (x *GetAllUrlsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUrlsResponse.ProtoReflect.Descriptor instead.
func (*GetAllUrlsResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_record_delivery_grpc_proto_record_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllUrlsResponse) GetUrls() []*ShortenURL {
	if x != nil {
		return x.Urls
	}
	return nil
}

type GetOriginalURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueId uint64 `protobuf:"varint,1,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
}

func (x *GetOriginalURLRequest) Reset() {
	*x = GetOriginalURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOriginalURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOriginalURLRequest) ProtoMessage() {}

func (x *GetOriginalURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOriginalURLRequest.ProtoReflect.Descriptor instead.
func (*GetOriginalURLRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_record_delivery_grpc_proto_record_proto_rawDescGZIP(), []int{3}
}

func (x *GetOriginalURLRequest) GetUniqueId() uint64 {
	if x != nil {
		return x.UniqueId
	}
	return 0
}

type GetOriginalURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *GetOriginalURLResponse) Reset() {
	*x = GetOriginalURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOriginalURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOriginalURLResponse) ProtoMessage() {}

func (x *GetOriginalURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOriginalURLResponse.ProtoReflect.Descriptor instead.
func (*GetOriginalURLResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_record_delivery_grpc_proto_record_proto_rawDescGZIP(), []int{4}
}

func (x *GetOriginalURLResponse) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type ShortenURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User        string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	OriginalUrl string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *ShortenURLRequest) Reset() {
	*x = ShortenURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLRequest) ProtoMessage() {}

func (x *ShortenURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLRequest.ProtoReflect.Descriptor instead.
func (*ShortenURLRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_record_delivery_grpc_proto_record_proto_rawDescGZIP(), []int{5}
}

func (x *ShortenURLRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ShortenURLRequest) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type ShortenURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueId uint64 `protobuf:"varint,1,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
}

func (x *ShortenURLResponse) Reset() {
	*x = ShortenURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLResponse) ProtoMessage() {}

func (x *ShortenURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLResponse.ProtoReflect.Descriptor instead.
func (*ShortenURLResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_record_delivery_grpc_proto_record_proto_rawDescGZIP(), []int{6}
}

func (x *ShortenURLResponse) GetUniqueId() uint64 {
	if x != nil {
		return x.UniqueId
	}
	return 0
}

var File_internal_app_record_delivery_grpc_proto_record_proto protoreflect.FileDescriptor

var file_internal_app_record_delivery_grpc_proto_record_proto_rawDesc = []byte{
	0x0a, 0x34, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x27, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x4c, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x55, 0x52, 0x4c, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x72, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x72, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x22,
	0x34, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x75, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55,
	0x72, 0x6c, 0x22, 0x4a, 0x0a, 0x11, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x31,
	0x0a, 0x12, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49,
	0x64, 0x32, 0xd8, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x3f, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x17, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c,
	0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3d, 0x5a, 0x3b,
	0x70, 0x72, 0x61, 0x6b, 0x74, 0x69, 0x6b, 0x75, 0x6d, 0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_internal_app_record_delivery_grpc_proto_record_proto_rawDescOnce sync.Once
	file_internal_app_record_delivery_grpc_proto_record_proto_rawDescData = file_internal_app_record_delivery_grpc_proto_record_proto_rawDesc
)

func file_internal_app_record_delivery_grpc_proto_record_proto_rawDescGZIP() []byte {
	file_internal_app_record_delivery_grpc_proto_record_proto_rawDescOnce.Do(func() {
		file_internal_app_record_delivery_grpc_proto_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_app_record_delivery_grpc_proto_record_proto_rawDescData)
	})
	return file_internal_app_record_delivery_grpc_proto_record_proto_rawDescData
}

var file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_app_record_delivery_grpc_proto_record_proto_goTypes = []interface{}{
	(*GetAllUrlsRequest)(nil),      // 0: grpc.GetAllUrlsRequest
	(*ShortenURL)(nil),             // 1: grpc.ShortenURL
	(*GetAllUrlsResponse)(nil),     // 2: grpc.GetAllUrlsResponse
	(*GetOriginalURLRequest)(nil),  // 3: grpc.GetOriginalURLRequest
	(*GetOriginalURLResponse)(nil), // 4: grpc.GetOriginalURLResponse
	(*ShortenURLRequest)(nil),      // 5: grpc.ShortenURLRequest
	(*ShortenURLResponse)(nil),     // 6: grpc.ShortenURLResponse
}
var file_internal_app_record_delivery_grpc_proto_record_proto_depIdxs = []int32{
	1, // 0: grpc.GetAllUrlsResponse.urls:type_name -> grpc.ShortenURL
	0, // 1: grpc.Records.GetAllUrls:input_type -> grpc.GetAllUrlsRequest
	3, // 2: grpc.Records.GetOriginalURL:input_type -> grpc.GetOriginalURLRequest
	5, // 3: grpc.Records.ShortenURL:input_type -> grpc.ShortenURLRequest
	2, // 4: grpc.Records.GetAllUrls:output_type -> grpc.GetAllUrlsResponse
	4, // 5: grpc.Records.GetOriginalURL:output_type -> grpc.GetOriginalURLResponse
	6, // 6: grpc.Records.ShortenURL:output_type -> grpc.ShortenURLResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_app_record_delivery_grpc_proto_record_proto_init() }
func file_internal_app_record_delivery_grpc_proto_record_proto_init() {
	if File_internal_app_record_delivery_grpc_proto_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUrlsRequest); i {
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
		file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURL); i {
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
		file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUrlsResponse); i {
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
		file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOriginalURLRequest); i {
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
		file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOriginalURLResponse); i {
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
		file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURLRequest); i {
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
		file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURLResponse); i {
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
			RawDescriptor: file_internal_app_record_delivery_grpc_proto_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_app_record_delivery_grpc_proto_record_proto_goTypes,
		DependencyIndexes: file_internal_app_record_delivery_grpc_proto_record_proto_depIdxs,
		MessageInfos:      file_internal_app_record_delivery_grpc_proto_record_proto_msgTypes,
	}.Build()
	File_internal_app_record_delivery_grpc_proto_record_proto = out.File
	file_internal_app_record_delivery_grpc_proto_record_proto_rawDesc = nil
	file_internal_app_record_delivery_grpc_proto_record_proto_goTypes = nil
	file_internal_app_record_delivery_grpc_proto_record_proto_depIdxs = nil
}