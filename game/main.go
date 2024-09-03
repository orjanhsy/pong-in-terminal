package game

import (
	"log"
	"os"

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
	go game.listenForClose()
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

func (game *Game) listenForClose() {
	for {
		ev := game.Screen.PollEvent()

		switch ev := ev.(type){
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC || ev.Rune() == 'q'{
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
