package main

import (
	"context"
	"log"
	"time"

	gamingstats "github.com/rickschubert/grpc-game-title-server/gamingstats"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := gamingstats.NewGamingStatsClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.GetGame(ctx, &gamingstats.GameRequest{
		Title: "Uncharted 4",
	})
	if err != nil {
		log.Fatalf("Received error from server: %v", err)
	}
	log.Printf("Received game from server: %s - %v", response.GetTitle(), response.GetPlatforms())
}
