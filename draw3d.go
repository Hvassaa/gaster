package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hvassaa/gaster/raycasting"
)

func (g *Game) draw3d(screen *ebiten.Image, heights []float32, directions []raycasting.Direction) {
	c := color.RGBA{20, 20, 0, 0}
	screen.Fill(c)

	xStart := screen.Bounds().Min.X
	screenSize := screen.Bounds().Size()
	screnWidth := screenSize.X
	screenHeight := screenSize.Y
	columnWidth := float32(screnWidth) / 60.0

	for i, height := range heights {

		columnHeight := float32(screenHeight) - height
		columnColor := color.RGBA{255, 0, 0, 0}
		if directions[i] == raycasting.HORIZONTAL {
			columnColor = color.RGBA{255, 50, 0, 0}
		}

		x := float32(xStart) + float32(i)*columnWidth
		var y1 float32 = float32(screenHeight)/2 - columnHeight/2
		var y2 float32 = float32(screenHeight)/2 + columnHeight/2
		vector.StrokeLine(screen, x, y1, x, y2, columnWidth, columnColor, false)
	}
}
