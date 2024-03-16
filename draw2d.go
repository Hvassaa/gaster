package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hvassaa/gaster/raycasting"
)

type Renderer2D struct {
	UnitX, UnitY                                     float64
	Screen                                           *ebiten.Image
	ScreenWidth, ScreenHeight                        float64
	PlayerColor, RayColor, DirectionColor, WallColor color.Color
}

func (r2d *Renderer2D) translateX(screen *ebiten.Image, x float64) float32 {
	return float32(float64(screen.Bounds().Min.X) + x*float64(r2d.UnitX))
}

func (r2d *Renderer2D) translateY(screen *ebiten.Image, y float64) float32 {
	return float32(float64(screen.Bounds().Min.Y) + y*float64(r2d.UnitY))
}

func NewRenderer2D(screen *ebiten.Image) *Renderer2D {
	return &Renderer2D{
		UnitX:          float64(screen.Bounds().Dx()) / WORLD_WIDTH,
		UnitY:          float64(screen.Bounds().Dy()) / WORLD_HEIGHT,
		Screen:         screen,
		ScreenWidth:    float64(screen.Bounds().Dx()),
		ScreenHeight:   float64(screen.Bounds().Dy()),
		PlayerColor:    color.RGBA{200, 0, 0, 255},
		RayColor:       color.RGBA{0, 200, 200, 255},
		DirectionColor: color.RGBA{0, 200, 0, 255},
		WallColor:      color.RGBA{0, 50, 50, 255},
	}
}

func (g *Game) Render2D(raysHits []raycasting.Coordinate) {
	r2d := g.r2d
	screen := r2d.Screen
	screen.Fill(color.Black)
	xBlockWidth := r2d.translateX(screen, BLOCK_SIZE)
	yBlockWidth := r2d.translateY(screen, BLOCK_SIZE)

	for y, yv := range g.mab {
		yp := r2d.translateY(screen, float64(y*BLOCK_SIZE))
		for x, wallType := range yv {
			xp := r2d.translateX(screen, float64(x*BLOCK_SIZE))
			if wallType != 0 {
				vector.DrawFilledRect(screen, xp, yp, xBlockWidth, yBlockWidth, r2d.WallColor, false)
			}
		}
	}

	for y, yv := range g.mab {
		yp := r2d.translateY(screen, float64(y*BLOCK_SIZE))
		for x := range yv {
			xp := r2d.translateX(screen, float64(x*BLOCK_SIZE))
			if y == 0 {
				vector.StrokeLine(screen, xp, 0, xp, float32(r2d.ScreenHeight), 1, color.RGBA{70, 10, 10, 255}, false)
			}
		}
		vector.StrokeLine(screen, 0, yp, float32(r2d.ScreenWidth), yp, 1, color.RGBA{70, 10, 10, 255}, false)
	}

	radius := 20 * (r2d.UnitX + r2d.UnitY) / 2

	playerX := r2d.translateX(screen, float64(g.player.coordinate.X))
	playerY := r2d.translateY(screen, float64(g.player.coordinate.Y))
	vector.DrawFilledCircle(screen, playerX, playerY, float32(radius), r2d.PlayerColor, false)
	for _, coordinate := range raysHits {
		if !coordinate.IsInvalid() {
			x1 := r2d.translateX(screen, g.player.coordinate.X)
			y1 := r2d.translateY(screen, g.player.coordinate.Y)
			x2 := r2d.translateX(screen, coordinate.X)
			y2 := r2d.translateY(screen, coordinate.Y)
			vector.StrokeLine(screen, x1, y1, x2, y2, 1, r2d.PlayerColor, false)
		}
	}
	directionRayX := r2d.translateX(screen, g.player.coordinate.X+math.Cos(g.player.Angle)*radius*10)
	directionRayY := r2d.translateY(screen, g.player.coordinate.Y+math.Sin(g.player.Angle)*radius*10)
	vector.StrokeLine(screen, playerX, playerY, directionRayX, directionRayY, 1, r2d.DirectionColor, false)
}
