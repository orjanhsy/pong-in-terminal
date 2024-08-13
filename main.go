package main

import (

	"github.com/orjanhsy/pong-in-terminal/backend"
	"github.com/orjanhsy/pong-in-terminal/frontend"
)

func main() {
	game := backend.NewGame()
	view := frontend.NewGameInterface(game)
	defer game.Quit()
	
	go view.ListenForInput()

	for {
		view.Draw()
	}
}

