package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hvassaa/gaster/raycasting"
)

func (g *Game) draw3d(screen *ebiten.Image, rayDistances []float32, directions []raycasting.Direction) {
	backColor := color.RGBA{200, 200, 200, 255}
	screen.Fill(backColor)

	xStart := screen.Bounds().Min.X
	screenSize := screen.Bounds().Size()
	screnWidth := screenSize.X
	screenHeight := screenSize.Y
	columnWidth := float32(screnWidth) / NO_OF_RAYS

	topColor := color.RGBA{50, 150, 150, 255}
	vector.DrawFilledRect(screen, 0, float32(screenHeight) / 2, float32(screnWidth), -float32(screenHeight) / 2, topColor, false)

	for i, rayDist := range rayDistances {
		columnColor := color.RGBA{255, 0, 0, 255}
		if directions[i] == raycasting.HORIZONTAL {
			columnColor.R = 150
		}

		// this avoid fisheye on "side walls"
		columnHeight := float32(screenHeight) / rayDist
		columnHeight *= 70
		x := float32(xStart) + float32(i)*columnWidth
		var y1 float32 = float32(screenHeight)/2 - columnHeight/2
		var y2 float32 = float32(screenHeight)/2 + columnHeight/2
		vector.StrokeLine(screen, x, y1, x, y2, columnWidth, columnColor, false)
	}
}
