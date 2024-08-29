package backend

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
)


type Game struct {
	Player1 *Player
	Player2 *Player
	
	Screen tcell.Screen
}

func (g *Game) Quit() {
	g.Screen.Fini()
	os.Exit(0)
}

func NewGame() *Game {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("failed creating screen: %v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("failed initializing screen: %v", err)
	}

	return &Game{
		Player1: NewPlayer(),
		Player2: NewPlayer(),
		Screen: s,
	}
}

type Player struct {
	Paddle *Paddle
	Score int
}

func NewPlayer() *Player {
	return &Player{
		Paddle: &Paddle{Pos: 25}, // 25 is placeholder, needs to be centre of y axis of screen
		Score: 0,
	}
}


type Paddle struct {
	Pos int // moves only across y axis
}

type Direction int

const (
	UP Direction = iota
	DOWN
)

func (p *Paddle) Move(dir Direction) {
	switch dir {
	case UP:
		p.Pos--	
	case DOWN:
		p.Pos++
	}
}

