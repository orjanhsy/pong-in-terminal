package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/orjanhsy/pong-in-terminal/game"
	pb "github.com/orjanhsy/pong-in-terminal/proto"
	"google.golang.org/grpc"
)

const (
	maxClients = 2
	port = ":50051"
)

type client struct {
	streamServer pb.PongService_StreamGameStateServer
	id uuid.UUID
}

type GameServer struct {
	pb.UnimplementedPongServiceServer
	game *game.Game
	clients map[uuid.UUID]*client
}

func NewGameServer() *GameServer {
	game := game.NewGame()
	game.Init()
	gs := &GameServer{
		game: game,
		clients: make(map[uuid.UUID]*client),
	}
	return gs
}

func (gs *GameServer) StreamGameState(req *pb.GameStateRequest, stream pb.PongService_StreamGameStateServer) error {
	clientId, err := uuid.Parse(req.PlayerId)
	if err != nil {
		log.Fatalf("Could not parse client ID when streaming state: %v", err)
	}

	if len(gs.clients) >= maxClients {
		return fmt.Errorf("Server if full (%d/2)", len(gs.clients))
	}

	gs.clients[clientId] = &client{
		streamServer: stream,
		id: clientId, 
	}

	if gs.game.P1 == uuid.Nil {
		gs.game.P1 = clientId
	} else {
		gs.game.P2 = clientId
	}

	for {
		response := &pb.GameStateResponse{
			BallPos: &pb.Coordinate{
				X: int32(gs.game.BallPos.X),
				Y: int32(gs.game.BallPos.Y) ,
			},
			P1Pos: &pb.Coordinate{
				X: int32(gs.game.P1Pos.X),
				Y: int32(gs.game.P1Pos.Y) ,
			},
			P2Pos: &pb.Coordinate{
				X: int32(gs.game.P2Pos.X),
				Y: int32(gs.game.P2Pos.Y) ,
			},
			P1Score: int32(gs.game.P1Score),
			P2Score: int32(gs.game.P2Score),
		}

		if err := stream.Send(response); err != nil {
			log.Printf("Error sending game state to client %s: %v", clientId, err)
			delete(gs.clients, clientId)
			return err
		}

		time.Sleep(time.Second /60)
	}
}


func (gs *GameServer) UpdatePaddlePosition(ctx context.Context, req *pb.PaddleUpdateRequest) (*pb.PaddleUpdateResponse, error) {
	id, err :=	uuid.Parse(req.PlayerId)
	if err != nil {
		log.Fatalf("Failed to parse ID when updateing paddle: %v", err)
	}
  gs.game.MoveChannel <- game.Move{
      PlayerID: id,
      Direction: req.Direction,
  }

  return &pb.PaddleUpdateResponse{Status: "Paddle position updated successfully"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterPongServiceServer(grpcServer, NewGameServer())

	log.Printf("gRPC server running on port: %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
