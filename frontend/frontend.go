package frontend

import (
	"time"

	"github.com/gdamore/tcell"
	"github.com/orjanhsy/pong-in-terminal/backend"
)

type GameInterface struct {
	Game *backend.Game
}

func NewGameInterface(game *backend.Game) *GameInterface {
	return &GameInterface{
		Game: game,
	}
}


func (i *GameInterface) Draw() {
	i.Game.Screen.Clear()

	i.Game.Screen.SetContent(10, i.Game.Player1.Paddle.Pos, '|', nil, tcell.StyleDefault)

	i.Game.Screen.Show()

	time.Sleep(time.Millisecond * 50)
}

func (i *GameInterface) ListenForInput() {
	defer i.Game.Quit()
	for {
		ev := i.Game.Screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyUp || ev.Rune() == 'k' {
				i.Game.Player1.Paddle.Move(backend.UP)
			} else if ev.Key() == tcell.KeyDown || ev.Rune() == 'j' {
				i.Game.Player1.Paddle.Move(backend.DOWN)
			} else if ev.Key() == tcell.KeyCtrlC || ev.Rune() == 'q' {
				i.Game.Quit()
			}
		}
	}
}


