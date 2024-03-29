// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: pkg/rating/rating.proto

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

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingServiceClient interface {
	FindByUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*RatingResponse, error)
	FindBySong(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*RatingResponse, error)
	FindByAlbum(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*RatingResponse, error)
	FindByUserSong(ctx context.Context, in *UserSongRequest, opts ...grpc.CallOption) (*RatingData, error)
	FindByUserAlbum(ctx context.Context, in *UserAlbumRequest, opts ...grpc.CallOption) (*RatingData, error)
	RateSong(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error)
	RateAlbum(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error)
	Rate(ctx context.Context, in *NewRateRequest, opts ...grpc.CallOption) (*RateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*RatingResponse, error)
}

type ratingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingServiceClient(cc grpc.ClientConnInterface) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) FindByUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*RatingResponse, error) {
	out := new(RatingResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/FindByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) FindBySong(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*RatingResponse, error) {
	out := new(RatingResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/FindBySong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) FindByAlbum(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*RatingResponse, error) {
	out := new(RatingResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/FindByAlbum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) FindByUserSong(ctx context.Context, in *UserSongRequest, opts ...grpc.CallOption) (*RatingData, error) {
	out := new(RatingData)
	err := c.cc.Invoke(ctx, "/rating.RatingService/FindByUserSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) FindByUserAlbum(ctx context.Context, in *UserAlbumRequest, opts ...grpc.CallOption) (*RatingData, error) {
	out := new(RatingData)
	err := c.cc.Invoke(ctx, "/rating.RatingService/FindByUserAlbum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) RateSong(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/RateSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) RateAlbum(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/RateAlbum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) Rate(ctx context.Context, in *NewRateRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/Rate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*RatingResponse, error) {
	out := new(RatingResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
// All implementations must embed UnimplementedRatingServiceServer
// for forward compatibility
type RatingServiceServer interface {
	FindByUser(context.Context, *IdRequest) (*RatingResponse, error)
	FindBySong(context.Context, *IdRequest) (*RatingResponse, error)
	FindByAlbum(context.Context, *IdRequest) (*RatingResponse, error)
	FindByUserSong(context.Context, *UserSongRequest) (*RatingData, error)
	FindByUserAlbum(context.Context, *UserAlbumRequest) (*RatingData, error)
	RateSong(context.Context, *RateRequest) (*RateResponse, error)
	RateAlbum(context.Context, *RateRequest) (*RateResponse, error)
	Rate(context.Context, *NewRateRequest) (*RateResponse, error)
	Delete(context.Context, *DeleteRequest) (*RatingResponse, error)
	mustEmbedUnimplementedRatingServiceServer()
}

// UnimplementedRatingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRatingServiceServer struct {
}

func (UnimplementedRatingServiceServer) FindByUser(context.Context, *IdRequest) (*RatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByUser not implemented")
}
func (UnimplementedRatingServiceServer) FindBySong(context.Context, *IdRequest) (*RatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBySong not implemented")
}
func (UnimplementedRatingServiceServer) FindByAlbum(context.Context, *IdRequest) (*RatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByAlbum not implemented")
}
func (UnimplementedRatingServiceServer) FindByUserSong(context.Context, *UserSongRequest) (*RatingData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByUserSong not implemented")
}
func (UnimplementedRatingServiceServer) FindByUserAlbum(context.Context, *UserAlbumRequest) (*RatingData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByUserAlbum not implemented")
}
func (UnimplementedRatingServiceServer) RateSong(context.Context, *RateRequest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateSong not implemented")
}
func (UnimplementedRatingServiceServer) RateAlbum(context.Context, *RateRequest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateAlbum not implemented")
}
func (UnimplementedRatingServiceServer) Rate(context.Context, *NewRateRequest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rate not implemented")
}
func (UnimplementedRatingServiceServer) Delete(context.Context, *DeleteRequest) (*RatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRatingServiceServer) mustEmbedUnimplementedRatingServiceServer() {}

// UnsafeRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServiceServer will
// result in compilation errors.
type UnsafeRatingServiceServer interface {
	mustEmbedUnimplementedRatingServiceServer()
}

func RegisterRatingServiceServer(s grpc.ServiceRegistrar, srv RatingServiceServer) {
	s.RegisterService(&RatingService_ServiceDesc, srv)
}

func _RatingService_FindByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).FindByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/FindByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).FindByUser(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_FindBySong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).FindBySong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/FindBySong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).FindBySong(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_FindByAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).FindByAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/FindByAlbum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).FindByAlbum(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_FindByUserSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).FindByUserSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/FindByUserSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).FindByUserSong(ctx, req.(*UserSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_FindByUserAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAlbumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).FindByUserAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/FindByUserAlbum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).FindByUserAlbum(ctx, req.(*UserAlbumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_RateSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).RateSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/RateSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).RateSong(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_RateAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).RateAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/RateAlbum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).RateAlbum(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_Rate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).Rate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/Rate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).Rate(ctx, req.(*NewRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingService_ServiceDesc is the grpc.ServiceDesc for RatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rating.RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindByUser",
			Handler:    _RatingService_FindByUser_Handler,
		},
		{
			MethodName: "FindBySong",
			Handler:    _RatingService_FindBySong_Handler,
		},
		{
			MethodName: "FindByAlbum",
			Handler:    _RatingService_FindByAlbum_Handler,
		},
		{
			MethodName: "FindByUserSong",
			Handler:    _RatingService_FindByUserSong_Handler,
		},
		{
			MethodName: "FindByUserAlbum",
			Handler:    _RatingService_FindByUserAlbum_Handler,
		},
		{
			MethodName: "RateSong",
			Handler:    _RatingService_RateSong_Handler,
		},
		{
			MethodName: "RateAlbum",
			Handler:    _RatingService_RateAlbum_Handler,
		},
		{
			MethodName: "Rate",
			Handler:    _RatingService_Rate_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RatingService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/rating/rating.proto",
}
