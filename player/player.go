package player

import (
	"math"

	"github.com/hvassaa/gaster/raycasting"
)

type Player struct {
	Coord *raycasting.Coordinate
	Angle      float64
	Speed      float64
}

// TODO fix strafing and moving extra speed

func (p *Player) IncreaseAngle(inc float64) {
	p.Angle = raycasting.NormalizeAngle(p.Angle + inc)
}

func (p *Player) Move(multiplier float64) {
	p.Coord.X += math.Cos(p.Angle) * p.Speed * multiplier
	p.Coord.Y += math.Sin(p.Angle) * p.Speed * multiplier
}

func (p *Player) MoveWithAngle(multiplier, angle float64) {
	p.Coord.X += math.Cos(raycasting.NormalizeAngle(p.Angle + angle)) * p.Speed * multiplier
	p.Coord.Y += math.Sin(raycasting.NormalizeAngle(p.Angle + angle)) * p.Speed * multiplier
}
