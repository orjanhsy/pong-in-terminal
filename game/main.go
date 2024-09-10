package game

import (
	"log"
	"math/rand"
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
	P1Score int
	P2Score int
	P1 uuid.UUID
	P2 uuid.UUID

	ScreenH float64
	ScreenW float64

	MoveChannel chan Move
	Screen tcell.Screen
}

type Ball struct {
	Pos Vector
	Velo Vector
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

func (b *Ball) Init(x float64, y float64) {
	b.Pos = Vector{X: x, Y: y} // this needs to be middle of screen, eventually
	veloX := float64((2 * rand.Intn(2) - 1) * 50)
	veloY := float64(rand.Intn(100) - 50)
	b.Velo = Vector{X: veloX, Y: veloY}
}

func (b *Ball) Move() {
	lastUpdate := time.Now()
	for {
		time.Sleep(time.Millisecond * 16)
		now := time.Now()
		deltaTime := now.Sub(lastUpdate).Seconds()
		lastUpdate = now

		b.Pos.X += b.Velo.X * deltaTime
		b.Pos.Y += b.Velo.Y * deltaTime
	}
}

func (b *Ball) ChangeDir() {
	b.Velo = Vector{X: -b.Velo.X, Y: -b.Velo.Y} 	
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
		// score
		case g.Ball.Pos.X < g.P1Pos.X:
			g.P1Score ++
			g.Ball.Init(g.ScreenW/2, g.ScreenH/2)
		case g.Ball.Pos.X > g.P2Pos.X:
			g.P2Score ++ 
			g.Ball.Init(g.ScreenW/2, g.ScreenH/2)
		// ball has hit paddle
		case g.Ball.Pos.Equals(g.P1Pos) || g.Ball.Pos.Equals(g.P2Pos):
			g.Ball.ChangeDir()
		// ball has hit wall, there is no xBound as that would mean someone have scored
		case int(g.Ball.Pos.Y) == int(g.ScreenH) || int(g.Ball.Pos.Y) == 0: 
			g.Ball.ChangeDir()
		}
	}
}
