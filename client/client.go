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
	CurrentPlayer uuid.UUID
	Game string
	Stream pb.Game_StreamClient
}

func NewGameClient() *GameClient {
	return &GameClient{
		Game: "pong",
	}
}

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	gc := pb.NewGameClient(conn)
	client := NewGameClient()

	playerId := uuid.New()
	err = client.Connect(gc, playerId)
	if err != nil {
		log.Fatalf("did not fulfill connectRequest: %v", err)
	}

	client.Start()
}


func (c *GameClient) Connect(gc pb.GameClient, playerId uuid.UUID) error {
	req := &pb.ConnectRequest{
		Id: playerId.String(),
	}

	r, err := gc.Connect(context.Background(), req)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Printf("Connected with response: %s", r.String())

	stream, err := gc.Stream(context.Background())
	if err != nil {
		return err 
	}

	c.Stream = stream
	return nil
}

func (c *GameClient) Start() {
	go func() {
		for {
			resp, err := c.Stream.Recv()
			if err != nil {
				log.Printf("Don't work lol")
			}
			log.Printf("Response: %s", resp.Token)
		}
	}()
}





