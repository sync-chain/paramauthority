// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: noble/authority/v1/query.proto

package authorityv1

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

const (
	Query_Authority_FullMethodName = "/noble.authority.v1.Query/Authority"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	Authority(ctx context.Context, in *QueryAuthority, opts ...grpc.CallOption) (*QueryAuthorityResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Authority(ctx context.Context, in *QueryAuthority, opts ...grpc.CallOption) (*QueryAuthorityResponse, error) {
	out := new(QueryAuthorityResponse)
	err := c.cc.Invoke(ctx, Query_Authority_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	Authority(context.Context, *QueryAuthority) (*QueryAuthorityResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Authority(context.Context, *QueryAuthority) (*QueryAuthorityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authority not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Authority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAuthority)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Authority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Authority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Authority(ctx, req.(*QueryAuthority))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "noble.authority.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authority",
			Handler:    _Query_Authority_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "noble/authority/v1/query.proto",
}
