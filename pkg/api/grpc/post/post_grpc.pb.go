// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package post

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

// PostStorageClient is the client API for PostStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostStorageClient interface {
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
}

type postStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewPostStorageClient(cc grpc.ClientConnInterface) PostStorageClient {
	return &postStorageClient{cc}
}

func (c *postStorageClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, "/post.PostStorage/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostStorageServer is the server API for PostStorage service.
// All implementations must embed UnimplementedPostStorageServer
// for forward compatibility
type PostStorageServer interface {
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	mustEmbedUnimplementedPostStorageServer()
}

// UnimplementedPostStorageServer must be embedded to have forward compatible implementations.
type UnimplementedPostStorageServer struct {
}

func (UnimplementedPostStorageServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedPostStorageServer) mustEmbedUnimplementedPostStorageServer() {}

// UnsafePostStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostStorageServer will
// result in compilation errors.
type UnsafePostStorageServer interface {
	mustEmbedUnimplementedPostStorageServer()
}

func RegisterPostStorageServer(s grpc.ServiceRegistrar, srv PostStorageServer) {
	s.RegisterService(&PostStorage_ServiceDesc, srv)
}

func _PostStorage_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostStorageServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostStorage/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostStorageServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostStorage_ServiceDesc is the grpc.ServiceDesc for PostStorage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostStorage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "post.PostStorage",
	HandlerType: (*PostStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _PostStorage_CreatePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/post/post.proto",
}
