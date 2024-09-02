package game

import (
	"log"
	"os"
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/orjanhsy/pong-in-terminal/proto"
)

type Game struct {
	BallPos Coordinate 
	P1Pos Coordinate 
	P2Pos Coordinate 
	P1Score int
	P2Score int

	MoveChannel chan Move
	Screen tcell.Screen
}

type Coordinate struct {
	X int
	Y int
}

type Move struct {
	PlayerID string
	Direction pb.Direction
}

type Direction int

const (
	UP Direction = iota
	DOWN
	STOP
)

func NewGame() *Game {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Failed to create new screen: %v", err)
	}

	newGame := &Game{
		BallPos: Coordinate{0, 0},
		P1Pos: Coordinate{0, 0},
		P2Pos: Coordinate{0, 0},
		P1Score: 0,
		P2Score: 0,
		MoveChannel: make(chan Move),
		Screen: screen,
	}

	return newGame
}

func (game *Game) Init() {
	log.Print("Starting new game!")
	if err := game.Screen.Init(); err != nil {
		log.Fatalf("Failed to init screen: %v", err)
	}

	w, h := game.Screen.Size()

	// initial positions for moving objects
	game.BallPos = Coordinate{w/2, h/2}
	game.P1Pos = Coordinate{w/8, h/2}
	game.P2Pos = Coordinate{(w/8)*7, h/2}

	game.P1Score = 0
	game.P2Score = 0

	game.Start()
}

func (game *Game) Start() {
	go game.watchPlayerInput()
	go game.performPlayerMoves()
}

func (game *Game) Quit() {
	game.Screen.Fini()
	os.Exit(0)
}

func (game *Game) performPlayerMoves() {
	for {
		move := <- game.MoveChannel
		game.MovePaddle(move.PlayerID, move.Direction)
	}
}

func (game *Game) watchPlayerInput() {
	for {
		ev := game.Screen.PollEvent()

		switch ev := ev.(type){
		case *tcell.EventResize:
		case *tcell.EventKey:
			switch {  
			case	ev.Key() == tcell.KeyUp, ev.Rune() == 'k':
				game.MoveChannel <- Move{
					PlayerID: "1",
					Direction: pb.Direction_DOWN,
				}
			case ev.Key() == tcell.KeyDown, ev.Rune() == 'j':
				game.MoveChannel <- Move{
					PlayerID: "1",
					Direction: pb.Direction_DOWN,
				}
			case ev.Key() == tcell.KeyCtrlC, ev.Rune() == 'q':
				game.Quit()
			}
		}
	}
}

func (game *Game) MovePaddle(playerID string, dir pb.Direction) {
	switch playerID {
	case "1":
		switch dir {
		case pb.Direction_UP:
			game.P1Pos.Y --
		case pb.Direction_DOWN:
			game.P1Pos.Y ++
		default:
		}
	case "2":
		switch dir {
		case pb.Direction_UP:
			game.P2Pos.Y --
		case pb.Direction_DOWN:
			game.P2Pos.Y ++
		default:
		}
	}
}

func (game *Game) DrawObjects() {
	game.Screen.Clear()

	// menu
	ballPos := fmt.Sprintf("Current ball-position: (%d, %d)\n", game.BallPos.X, game.BallPos.Y)
	for i, r := range ballPos {
		game.Screen.SetContent(i, 0, r, nil, tcell.StyleDefault)
	}

	p1Pos := fmt.Sprintf("Current p1-position: (%d, %d)\n", game.P1Pos.X, game.P1Pos.Y)
	for i, r := range p1Pos {
		game.Screen.SetContent(i, 1, r, nil, tcell.StyleDefault)
	}

	p2Pos := fmt.Sprintf("Current p2-position: (%d, %d)\n", game.P2Pos.X, game.P2Pos.Y)
	for i, r := range p2Pos {
		game.Screen.SetContent(i, 2, r, nil, tcell.StyleDefault)
	}

	// ball
	x,y := game.BallPos.X, game.BallPos.Y
	game.Screen.SetContent(x, y, '*', nil, tcell.StyleDefault)

	// paddles 
	// p1
	x, y = game.P1Pos.X, game.P1Pos.Y
	for i := -1; i <= 1; i++ {
		game.Screen.SetContent(x, y+i, '|', nil, tcell.StyleDefault)
	}
	//p2
	x, y = game.P2Pos.X, game.P2Pos.Y
	for i := -1; i <= 1; i++ {
		game.Screen.SetContent(x, y+i, '|', nil, tcell.StyleDefault)
	}

	game.Screen.Show()
}


