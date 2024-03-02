package main

import (
	"math"

	"github.com/hvassaa/gaster/raycasting"
)

type Player struct {
	coordinate *raycasting.Coordinate
	Angle      float64
	Speed      float64
}

func (p *Player) IncreaseAngle(inc float64) {
	p.Angle = raycasting.NormalizeAngle(p.Angle + inc)
}

func (p *Player) Move(multiplier float64) {
	p.coordinate.X += math.Cos(p.Angle) * p.Speed * multiplier
	p.coordinate.Y += math.Sin(p.Angle) * p.Speed * multiplier
}
