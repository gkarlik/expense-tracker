// Code generated by protoc-gen-go.
// source: user-service/v1/user-service.proto
// DO NOT EDIT!

/*
Package proxy is a generated protocol buffer package.

It is generated from these files:
	user-service/v1/user-service.proto

It has these top-level messages:
	UserRequest
	UserResponse
	UserIDRequest
	UserLoginRequest
	RegisterUserResponse
	UserCredentialsRequest
	EmptyResponse
*/
package proxy

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

type UserRequest struct {
	FirstName string `protobuf:"bytes,1,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=LastName" json:"LastName,omitempty"`
	Login     string `protobuf:"bytes,3,opt,name=Login" json:"Login,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=Password" json:"Password,omitempty"`
	Pin       string `protobuf:"bytes,5,opt,name=Pin" json:"Pin,omitempty"`
}

func (m *UserRequest) Reset()                    { *m = UserRequest{} }
func (m *UserRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()               {}
func (*UserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *UserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserRequest) GetPin() string {
	if m != nil {
		return m.Pin
	}
	return ""
}

type UserResponse struct {
	ID        uint32 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=LastName" json:"LastName,omitempty"`
	Login     string `protobuf:"bytes,4,opt,name=Login" json:"Login,omitempty"`
}

func (m *UserResponse) Reset()                    { *m = UserResponse{} }
func (m *UserResponse) String() string            { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()               {}
func (*UserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserResponse) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *UserResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserResponse) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserResponse) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

type UserIDRequest struct {
	ID uint32 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *UserIDRequest) Reset()                    { *m = UserIDRequest{} }
func (m *UserIDRequest) String() string            { return proto.CompactTextString(m) }
func (*UserIDRequest) ProtoMessage()               {}
func (*UserIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserIDRequest) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type UserLoginRequest struct {
	Login string `protobuf:"bytes,1,opt,name=Login" json:"Login,omitempty"`
}

