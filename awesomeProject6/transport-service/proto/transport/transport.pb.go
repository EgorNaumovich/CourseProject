// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.1
// source: proto/transport/transport.proto

package transport

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

type Transport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ContainerCapacity int32  `protobuf:"varint,3,opt,name=container_capacity,json=containerCapacity,proto3" json:"container_capacity,omitempty"`
	Weight            int32  `protobuf:"varint,4,opt,name=weight,proto3" json:"weight,omitempty"`
	Available         bool   `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwnerId           string `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *Transport) Reset() {
	*x = Transport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transport_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transport) ProtoMessage() {}

func (x *Transport) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transport.ProtoReflect.Descriptor instead.
func (*Transport) Descriptor() ([]byte, []int) {
	return file_proto_transport_transport_proto_rawDescGZIP(), []int{0}
}

func (x *Transport) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transport) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Transport) GetContainerCapacity() int32 {
	if x != nil {
		return x.ContainerCapacity
	}
	return 0
}

func (x *Transport) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Transport) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

func (x *Transport) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type Description struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerCapacity int32 `protobuf:"varint,1,opt,name=container_capacity,json=containerCapacity,proto3" json:"container_capacity,omitempty"`
	Weight            int32 `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *Description) Reset() {
	*x = Description{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transport_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Description) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Description) ProtoMessage() {}

func (x *Description) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Description.ProtoReflect.Descriptor instead.
func (*Description) Descriptor() ([]byte, []int) {
	return file_proto_transport_transport_proto_rawDescGZIP(), []int{1}
}

func (x *Description) GetContainerCapacity() int32 {
	if x != nil {
		return x.ContainerCapacity
	}
	return 0
}

func (x *Description) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transport  *Transport   `protobuf:"bytes,1,opt,name=transport,proto3" json:"transport,omitempty"`
	Transports []*Transport `protobuf:"bytes,2,rep,name=transports,proto3" json:"transports,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transport_transport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_transport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_transport_transport_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetTransport() *Transport {
	if x != nil {
		return x.Transport
	}
	return nil
}

func (x *Response) GetTransports() []*Transport {
	if x != nil {
		return x.Transports
	}
	return nil
}

var File_proto_transport_transport_proto protoreflect.FileDescriptor

var file_proto_transport_transport_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x22, 0xaf, 0x01, 0x0a,
	0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d,
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x70, 0x61,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x54,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a,
	0x12, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x22, 0x74, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x32, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x34, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0a,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x32, 0x4e, 0x0a, 0x10, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a,
	0x0a, 0x09, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_transport_transport_proto_rawDescOnce sync.Once
	file_proto_transport_transport_proto_rawDescData = file_proto_transport_transport_proto_rawDesc
)

func file_proto_transport_transport_proto_rawDescGZIP() []byte {
	file_proto_transport_transport_proto_rawDescOnce.Do(func() {
		file_proto_transport_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_transport_transport_proto_rawDescData)
	})
	return file_proto_transport_transport_proto_rawDescData
}

var file_proto_transport_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_transport_transport_proto_goTypes = []interface{}{
	(*Transport)(nil),   // 0: transport.Transport
	(*Description)(nil), // 1: transport.Description
	(*Response)(nil),    // 2: transport.Response
}
var file_proto_transport_transport_proto_depIdxs = []int32{
	0, // 0: transport.Response.transport:type_name -> transport.Transport
	0, // 1: transport.Response.transports:type_name -> transport.Transport
	1, // 2: transport.TransportService.Available:input_type -> transport.Description
	2, // 3: transport.TransportService.Available:output_type -> transport.Response
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_transport_transport_proto_init() }
func file_proto_transport_transport_proto_init() {
	if File_proto_transport_transport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_transport_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transport); i {
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
		file_proto_transport_transport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Description); i {
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
		file_proto_transport_transport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_proto_transport_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_transport_transport_proto_goTypes,
		DependencyIndexes: file_proto_transport_transport_proto_depIdxs,
		MessageInfos:      file_proto_transport_transport_proto_msgTypes,
	}.Build()
	File_proto_transport_transport_proto = out.File
	file_proto_transport_transport_proto_rawDesc = nil
	file_proto_transport_transport_proto_goTypes = nil
	file_proto_transport_transport_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TransportServiceClient is the client API for TransportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransportServiceClient interface {
	Available(ctx context.Context, in *Description, opts ...grpc.CallOption) (*Response, error)
}

type transportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransportServiceClient(cc grpc.ClientConnInterface) TransportServiceClient {
	return &transportServiceClient{cc}
}

func (c *transportServiceClient) Available(ctx context.Context, in *Description, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/transport.TransportService/Available", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransportServiceServer is the server API for TransportService service.
type TransportServiceServer interface {
	Available(context.Context, *Description) (*Response, error)
}

// UnimplementedTransportServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTransportServiceServer struct {
}

func (*UnimplementedTransportServiceServer) Available(context.Context, *Description) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Available not implemented")
}

func RegisterTransportServiceServer(s *grpc.Server, srv TransportServiceServer) {
	s.RegisterService(&_TransportService_serviceDesc, srv)
}

func _TransportService_Available_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Description)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).Available(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.TransportService/Available",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).Available(ctx, req.(*Description))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transport.TransportService",
	HandlerType: (*TransportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Available",
			Handler:    _TransportService_Available_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/transport/transport.proto",
}