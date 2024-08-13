package main

import (
	"context"
	"flag"
	"log"

	"github.com/google/uuid"
	pb "github.com/orjanhsy/pong-in-terminal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "Client_name"
)

var(
	addr = flag.String("addr", "localhost:50052", "the adress to connect to")
	name = flag.String("name", defaultName, "Client to control bat")
)

type GameClient struct {
	currentPlayer uuid.UUID
	game string
}

func NewGameClient() *GameClient {
	return &GameClient{
		game: "pong",
	}
}

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGameClient(conn)
	client := NewGameClient()

	playerId := uuid.New()
	err = client.Connect(c, playerId)
	if err != nil {
		log.Fatalf("did not fulfill connectRequest: %v", err)
	}
}


func (cc *GameClient) Connect(gc pb.GameClient, playerId uuid.UUID) error {
	req := &pb.ConnectRequest{
		Id: playerId.String(),
	}

	r, err := gc.Connect(context.Background(), req)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Printf("Connected with response: %s", r.String())
	return nil
}


// func (cc *GameClient) Stream(ctx context.Context, req *pb.ConnectRequest) (*pb.ConnectResponse, error) {

// }
