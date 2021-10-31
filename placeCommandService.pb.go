// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: placeCommandService.proto

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

type PlaceCommandCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaceName        string `protobuf:"bytes,1,opt,name=placeName,proto3" json:"placeName,omitempty"`
	TelephoneNumber1 string `protobuf:"bytes,2,opt,name=telephoneNumber1,proto3" json:"telephoneNumber1,omitempty"`
	TelephoneNumber2 string `protobuf:"bytes,3,opt,name=telephoneNumber2,proto3" json:"telephoneNumber2,omitempty"`
	TelephoneNumber3 string `protobuf:"bytes,4,opt,name=telephoneNumber3,proto3" json:"telephoneNumber3,omitempty"`
	CheckInAt        string `protobuf:"bytes,5,opt,name=checkInAt,proto3" json:"checkInAt,omitempty"`
	CheckOutAt       int64  `protobuf:"varint,6,opt,name=checkOutAt,proto3" json:"checkOutAt,omitempty"`
	Address          string `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Email            string `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	Introduction     string `protobuf:"bytes,9,opt,name=introduction,proto3" json:"introduction,omitempty"`
}

func (x *PlaceCommandCreateRequest) Reset() {
	*x = PlaceCommandCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_placeCommandService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceCommandCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceCommandCreateRequest) ProtoMessage() {}

func (x *PlaceCommandCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_placeCommandService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceCommandCreateRequest.ProtoReflect.Descriptor instead.
func (*PlaceCommandCreateRequest) Descriptor() ([]byte, []int) {
	return file_placeCommandService_proto_rawDescGZIP(), []int{0}
}

func (x *PlaceCommandCreateRequest) GetPlaceName() string {
	if x != nil {
		return x.PlaceName
	}
	return ""
}

func (x *PlaceCommandCreateRequest) GetTelephoneNumber1() string {
	if x != nil {
		return x.TelephoneNumber1
	}
	return ""
}

func (x *PlaceCommandCreateRequest) GetTelephoneNumber2() string {
	if x != nil {
		return x.TelephoneNumber2
	}
	return ""
}

func (x *PlaceCommandCreateRequest) GetTelephoneNumber3() string {
	if x != nil {
		return x.TelephoneNumber3
	}
	return ""
}

func (x *PlaceCommandCreateRequest) GetCheckInAt() string {
	if x != nil {
		return x.CheckInAt
	}
	return ""
}

func (x *PlaceCommandCreateRequest) GetCheckOutAt() int64 {
	if x != nil {
		return x.CheckOutAt
	}
	return 0
}

func (x *PlaceCommandCreateRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PlaceCommandCreateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PlaceCommandCreateRequest) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

type PlaceCommandCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlaceCommandCreateResponse) Reset() {
	*x = PlaceCommandCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_placeCommandService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceCommandCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceCommandCreateResponse) ProtoMessage() {}

func (x *PlaceCommandCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_placeCommandService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceCommandCreateResponse.ProtoReflect.Descriptor instead.
func (*PlaceCommandCreateResponse) Descriptor() ([]byte, []int) {
	return file_placeCommandService_proto_rawDescGZIP(), []int{1}
}

var File_placeCommandService_proto protoreflect.FileDescriptor

var file_placeCommandService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e, 0x6f, 0x77,
	0x69, 0x6e, 0x22, 0xcf, 0x02, 0x0a, 0x19, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x31, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x33, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x33, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x41, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x41, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x41, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x41, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1c, 0x0a, 0x1a, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x66, 0x0a, 0x13, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x2e, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x66, 0x74, 0x69, 0x61, 0x2d,
	0x69, 0x6e, 0x63, 0x2f, 0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x6e, 0x6f, 0x77, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_placeCommandService_proto_rawDescOnce sync.Once
	file_placeCommandService_proto_rawDescData = file_placeCommandService_proto_rawDesc
)

func file_placeCommandService_proto_rawDescGZIP() []byte {
	file_placeCommandService_proto_rawDescOnce.Do(func() {
		file_placeCommandService_proto_rawDescData = protoimpl.X.CompressGZIP(file_placeCommandService_proto_rawDescData)
	})
	return file_placeCommandService_proto_rawDescData
}

var file_placeCommandService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_placeCommandService_proto_goTypes = []interface{}{
	(*PlaceCommandCreateRequest)(nil),  // 0: nowin.PlaceCommandCreateRequest
	(*PlaceCommandCreateResponse)(nil), // 1: nowin.PlaceCommandCreateResponse
}
var file_placeCommandService_proto_depIdxs = []int32{
	0, // 0: nowin.PlaceCommandService.Create:input_type -> nowin.PlaceCommandCreateRequest
	1, // 1: nowin.PlaceCommandService.Create:output_type -> nowin.PlaceCommandCreateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_placeCommandService_proto_init() }
func file_placeCommandService_proto_init() {
	if File_placeCommandService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_placeCommandService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceCommandCreateRequest); i {
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
		file_placeCommandService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceCommandCreateResponse); i {
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
			RawDescriptor: file_placeCommandService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_placeCommandService_proto_goTypes,
		DependencyIndexes: file_placeCommandService_proto_depIdxs,
		MessageInfos:      file_placeCommandService_proto_msgTypes,
	}.Build()
	File_placeCommandService_proto = out.File
	file_placeCommandService_proto_rawDesc = nil
	file_placeCommandService_proto_goTypes = nil
	file_placeCommandService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PlaceCommandServiceClient is the client API for PlaceCommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlaceCommandServiceClient interface {
	Create(ctx context.Context, in *PlaceCommandCreateRequest, opts ...grpc.CallOption) (*PlaceCommandCreateResponse, error)
}

type placeCommandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaceCommandServiceClient(cc grpc.ClientConnInterface) PlaceCommandServiceClient {
	return &placeCommandServiceClient{cc}
}

func (c *placeCommandServiceClient) Create(ctx context.Context, in *PlaceCommandCreateRequest, opts ...grpc.CallOption) (*PlaceCommandCreateResponse, error) {
	out := new(PlaceCommandCreateResponse)
	err := c.cc.Invoke(ctx, "/nowin.PlaceCommandService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaceCommandServiceServer is the server API for PlaceCommandService service.
type PlaceCommandServiceServer interface {
	Create(context.Context, *PlaceCommandCreateRequest) (*PlaceCommandCreateResponse, error)
}

// UnimplementedPlaceCommandServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPlaceCommandServiceServer struct {
}

func (*UnimplementedPlaceCommandServiceServer) Create(context.Context, *PlaceCommandCreateRequest) (*PlaceCommandCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterPlaceCommandServiceServer(s *grpc.Server, srv PlaceCommandServiceServer) {
	s.RegisterService(&_PlaceCommandService_serviceDesc, srv)
}

func _PlaceCommandService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceCommandCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaceCommandServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nowin.PlaceCommandService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaceCommandServiceServer).Create(ctx, req.(*PlaceCommandCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlaceCommandService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nowin.PlaceCommandService",
	HandlerType: (*PlaceCommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PlaceCommandService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "placeCommandService.proto",
}
