// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0
// source: common/mgrpc/account/account.proto

package account

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AccountProto_GetAccountsByIds_FullMethodName     = "/account.AccountProto/GetAccountsByIds"
	AccountProto_GetAccountsByAliases_FullMethodName = "/account.AccountProto/GetAccountsByAliases"
	AccountProto_UpdateAccounts_FullMethodName       = "/account.AccountProto/UpdateAccounts"
)

// AccountProtoClient is the client API for AccountProto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountProtoClient interface {
	GetAccountsByIds(ctx context.Context, in *AccountsID, opts ...grpc.CallOption) (*AccountsResponse, error)
	GetAccountsByAliases(ctx context.Context, in *AccountsAlias, opts ...grpc.CallOption) (*AccountsResponse, error)
	UpdateAccounts(ctx context.Context, in *AccountsRequest, opts ...grpc.CallOption) (*AccountsResponse, error)
}

type accountProtoClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountProtoClient(cc grpc.ClientConnInterface) AccountProtoClient {
	return &accountProtoClient{cc}
}

func (c *accountProtoClient) GetAccountsByIds(ctx context.Context, in *AccountsID, opts ...grpc.CallOption) (*AccountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountsResponse)
	err := c.cc.Invoke(ctx, AccountProto_GetAccountsByIds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountProtoClient) GetAccountsByAliases(ctx context.Context, in *AccountsAlias, opts ...grpc.CallOption) (*AccountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountsResponse)
	err := c.cc.Invoke(ctx, AccountProto_GetAccountsByAliases_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountProtoClient) UpdateAccounts(ctx context.Context, in *AccountsRequest, opts ...grpc.CallOption) (*AccountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountsResponse)
	err := c.cc.Invoke(ctx, AccountProto_UpdateAccounts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountProtoServer is the server API for AccountProto service.
// All implementations must embed UnimplementedAccountProtoServer
// for forward compatibility.
type AccountProtoServer interface {
	GetAccountsByIds(context.Context, *AccountsID) (*AccountsResponse, error)
	GetAccountsByAliases(context.Context, *AccountsAlias) (*AccountsResponse, error)
	UpdateAccounts(context.Context, *AccountsRequest) (*AccountsResponse, error)
	mustEmbedUnimplementedAccountProtoServer()
}

// UnimplementedAccountProtoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAccountProtoServer struct{}

func (UnimplementedAccountProtoServer) GetAccountsByIds(context.Context, *AccountsID) (*AccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountsByIds not implemented")
}
func (UnimplementedAccountProtoServer) GetAccountsByAliases(context.Context, *AccountsAlias) (*AccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountsByAliases not implemented")
}
func (UnimplementedAccountProtoServer) UpdateAccounts(context.Context, *AccountsRequest) (*AccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccounts not implemented")
}
func (UnimplementedAccountProtoServer) mustEmbedUnimplementedAccountProtoServer() {}
func (UnimplementedAccountProtoServer) testEmbeddedByValue()                      {}

// UnsafeAccountProtoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountProtoServer will
// result in compilation errors.
type UnsafeAccountProtoServer interface {
	mustEmbedUnimplementedAccountProtoServer()
}

func RegisterAccountProtoServer(s grpc.ServiceRegistrar, srv AccountProtoServer) {
	// If the following call pancis, it indicates UnimplementedAccountProtoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AccountProto_ServiceDesc, srv)
}

func _AccountProto_GetAccountsByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountsID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountProtoServer).GetAccountsByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountProto_GetAccountsByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountProtoServer).GetAccountsByIds(ctx, req.(*AccountsID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountProto_GetAccountsByAliases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountsAlias)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountProtoServer).GetAccountsByAliases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountProto_GetAccountsByAliases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountProtoServer).GetAccountsByAliases(ctx, req.(*AccountsAlias))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountProto_UpdateAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountProtoServer).UpdateAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountProto_UpdateAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountProtoServer).UpdateAccounts(ctx, req.(*AccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountProto_ServiceDesc is the grpc.ServiceDesc for AccountProto service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountProto_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.AccountProto",
	HandlerType: (*AccountProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountsByIds",
			Handler:    _AccountProto_GetAccountsByIds_Handler,
		},
		{
			MethodName: "GetAccountsByAliases",
			Handler:    _AccountProto_GetAccountsByAliases_Handler,
		},
		{
			MethodName: "UpdateAccounts",
			Handler:    _AccountProto_UpdateAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/mgrpc/account/account.proto",
}
