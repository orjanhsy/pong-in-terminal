package client

import (
	"context"
	"log"

	pb "github.com/orjanhsy/pong-in-terminal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const adress = "localhost:50051"

type GameClient struct {
}



func main() {
	conn, err := grpc.NewClient(adress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewPongServiceClient(conn)
	gameClient := GameClient{}

	gameClient.updatePaddlePosition(grpcClient, "2", 2, 3)
	gameClient.recieveGameState(grpcClient, "2")
}

func (gc *GameClient) updatePaddlePosition(grpcClient pb.PongServiceClient, playerID string, x int, y int) {
	req := &pb.PaddleUpdateRequest{
		PlayerId: playerID,
		PaddlePosition: &pb.Coordinate{
			X: int32(x),
			Y: int32(y),
		},
	}

	resp, err := grpcClient.UpdatePaddlePosition(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not update paddle position: %v", err)
	}
	log.Printf("Reponse of UpdatePaddlePosition: %s", resp.Status)
}

func (gc *GameClient) recieveGameState(grpcClient pb.PongServiceClient, playerID string) {
	req := &pb.GameStateRequest{PlayerId: playerID}
	stream, err := grpcClient.StreamGameState(context.Background(), req) 
	if err != nil {
		log.Fatalf("Error while recieving stream: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error while recieving stream: %v", err)
			break
		}
		log.Printf("GameState: BallPos(%d, %d), P1Pos(%d, %d), P2Pos(%d, %d)", 
			resp.BallPos.X, resp.BallPos.Y, resp.P1Pos.X, resp.P1Pos.Y, resp.P2Pos.X, resp.P2Pos.Y)
	}
}
