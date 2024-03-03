package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) draw3d(screen *ebiten.Image) {
	c := color.RGBA{20, 20, 0, 0}
	screen.Fill(c)

	c = color.RGBA{255, 0, 0, 0}
	xStart := screen.Bounds().Min.X
	screenSize := screen.Bounds().Size()
	width := screenSize.X
	height := screenSize.Y
	_, _, _ = xStart, width, height
	// columnWidth := 

	vector.StrokeLine(screen, float32(xStart), float32(height)/2, float32(xStart) + float32(width) -20, float32(height) / 2, 30, c, false)
	vector.DrawFilledCircle(screen, float32(xStart), 10, 200, c, false)
}
