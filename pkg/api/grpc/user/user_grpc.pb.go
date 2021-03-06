// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserStorageClient is the client API for UserStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserStorageClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetUserByIdent(ctx context.Context, in *GetUserByIdentRequest, opts ...grpc.CallOption) (*GetUserByIdentResponse, error)
}

type userStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewUserStorageClient(cc grpc.ClientConnInterface) UserStorageClient {
	return &userStorageClient{cc}
}

func (c *userStorageClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserStorage/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStorageClient) GetUserByIdent(ctx context.Context, in *GetUserByIdentRequest, opts ...grpc.CallOption) (*GetUserByIdentResponse, error) {
	out := new(GetUserByIdentResponse)
	err := c.cc.Invoke(ctx, "/user.UserStorage/GetUserByIdent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserStorageServer is the server API for UserStorage service.
// All implementations must embed UnimplementedUserStorageServer
// for forward compatibility
type UserStorageServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUserByIdent(context.Context, *GetUserByIdentRequest) (*GetUserByIdentResponse, error)
	mustEmbedUnimplementedUserStorageServer()
}

// UnimplementedUserStorageServer must be embedded to have forward compatible implementations.
type UnimplementedUserStorageServer struct {
}

func (UnimplementedUserStorageServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserStorageServer) GetUserByIdent(context.Context, *GetUserByIdentRequest) (*GetUserByIdentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByIdent not implemented")
}
func (UnimplementedUserStorageServer) mustEmbedUnimplementedUserStorageServer() {}

// UnsafeUserStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserStorageServer will
// result in compilation errors.
type UnsafeUserStorageServer interface {
	mustEmbedUnimplementedUserStorageServer()
}

func RegisterUserStorageServer(s grpc.ServiceRegistrar, srv UserStorageServer) {
	s.RegisterService(&UserStorage_ServiceDesc, srv)
}

func _UserStorage_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserStorageServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserStorage/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserStorageServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserStorage_GetUserByIdent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserStorageServer).GetUserByIdent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserStorage/GetUserByIdent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserStorageServer).GetUserByIdent(ctx, req.(*GetUserByIdentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserStorage_ServiceDesc is the grpc.ServiceDesc for UserStorage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserStorage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserStorage",
	HandlerType: (*UserStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserStorage_CreateUser_Handler,
		},
		{
			MethodName: "GetUserByIdent",
			Handler:    _UserStorage_GetUserByIdent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/user/user.proto",
}
