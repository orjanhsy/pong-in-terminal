package main

import (
	"time"

	game "github.com/orjanhsy/pong-in-terminal/game"
)

func main() {
	game := game.NewGame()
	game.Init()

	defer game.Quit()
	for {
		game.DrawObjects()
		time.Sleep(time.Millisecond * 50)
	}
}

