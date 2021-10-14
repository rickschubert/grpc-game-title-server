package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"

	"github.com/rickschubert/grpc-game-title-server/gametitles"
	"google.golang.org/grpc"
)

type server struct {
	gametitles.UnimplementedGameTitlesServer
}

func readAllGames() ([]gametitles.GameStats, error) {
	bytes, err := ioutil.ReadFile("server/games.json")
	if err != nil {
		return []gametitles.GameStats{}, fmt.Errorf("unable to read JSON games file: %w", err)
	}
	var games []gametitles.GameStats
	err = json.Unmarshal(bytes, &games)
	if err != nil {
		return []gametitles.GameStats{}, fmt.Errorf("unable to unmarshal JSON games file: %w", err)
	}
	return games, err
}

// A real project would find this game in a database. As this is just an example,
// we are instead parsing a JSON and find a suitable result by string match.
func FindGameInDatabase(title string) (gametitles.GameStats, error) {
	games, err := readAllGames()
	if err != nil {
		return gametitles.GameStats{}, fmt.Errorf("unable to read games: %w", err)
	}
	for _, game := range games {
		desiredTitle := strings.ToLower(title)
		desiredTitle = strings.TrimSpace(desiredTitle)
		gameTitle := strings.ToLower(game.GetTitle())
		gameTitle = strings.TrimSpace(gameTitle)
		if strings.Contains(gameTitle, desiredTitle) {
			return game, nil
		}
	}
	return gametitles.GameStats{}, fmt.Errorf("cannot find game with title %s", title)
}

func (s *server) GetGame(ctx context.Context, request *gametitles.GameRequest) (*gametitles.GameStats, error) {
	game, err := FindGameInDatabase(request.Title)
	if err != nil {
		return &gametitles.GameStats{}, fmt.Errorf("unable to find game: %w", err)
	}
	return &game, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gametitles.RegisterGameTitlesServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
