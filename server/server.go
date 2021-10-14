package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/rickschubert/grpctaster/gamingstats"
	"google.golang.org/grpc"
)

type server struct {
	gamingstats.UnimplementedGamingStatsServer
}

func (s *server) GetGame(ctx context.Context, request *gamingstats.Game) (*gamingstats.GameStats, error) {
	return &gamingstats.GameStats{
		Title: "Uncharted 4",
		Platforms: []gamingstats.Platform{
			gamingstats.Platform_PLAYSTATION_4,
			gamingstats.Platform_PLAYSTATION_5,
			gamingstats.Platform_PC,
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gamingstats.RegisterGamingStatsServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
