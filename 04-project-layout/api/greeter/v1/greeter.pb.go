// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: api/greeter/v1/greeter.proto

package v1

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

//
type GetGreeterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGreeterRequest) Reset() {
	*x = GetGreeterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_greeter_v1_greeter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGreeterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGreeterRequest) ProtoMessage() {}

func (x *GetGreeterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_greeter_v1_greeter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGreeterRequest.ProtoReflect.Descriptor instead.
func (*GetGreeterRequest) Descriptor() ([]byte, []int) {
	return file_api_greeter_v1_greeter_proto_rawDescGZIP(), []int{0}
}

func (x *GetGreeterRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetGreeterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Job  string `protobuf:"bytes,2,opt,name=job,proto3" json:"job,omitempty"`
	Word string `protobuf:"bytes,3,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *GetGreeterReply) Reset() {
	*x = GetGreeterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_greeter_v1_greeter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGreeterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGreeterReply) ProtoMessage() {}

func (x *GetGreeterReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_greeter_v1_greeter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGreeterReply.ProtoReflect.Descriptor instead.
func (*GetGreeterReply) Descriptor() ([]byte, []int) {
	return file_api_greeter_v1_greeter_proto_rawDescGZIP(), []int{1}
}

func (x *GetGreeterReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetGreeterReply) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *GetGreeterReply) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

//
type ListGreeterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListGreeterRequest) Reset() {
	*x = ListGreeterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_greeter_v1_greeter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGreeterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGreeterRequest) ProtoMessage() {}

func (x *ListGreeterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_greeter_v1_greeter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGreeterRequest.ProtoReflect.Descriptor instead.
func (*ListGreeterRequest) Descriptor() ([]byte, []int) {
	return file_api_greeter_v1_greeter_proto_rawDescGZIP(), []int{2}
}

type ListGt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Job  string `protobuf:"bytes,2,opt,name=job,proto3" json:"job,omitempty"`
	Word string `protobuf:"bytes,3,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *ListGt) Reset() {
	*x = ListGt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_greeter_v1_greeter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGt) ProtoMessage() {}

func (x *ListGt) ProtoReflect() protoreflect.Message {
	mi := &file_api_greeter_v1_greeter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGt.ProtoReflect.Descriptor instead.
func (*ListGt) Descriptor() ([]byte, []int) {
	return file_api_greeter_v1_greeter_proto_rawDescGZIP(), []int{3}
}

func (x *ListGt) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListGt) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *ListGt) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

type ListGreeterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ListGt `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListGreeterReply) Reset() {
	*x = ListGreeterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_greeter_v1_greeter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGreeterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGreeterReply) ProtoMessage() {}

func (x *ListGreeterReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_greeter_v1_greeter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGreeterReply.ProtoReflect.Descriptor instead.
func (*ListGreeterReply) Descriptor() ([]byte, []int) {
	return file_api_greeter_v1_greeter_proto_rawDescGZIP(), []int{4}
}

func (x *ListGreeterReply) GetData() []*ListGt {
	if x != nil {
		return x.Data
	}
	return nil
}

//
type CreateGreeterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job  string `protobuf:"bytes,2,opt,name=job,proto3" json:"job,omitempty"`
	Word string `protobuf:"bytes,3,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *CreateGreeterRequest) Reset() {
	*x = CreateGreeterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_greeter_v1_greeter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGreeterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGreeterRequest) ProtoMessage() {}

func (x *CreateGreeterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_greeter_v1_greeter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGreeterRequest.ProtoReflect.Descriptor instead.
func (*CreateGreeterRequest) Descriptor() ([]byte, []int) {
	return file_api_greeter_v1_greeter_proto_rawDescGZIP(), []int{5}
}

func (x *CreateGreeterRequest) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *CreateGreeterRequest) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

type CreateGreeterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateGreeterReply) Reset() {
	*x = CreateGreeterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_greeter_v1_greeter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGreeterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGreeterReply) ProtoMessage() {}

