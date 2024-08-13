package backend

import "github.com/gdamore/tcell"


type Game struct {
	player1 *Player
	player2 *Player

	screen tcell.Screen
}


func NewGame() *Game {
	return &Game{
		player1: NewPlayer(),
		player2: NewPlayer(),
	}
}

type Player struct {
	paddle *Paddle
	score int
}

func NewPlayer() *Player {
	return &Player{
		paddle: &Paddle{pos: 25}, // 25 is placeholder, needs to be centre of y axis of screen
		score: 0,
	}
}


type Paddle struct {
	pos int // moves only across y axis
}

type Direction int

const (
	UP Direction = iota
	DOWN
)


func (p *Paddle) Move(dir Direction) {
	switch dir {
	case UP:
		p.pos--	
	case DOWN:
		p.pos++
	}
}

