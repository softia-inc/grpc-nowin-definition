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

type HostCommandCreateFirstRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostGroupUUID string `protobuf:"bytes,1,opt,name=hostGroupUUID,proto3" json:"hostGroupUUID,omitempty"`
	Email         string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FirstName     string `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName      string `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Password      string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *HostCommandCreateFirstRequest) Reset() {
	*x = HostCommandCreateFirstRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostCommandService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostCommandCreateFirstRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostCommandCreateFirstRequest) ProtoMessage() {}

func (x *HostCommandCreateFirstRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use HostCommandCreateFirstRequest.ProtoReflect.Descriptor instead.
func (*HostCommandCreateFirstRequest) Descriptor() ([]byte, []int) {
	return file_hostCommandService_proto_rawDescGZIP(), []int{0}
}

func (x *HostCommandCreateFirstRequest) GetHostGroupUUID() string {
	if x != nil {
		return x.HostGroupUUID
	}
	return ""
}

func (x *HostCommandCreateFirstRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *HostCommandCreateFirstRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *HostCommandCreateFirstRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *HostCommandCreateFirstRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type HostCommandCreateFirstResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JwtToken  string `protobuf:"bytes,1,opt,name=jwtToken,proto3" json:"jwtToken,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
}

func (x *HostCommandCreateFirstResponse) Reset() {
	*x = HostCommandCreateFirstResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostCommandService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostCommandCreateFirstResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostCommandCreateFirstResponse) ProtoMessage() {}

func (x *HostCommandCreateFirstResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use HostCommandCreateFirstResponse.ProtoReflect.Descriptor instead.
func (*HostCommandCreateFirstResponse) Descriptor() ([]byte, []int) {
	return file_hostCommandService_proto_rawDescGZIP(), []int{1}
}

func (x *HostCommandCreateFirstResponse) GetJwtToken() string {
	if x != nil {
		return x.JwtToken
	}
	return ""
}

func (x *HostCommandCreateFirstResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *HostCommandCreateFirstResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *HostCommandCreateFirstResponse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

var File_hostCommandService_proto protoreflect.FileDescriptor

var file_hostCommandService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x68, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e, 0x6f, 0x77, 0x69,
	0x6e, 0x22, 0xb1, 0x01, 0x0a, 0x1d, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x6f, 0x73, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x1e, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x77, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x77, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x32, 0x72, 0x0a, 0x12, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x6e, 0x6f, 0x77, 0x69,
	0x6e, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x52, 0x65,
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

var file_hostCommandService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_hostCommandService_proto_goTypes = []interface{}{
	(*HostCommandCreateFirstRequest)(nil),  // 0: nowin.HostCommandCreateFirstRequest
	(*HostCommandCreateFirstResponse)(nil), // 1: nowin.HostCommandCreateFirstResponse
}
var file_hostCommandService_proto_depIdxs = []int32{
	0, // 0: nowin.HostCommandService.CreateFirst:input_type -> nowin.HostCommandCreateFirstRequest
	1, // 1: nowin.HostCommandService.CreateFirst:output_type -> nowin.HostCommandCreateFirstResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
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
			switch v := v.(*HostCommandCreateFirstRequest); i {
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
			switch v := v.(*HostCommandCreateFirstResponse); i {
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
			NumMessages:   2,
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
	CreateFirst(ctx context.Context, in *HostCommandCreateFirstRequest, opts ...grpc.CallOption) (*HostCommandCreateFirstResponse, error)
}

type hostCommandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHostCommandServiceClient(cc grpc.ClientConnInterface) HostCommandServiceClient {
	return &hostCommandServiceClient{cc}
}

func (c *hostCommandServiceClient) CreateFirst(ctx context.Context, in *HostCommandCreateFirstRequest, opts ...grpc.CallOption) (*HostCommandCreateFirstResponse, error) {
	out := new(HostCommandCreateFirstResponse)
	err := c.cc.Invoke(ctx, "/nowin.HostCommandService/CreateFirst", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostCommandServiceServer is the server API for HostCommandService service.
type HostCommandServiceServer interface {
	CreateFirst(context.Context, *HostCommandCreateFirstRequest) (*HostCommandCreateFirstResponse, error)
}

// UnimplementedHostCommandServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHostCommandServiceServer struct {
}

func (*UnimplementedHostCommandServiceServer) CreateFirst(context.Context, *HostCommandCreateFirstRequest) (*HostCommandCreateFirstResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFirst not implemented")
}

func RegisterHostCommandServiceServer(s *grpc.Server, srv HostCommandServiceServer) {
	s.RegisterService(&_HostCommandService_serviceDesc, srv)
}

func _HostCommandService_CreateFirst_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostCommandCreateFirstRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostCommandServiceServer).CreateFirst(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nowin.HostCommandService/CreateFirst",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostCommandServiceServer).CreateFirst(ctx, req.(*HostCommandCreateFirstRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HostCommandService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nowin.HostCommandService",
	HandlerType: (*HostCommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFirst",
			Handler:    _HostCommandService_CreateFirst_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hostCommandService.proto",
}