func (x *CreateGreeterReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_greeter_v1_greeter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGreeterReply.ProtoReflect.Descriptor instead.
func (*CreateGreeterReply) Descriptor() ([]byte, []int) {
	return file_api_greeter_v1_greeter_proto_rawDescGZIP(), []int{6}
}

//
type UpdateGreeterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Job  string `protobuf:"bytes,2,opt,name=job,proto3" json:"job,omitempty"`
	Word string `protobuf:"bytes,3,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *UpdateGreeterRequest) Reset() {
	*x = UpdateGreeterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_greeter_v1_greeter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGreeterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGreeterRequest) ProtoMessage() {}

func (x *UpdateGreeterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_greeter_v1_greeter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGreeterRequest.ProtoReflect.Descriptor instead.
func (*UpdateGreeterRequest) Descriptor() ([]byte, []int) {
	return file_api_greeter_v1_greeter_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateGreeterRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateGreeterRequest) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *UpdateGreeterRequest) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

type UpdateGreeterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateGreeterReply) Reset() {
	*x = UpdateGreeterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_greeter_v1_greeter_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGreeterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGreeterReply) ProtoMessage() {}

func (x *UpdateGreeterReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_greeter_v1_greeter_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGreeterReply.ProtoReflect.Descriptor instead.
func (*UpdateGreeterReply) Descriptor() ([]byte, []int) {
	return file_api_greeter_v1_greeter_proto_rawDescGZIP(), []int{8}
}

//
type DeleteGreeterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGreeterRequest) Reset() {
	*x = DeleteGreeterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_greeter_v1_greeter_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGreeterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGreeterRequest) ProtoMessage() {}

func (x *DeleteGreeterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_greeter_v1_greeter_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGreeterRequest.ProtoReflect.Descriptor instead.
func (*DeleteGreeterRequest) Descriptor() ([]byte, []int) {
	return file_api_greeter_v1_greeter_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteGreeterRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteGreeterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteGreeterReply) Reset() {
	*x = DeleteGreeterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_greeter_v1_greeter_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGreeterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGreeterReply) ProtoMessage() {}

func (x *DeleteGreeterReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_greeter_v1_greeter_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGreeterReply.ProtoReflect.Descriptor instead.
func (*DeleteGreeterReply) Descriptor() ([]byte, []int) {
	return file_api_greeter_v1_greeter_proto_rawDescGZIP(), []int{10}
}

var File_api_greeter_v1_greeter_proto protoreflect.FileDescriptor

var file_api_greeter_v1_greeter_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x47, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6a, 0x6f, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3e,
	0x0a, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3a,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3c, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6a, 0x6f, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x4c,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x14, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x32, 0xa3, 0x03, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x20, 0x2e,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x53, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x12, 0x20, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_greeter_v1_greeter_proto_rawDescOnce sync.Once
	file_api_greeter_v1_greeter_proto_rawDescData = file_api_greeter_v1_greeter_proto_rawDesc
)

func file_api_greeter_v1_greeter_proto_rawDescGZIP() []byte {
	file_api_greeter_v1_greeter_proto_rawDescOnce.Do(func() {
		file_api_greeter_v1_greeter_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_greeter_v1_greeter_proto_rawDescData)
	})
	return file_api_greeter_v1_greeter_proto_rawDescData
}

var file_api_greeter_v1_greeter_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_greeter_v1_greeter_proto_goTypes = []interface{}{
	(*GetGreeterRequest)(nil),    // 0: greeter.v1.GetGreeterRequest
	(*GetGreeterReply)(nil),      // 1: greeter.v1.GetGreeterReply
	(*ListGreeterRequest)(nil),   // 2: greeter.v1.ListGreeterRequest
	(*ListGt)(nil),               // 3: greeter.v1.ListGt
	(*ListGreeterReply)(nil),     // 4: greeter.v1.ListGreeterReply
	(*CreateGreeterRequest)(nil), // 5: greeter.v1.CreateGreeterRequest
	(*CreateGreeterReply)(nil),   // 6: greeter.v1.CreateGreeterReply
	(*UpdateGreeterRequest)(nil), // 7: greeter.v1.UpdateGreeterRequest
	(*UpdateGreeterReply)(nil),   // 8: greeter.v1.UpdateGreeterReply
	(*DeleteGreeterRequest)(nil), // 9: greeter.v1.DeleteGreeterRequest
	(*DeleteGreeterReply)(nil),   // 10: greeter.v1.DeleteGreeterReply
}
var file_api_greeter_v1_greeter_proto_depIdxs = []int32{
	3,  // 0: greeter.v1.ListGreeterReply.data:type_name -> greeter.v1.ListGt
	0,  // 1: greeter.v1.Greeter.GetGreeter:input_type -> greeter.v1.GetGreeterRequest
	2,  // 2: greeter.v1.Greeter.ListGreeter:input_type -> greeter.v1.ListGreeterRequest
	5,  // 3: greeter.v1.Greeter.CreateGreeter:input_type -> greeter.v1.CreateGreeterRequest
	7,  // 4: greeter.v1.Greeter.UpdateGreeter:input_type -> greeter.v1.UpdateGreeterRequest
	9,  // 5: greeter.v1.Greeter.DeleteGreeter:input_type -> greeter.v1.DeleteGreeterRequest
	1,  // 6: greeter.v1.Greeter.GetGreeter:output_type -> greeter.v1.GetGreeterReply
	4,  // 7: greeter.v1.Greeter.ListGreeter:output_type -> greeter.v1.ListGreeterReply
	6,  // 8: greeter.v1.Greeter.CreateGreeter:output_type -> greeter.v1.CreateGreeterReply
	8,  // 9: greeter.v1.Greeter.UpdateGreeter:output_type -> greeter.v1.UpdateGreeterReply
	10, // 10: greeter.v1.Greeter.DeleteGreeter:output_type -> greeter.v1.DeleteGreeterReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_greeter_v1_greeter_proto_init() }
func file_api_greeter_v1_greeter_proto_init() {
	if File_api_greeter_v1_greeter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_greeter_v1_greeter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGreeterRequest); i {
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
		file_api_greeter_v1_greeter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGreeterReply); i {
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
		file_api_greeter_v1_greeter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGreeterRequest); i {
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
		file_api_greeter_v1_greeter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGt); i {
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
		file_api_greeter_v1_greeter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGreeterReply); i {
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
		file_api_greeter_v1_greeter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGreeterRequest); i {
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
		file_api_greeter_v1_greeter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGreeterReply); i {
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
		file_api_greeter_v1_greeter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGreeterRequest); i {
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
		file_api_greeter_v1_greeter_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGreeterReply); i {
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
		file_api_greeter_v1_greeter_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGreeterRequest); i {
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
		file_api_greeter_v1_greeter_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGreeterReply); i {
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
			RawDescriptor: file_api_greeter_v1_greeter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_greeter_v1_greeter_proto_goTypes,
		DependencyIndexes: file_api_greeter_v1_greeter_proto_depIdxs,
		MessageInfos:      file_api_greeter_v1_greeter_proto_msgTypes,
	}.Build()
	File_api_greeter_v1_greeter_proto = out.File
	file_api_greeter_v1_greeter_proto_rawDesc = nil
	file_api_greeter_v1_greeter_proto_goTypes = nil
	file_api_greeter_v1_greeter_proto_depIdxs = nil
}
