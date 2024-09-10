package backend

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
	"github.com/google/uuid"
	"github.com/orjanhsy/pong-in-terminal/proto"
)

type Game struct {
	Ball *Ball
	P1Pos Vector 
	P2Pos Vector 
	P1Score int
	P2Score int
	P1 uuid.UUID
	P2 uuid.UUID

	ScreenH float64
	ScreenW float64

	MoveChannel chan Move
	Screen tcell.Screen
}

type Vector struct {
	X float64
	Y float64
}

type Move struct {
	PlayerID uuid.UUID
	Direction pb.Direction
}

type Direction int

const (
	UP Direction = iota
	DOWN
	STOP
)

func (v Vector) Equals(vec Vector) bool {
	return int(v.X) == int(vec.X) && int(v.Y) == int(vec.Y)
}

func NewGame() *Game {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Failed to create new screen: %v", err)
	}
	
	ball := &Ball{}

	newGame := &Game{
		Ball: ball,
		P1Pos: Vector{0, 0},
		P2Pos: Vector{0, 0},
		P1Score: 0,
		P2Score: 0,
		MoveChannel: make(chan Move),
		Screen: screen,
	}

	return newGame
}

func (g *Game) Init() {
	log.Print("Starting new game!")
	if err := g.Screen.Init(); err != nil {
		log.Fatalf("Failed to init screen: %v", err)
	}

	g.Screen.HideCursor()
	w, h := g.Screen.Size()
	g.ScreenW = float64(w)
	g.ScreenH = float64(h)

	// initial positions for moving objects
	g.Ball.Init(g.ScreenW/2, g.ScreenH/2)
	g.P1Pos = Vector{g.ScreenW/8, g.ScreenH/2}
	g.P2Pos = Vector{(g.ScreenW/8)*7, g.ScreenH/2}

	g.P1Score = 0
	g.P2Score = 0


	g.Start()
}

func (game *Game) Start() {
	go game.Ball.Move()
	go game.checkForCollitions()
	go game.listenForClose()
	go game.performPlayerMoves()
}

func (game *Game) Quit() {
	game.Screen.ShowCursor(int(game.ScreenW/2), int(game.ScreenH/2))
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

func (game *Game) MovePaddle(playerID uuid.UUID, dir pb.Direction) {
	switch playerID {
	case game.P2:
		switch dir {
		case pb.Direction_UP:
			game.P1Pos.Y --
		case pb.Direction_DOWN:
			game.P1Pos.Y ++
		default:
		}
	case game.P1:
		switch dir {
		case pb.Direction_UP:
			game.P2Pos.Y --
		case pb.Direction_DOWN:
			game.P2Pos.Y ++
		default:
		}
	}
}

func (g *Game) checkForCollitions() {
	for {
		switch {
		//player goals
		case g.Ball.Pos.X > g.ScreenW: 
			g.P2Score ++ 
			g.Ball.Init(g.ScreenW/2, g.ScreenH/2)
		case g.Ball.Pos.X < 0:
			g.P1Score ++
			g.Ball.Init(g.ScreenW/2, g.ScreenH/2)

		//player paddles
		case g.Ball.Pos.Equals(g.P1Pos) && g.Ball.LastHit != "p1":
			g.Ball.ChangeDir("x")
			g.Ball.LastHit = "p1"
		case g.Ball.Pos.Equals(g.P2Pos) && g.Ball.LastHit != "p2":
			g.Ball.ChangeDir("x")
			g.Ball.LastHit = "p2"

		// bounds of play area
		case int(g.Ball.Pos.Y) >= int(g.ScreenH) && g.Ball.LastHit != "floor": 
			g.Ball.ChangeDir("y")
			g.Ball.LastHit = "floor"
		case int(g.Ball.Pos.Y) <= 0 && g.Ball.LastHit != "roof": 
			g.Ball.ChangeDir("y")
			g.Ball.LastHit = "roof"
		}
	}
}
