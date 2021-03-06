// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client-stream-grpc/proto/user.proto

package proto

/*
客户端流式gRPC
客户端将一段连续的数据流发送到服务端，服务端返回一个响应
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 客户端请求的格式
type UserRequest struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_0e719fcfbaf89c86, []int{0}
}
func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (dst *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(dst, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

// 服务端响应的方式
type UserResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_0e719fcfbaf89c86, []int{1}
}
func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (dst *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(dst, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserResponse) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "proto.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "proto.UserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	// 微服务中获取用户信息的RPC函数
	// stram关键字表示此函数请求时将发送数据流
	GetUserInfo(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserInfoClient, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserInfo(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[0], "/proto.UserService/GetUserInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetUserInfoClient{stream}
	return x, nil
}

type UserService_GetUserInfoClient interface {
	Send(*UserRequest) error
	CloseAndRecv() (*UserResponse, error)
	grpc.ClientStream
}

type userServiceGetUserInfoClient struct {
	grpc.ClientStream
}

func (x *userServiceGetUserInfoClient) Send(m *UserRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceGetUserInfoClient) CloseAndRecv() (*UserResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	// 微服务中获取用户信息的RPC函数
	// stram关键字表示此函数请求时将发送数据流
	GetUserInfo(UserService_GetUserInfoServer) error
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUserInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).GetUserInfo(&userServiceGetUserInfoServer{stream})
}

type UserService_GetUserInfoServer interface {
	SendAndClose(*UserResponse) error
	Recv() (*UserRequest, error)
	grpc.ServerStream
}

type userServiceGetUserInfoServer struct {
	grpc.ServerStream
}

func (x *userServiceGetUserInfoServer) SendAndClose(m *UserResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceGetUserInfoServer) Recv() (*UserRequest, error) {
	m := new(UserRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUserInfo",
			Handler:       _UserService_GetUserInfo_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "client-stream-grpc/proto/user.proto",
}

func init() {
	proto.RegisterFile("client-stream-grpc/proto/user.proto", fileDescriptor_user_0e719fcfbaf89c86)
}

var fileDescriptor_user_0e719fcfbaf89c86 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0x2d, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xd5, 0x4d, 0x2f, 0x2a, 0x48, 0xd6, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x03, 0x33, 0x85, 0x58, 0xc1, 0x94,
	0x92, 0x2c, 0x17, 0x77, 0x68, 0x71, 0x6a, 0x51, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x10,
	0x1f, 0x17, 0x93, 0xa7, 0x8b, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x93, 0xa7, 0x8b, 0x92,
	0x09, 0x17, 0x0f, 0x44, 0xba, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x88, 0x8b, 0x25, 0x2f,
	0x31, 0x37, 0x15, 0xac, 0x82, 0x33, 0x08, 0xcc, 0x16, 0x12, 0xe0, 0x62, 0x4e, 0x4c, 0x4f, 0x95,
	0x60, 0x02, 0x6b, 0x02, 0x31, 0x8d, 0xdc, 0x21, 0x86, 0x06, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7,
	0x0a, 0x59, 0x70, 0x71, 0xbb, 0xa7, 0x96, 0x80, 0x44, 0x3c, 0xf3, 0xd2, 0xf2, 0x85, 0x84, 0x20,
	0x2e, 0xd0, 0x43, 0xb2, 0x57, 0x4a, 0x18, 0x45, 0x0c, 0x62, 0x99, 0x06, 0x63, 0x12, 0x1b, 0x58,
	0xd4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x85, 0x9b, 0x0e, 0xbc, 0xd2, 0x00, 0x00, 0x00,
}
