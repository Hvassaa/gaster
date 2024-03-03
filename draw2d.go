package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hvassaa/gaster/raycasting"
)

func translateX(screen *ebiten.Image, x float64) float32 {
	unitX := float64(screen.Bounds().Dx()) / WORLD_WIDTH
	return float32(x * unitX)
}

func translateY(screen *ebiten.Image, y float64) float32 {
	unitY := float64(screen.Bounds().Dy()) / WORLD_HEIGHT
	return float32(y * unitY)
}

func (g *Game) Draw2DPlayer(screen *ebiten.Image) {
	playerColor := color.RGBA{200, 0, 0, 1}
	yellow := color.RGBA{200, 200, 0, 1}
	unitX := float64(screen.Bounds().Dx()) / WORLD_WIDTH
	unitY := float64(screen.Bounds().Dy()) / WORLD_HEIGHT
	compUnit := (unitX + unitY) / 2

	radius := 10* compUnit
	playerX := translateX(screen, float64(g.player.coordinate.X))
	playerY := translateX(screen, float64(g.player.coordinate.Y))
	vector.DrawFilledCircle(screen, playerX, playerY, float32(radius), playerColor, false)
	directionRayX := translateX(screen, g.player.coordinate.X + math.Cos(g.player.Angle)*50)
	directionRayY := translateY(screen, g.player.coordinate.Y + math.Sin(g.player.Angle)*50)
	vector.StrokeLine(screen, playerX, playerY, directionRayX, directionRayY, 2, yellow, false)
	for i := -30; i <= 30; i++ {
		rayAngle := raycasting.NormalizeAngle(g.player.Angle + (float64(i) * raycasting.DEG_TO_RAD))
		coordinate, _, _ := raycasting.CastRay(*g.player.coordinate, rayAngle, BLOCK_SIZE, g.mab)
		if !coordinate.IsInvalid() {
			x1 := translateX(screen, g.player.coordinate.X)
			y1 := translateX(screen, g.player.coordinate.Y)
			x2 := translateX(screen, coordinate.X)
			y2 := translateX(screen, coordinate.Y)
			vector.StrokeLine(screen, x1, y1, x2, y2, 2, playerColor, false)
		}
	}
}

func (g *Game) Draw2DWalls(screen *ebiten.Image) {
	wallColor := color.RGBA{0, 50, 50, 0}
	for y, yv := range g.mab {
		yp := translateY(screen, float64(y * BLOCK_SIZE))
		vector.StrokeLine(screen, 0, yp, float32(WORLD_WIDTH), yp, 1, wallColor, false)
		for x, wallType := range yv {
			xp := translateX(screen, float64(x * BLOCK_SIZE))
			if y == 0 {
				vector.StrokeLine(screen, xp, 0, xp, translateX(screen, WORLD_HEIGHT), 1, wallColor, false)
			}
			if wallType != 0 {
				xBlockWidth := translateX(screen, BLOCK_SIZE)
				yBlockWidth := translateX(screen, BLOCK_SIZE)
				vector.DrawFilledRect(screen, xp, yp, xBlockWidth, yBlockWidth, wallColor, false)
			}
		}
	}
}
