package game

import "github.com/gdamore/tcell"

type Game struct {
	ballPos Coordinate 
	p1Pos Coordinate 
	p2Pos Coordinate 
	p1Score int32
	p2Score int32

	MoveChannel chan Move
	Screen tcell.Screen
}

type Coordinate struct {
	X int32
	Y int32
}

type Move struct {
	PlayerID string
	Direction Direction
}

type Direction int

const (
	UP Direction = iota
	DOWN
	STOP
)

func (game *Game) performPlayerMoves() {
	for {
		move := <- game.MoveChannel
		game.MovePaddle(move.PlayerID, move.Direction)
	}
}

func (game *Game) watchPlayerMoves() {
	for {
		ev := game.Screen.PollEvent()

		switch ev := ev.(type){
		case *tcell.EventResize:
		case *tcell.EventKey:
			switch ev.Key() {  
			case	tcell.KeyUp || ev.Rune() == 'k':
				game.MoveChannel <- &Move{
					PlayerID: "1",
					Direction: UP,
				}
			case tcell.KeyDown || ev.Rune() == 'j':
				game.MoveChannel <- &Move{
					PlayerId: "1",
					Direction: DOWN,
				}
			}
		}
	}
}

func (game *Game) MovePaddle(playerID string, dir Direction) {
	switch playerID {
	case "1":
		switch dir {
		case UP:
			game.p1Pos.X ++
		case DOWN:
			game.p1Pos.X --
		default:
		}
	case "2":
		switch dir {
		case UP:
			game.p2Pos.X ++
		case DOWN:
			game.p2Pos.X --
		default:
		}
	}

}
