package main

import (
	"log"
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"sync"

	"github.com/google/uuid"
	pb "github.com/orjanhsy/pong-in-terminal/proto"
	"google.golang.org/grpc"
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
	// password string // (not MVP)

	pb.UnimplementedGameServer
}


func (s *GameServer) Connect(ctx context.Context, req *pb.ConnectRequest) (*pb.ConnectResponse, error) {
	if len(s.clients) >= maxClients {
		return nil, errors.New("Server is full")
	}

	// add new the newly connected client 
	s.mu.Lock()
	token := uuid.New()
	s.clients[token] = &client{
		id: token,
	}
	s.mu.Unlock()

	log.Printf("%s connected to server", token.String())

	return &pb.ConnectResponse{
		Token: token.String(),
	}, nil
}

func (s *GameServer) Stream(ctx context.Context, req *pb.StreamRequest) (*pb.StreamResponse, error) {
	return &pb.StreamResponse{
		Token: uuid.New().String(),
	}, nil
}

func NewGameServer() *GameServer {
	s := &GameServer{
		clients: make(map[uuid.UUID]*client),
		mu: sync.Mutex{},
	}

	return s
}


var (
	port = flag.Int("port", 50052, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening on: %d", port)

	s := grpc.NewServer()
	server := NewGameServer()
	pb.RegisterGameServer(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
