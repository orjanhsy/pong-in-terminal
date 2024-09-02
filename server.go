package main

import (
	"github.com/google/uuid"
	"github.com/orjanhsy/pong-in-terminal/game"
	"context"
	pb "github.com/orjanhsy/pong-in-terminal/proto"
)

type client struct {
	streamServer pb.PongService_StreamGameStateServer
	id uuid.UUID
}

type GameServer struct {
	pb.UnimplementedPongServiceServer
	game *game.Game
}

func (gs *GameServer) StreamGameState(req *pb.GameStateRequest, ps pb.PongService_StreamGameStateServer) error {
	return nil
}

func (gs *GameServer ) UpdatePaddlePosition(ctx context.Context, req *pb.PaddleUpdateRequest) (*pb.PaddleUpdateResponse, error) {
	return nil, nil
}

