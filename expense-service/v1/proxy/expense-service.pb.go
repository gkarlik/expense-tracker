// Code generated by protoc-gen-go.
// source: expense-service/v1/expense-service.proto
// DO NOT EDIT!

/*
Package proxy is a generated protocol buffer package.

It is generated from these files:
	expense-service/v1/expense-service.proto

It has these top-level messages:
	UserPagingRequest
	ExpensesResponse
	CategoriesResponse
	ExpenseRequest
	ExpenseResponse
	ExpenseIDRequest
	CategoryRequest
	CategoryResponse
	CategoryIDRequest
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

type UserPagingRequest struct {
	UserID string `protobuf:"bytes,1,opt,name=UserID" json:"UserID,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=Offset" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,3,opt,name=Limit" json:"Limit,omitempty"`
}

func (m *UserPagingRequest) Reset()                    { *m = UserPagingRequest{} }
func (m *UserPagingRequest) String() string            { return proto.CompactTextString(m) }
func (*UserPagingRequest) ProtoMessage()               {}
func (*UserPagingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserPagingRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UserPagingRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *UserPagingRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ExpensesResponse struct {
	Expenses []*ExpenseResponse `protobuf:"bytes,1,rep,name=Expenses" json:"Expenses,omitempty"`
}

func (m *ExpensesResponse) Reset()                    { *m = ExpensesResponse{} }
func (m *ExpensesResponse) String() string            { return proto.CompactTextString(m) }
func (*ExpensesResponse) ProtoMessage()               {}
func (*ExpensesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ExpensesResponse) GetExpenses() []*ExpenseResponse {
	if m != nil {
		return m.Expenses
	}
	return nil
}

type CategoriesResponse struct {
	Categories []*CategoryResponse `protobuf:"bytes,1,rep,name=Categories" json:"Categories,omitempty"`
}

func (m *CategoriesResponse) Reset()                    { *m = CategoriesResponse{} }
func (m *CategoriesResponse) String() string            { return proto.CompactTextString(m) }
func (*CategoriesResponse) ProtoMessage()               {}
func (*CategoriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CategoriesResponse) GetCategories() []*CategoryResponse {
	if m != nil {
		return m.Categories
	}
	return nil
}

type ExpenseRequest struct {
	ID         string  `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Date       int64   `protobuf:"varint,2,opt,name=Date" json:"Date,omitempty"`
	Value      float32 `protobuf:"fixed32,3,opt,name=Value" json:"Value,omitempty"`
	CategoryID string  `protobuf:"bytes,4,opt,name=CategoryID" json:"CategoryID,omitempty"`
	UserID     string  `protobuf:"bytes,5,opt,name=UserID" json:"UserID,omitempty"`
}

func (m *ExpenseRequest) Reset()                    { *m = ExpenseRequest{} }
func (m *ExpenseRequest) String() string            { return proto.CompactTextString(m) }
func (*ExpenseRequest) ProtoMessage()               {}
func (*ExpenseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ExpenseRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ExpenseRequest) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *ExpenseRequest) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *ExpenseRequest) GetCategoryID() string {
	if m != nil {
		return m.CategoryID
	}
	return ""
}

func (m *ExpenseRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type ExpenseResponse struct {
	ID         string  `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Date       int64   `protobuf:"varint,2,opt,name=Date" json:"Date,omitempty"`
	Value      float32 `protobuf:"fixed32,3,opt,name=Value" json:"Value,omitempty"`
	CategoryID string  `protobuf:"bytes,4,opt,name=CategoryID" json:"CategoryID,omitempty"`
}

func (m *ExpenseResponse) Reset()                    { *m = ExpenseResponse{} }
func (m *ExpenseResponse) String() string            { return proto.CompactTextString(m) }
func (*ExpenseResponse) ProtoMessage()               {}
func (*ExpenseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ExpenseResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ExpenseResponse) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *ExpenseResponse) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *ExpenseResponse) GetCategoryID() string {
	if m != nil {
		return m.CategoryID
	}
	return ""
}

type ExpenseIDRequest struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *ExpenseIDRequest) Reset()                    { *m = ExpenseIDRequest{} }
func (m *ExpenseIDRequest) String() string            { return proto.CompactTextString(m) }
func (*ExpenseIDRequest) ProtoMessage()               {}
func (*ExpenseIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ExpenseIDRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type CategoryRequest struct {
	ID     string  `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Limit  float32 `protobuf:"fixed32,2,opt,name=Limit" json:"Limit,omitempty"`
	Name   string  `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	UserID string  `protobuf:"bytes,4,opt,name=UserID" json:"UserID,omitempty"`
}

func (m *CategoryRequest) Reset()                    { *m = CategoryRequest{} }
func (m *CategoryRequest) String() string            { return proto.CompactTextString(m) }
func (*CategoryRequest) ProtoMessage()               {}
func (*CategoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CategoryRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *CategoryRequest) GetLimit() float32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *CategoryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CategoryRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type CategoryResponse struct {
	ID    string  `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Limit float32 `protobuf:"fixed32,2,opt,name=Limit" json:"Limit,omitempty"`
	Name  string  `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
}

func (m *CategoryResponse) Reset()                    { *m = CategoryResponse{} }
func (m *CategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*CategoryResponse) ProtoMessage()               {}
func (*CategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CategoryResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *CategoryResponse) GetLimit() float32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *CategoryResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CategoryIDRequest struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *CategoryIDRequest) Reset()                    { *m = CategoryIDRequest{} }
func (m *CategoryIDRequest) String() string            { return proto.CompactTextString(m) }
func (*CategoryIDRequest) ProtoMessage()               {}
func (*CategoryIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CategoryIDRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type EmptyResponse struct {
}

func (m *EmptyResponse) Reset()                    { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string            { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()               {}
func (*EmptyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func init() {
	proto.RegisterType((*UserPagingRequest)(nil), "UserPagingRequest")
	proto.RegisterType((*ExpensesResponse)(nil), "ExpensesResponse")
	proto.RegisterType((*CategoriesResponse)(nil), "CategoriesResponse")
	proto.RegisterType((*ExpenseRequest)(nil), "ExpenseRequest")
	proto.RegisterType((*ExpenseResponse)(nil), "ExpenseResponse")
	proto.RegisterType((*ExpenseIDRequest)(nil), "ExpenseIDRequest")
	proto.RegisterType((*CategoryRequest)(nil), "CategoryRequest")
	proto.RegisterType((*CategoryResponse)(nil), "CategoryResponse")
	proto.RegisterType((*CategoryIDRequest)(nil), "CategoryIDRequest")
	proto.RegisterType((*EmptyResponse)(nil), "EmptyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ExpenseService service

type ExpenseServiceClient interface {
	GetExpense(ctx context.Context, in *ExpenseIDRequest, opts ...grpc.CallOption) (*ExpenseResponse, error)
	UpdateExpense(ctx context.Context, in *ExpenseRequest, opts ...grpc.CallOption) (*ExpenseResponse, error)
	RemoveExpense(ctx context.Context, in *ExpenseIDRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetUserExpenses(ctx context.Context, in *UserPagingRequest, opts ...grpc.CallOption) (*ExpensesResponse, error)
	GetCategory(ctx context.Context, in *CategoryIDRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	UpdateCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	RemoveCategory(ctx context.Context, in *CategoryIDRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetUserCategories(ctx context.Context, in *UserPagingRequest, opts ...grpc.CallOption) (*CategoriesResponse, error)
}

type expenseServiceClient struct {
	cc *grpc.ClientConn
}

func NewExpenseServiceClient(cc *grpc.ClientConn) ExpenseServiceClient {
	return &expenseServiceClient{cc}
}

func (c *expenseServiceClient) GetExpense(ctx context.Context, in *ExpenseIDRequest, opts ...grpc.CallOption) (*ExpenseResponse, error) {
	out := new(ExpenseResponse)
	err := grpc.Invoke(ctx, "/ExpenseService/GetExpense", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) UpdateExpense(ctx context.Context, in *ExpenseRequest, opts ...grpc.CallOption) (*ExpenseResponse, error) {
	out := new(ExpenseResponse)
	err := grpc.Invoke(ctx, "/ExpenseService/UpdateExpense", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) RemoveExpense(ctx context.Context, in *ExpenseIDRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/ExpenseService/RemoveExpense", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) GetUserExpenses(ctx context.Context, in *UserPagingRequest, opts ...grpc.CallOption) (*ExpensesResponse, error) {
	out := new(ExpensesResponse)
	err := grpc.Invoke(ctx, "/ExpenseService/GetUserExpenses", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) GetCategory(ctx context.Context, in *CategoryIDRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	out := new(CategoryResponse)
	err := grpc.Invoke(ctx, "/ExpenseService/GetCategory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) UpdateCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	out := new(CategoryResponse)
	err := grpc.Invoke(ctx, "/ExpenseService/UpdateCategory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) RemoveCategory(ctx context.Context, in *CategoryIDRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/ExpenseService/RemoveCategory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) GetUserCategories(ctx context.Context, in *UserPagingRequest, opts ...grpc.CallOption) (*CategoriesResponse, error) {
	out := new(CategoriesResponse)
	err := grpc.Invoke(ctx, "/ExpenseService/GetUserCategories", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ExpenseService service

type ExpenseServiceServer interface {
	GetExpense(context.Context, *ExpenseIDRequest) (*ExpenseResponse, error)
	UpdateExpense(context.Context, *ExpenseRequest) (*ExpenseResponse, error)
	RemoveExpense(context.Context, *ExpenseIDRequest) (*EmptyResponse, error)
	GetUserExpenses(context.Context, *UserPagingRequest) (*ExpensesResponse, error)
	GetCategory(context.Context, *CategoryIDRequest) (*CategoryResponse, error)
	UpdateCategory(context.Context, *CategoryRequest) (*CategoryResponse, error)
	RemoveCategory(context.Context, *CategoryIDRequest) (*EmptyResponse, error)
	GetUserCategories(context.Context, *UserPagingRequest) (*CategoriesResponse, error)
}

func RegisterExpenseServiceServer(s *grpc.Server, srv ExpenseServiceServer) {
	s.RegisterService(&_ExpenseService_serviceDesc, srv)
}

func _ExpenseService_GetExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpenseIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).GetExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpenseService/GetExpense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).GetExpense(ctx, req.(*ExpenseIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_UpdateExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).UpdateExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpenseService/UpdateExpense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).UpdateExpense(ctx, req.(*ExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_RemoveExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpenseIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).RemoveExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpenseService/RemoveExpense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).RemoveExpense(ctx, req.(*ExpenseIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_GetUserExpenses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).GetUserExpenses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpenseService/GetUserExpenses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).GetUserExpenses(ctx, req.(*UserPagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpenseService/GetCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).GetCategory(ctx, req.(*CategoryIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpenseService/UpdateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).UpdateCategory(ctx, req.(*CategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_RemoveCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).RemoveCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpenseService/RemoveCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).RemoveCategory(ctx, req.(*CategoryIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_GetUserCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).GetUserCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpenseService/GetUserCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).GetUserCategories(ctx, req.(*UserPagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExpenseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ExpenseService",
	HandlerType: (*ExpenseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExpense",
			Handler:    _ExpenseService_GetExpense_Handler,
		},
		{
			MethodName: "UpdateExpense",
			Handler:    _ExpenseService_UpdateExpense_Handler,
		},
		{
			MethodName: "RemoveExpense",
			Handler:    _ExpenseService_RemoveExpense_Handler,
		},
		{
			MethodName: "GetUserExpenses",
			Handler:    _ExpenseService_GetUserExpenses_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _ExpenseService_GetCategory_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _ExpenseService_UpdateCategory_Handler,
		},
		{
			MethodName: "RemoveCategory",
			Handler:    _ExpenseService_RemoveCategory_Handler,
		},
		{
			MethodName: "GetUserCategories",
			Handler:    _ExpenseService_GetUserCategories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "expense-service/v1/expense-service.proto",
}

func init() { proto.RegisterFile("expense-service/v1/expense-service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x94, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x9b, 0xb4, 0x1d, 0xec, 0x4d, 0x4b, 0x9a, 0x07, 0x42, 0x51, 0x0f, 0xa8, 0x32, 0x97,
	0x1e, 0xc0, 0xd3, 0x36, 0x34, 0x24, 0x0e, 0x08, 0x41, 0xa6, 0xaa, 0xd2, 0x04, 0xc8, 0x68, 0x48,
	0x70, 0x0b, 0xe3, 0xad, 0x8a, 0xa0, 0x4d, 0x88, 0xbd, 0x6a, 0xbd, 0xf2, 0x47, 0xf2, 0xf7, 0xa0,
	0xd8, 0x89, 0x9b, 0x26, 0xed, 0x0e, 0x48, 0xbb, 0xe5, 0x7d, 0xf2, 0x7b, 0x9f, 0x7f, 0x9f, 0xed,
	0xc0, 0x98, 0x6e, 0x33, 0x5a, 0x48, 0x7a, 0x21, 0x29, 0x5f, 0x26, 0x57, 0x74, 0xb4, 0x3c, 0x3e,
	0x6a, 0x48, 0x3c, 0xcb, 0x53, 0x95, 0xb2, 0xaf, 0x10, 0x5c, 0x4a, 0xca, 0x3f, 0xc5, 0xb3, 0x64,
	0x31, 0x13, 0xf4, 0xfb, 0x86, 0xa4, 0xc2, 0x27, 0xb0, 0x57, 0x88, 0xd3, 0x28, 0x74, 0x46, 0xce,
	0x78, 0x5f, 0x94, 0x55, 0xa1, 0x7f, 0xbc, 0xbe, 0x96, 0xa4, 0x42, 0x77, 0xe4, 0x8c, 0xfb, 0xa2,
	0xac, 0xf0, 0x31, 0xf4, 0x2f, 0x92, 0x79, 0xa2, 0xc2, 0xae, 0x96, 0x4d, 0xc1, 0xde, 0xc2, 0xe0,
	0xdc, 0x78, 0x4a, 0x41, 0x32, 0x4b, 0x17, 0x92, 0xf0, 0x39, 0x3c, 0xac, 0xb4, 0xd0, 0x19, 0x75,
	0xc7, 0x07, 0x27, 0x03, 0x5e, 0x0a, 0xd5, 0x1a, 0x61, 0x57, 0xb0, 0x09, 0xe0, 0xfb, 0x58, 0xd1,
	0x2c, 0xcd, 0x93, 0xda, 0x8c, 0x63, 0x80, 0xb5, 0x5a, 0x4e, 0x09, 0x78, 0x29, 0xad, 0xec, 0x98,
	0xda, 0x22, 0xf6, 0xc7, 0x01, 0xcf, 0xda, 0x18, 0x46, 0x0f, 0x5c, 0xcb, 0xe7, 0x4e, 0x23, 0x44,
	0xe8, 0x45, 0xb1, 0x22, 0x4d, 0xd6, 0x15, 0xfa, 0xbb, 0xe0, 0xfa, 0x12, 0xff, 0xba, 0x21, 0xcd,
	0xe5, 0x0a, 0x53, 0xe0, 0x53, 0xeb, 0xbf, 0x9a, 0x46, 0x61, 0x4f, 0x4f, 0xa8, 0x29, 0xb5, 0xf4,
	0xfa, 0xf5, 0xf4, 0xd8, 0x4f, 0xf0, 0x1b, 0xa8, 0xf7, 0xb7, 0x09, 0xc6, 0x6c, 0xf8, 0xd3, 0x68,
	0x07, 0x32, 0xbb, 0x02, 0x7f, 0x9d, 0xda, 0xf6, 0x54, 0xec, 0xc9, 0xba, 0xc6, 0x5c, 0x17, 0xc5,
	0x36, 0x3f, 0xc4, 0x73, 0xb3, 0xa3, 0x7d, 0xa1, 0xbf, 0x6b, 0xd4, 0xbd, 0x0d, 0xea, 0x0b, 0x18,
	0x34, 0x8f, 0xe6, 0xff, 0x5d, 0xd8, 0x33, 0x08, 0xd6, 0x90, 0xbb, 0xb8, 0x7c, 0x38, 0x3c, 0x9f,
	0x67, 0xca, 0xfa, 0x9d, 0xfc, 0xed, 0xda, 0xe3, 0xff, 0x6c, 0x6e, 0x3f, 0x9e, 0x02, 0x4c, 0x48,
	0x95, 0x22, 0x06, 0xbc, 0x19, 0xd6, 0xb0, 0x75, 0x2f, 0x59, 0x07, 0x5f, 0xc2, 0xe1, 0x65, 0xf6,
	0x23, 0x56, 0x54, 0xf5, 0xf9, 0x7c, 0xf3, 0x56, 0xed, 0xea, 0x12, 0x34, 0x4f, 0x97, 0x74, 0x87,
	0x9b, 0xc7, 0x37, 0x76, 0xcc, 0x3a, 0xf8, 0x1a, 0xfc, 0x09, 0xa9, 0x22, 0xc4, 0xea, 0x39, 0x20,
	0xf2, 0xd6, 0x53, 0x1d, 0xda, 0x59, 0xb2, 0xd6, 0x7b, 0x06, 0x07, 0x13, 0x52, 0x55, 0x50, 0x88,
	0xbc, 0x95, 0xd9, 0xb0, 0xfd, 0x60, 0x58, 0x07, 0x5f, 0x81, 0x67, 0xf8, 0x6c, 0xeb, 0x80, 0x37,
	0x6e, 0xc8, 0xf6, 0xc6, 0x33, 0xf0, 0x0c, 0xe2, 0x9d, 0x9e, 0x6d, 0xc8, 0x37, 0x10, 0x94, 0x90,
	0xeb, 0xc7, 0xba, 0x15, 0xf3, 0x11, 0x6f, 0xff, 0x08, 0x58, 0xe7, 0xdd, 0x83, 0x6f, 0xfd, 0x2c,
	0x4f, 0x6f, 0x57, 0xdf, 0xf7, 0xf4, 0xdf, 0xec, 0xf4, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75,
	0xc6, 0xec, 0x3c, 0xf9, 0x04, 0x00, 0x00,
}
