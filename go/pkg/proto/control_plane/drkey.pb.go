// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.3
// source: proto/control_plane/v1/drkey.proto

package control_plane

import (
	context "context"
	drkey "github.com/scionproto/scion/go/pkg/proto/drkey"
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

type DRKeyLvl2Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseReq *drkey.DRKeyLvl2Request `protobuf:"bytes,1,opt,name=base_req,json=baseReq,proto3" json:"base_req,omitempty"`
}

func (x *DRKeyLvl2Request) Reset() {
	*x = DRKeyLvl2Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_control_plane_v1_drkey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DRKeyLvl2Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DRKeyLvl2Request) ProtoMessage() {}

func (x *DRKeyLvl2Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_control_plane_v1_drkey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DRKeyLvl2Request.ProtoReflect.Descriptor instead.
func (*DRKeyLvl2Request) Descriptor() ([]byte, []int) {
	return file_proto_control_plane_v1_drkey_proto_rawDescGZIP(), []int{0}
}

func (x *DRKeyLvl2Request) GetBaseReq() *drkey.DRKeyLvl2Request {
	if x != nil {
		return x.BaseReq
	}
	return nil
}

type DRKeyLvl2Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRep *drkey.DRKeyLvl2Response `protobuf:"bytes,1,opt,name=base_rep,json=baseRep,proto3" json:"base_rep,omitempty"`
}

func (x *DRKeyLvl2Response) Reset() {
	*x = DRKeyLvl2Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_control_plane_v1_drkey_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DRKeyLvl2Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DRKeyLvl2Response) ProtoMessage() {}

func (x *DRKeyLvl2Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_control_plane_v1_drkey_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DRKeyLvl2Response.ProtoReflect.Descriptor instead.
func (*DRKeyLvl2Response) Descriptor() ([]byte, []int) {
	return file_proto_control_plane_v1_drkey_proto_rawDescGZIP(), []int{1}
}

func (x *DRKeyLvl2Response) GetBaseRep() *drkey.DRKeyLvl2Response {
	if x != nil {
		return x.BaseRep
	}
	return nil
}

var File_proto_control_plane_v1_drkey_proto protoreflect.FileDescriptor

var file_proto_control_plane_v1_drkey_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x72, 0x6b, 0x65, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x72, 0x6b, 0x65, 0x79, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x10,
	0x44, 0x52, 0x4b, 0x65, 0x79, 0x4c, 0x76, 0x6c, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x40, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x72, 0x6b, 0x65, 0x79,
	0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x52, 0x4b, 0x65, 0x79, 0x4c, 0x76,
	0x6c, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x22, 0x56, 0x0a, 0x11, 0x44, 0x52, 0x4b, 0x65, 0x79, 0x4c, 0x76, 0x6c, 0x32, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x72, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x64, 0x72, 0x6b, 0x65, 0x79, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x52, 0x4b, 0x65, 0x79, 0x4c, 0x76, 0x6c, 0x32, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x70, 0x32, 0x70, 0x0a, 0x10, 0x44, 0x52,
	0x4b, 0x65, 0x79, 0x4c, 0x76, 0x6c, 0x31, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c,
	0x0a, 0x09, 0x44, 0x52, 0x4b, 0x65, 0x79, 0x4c, 0x76, 0x6c, 0x31, 0x12, 0x25, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x72, 0x6b, 0x65, 0x79, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x52, 0x4b, 0x65, 0x79, 0x4c, 0x76, 0x6c, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x72, 0x6b, 0x65, 0x79,
	0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x52, 0x4b, 0x65, 0x79, 0x4c, 0x76,
	0x6c, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x76, 0x0a, 0x10,
	0x44, 0x52, 0x4b, 0x65, 0x79, 0x4c, 0x76, 0x6c, 0x32, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x62, 0x0a, 0x09, 0x44, 0x52, 0x4b, 0x65, 0x79, 0x4c, 0x76, 0x6c, 0x32, 0x12, 0x28, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x52, 0x4b, 0x65, 0x79, 0x4c, 0x76, 0x6c, 0x32,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x52, 0x4b, 0x65, 0x79, 0x4c, 0x76, 0x6c, 0x32, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x69, 0x6f, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63,
	0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_control_plane_v1_drkey_proto_rawDescOnce sync.Once
	file_proto_control_plane_v1_drkey_proto_rawDescData = file_proto_control_plane_v1_drkey_proto_rawDesc
)

