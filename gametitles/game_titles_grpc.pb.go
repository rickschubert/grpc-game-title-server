// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gametitles

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

// GameTitlesClient is the client API for GameTitles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameTitlesClient interface {
	GetGame(ctx context.Context, in *GameRequest, opts ...grpc.CallOption) (*GameStats, error)
}

type gameTitlesClient struct {
	cc grpc.ClientConnInterface
}

func NewGameTitlesClient(cc grpc.ClientConnInterface) GameTitlesClient {
	return &gameTitlesClient{cc}
}

func (c *gameTitlesClient) GetGame(ctx context.Context, in *GameRequest, opts ...grpc.CallOption) (*GameStats, error) {
	out := new(GameStats)
	err := c.cc.Invoke(ctx, "/gametitles.GameTitles/GetGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameTitlesServer is the server API for GameTitles service.
// All implementations must embed UnimplementedGameTitlesServer
// for forward compatibility
type GameTitlesServer interface {
	GetGame(context.Context, *GameRequest) (*GameStats, error)
	mustEmbedUnimplementedGameTitlesServer()
}

// UnimplementedGameTitlesServer must be embedded to have forward compatible implementations.
type UnimplementedGameTitlesServer struct {
}

func (UnimplementedGameTitlesServer) GetGame(context.Context, *GameRequest) (*GameStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGame not implemented")
}
func (UnimplementedGameTitlesServer) mustEmbedUnimplementedGameTitlesServer() {}

// UnsafeGameTitlesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameTitlesServer will
// result in compilation errors.
type UnsafeGameTitlesServer interface {
	mustEmbedUnimplementedGameTitlesServer()
}

func RegisterGameTitlesServer(s grpc.ServiceRegistrar, srv GameTitlesServer) {
	s.RegisterService(&GameTitles_ServiceDesc, srv)
}

func _GameTitles_GetGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameTitlesServer).GetGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gametitles.GameTitles/GetGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameTitlesServer).GetGame(ctx, req.(*GameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GameTitles_ServiceDesc is the grpc.ServiceDesc for GameTitles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameTitles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gametitles.GameTitles",
	HandlerType: (*GameTitlesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGame",
			Handler:    _GameTitles_GetGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gametitles/game_titles.proto",
}