func (m *UserLoginRequest) Reset()                    { *m = UserLoginRequest{} }
func (m *UserLoginRequest) String() string            { return proto.CompactTextString(m) }
func (*UserLoginRequest) ProtoMessage()               {}
func (*UserLoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UserLoginRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

type RegisterUserResponse struct {
	VerificationLink string `protobuf:"bytes,1,opt,name=VerificationLink" json:"VerificationLink,omitempty"`
}

func (m *RegisterUserResponse) Reset()                    { *m = RegisterUserResponse{} }
func (m *RegisterUserResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterUserResponse) ProtoMessage()               {}
func (*RegisterUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RegisterUserResponse) GetVerificationLink() string {
	if m != nil {
		return m.VerificationLink
	}
	return ""
}

type UserCredentialsRequest struct {
	Login    string `protobuf:"bytes,1,opt,name=Login" json:"Login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password" json:"Password,omitempty"`
	Pin      string `protobuf:"bytes,3,opt,name=Pin" json:"Pin,omitempty"`
}

func (m *UserCredentialsRequest) Reset()                    { *m = UserCredentialsRequest{} }
func (m *UserCredentialsRequest) String() string            { return proto.CompactTextString(m) }
func (*UserCredentialsRequest) ProtoMessage()               {}
func (*UserCredentialsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserCredentialsRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *UserCredentialsRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserCredentialsRequest) GetPin() string {
	if m != nil {
		return m.Pin
	}
	return ""
}

type EmptyResponse struct {
}

func (m *EmptyResponse) Reset()                    { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string            { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()               {}
func (*EmptyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*UserRequest)(nil), "UserRequest")
	proto.RegisterType((*UserResponse)(nil), "UserResponse")
	proto.RegisterType((*UserIDRequest)(nil), "UserIDRequest")
	proto.RegisterType((*UserLoginRequest)(nil), "UserLoginRequest")
	proto.RegisterType((*RegisterUserResponse)(nil), "RegisterUserResponse")
	proto.RegisterType((*UserCredentialsRequest)(nil), "UserCredentialsRequest")
	proto.RegisterType((*EmptyResponse)(nil), "EmptyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserService service

type UserServiceClient interface {
	RegisterUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error)
	AuthenticateUser(ctx context.Context, in *UserCredentialsRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetUserByID(ctx context.Context, in *UserIDRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUserByLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) RegisterUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	out := new(RegisterUserResponse)
	err := grpc.Invoke(ctx, "/UserService/RegisterUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AuthenticateUser(ctx context.Context, in *UserCredentialsRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/UserService/AuthenticateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByID(ctx context.Context, in *UserIDRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/UserService/GetUserByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/UserService/GetUserByLogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	RegisterUser(context.Context, *UserRequest) (*RegisterUserResponse, error)
	AuthenticateUser(context.Context, *UserCredentialsRequest) (*EmptyResponse, error)
	GetUserByID(context.Context, *UserIDRequest) (*UserResponse, error)
	GetUserByLogin(context.Context, *UserLoginRequest) (*UserResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RegisterUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AuthenticateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AuthenticateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/AuthenticateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AuthenticateUser(ctx, req.(*UserCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/GetUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByID(ctx, req.(*UserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/GetUserByLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _UserService_RegisterUser_Handler,
		},
		{
			MethodName: "AuthenticateUser",
			Handler:    _UserService_AuthenticateUser_Handler,
		},
		{
			MethodName: "GetUserByID",
			Handler:    _UserService_GetUserByID_Handler,
		},
		{
			MethodName: "GetUserByLogin",
			Handler:    _UserService_GetUserByLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user-service/v1/user-service.proto",
}

func init() { proto.RegisterFile("user-service/v1/user-service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x52, 0x5d, 0x4f, 0xea, 0x40,
	0x10, 0xa5, 0x2d, 0xdc, 0x7b, 0x19, 0x68, 0x6f, 0xdd, 0xa0, 0x36, 0x8d, 0x89, 0x66, 0x9f, 0x88,
	0x89, 0x4b, 0xfc, 0x7a, 0xf4, 0x41, 0xac, 0x9a, 0x26, 0xc4, 0x90, 0x1a, 0x7d, 0x30, 0xbe, 0x54,
	0x58, 0x71, 0xa3, 0xb4, 0x75, 0x77, 0x41, 0xf9, 0x0d, 0xfe, 0x52, 0xff, 0x85, 0x69, 0x4b, 0xa1,
	0x05, 0xe4, 0x6d, 0x67, 0xe6, 0x9c, 0x9d, 0x73, 0x66, 0x06, 0xf0, 0x48, 0x50, 0x7e, 0x20, 0x28,
	0x1f, 0xb3, 0x1e, 0x6d, 0x8d, 0x0f, 0x5b, 0xf9, 0x98, 0x44, 0x3c, 0x94, 0x21, 0xfe, 0x52, 0xa0,
	0x76, 0x27, 0x28, 0xf7, 0xe8, 0xfb, 0x88, 0x0a, 0x89, 0x76, 0xa0, 0x7a, 0xc5, 0xb8, 0x90, 0x37,
	0xfe, 0x90, 0x5a, 0xca, 0x9e, 0xd2, 0xac, 0x7a, 0xf3, 0x04, 0xb2, 0xe1, 0x5f, 0xc7, 0x9f, 0x16,
	0xd5, 0xa4, 0x38, 0x8b, 0x51, 0x03, 0x2a, 0x9d, 0x70, 0xc0, 0x02, 0x4b, 0x4b, 0x0a, 0x69, 0x10,
	0x33, 0xba, 0xbe, 0x10, 0x1f, 0x21, 0xef, 0x5b, 0xe5, 0x94, 0x91, 0xc5, 0xc8, 0x04, 0xad, 0xcb,
	0x02, 0xab, 0x92, 0xa4, 0xe3, 0x27, 0x0e, 0xa0, 0x9e, 0x8a, 0x11, 0x51, 0x18, 0x08, 0x8a, 0x0c,
	0x50, 0x5d, 0x27, 0x91, 0xa1, 0x7b, 0xaa, 0xeb, 0x14, 0xd5, 0xa9, 0xeb, 0xd4, 0x69, 0xbf, 0xa9,
	0x2b, 0xe7, 0xd4, 0xe1, 0x5d, 0xd0, 0xe3, 0x7e, 0xae, 0x93, 0xd9, 0x5f, 0x68, 0x88, 0x9b, 0x60,
	0xc6, 0x80, 0x04, 0x9d, 0x61, 0x66, 0x5f, 0x29, 0xf9, 0xaf, 0xda, 0xd0, 0xf0, 0xe8, 0x80, 0x09,
	0x49, 0x79, 0xc1, 0xc2, 0x3e, 0x98, 0xf7, 0x94, 0xb3, 0x67, 0xd6, 0xf3, 0x25, 0x0b, 0x83, 0x0e,
	0x0b, 0x5e, 0xa7, 0xc4, 0xa5, 0x3c, 0x7e, 0x84, 0xad, 0x98, 0x7b, 0xc1, 0x69, 0x9f, 0x06, 0x92,
	0xf9, 0x6f, 0x62, 0x6d, 0xcf, 0xc2, 0x70, 0xd5, 0xd5, 0xc3, 0xd5, 0xe6, 0xc3, 0xfd, 0x0f, 0xfa,
	0xe5, 0x30, 0x92, 0x93, 0x4c, 0xda, 0xd1, 0xf7, 0x74, 0xf7, 0xb7, 0xe9, 0x45, 0xa0, 0x53, 0xa8,
	0xe7, 0x2d, 0xa0, 0x3a, 0xc9, 0x5d, 0x86, 0xbd, 0x49, 0x56, 0xf9, 0xc3, 0x25, 0x74, 0x06, 0xe6,
	0xf9, 0x48, 0xbe, 0xc4, 0x8a, 0x7b, 0xbe, 0xa4, 0x09, 0x75, 0x9b, 0xac, 0x36, 0x62, 0x1b, 0xa4,
	0xa0, 0x01, 0x97, 0x10, 0x81, 0xda, 0x35, 0x95, 0x31, 0xbc, 0x3d, 0x71, 0x1d, 0x64, 0x90, 0xc2,
	0x46, 0x6c, 0x9d, 0x2c, 0xb4, 0x3b, 0x01, 0x63, 0x86, 0x4f, 0xc7, 0xb0, 0x41, 0x16, 0x77, 0xb4,
	0xc4, 0x6a, 0xff, 0x7d, 0xa8, 0x44, 0x3c, 0xfc, 0x9c, 0x3c, 0xfd, 0x49, 0xee, 0xfe, 0xf8, 0x27,
	0x00, 0x00, 0xff, 0xff, 0x79, 0xc7, 0x24, 0xc4, 0x1d, 0x03, 0x00, 0x00,
}
