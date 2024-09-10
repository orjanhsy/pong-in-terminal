package backend

import (
	"math/rand"
	"time"
)

type Ball struct {
	Pos Vector
	Velo Vector
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