func file_proto_control_plane_v1_drkey_proto_rawDescGZIP() []byte {
	file_proto_control_plane_v1_drkey_proto_rawDescOnce.Do(func() {
		file_proto_control_plane_v1_drkey_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_control_plane_v1_drkey_proto_rawDescData)
	})
	return file_proto_control_plane_v1_drkey_proto_rawDescData
}

var file_proto_control_plane_v1_drkey_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_control_plane_v1_drkey_proto_goTypes = []interface{}{
	(*DRKeyLvl2Request)(nil),        // 0: proto.control_plane.v1.DRKeyLvl2Request
	(*DRKeyLvl2Response)(nil),       // 1: proto.control_plane.v1.DRKeyLvl2Response
	(*drkey.DRKeyLvl2Request)(nil),  // 2: proto.drkey.mgmt.v1.DRKeyLvl2Request
	(*drkey.DRKeyLvl2Response)(nil), // 3: proto.drkey.mgmt.v1.DRKeyLvl2Response
	(*drkey.DRKeyLvl1Request)(nil),  // 4: proto.drkey.mgmt.v1.DRKeyLvl1Request
	(*drkey.DRKeyLvl1Response)(nil), // 5: proto.drkey.mgmt.v1.DRKeyLvl1Response
}
var file_proto_control_plane_v1_drkey_proto_depIdxs = []int32{
	2, // 0: proto.control_plane.v1.DRKeyLvl2Request.base_req:type_name -> proto.drkey.mgmt.v1.DRKeyLvl2Request
	3, // 1: proto.control_plane.v1.DRKeyLvl2Response.base_rep:type_name -> proto.drkey.mgmt.v1.DRKeyLvl2Response
	4, // 2: proto.control_plane.v1.DRKeyLvl1Service.DRKeyLvl1:input_type -> proto.drkey.mgmt.v1.DRKeyLvl1Request
	0, // 3: proto.control_plane.v1.DRKeyLvl2Service.DRKeyLvl2:input_type -> proto.control_plane.v1.DRKeyLvl2Request
	5, // 4: proto.control_plane.v1.DRKeyLvl1Service.DRKeyLvl1:output_type -> proto.drkey.mgmt.v1.DRKeyLvl1Response
	1, // 5: proto.control_plane.v1.DRKeyLvl2Service.DRKeyLvl2:output_type -> proto.control_plane.v1.DRKeyLvl2Response
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_control_plane_v1_drkey_proto_init() }
func file_proto_control_plane_v1_drkey_proto_init() {
	if File_proto_control_plane_v1_drkey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_control_plane_v1_drkey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DRKeyLvl2Request); i {
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
		file_proto_control_plane_v1_drkey_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DRKeyLvl2Response); i {
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
			RawDescriptor: file_proto_control_plane_v1_drkey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_control_plane_v1_drkey_proto_goTypes,
		DependencyIndexes: file_proto_control_plane_v1_drkey_proto_depIdxs,
		MessageInfos:      file_proto_control_plane_v1_drkey_proto_msgTypes,
	}.Build()
	File_proto_control_plane_v1_drkey_proto = out.File
	file_proto_control_plane_v1_drkey_proto_rawDesc = nil
	file_proto_control_plane_v1_drkey_proto_goTypes = nil
	file_proto_control_plane_v1_drkey_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DRKeyLvl1ServiceClient is the client API for DRKeyLvl1Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DRKeyLvl1ServiceClient interface {
	DRKeyLvl1(ctx context.Context, in *drkey.DRKeyLvl1Request, opts ...grpc.CallOption) (*drkey.DRKeyLvl1Response, error)
}

type dRKeyLvl1ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDRKeyLvl1ServiceClient(cc grpc.ClientConnInterface) DRKeyLvl1ServiceClient {
	return &dRKeyLvl1ServiceClient{cc}
}

