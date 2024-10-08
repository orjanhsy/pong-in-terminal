package backend

import (
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/google/uuid"
	"github.com/orjanhsy/pong-in-terminal/proto"
)

type Game struct {
	Ball *Ball
	P1Pos Vector 
	P2Pos Vector 
	P1Dir Direction
	P2Dir Direction
	P1Score int
	P2Score int
	P1 uuid.UUID
	P2 uuid.UUID

	DirUpdates chan Move

	ScreenH float64
	ScreenW float64
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
		P1Dir: STOP,
		P2Dir: STOP,
		P1Score: 0,
		P2Score: 0,
		DirUpdates: make(chan Move),
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
	g.ResetValues()

	g.Start()
}

func (g *Game) ResetValues() {
	w, h := g.Screen.Size()
	g.ScreenW = float64(w)
	g.ScreenH = float64(h)

	// initial positions for moving objects
	g.Ball.Init(g.ScreenW/2, g.ScreenH/2)
	g.P1Pos = Vector{g.ScreenW/8, g.ScreenH/2}
	g.P2Pos = Vector{(g.ScreenW/8)*7, g.ScreenH/2}

	g.P1Score = 0
	g.P2Score = 0
}

func (game *Game) Start() {
	go game.Ball.Move()
	go game.checkForCollitions()
	go game.UpdatePaddleDirections()
	go game.listenForClose()
	go game.MovePaddles()
}

func (game *Game) Quit() {
	game.Screen.ShowCursor(int(game.ScreenW/2), int(game.ScreenH/2))
	game.Screen.Fini()
	os.Exit(0)
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

func (g *Game) UpdatePaddleDirections() {
	for	{
		move := <- g.DirUpdates
		switch move.PlayerID {
		case g.P1:
			g.P1Dir = Direction(move.Direction)
		case g.P2:
			g.P2Dir = Direction(move.Direction)
		}
	}
}


func (game *Game) MovePaddles() {
	for {
		switch game.P1Dir {
		case UP:
			if game.P1Pos.Y > 4 {
				game.P1Pos.Y --
			}
		case DOWN:
			if game.P1Pos.Y < game.ScreenH - 3{
				game.P1Pos.Y ++
			}
		}
		switch game.P2Dir {
		case UP:
			if game.P2Pos.Y > 4 {
				game.P2Pos.Y --
			}
		case DOWN:
			if game.P2Pos.Y < game.ScreenH - 3{
				game.P2Pos.Y ++
			}
		}

		time.Sleep(time.Second / 40)
	}
}

func (g *Game) checkForCollitions() {
	for {
		switch {
		//player goals
		case g.Ball.Pos.X > g.ScreenW: 
			g.P2Score ++ 
			if g.P2Score == 10 {
				g.ResetValues()
				continue
			}

			g.Ball.Init(g.ScreenW/2, g.ScreenH/2)
		case g.Ball.Pos.X < 0:
			g.P1Score ++
			if g.P1Score == 10 {
				g.ResetValues()
				continue
			}

			g.Ball.Init(g.ScreenW/2, g.ScreenH/2)

		//player paddles
		case onPaddle(g.Ball.Pos, g.P1Pos) && g.Ball.LastHit != "p1":
			g.Ball.ChangeDir("x")
			g.Ball.LastHit = "p1"
		case onPaddle(g.Ball.Pos, g.P2Pos) && g.Ball.LastHit != "p2":
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

func onPaddle(ball Vector, pad Vector) bool {
	return ball.Y <= pad.Y + 2 && ball.Y >= pad.Y - 2 && ball.X >= pad.X - 1 && ball.X <= pad.X + 1
}
