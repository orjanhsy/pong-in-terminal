package backend

import (
	"math"
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
	b.Velo = Vector{X: veloX, Y: veloY}

	b.LastHit = ""
}

func (b *Ball) Move() {
	for {
		if b.Velo.Y > b.Velo.X {
			// Account for different speed when moving diagonally
			sleepFactor := math.Sqrt(math.Pow(b.Velo.X, 2) + math.Pow(b.Velo.Y, 2))
			time.Sleep((time.Second / 40) * time.Duration(sleepFactor))
		} else {
			time.Sleep((time.Second / 40))
		}

		b.Pos.X += b.Velo.X
		b.Pos.Y += b.Velo.Y
	}
}

func (b *Ball) ChangeDir(dir string) {
	switch dir{
	case "y":
		b.Velo.Y = -b.Velo.Y
	case "x":
		yRand := float64(rand.Intn(3) - 1)
		b.Velo.X = -b.Velo.X  
		b.Velo.Y += float64(yRand)
	}
}

