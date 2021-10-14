package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/rickschubert/grpctaster/gamingstats"
	"google.golang.org/grpc"
)

type server struct {
	gamingstats.UnimplementedGamingStatsServer
}

func readAllGames() ([]gamingstats.GameStats, error) {
	bytes, err := ioutil.ReadFile("server/games.json")
	if err != nil {
		return []gamingstats.GameStats{}, fmt.Errorf("unable to read JSON games file: %w", err)
	}
	var games []gamingstats.GameStats
	err = json.Unmarshal(bytes, &games)
	if err != nil {
		return []gamingstats.GameStats{}, fmt.Errorf("unable to unmarshal JSON games file: %w", err)
	}
	return games, err
}

// A real project would find this game in a database. As this is just an eample,
// we are instead just parsing a JSON and finding the results from there.
func FindGameInDatabase(title string) (gamingstats.GameStats, error) {
	games, err := readAllGames()
	if err != nil {
		return gamingstats.GameStats{}, fmt.Errorf("unable to read games: %w", err)
	}
	return games[0], nil
}

func (s *server) GetGame(ctx context.Context, request *gamingstats.GameRequest) (*gamingstats.GameStats, error) {
	game, err := FindGameInDatabase(request.Title)
	if err != nil {
		return &gamingstats.GameStats{}, fmt.Errorf("Unable to find game: %w", err)
	}
	return &game, nil
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