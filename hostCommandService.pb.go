// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: hostCommandService.proto

package nowin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type HostCommandCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *HostCommandCreateRequest) Reset() {
	*x = HostCommandCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostCommandService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostCommandCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostCommandCreateRequest) ProtoMessage() {}

func (x *HostCommandCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hostCommandService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostCommandCreateRequest.ProtoReflect.Descriptor instead.
func (*HostCommandCreateRequest) Descriptor() ([]byte, []int) {
	return file_hostCommandService_proto_rawDescGZIP(), []int{0}
}

func (x *HostCommandCreateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *HostCommandCreateRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *HostCommandCreateRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *HostCommandCreateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type HostCommandCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *HostCommandCreateResponse) Reset() {
	*x = HostCommandCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostCommandService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostCommandCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostCommandCreateResponse) ProtoMessage() {}

func (x *HostCommandCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hostCommandService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostCommandCreateResponse.ProtoReflect.Descriptor instead.
func (*HostCommandCreateResponse) Descriptor() ([]byte, []int) {
	return file_hostCommandService_proto_rawDescGZIP(), []int{1}
}

func (x *HostCommandCreateResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type HostCommandUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostUUID  string `protobuf:"bytes,1,opt,name=hostUUID,proto3" json:"hostUUID,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *HostCommandUpdateRequest) Reset() {
	*x = HostCommandUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostCommandService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostCommandUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostCommandUpdateRequest) ProtoMessage() {}

func (x *HostCommandUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hostCommandService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostCommandUpdateRequest.ProtoReflect.Descriptor instead.
func (*HostCommandUpdateRequest) Descriptor() ([]byte, []int) {
	return file_hostCommandService_proto_rawDescGZIP(), []int{2}
}

func (x *HostCommandUpdateRequest) GetHostUUID() string {
	if x != nil {
		return x.HostUUID
	}
	return ""
}

func (x *HostCommandUpdateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *HostCommandUpdateRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *HostCommandUpdateRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *HostCommandUpdateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type HostCommandUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HostCommandUpdateResponse) Reset() {
	*x = HostCommandUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostCommandService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostCommandUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostCommandUpdateResponse) ProtoMessage() {}

func (x *HostCommandUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hostCommandService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostCommandUpdateResponse.ProtoReflect.Descriptor instead.
func (*HostCommandUpdateResponse) Descriptor() ([]byte, []int) {
	return file_hostCommandService_proto_rawDescGZIP(), []int{3}
}

type HostCommandDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostUUID string `protobuf:"bytes,1,opt,name=hostUUID,proto3" json:"hostUUID,omitempty"`
}

func (x *HostCommandDeleteRequest) Reset() {
	*x = HostCommandDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostCommandService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostCommandDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostCommandDeleteRequest) ProtoMessage() {}

func (x *HostCommandDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hostCommandService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostCommandDeleteRequest.ProtoReflect.Descriptor instead.
func (*HostCommandDeleteRequest) Descriptor() ([]byte, []int) {
	return file_hostCommandService_proto_rawDescGZIP(), []int{4}
}

func (x *HostCommandDeleteRequest) GetHostUUID() string {
	if x != nil {
		return x.HostUUID
	}
	return ""
}

type HostCommandDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HostCommandDeleteResponse) Reset() {
	*x = HostCommandDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostCommandService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostCommandDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostCommandDeleteResponse) ProtoMessage() {}

func (x *HostCommandDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hostCommandService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostCommandDeleteResponse.ProtoReflect.Descriptor instead.
func (*HostCommandDeleteResponse) Descriptor() ([]byte, []int) {
	return file_hostCommandService_proto_rawDescGZIP(), []int{5}
}

var File_hostCommandService_proto protoreflect.FileDescriptor

var file_hostCommandService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x68, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e, 0x6f, 0x77, 0x69,
	0x6e, 0x22, 0x86, 0x01, 0x0a, 0x18, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x31, 0x0a, 0x19, 0x48, 0x6f,
	0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa2, 0x01,
	0x0a, 0x18, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x36, 0x0a, 0x18, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x55, 0x55, 0x49, 0x44, 0x22, 0x1b, 0x0a, 0x19, 0x48, 0x6f, 0x73, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x81, 0x02, 0x0a, 0x12, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x2e, 0x48, 0x6f,
	0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x2e, 0x48,
	0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x2e, 0x48, 0x6f, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x2e, 0x48, 0x6f,
	0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x2e, 0x48, 0x6f, 0x73, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x2e, 0x48, 0x6f, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x66, 0x74, 0x69, 0x61, 0x2d, 0x69, 0x6e,
	0x63, 0x2f, 0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x3b, 0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hostCommandService_proto_rawDescOnce sync.Once
	file_hostCommandService_proto_rawDescData = file_hostCommandService_proto_rawDesc
)

func file_hostCommandService_proto_rawDescGZIP() []byte {
	file_hostCommandService_proto_rawDescOnce.Do(func() {
		file_hostCommandService_proto_rawDescData = protoimpl.X.CompressGZIP(file_hostCommandService_proto_rawDescData)
	})
	return file_hostCommandService_proto_rawDescData
}

var file_hostCommandService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_hostCommandService_proto_goTypes = []interface{}{
	(*HostCommandCreateRequest)(nil),  // 0: nowin.HostCommandCreateRequest
	(*HostCommandCreateResponse)(nil), // 1: nowin.HostCommandCreateResponse
	(*HostCommandUpdateRequest)(nil),  // 2: nowin.HostCommandUpdateRequest
	(*HostCommandUpdateResponse)(nil), // 3: nowin.HostCommandUpdateResponse
	(*HostCommandDeleteRequest)(nil),  // 4: nowin.HostCommandDeleteRequest
	(*HostCommandDeleteResponse)(nil), // 5: nowin.HostCommandDeleteResponse
}
var file_hostCommandService_proto_depIdxs = []int32{
	0, // 0: nowin.HostCommandService.Create:input_type -> nowin.HostCommandCreateRequest
	2, // 1: nowin.HostCommandService.Update:input_type -> nowin.HostCommandUpdateRequest
	4, // 2: nowin.HostCommandService.Delete:input_type -> nowin.HostCommandDeleteRequest
	1, // 3: nowin.HostCommandService.Create:output_type -> nowin.HostCommandCreateResponse
	3, // 4: nowin.HostCommandService.Update:output_type -> nowin.HostCommandUpdateResponse
	5, // 5: nowin.HostCommandService.Delete:output_type -> nowin.HostCommandDeleteResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hostCommandService_proto_init() }
func file_hostCommandService_proto_init() {
	if File_hostCommandService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hostCommandService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostCommandCreateRequest); i {
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
		file_hostCommandService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostCommandCreateResponse); i {
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
		file_hostCommandService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostCommandUpdateRequest); i {
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
		file_hostCommandService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostCommandUpdateResponse); i {
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
		file_hostCommandService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostCommandDeleteRequest); i {
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
		file_hostCommandService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostCommandDeleteResponse); i {
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
			RawDescriptor: file_hostCommandService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hostCommandService_proto_goTypes,
		DependencyIndexes: file_hostCommandService_proto_depIdxs,
		MessageInfos:      file_hostCommandService_proto_msgTypes,
	}.Build()
	File_hostCommandService_proto = out.File
	file_hostCommandService_proto_rawDesc = nil
	file_hostCommandService_proto_goTypes = nil
	file_hostCommandService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HostCommandServiceClient is the client API for HostCommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HostCommandServiceClient interface {
	Create(ctx context.Context, in *HostCommandCreateRequest, opts ...grpc.CallOption) (*HostCommandCreateResponse, error)
	Update(ctx context.Context, in *HostCommandUpdateRequest, opts ...grpc.CallOption) (*HostCommandUpdateResponse, error)
	Delete(ctx context.Context, in *HostCommandDeleteRequest, opts ...grpc.CallOption) (*HostCommandDeleteResponse, error)
}

type hostCommandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHostCommandServiceClient(cc grpc.ClientConnInterface) HostCommandServiceClient {
	return &hostCommandServiceClient{cc}
}

func (c *hostCommandServiceClient) Create(ctx context.Context, in *HostCommandCreateRequest, opts ...grpc.CallOption) (*HostCommandCreateResponse, error) {
	out := new(HostCommandCreateResponse)
	err := c.cc.Invoke(ctx, "/nowin.HostCommandService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostCommandServiceClient) Update(ctx context.Context, in *HostCommandUpdateRequest, opts ...grpc.CallOption) (*HostCommandUpdateResponse, error) {
	out := new(HostCommandUpdateResponse)
	err := c.cc.Invoke(ctx, "/nowin.HostCommandService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostCommandServiceClient) Delete(ctx context.Context, in *HostCommandDeleteRequest, opts ...grpc.CallOption) (*HostCommandDeleteResponse, error) {
	out := new(HostCommandDeleteResponse)
	err := c.cc.Invoke(ctx, "/nowin.HostCommandService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostCommandServiceServer is the server API for HostCommandService service.
type HostCommandServiceServer interface {
	Create(context.Context, *HostCommandCreateRequest) (*HostCommandCreateResponse, error)
	Update(context.Context, *HostCommandUpdateRequest) (*HostCommandUpdateResponse, error)
	Delete(context.Context, *HostCommandDeleteRequest) (*HostCommandDeleteResponse, error)
}

// UnimplementedHostCommandServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHostCommandServiceServer struct {
}

func (*UnimplementedHostCommandServiceServer) Create(context.Context, *HostCommandCreateRequest) (*HostCommandCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedHostCommandServiceServer) Update(context.Context, *HostCommandUpdateRequest) (*HostCommandUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedHostCommandServiceServer) Delete(context.Context, *HostCommandDeleteRequest) (*HostCommandDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterHostCommandServiceServer(s *grpc.Server, srv HostCommandServiceServer) {
	s.RegisterService(&_HostCommandService_serviceDesc, srv)
}

func _HostCommandService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostCommandCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostCommandServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nowin.HostCommandService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostCommandServiceServer).Create(ctx, req.(*HostCommandCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostCommandService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostCommandUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostCommandServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nowin.HostCommandService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostCommandServiceServer).Update(ctx, req.(*HostCommandUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostCommandService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostCommandDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostCommandServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nowin.HostCommandService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostCommandServiceServer).Delete(ctx, req.(*HostCommandDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HostCommandService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nowin.HostCommandService",
	HandlerType: (*HostCommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _HostCommandService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _HostCommandService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _HostCommandService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hostCommandService.proto",
}
