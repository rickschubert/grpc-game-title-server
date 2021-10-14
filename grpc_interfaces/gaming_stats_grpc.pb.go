// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpctaster

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

// GamingStatsClient is the client API for GamingStats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GamingStatsClient interface {
	GetGame(ctx context.Context, in *Game, opts ...grpc.CallOption) (*GameStats, error)
}

type gamingStatsClient struct {
	cc grpc.ClientConnInterface
}

func NewGamingStatsClient(cc grpc.ClientConnInterface) GamingStatsClient {
	return &gamingStatsClient{cc}
}

func (c *gamingStatsClient) GetGame(ctx context.Context, in *Game, opts ...grpc.CallOption) (*GameStats, error) {
	out := new(GameStats)
	err := c.cc.Invoke(ctx, "/grpctaster.GamingStats/GetGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GamingStatsServer is the server API for GamingStats service.
// All implementations must embed UnimplementedGamingStatsServer
// for forward compatibility
type GamingStatsServer interface {
	GetGame(context.Context, *Game) (*GameStats, error)
	mustEmbedUnimplementedGamingStatsServer()
}

// UnimplementedGamingStatsServer must be embedded to have forward compatible implementations.
type UnimplementedGamingStatsServer struct {
}

func (UnimplementedGamingStatsServer) GetGame(context.Context, *Game) (*GameStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGame not implemented")
}
func (UnimplementedGamingStatsServer) mustEmbedUnimplementedGamingStatsServer() {}

// UnsafeGamingStatsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GamingStatsServer will
// result in compilation errors.
type UnsafeGamingStatsServer interface {
	mustEmbedUnimplementedGamingStatsServer()
}

func RegisterGamingStatsServer(s grpc.ServiceRegistrar, srv GamingStatsServer) {
	s.RegisterService(&GamingStats_ServiceDesc, srv)
}

func _GamingStats_GetGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Game)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GamingStatsServer).GetGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpctaster.GamingStats/GetGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GamingStatsServer).GetGame(ctx, req.(*Game))
	}
	return interceptor(ctx, in, info, handler)
}

// GamingStats_ServiceDesc is the grpc.ServiceDesc for GamingStats service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GamingStats_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpctaster.GamingStats",
	HandlerType: (*GamingStatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGame",
			Handler:    _GamingStats_GetGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gaming_stats.proto",
}
