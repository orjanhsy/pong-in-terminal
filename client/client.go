package main

import ( 
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell"
	"github.com/google/uuid"
	pb "github.com/orjanhsy/pong-in-terminal/proto"
	"github.com/orjanhsy/pong-in-terminal/backend"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const adress = "localhost:50051"

type GameClient struct {
	screen tcell.Screen
	grpcClient pb.PongServiceClient
	playerId uuid.UUID
	ballPosHistory [3] backend.Vector 
}

func NewGameClient(playerId uuid.UUID, grpcClient pb.PongServiceClient) *GameClient {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Failed to create tcell screen: %v", err)
	}

	return &GameClient{
		screen: s,
		grpcClient: grpcClient,
		playerId: playerId,
	}
}

func (gc *GameClient) Start() {
	if err := gc.screen.Init(); err != nil {
		log.Fatalf("Failed to init screen: %v", err)
	}
	
	go gc.listenForPlayerInput()

	go func() {
		if err := gc.recieveGameState(); err != nil {
			log.Printf("Error recieving game state: %v", err)
		}
	}()
	gc.drawBorders(192, 49)
	select {}
}

func (gc *GameClient) Quit() {
	gc.screen.Fini()
	log.Println("Quit game")
	os.Exit(0)
}

func (gc *GameClient) listenForPlayerInput() {
	log.Println("Listening for input")
	for {
		ev := gc.screen.PollEvent()

		switch ev := ev.(type){
		case *tcell.EventResize:
		case *tcell.EventKey:
			switch {  
			case	ev.Key() == tcell.KeyUp, ev.Rune() == 'k':
				gc.sendPaddleUpdate(pb.Direction_UP)	
			case ev.Key() == tcell.KeyDown, ev.Rune() == 'j':
				gc.sendPaddleUpdate(pb.Direction_DOWN)
			case ev.Rune() == ' ':
				gc.sendPaddleUpdate(pb.Direction_STOP)
			case ev.Key() == tcell.KeyCtrlC, ev.Rune() == 'q':
				gc.Quit()
				return
			}
		}
	}
}

func (gc *GameClient) sendPaddleUpdate(dir pb.Direction) {
	req := &pb.PaddleUpdateRequest{
		PlayerId: gc.playerId.String(),
		Direction: dir,
	}

	_, err := gc.grpcClient.UpdatePaddleDirection(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not update paddle position: %v", err)
	}
}

func (gc *GameClient) recieveGameState() error {
	req := &pb.GameStateRequest{PlayerId: gc.playerId.String()}
	stream, err := gc.grpcClient.StreamGameState(context.Background(), req) 
	if err != nil {
		log.Fatalf("Error while recieving stream: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}
		gc.drawGameState(resp)
	}
	return nil
}

func (gc *GameClient) drawBorders(w int, h int) {
	// roof and floor borders
	for i := 0; i < int(w); i++ {
		gc.screen.SetContent(i, 0, '-', nil, tcell.StyleDefault.Bold(true))
		gc.screen.SetContent(i, h, '-', nil, tcell.StyleDefault.Bold(true))
	}
	for i := 0; i <= int(h); i++ {
		gc.screen.SetContent(0, i, '|', nil, tcell.StyleDefault.Bold(true))	
		gc.screen.SetContent(w, i, '|', nil, tcell.StyleDefault.Bold(true))	
	}
}


func (gc *GameClient) drawGameState(state *pb.GameStateResponse) {
	gc.screen.Clear()
	sw := int(state.GetScreenWidth())
	sh :=  int(state.GetScreenHeight())
	gc.drawBorders(sw, sh)

	// menu
	// ballPos := fmt.Sprintf("Current ball-position: (%d, %d)\n", state.BallPos.X, state.BallPos.Y)
	// for i, r := range ballPos {
	// 	gc.screen.SetContent(i, 0, r, nil, tcell.StyleDefault)
	// }
	//
	// p1Pos := fmt.Sprintf("Current p1-position: (%d, %d)\n", state.P1Pos.X, state.P1Pos.Y)
	// for i, r := range p1Pos {
	// 	gc.screen.SetContent(i, 1, r, nil, tcell.StyleDefault)
	// }
	//
	// p2Pos := fmt.Sprintf("Current p2-position: (%d, %d)\n", state.P2Pos.X, state.P2Pos.Y)
	// for i, r := range p2Pos {
	// 	gc.screen.SetContent(i, 2, r, nil, tcell.StyleDefault)
	// }
	//

	// screenPos := fmt.Sprintf("ScreenDim: %d, %d", state.ScreenWidth, state.ScreenHeight)
	// for i, r := range screenPos {
	// 	gc.screen.SetContent(i, 5, r, nil, tcell.StyleDefault)
	// }

	// scores
	p1Score := fmt.Sprintf("%d", state.P1Score)
	for i, r := range p1Score {
		gc.screen.SetContent((sw / 16) * 18 + i, 4, r, nil, tcell.StyleDefault)
	}
	p2Score := fmt.Sprintf("%d", state.P2Score)
	for i, r := range p2Score {
		gc.screen.SetContent(sw / 16 + i, 4, r, nil, tcell.StyleDefault)
	}

	// ball
	x,y := int(state.BallPos.X), int(state.BallPos.Y)
	gc.screen.SetContent(x, y, 'O', nil, tcell.StyleDefault)

	// paddles 
	// p1
	x, y = int(state.P1Pos.X), int(state.P1Pos.Y)
	for i := -2; i <= 2; i++ {
		gc.screen.SetContent(x, y+i, '|', nil, tcell.StyleDefault)
	}
	//p2
	x, y = int(state.P2Pos.X), int(state.P2Pos.Y)
	for i := -2; i <= 2; i++ {
		gc.screen.SetContent(x, y+i, '|', nil, tcell.StyleDefault)
	}

	gc.screen.Show()
}


func main() {
	conn, err := grpc.NewClient(adress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	clientId := uuid.New()

	grpcClient := pb.NewPongServiceClient(conn)
	client := NewGameClient(clientId, grpcClient)
	defer client.Quit()
	client.Start()
}

