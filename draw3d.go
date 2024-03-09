package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hvassaa/gaster/raycasting"
)

type Renderer3D struct {
	TopColor, BottomColor                             color.Color
	WallColors                                        map[raycasting.WallType]color.Color
	Screen                                            *ebiten.Image
	ScreenMid, ColumnWidth, ScreenWidth, ScreenHeight float32
}

func NewRenderer3D(screen *ebiten.Image, noOfRays int) *Renderer3D {
	wallColors := make(map[raycasting.WallType]color.Color)
	wallColors[0] = color.RGBA{255, 0, 0, 255}
	wallColors[1] = color.RGBA{155, 0, 0, 255}

	screenHeight := float32(screen.Bounds().Size().Y)
	screenWidth := float32(screen.Bounds().Size().X)

	return &Renderer3D{
		TopColor:     color.RGBA{50, 150, 150, 255},
		BottomColor:  color.RGBA{200, 200, 200, 255},
		WallColors:   wallColors,
		Screen:       screen,
		ScreenHeight: screenHeight,
		ScreenWidth:  screenWidth,
		ScreenMid:    screenHeight / 2.,
		ColumnWidth:  screenWidth / float32(noOfRays),
	}
}

func (g *Game) Render3D(rayDistances []float32, directions []raycasting.Direction) {
	r3d := g.r3d

	r3d.Screen.Fill(r3d.BottomColor)

	xStart := r3d.Screen.Bounds().Min.X
	// we render walls "half up and down" from this point
	// we initially set it to the middle of the screen
	renderMiddle := r3d.ScreenMid + g.yMod

	vector.DrawFilledRect(r3d.Screen, 0, renderMiddle, r3d.ScreenWidth, -r3d.ScreenHeight*4, r3d.TopColor, false)

	for i, rayDist := range rayDistances {
		columnColor := color.RGBA{255, 0, 0, 255}
		if directions[i] == raycasting.HORIZONTAL {
			columnColor.R = 150
		}

		// this avoid fisheye on "side walls"
		columnHeight := r3d.ScreenHeight / rayDist
		columnHeight *= 70
		x := float32(xStart) + float32(i)*r3d.ColumnWidth
		var y1 float32 = renderMiddle - columnHeight/2
		var y2 float32 = renderMiddle + columnHeight/2
		vector.StrokeLine(r3d.Screen, x, y1, x, y2, r3d.ColumnWidth, columnColor, false)
	}
}

func (g *Game) draw3d(screen *ebiten.Image, rayDistances []float32, directions []raycasting.Direction) {
	backColor := color.RGBA{200, 200, 200, 255}
	screen.Fill(backColor)

	xStart := screen.Bounds().Min.X
	screenSize := screen.Bounds().Size()
	screnWidth := screenSize.X
	screenHeight := screenSize.Y
	columnWidth := float32(screnWidth) / NO_OF_RAYS
	// we render walls "half up and down" from this point
	// we initially set it to the middle of the screen
	renderMiddle := (float32(screenHeight) / 2) + g.yMod

	topColor := color.RGBA{50, 150, 150, 255}
	vector.DrawFilledRect(screen, 0, renderMiddle, float32(screnWidth), -float32(screenHeight)*4, topColor, false)

	for i, rayDist := range rayDistances {
		columnColor := color.RGBA{255, 0, 0, 255}
		if directions[i] == raycasting.HORIZONTAL {
			columnColor.R = 150
		}

		// this avoid fisheye on "side walls"
		columnHeight := float32(screenHeight) / rayDist
		columnHeight *= 70
		x := float32(xStart) + float32(i)*columnWidth
		var y1 float32 = renderMiddle - columnHeight/2
		var y2 float32 = renderMiddle + columnHeight/2
		vector.StrokeLine(screen, x, y1, x, y2, columnWidth, columnColor, false)
	}
}
