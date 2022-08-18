// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: pkg/musician/musician.proto

package pb

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

// MusicianServiceClient is the client API for MusicianService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MusicianServiceClient interface {
	FindOne(ctx context.Context, in *FindOneRequest, opts ...grpc.CallOption) (*FindOneResponse, error)
	CreateMusician(ctx context.Context, in *MusicianData, opts ...grpc.CallOption) (*FindOneResponse, error)
	UpdateMusician(ctx context.Context, in *MusicianData, opts ...grpc.CallOption) (*FindOneResponse, error)
	DeleteMusician(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	SearchMusician(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type musicianServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMusicianServiceClient(cc grpc.ClientConnInterface) MusicianServiceClient {
	return &musicianServiceClient{cc}
}

func (c *musicianServiceClient) FindOne(ctx context.Context, in *FindOneRequest, opts ...grpc.CallOption) (*FindOneResponse, error) {
	out := new(FindOneResponse)
	err := c.cc.Invoke(ctx, "/auth.MusicianService/FindOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicianServiceClient) CreateMusician(ctx context.Context, in *MusicianData, opts ...grpc.CallOption) (*FindOneResponse, error) {
	out := new(FindOneResponse)
	err := c.cc.Invoke(ctx, "/auth.MusicianService/CreateMusician", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicianServiceClient) UpdateMusician(ctx context.Context, in *MusicianData, opts ...grpc.CallOption) (*FindOneResponse, error) {
	out := new(FindOneResponse)
	err := c.cc.Invoke(ctx, "/auth.MusicianService/UpdateMusician", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicianServiceClient) DeleteMusician(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/auth.MusicianService/DeleteMusician", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicianServiceClient) SearchMusician(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/auth.MusicianService/SearchMusician", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MusicianServiceServer is the server API for MusicianService service.
// All implementations must embed UnimplementedMusicianServiceServer
// for forward compatibility
type MusicianServiceServer interface {
	FindOne(context.Context, *FindOneRequest) (*FindOneResponse, error)
	CreateMusician(context.Context, *MusicianData) (*FindOneResponse, error)
	UpdateMusician(context.Context, *MusicianData) (*FindOneResponse, error)
	DeleteMusician(context.Context, *DeleteRequest) (*DeleteResponse, error)
	SearchMusician(context.Context, *SearchRequest) (*SearchResponse, error)
	mustEmbedUnimplementedMusicianServiceServer()
}

// UnimplementedMusicianServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMusicianServiceServer struct {
}

func (UnimplementedMusicianServiceServer) FindOne(context.Context, *FindOneRequest) (*FindOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (UnimplementedMusicianServiceServer) CreateMusician(context.Context, *MusicianData) (*FindOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMusician not implemented")
}
func (UnimplementedMusicianServiceServer) UpdateMusician(context.Context, *MusicianData) (*FindOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMusician not implemented")
}
func (UnimplementedMusicianServiceServer) DeleteMusician(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMusician not implemented")
}
func (UnimplementedMusicianServiceServer) SearchMusician(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMusician not implemented")
}
func (UnimplementedMusicianServiceServer) mustEmbedUnimplementedMusicianServiceServer() {}

// UnsafeMusicianServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MusicianServiceServer will
// result in compilation errors.
type UnsafeMusicianServiceServer interface {
	mustEmbedUnimplementedMusicianServiceServer()
}

func RegisterMusicianServiceServer(s grpc.ServiceRegistrar, srv MusicianServiceServer) {
	s.RegisterService(&MusicianService_ServiceDesc, srv)
}

func _MusicianService_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicianServiceServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.MusicianService/FindOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicianServiceServer).FindOne(ctx, req.(*FindOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicianService_CreateMusician_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MusicianData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicianServiceServer).CreateMusician(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.MusicianService/CreateMusician",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicianServiceServer).CreateMusician(ctx, req.(*MusicianData))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicianService_UpdateMusician_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MusicianData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicianServiceServer).UpdateMusician(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.MusicianService/UpdateMusician",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicianServiceServer).UpdateMusician(ctx, req.(*MusicianData))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicianService_DeleteMusician_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicianServiceServer).DeleteMusician(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.MusicianService/DeleteMusician",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicianServiceServer).DeleteMusician(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicianService_SearchMusician_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicianServiceServer).SearchMusician(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.MusicianService/SearchMusician",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicianServiceServer).SearchMusician(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MusicianService_ServiceDesc is the grpc.ServiceDesc for MusicianService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MusicianService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.MusicianService",
	HandlerType: (*MusicianServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindOne",
			Handler:    _MusicianService_FindOne_Handler,
		},
		{
			MethodName: "CreateMusician",
			Handler:    _MusicianService_CreateMusician_Handler,
		},
		{
			MethodName: "UpdateMusician",
			Handler:    _MusicianService_UpdateMusician_Handler,
		},
		{
			MethodName: "DeleteMusician",
			Handler:    _MusicianService_DeleteMusician_Handler,
		},
		{
			MethodName: "SearchMusician",
			Handler:    _MusicianService_SearchMusician_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/musician/musician.proto",
}