func (c *dRKeyLvl1ServiceClient) DRKeyLvl1(ctx context.Context, in *drkey.DRKeyLvl1Request, opts ...grpc.CallOption) (*drkey.DRKeyLvl1Response, error) {
	out := new(drkey.DRKeyLvl1Response)
	err := c.cc.Invoke(ctx, "/proto.control_plane.v1.DRKeyLvl1Service/DRKeyLvl1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DRKeyLvl1ServiceServer is the server API for DRKeyLvl1Service service.
type DRKeyLvl1ServiceServer interface {
	DRKeyLvl1(context.Context, *drkey.DRKeyLvl1Request) (*drkey.DRKeyLvl1Response, error)
}

// UnimplementedDRKeyLvl1ServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDRKeyLvl1ServiceServer struct {
}

func (*UnimplementedDRKeyLvl1ServiceServer) DRKeyLvl1(context.Context, *drkey.DRKeyLvl1Request) (*drkey.DRKeyLvl1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DRKeyLvl1 not implemented")
}

func RegisterDRKeyLvl1ServiceServer(s *grpc.Server, srv DRKeyLvl1ServiceServer) {
	s.RegisterService(&_DRKeyLvl1Service_serviceDesc, srv)
}

func _DRKeyLvl1Service_DRKeyLvl1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(drkey.DRKeyLvl1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DRKeyLvl1ServiceServer).DRKeyLvl1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.control_plane.v1.DRKeyLvl1Service/DRKeyLvl1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DRKeyLvl1ServiceServer).DRKeyLvl1(ctx, req.(*drkey.DRKeyLvl1Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _DRKeyLvl1Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.control_plane.v1.DRKeyLvl1Service",
	HandlerType: (*DRKeyLvl1ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DRKeyLvl1",
			Handler:    _DRKeyLvl1Service_DRKeyLvl1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/control_plane/v1/drkey.proto",
}

// DRKeyLvl2ServiceClient is the client API for DRKeyLvl2Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DRKeyLvl2ServiceClient interface {
	DRKeyLvl2(ctx context.Context, in *DRKeyLvl2Request, opts ...grpc.CallOption) (*DRKeyLvl2Response, error)
}

type dRKeyLvl2ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDRKeyLvl2ServiceClient(cc grpc.ClientConnInterface) DRKeyLvl2ServiceClient {
	return &dRKeyLvl2ServiceClient{cc}
}

func (c *dRKeyLvl2ServiceClient) DRKeyLvl2(ctx context.Context, in *DRKeyLvl2Request, opts ...grpc.CallOption) (*DRKeyLvl2Response, error) {
	out := new(DRKeyLvl2Response)
	err := c.cc.Invoke(ctx, "/proto.control_plane.v1.DRKeyLvl2Service/DRKeyLvl2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DRKeyLvl2ServiceServer is the server API for DRKeyLvl2Service service.
type DRKeyLvl2ServiceServer interface {
	DRKeyLvl2(context.Context, *DRKeyLvl2Request) (*DRKeyLvl2Response, error)
}

// UnimplementedDRKeyLvl2ServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDRKeyLvl2ServiceServer struct {
}

func (*UnimplementedDRKeyLvl2ServiceServer) DRKeyLvl2(context.Context, *DRKeyLvl2Request) (*DRKeyLvl2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DRKeyLvl2 not implemented")
}

func RegisterDRKeyLvl2ServiceServer(s *grpc.Server, srv DRKeyLvl2ServiceServer) {
	s.RegisterService(&_DRKeyLvl2Service_serviceDesc, srv)
}

func _DRKeyLvl2Service_DRKeyLvl2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DRKeyLvl2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DRKeyLvl2ServiceServer).DRKeyLvl2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.control_plane.v1.DRKeyLvl2Service/DRKeyLvl2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DRKeyLvl2ServiceServer).DRKeyLvl2(ctx, req.(*DRKeyLvl2Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _DRKeyLvl2Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.control_plane.v1.DRKeyLvl2Service",
	HandlerType: (*DRKeyLvl2ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DRKeyLvl2",
			Handler:    _DRKeyLvl2Service_DRKeyLvl2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/control_plane/v1/drkey.proto",
}
