// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

package store

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ProtoIsAllowedStoreIDRequest struct {
	RunMode              uint32   `protobuf:"varint,1,opt,name=run_mode,json=runMode,proto3" json:"run_mode,omitempty"`
	StoreID              uint32   `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoIsAllowedStoreIDRequest) Reset()         { *m = ProtoIsAllowedStoreIDRequest{} }
func (m *ProtoIsAllowedStoreIDRequest) String() string { return proto.CompactTextString(m) }
func (*ProtoIsAllowedStoreIDRequest) ProtoMessage()    {}
func (*ProtoIsAllowedStoreIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}
func (m *ProtoIsAllowedStoreIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoIsAllowedStoreIDRequest.Unmarshal(m, b)
}
func (m *ProtoIsAllowedStoreIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoIsAllowedStoreIDRequest.Marshal(b, m, deterministic)
}
func (m *ProtoIsAllowedStoreIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoIsAllowedStoreIDRequest.Merge(m, src)
}
func (m *ProtoIsAllowedStoreIDRequest) XXX_Size() int {
	return xxx_messageInfo_ProtoIsAllowedStoreIDRequest.Size(m)
}
func (m *ProtoIsAllowedStoreIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoIsAllowedStoreIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoIsAllowedStoreIDRequest proto.InternalMessageInfo

func (m *ProtoIsAllowedStoreIDRequest) GetRunMode() uint32 {
	if m != nil {
		return m.RunMode
	}
	return 0
}

func (m *ProtoIsAllowedStoreIDRequest) GetStoreID() uint32 {
	if m != nil {
		return m.StoreID
	}
	return 0
}

type ProtoIsAllowedStoreIDResponse struct {
	IsAllowed            bool     `protobuf:"varint,1,opt,name=is_allowed,json=isAllowed,proto3" json:"is_allowed,omitempty"`
	StoreCode            string   `protobuf:"bytes,2,opt,name=store_code,json=storeCode,proto3" json:"store_code,omitempty"`
	Error                []byte   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoIsAllowedStoreIDResponse) Reset()         { *m = ProtoIsAllowedStoreIDResponse{} }
func (m *ProtoIsAllowedStoreIDResponse) String() string { return proto.CompactTextString(m) }
func (*ProtoIsAllowedStoreIDResponse) ProtoMessage()    {}
func (*ProtoIsAllowedStoreIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}
func (m *ProtoIsAllowedStoreIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoIsAllowedStoreIDResponse.Unmarshal(m, b)
}
func (m *ProtoIsAllowedStoreIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoIsAllowedStoreIDResponse.Marshal(b, m, deterministic)
}
func (m *ProtoIsAllowedStoreIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoIsAllowedStoreIDResponse.Merge(m, src)
}
func (m *ProtoIsAllowedStoreIDResponse) XXX_Size() int {
	return xxx_messageInfo_ProtoIsAllowedStoreIDResponse.Size(m)
}
func (m *ProtoIsAllowedStoreIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoIsAllowedStoreIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoIsAllowedStoreIDResponse proto.InternalMessageInfo

func (m *ProtoIsAllowedStoreIDResponse) GetIsAllowed() bool {
	if m != nil {
		return m.IsAllowed
	}
	return false
}

func (m *ProtoIsAllowedStoreIDResponse) GetStoreCode() string {
	if m != nil {
		return m.StoreCode
	}
	return ""
}

func (m *ProtoIsAllowedStoreIDResponse) GetError() []byte {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtoIsAllowedStoreIDRequest)(nil), "store.ProtoIsAllowedStoreIDRequest")
	proto.RegisterType((*ProtoIsAllowedStoreIDResponse)(nil), "store.ProtoIsAllowedStoreIDResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xb5, 0x5c, 0x5c, 0x59, 0x5b, 0x0b, 0xca, 0x1e, 0x8a, 0x50, 0x5d, 0x64, 0xd4, 0x52, 0x4c,
	0x69, 0xa5, 0x7e, 0x40, 0x0f, 0xbd, 0x59, 0xad, 0x0f, 0x2e, 0x2d, 0x14, 0xf9, 0xd6, 0x8b, 0x90,
	0xad, 0x89, 0xb2, 0x60, 0xef, 0x2a, 0xbb, 0x2b, 0x87, 0x5c, 0xf3, 0x17, 0x72, 0xcd, 0xbf, 0xca,
	0xdd, 0x07, 0x93, 0x7f, 0x90, 0x3f, 0x10, 0x76, 0x57, 0x8a, 0x49, 0xc0, 0xbe, 0xcd, 0xcc, 0x7b,
	0x6f, 0xde, 0xac, 0x9e, 0x90, 0x2b, 0x80, 0x6f, 0xc8, 0x12, 0xa2, 0x8a, 0x33, 0xc9, 0x70, 0x4f,
	0x48, 0xc6, 0xc1, 0xc7, 0x40, 0x25, 0x91, 0x04, 0x44, 0x56, 0x02, 0x35, 0x90, 0xff, 0xa9, 0x24,
	0xf2, 0xb4, 0x5e, 0x44, 0x4b, 0xb6, 0x8e, 0x4b, 0x56, 0xb2, 0x58, 0x8f, 0x17, 0xf5, 0x89, 0xee,
	0x74, 0xa3, 0xab, 0x86, 0xfe, 0xba, 0x64, 0xac, 0x5c, 0xc1, 0x9e, 0x05, 0xeb, 0x4a, 0x5e, 0x34,
	0xe0, 0xb0, 0x01, 0xf3, 0x8a, 0xc4, 0x39, 0xa5, 0x4c, 0xe6, 0x92, 0x30, 0x2a, 0x0c, 0x1a, 0x52,
	0x34, 0xfc, 0xa7, 0x8a, 0x99, 0x98, 0xac, 0x56, 0xec, 0x1c, 0x8a, 0xb9, 0x3a, 0x6a, 0xf6, 0x2b,
	0x85, 0xb3, 0x1a, 0x84, 0xc4, 0xef, 0x51, 0x9f, 0xd7, 0x34, 0x5b, 0xb3, 0x02, 0x3c, 0x6b, 0x64,
	0x8d, 0xdd, 0xe4, 0xc5, 0x6e, 0x1b, 0xd8, 0x69, 0x4d, 0xff, 0xb2, 0x02, 0x52, 0x9b, 0x9b, 0x42,
	0xf1, 0xf4, 0x73, 0x32, 0x52, 0x78, 0xdd, 0x3d, 0xaf, 0xdd, 0x66, 0x6b, 0x70, 0x56, 0x84, 0xd7,
	0x16, 0x7a, 0x73, 0xc0, 0x50, 0x54, 0x8c, 0x0a, 0xc0, 0x1f, 0x11, 0x22, 0x22, 0xcb, 0x0d, 0xa8,
	0x3d, 0xfb, 0x89, 0xbb, 0xdb, 0x06, 0xce, 0x83, 0x22, 0x75, 0x48, 0x5b, 0x2a, 0xb6, 0xf1, 0x5d,
	0xaa, 0x0b, 0x95, 0xb3, 0x63, 0xd8, 0x7a, 0xed, 0x4f, 0x75, 0xa3, 0x23, 0xda, 0x12, 0x07, 0xa8,
	0x07, 0x9c, 0x33, 0xee, 0x3d, 0x1b, 0x59, 0xe3, 0x41, 0xe2, 0xec, 0xb6, 0x41, 0x6f, 0xaa, 0x06,
	0xa9, 0x99, 0x7f, 0xbd, 0xb3, 0xd0, 0x40, 0x2b, 0xe7, 0x26, 0x2a, 0x9c, 0xa1, 0x97, 0x4f, 0x2f,
	0xc5, 0x6f, 0x23, 0xbd, 0x31, 0x3a, 0xf6, 0xe1, 0xfc, 0x77, 0xc7, 0x49, 0xe6, 0xb1, 0x61, 0x07,
	0xff, 0x46, 0xfd, 0x49, 0x61, 0xe6, 0x78, 0xd0, 0x68, 0x74, 0xe7, 0xbf, 0x8a, 0x4c, 0x72, 0x51,
	0x1b, 0x6b, 0x34, 0x55, 0xb1, 0x86, 0xde, 0xe5, 0xcd, 0xed, 0x55, 0x17, 0x87, 0xae, 0x8e, 0x74,
	0xf3, 0x25, 0xae, 0x05, 0x70, 0xf1, 0xc3, 0xfa, 0x80, 0xbf, 0x23, 0xf4, 0x87, 0x08, 0xa9, 0xe5,
	0x02, 0x1f, 0xd0, 0xfb, 0x8f, 0x5c, 0xc2, 0xce, 0x67, 0x2b, 0xb1, 0xff, 0x9b, 0x7f, 0x71, 0xf1,
	0x5c, 0x53, 0xbf, 0xdd, 0x07, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x77, 0x88, 0xd3, 0xaa, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StoreServiceClient is the client API for StoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreServiceClient interface {
	IsAllowedStoreID(ctx context.Context, in *ProtoIsAllowedStoreIDRequest, opts ...grpc.CallOption) (*ProtoIsAllowedStoreIDResponse, error)
	AddStore(ctx context.Context, in *Store, opts ...grpc.CallOption) (*types.Empty, error)
	ListStores(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (StoreService_ListStoresClient, error)
}

type storeServiceClient struct {
	cc *grpc.ClientConn
}

func NewStoreServiceClient(cc *grpc.ClientConn) StoreServiceClient {
	return &storeServiceClient{cc}
}

func (c *storeServiceClient) IsAllowedStoreID(ctx context.Context, in *ProtoIsAllowedStoreIDRequest, opts ...grpc.CallOption) (*ProtoIsAllowedStoreIDResponse, error) {
	out := new(ProtoIsAllowedStoreIDResponse)
	err := c.cc.Invoke(ctx, "/store.StoreService/IsAllowedStoreID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) AddStore(ctx context.Context, in *Store, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/store.StoreService/AddStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) ListStores(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (StoreService_ListStoresClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StoreService_serviceDesc.Streams[0], "/store.StoreService/ListStores", opts...)
	if err != nil {
		return nil, err
	}
	x := &storeServiceListStoresClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StoreService_ListStoresClient interface {
	Recv() (*Store, error)
	grpc.ClientStream
}

type storeServiceListStoresClient struct {
	grpc.ClientStream
}

func (x *storeServiceListStoresClient) Recv() (*Store, error) {
	m := new(Store)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StoreServiceServer is the server API for StoreService service.
type StoreServiceServer interface {
	IsAllowedStoreID(context.Context, *ProtoIsAllowedStoreIDRequest) (*ProtoIsAllowedStoreIDResponse, error)
	AddStore(context.Context, *Store) (*types.Empty, error)
	ListStores(*types.Empty, StoreService_ListStoresServer) error
}

func RegisterStoreServiceServer(s *grpc.Server, srv StoreServiceServer) {
	s.RegisterService(&_StoreService_serviceDesc, srv)
}

func _StoreService_IsAllowedStoreID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProtoIsAllowedStoreIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).IsAllowedStoreID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.StoreService/IsAllowedStoreID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).IsAllowedStoreID(ctx, req.(*ProtoIsAllowedStoreIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_AddStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Store)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).AddStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.StoreService/AddStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).AddStore(ctx, req.(*Store))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_ListStores_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(types.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StoreServiceServer).ListStores(m, &storeServiceListStoresServer{stream})
}

type StoreService_ListStoresServer interface {
	Send(*Store) error
	grpc.ServerStream
}

type storeServiceListStoresServer struct {
	grpc.ServerStream
}

func (x *storeServiceListStoresServer) Send(m *Store) error {
	return x.ServerStream.SendMsg(m)
}

var _StoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "store.StoreService",
	HandlerType: (*StoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAllowedStoreID",
			Handler:    _StoreService_IsAllowedStoreID_Handler,
		},
		{
			MethodName: "AddStore",
			Handler:    _StoreService_AddStore_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListStores",
			Handler:       _StoreService_ListStores_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
