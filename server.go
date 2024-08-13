package main

import (
	"context"
	"errors"
	"net"
	"sync"

	"github.com/google/uuid"
	pb "github.com/orjanhsy/pong-in-terminal/proto"
)

const (
	maxClients = 2
)

type client struct {
	id uuid.UUID
}

type GameServer struct {
	clients map[uuid.UUID]*client
	mu sync.Mutex
	lis map[net.Listener]bool
	// password string // (not MVP)

	pb.UnimplementedGameServer
}


func (s *GameServer) Connect(ctx context.Context, req *pb.ConnectRequest) (*pb.ConnectResponse, error) {
	if len(s.clients) >= maxClients {
		return nil, errors.New("Server is full")
	}

	s.mu.Lock()
	token := uuid.New()
	s.clients[token] = &client{
		id: token,
	}
	s.mu.Unlock()

	return &pb.ConnectResponse{
		Token: token.String(),
	}, nil
}

func (s *GameServer) Stream() {

}


