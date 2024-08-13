package frontend

import "github.com/orjanhsy/pong-in-terminal/backend"

type GameInterface struct {
	Game *backend.Game
}

func NewGameInterface(game *backend.Game) *GameInterface {
	return &GameInterface{
		Game: game,
	}
}

