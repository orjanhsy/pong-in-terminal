package backend

import (
	"math/rand"
	"time"
)

type Ball struct {
	Pos Vector
	Velo Vector
	LastHit string 
}

func (b *Ball) Init(x float64, y float64) {
	b.Pos.X = x
	b.Pos.Y = y
	veloX := float64(2 * rand.Intn(2) - 1)
	veloY := float64(rand.Intn(3) - 1)
	/*
	veloX := float64((2 * rand.Intn(2) - 1) * 50)
	veloY := float64(rand.Intn(100) - 50)
	*/
	b.Velo = Vector{X: veloX, Y: veloY}

	b.LastHit = ""
}

func (b *Ball) Move() {
	// lastUpdate := time.Now()
	for {
		time.Sleep(time.Second / 40)
		b.Pos.X += b.Velo.X
		b.Pos.Y += b.Velo.Y
		/*
		now := time.Now()
		deltaTime := now.Sub(lastUpdate).Seconds()
		lastUpdate = now
		b.Pos.X += b.Velo.X * deltaTime
		b.Pos.Y += b.Velo.Y * deltaTime
		*/
	}
}

func (b *Ball) ChangeDir(dir string) {
	switch dir{
	case "y":
		b.Velo.Y = -b.Velo.Y
	case "x":
		yRand := rand.Intn(3) - 1
		b.Velo.X = -b.Velo.X
		b.Velo.Y += float64(yRand)
	}
}

