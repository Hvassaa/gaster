package rendering

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hvassaa/gaster/player"
	"github.com/hvassaa/gaster/raycasting"
)

type Renderer2D struct {
	UnitX, UnitY, BlockSize                          float64
	Screen                                           *ebiten.Image
	ScreenWidth, ScreenHeight                        float64
	player                                           *player.Player
	PlayerColor, RayColor, DirectionColor, WallColor color.Color
	mab                                              [][]raycasting.WallType
}

func (r2d *Renderer2D) translateX(screen *ebiten.Image, x float64) float32 {
	return float32(float64(screen.Bounds().Min.X) + x*float64(r2d.UnitX))
}

func (r2d *Renderer2D) translateY(screen *ebiten.Image, y float64) float32 {
	return float32(float64(screen.Bounds().Min.Y) + y*float64(r2d.UnitY))
}

func NewRenderer2D(screen *ebiten.Image, worldWidth, worldHeight, blockSize float64, player *player.Player, mab [][]raycasting.WallType) *Renderer2D {
	return &Renderer2D{
		UnitX:          float64(screen.Bounds().Dx()) / worldWidth,
		UnitY:          float64(screen.Bounds().Dy()) / worldHeight,
		Screen:         screen,
		ScreenWidth:    float64(screen.Bounds().Dx()),
		ScreenHeight:   float64(screen.Bounds().Dy()),
		BlockSize:      blockSize,
		PlayerColor:    color.RGBA{200, 0, 0, 255},
		RayColor:       color.RGBA{0, 200, 200, 255},
		DirectionColor: color.RGBA{0, 200, 0, 255},
		WallColor:      color.RGBA{0, 50, 50, 255},
		player:         player,
		mab:            mab,
	}
}

func (r2d *Renderer2D) Render2D(rays []raycasting.Ray) {
	screen := r2d.Screen
	screen.Fill(color.Black)
	xBlockWidth := r2d.translateX(screen, r2d.BlockSize)
	yBlockWidth := r2d.translateY(screen, r2d.BlockSize)

	for y, yv := range r2d.mab {
		yp := r2d.translateY(screen, float64(y)*r2d.BlockSize)
		for x, wallType := range yv {
			xp := r2d.translateX(screen, float64(x)*r2d.BlockSize)
			if wallType != 0 {
				vector.DrawFilledRect(screen, xp, yp, xBlockWidth, yBlockWidth, r2d.WallColor, false)
			}
		}
	}

	for y, yv := range r2d.mab {
		yp := r2d.translateY(screen, float64(y)*r2d.BlockSize)
		for x := range yv {
			xp := r2d.translateX(screen, float64(x)*r2d.BlockSize)
			if y == 0 {
				vector.StrokeLine(screen, xp, 0, xp, float32(r2d.ScreenHeight), 1, color.RGBA{70, 10, 10, 255}, false)
			}
		}
		vector.StrokeLine(screen, 0, yp, float32(r2d.ScreenWidth), yp, 1, color.RGBA{70, 10, 10, 255}, false)
	}

	radius := 20 * (r2d.UnitX + r2d.UnitY) / 2

	playerX := r2d.translateX(screen, float64(r2d.player.Coord.X))
	playerY := r2d.translateY(screen, float64(r2d.player.Coord.Y))
	vector.DrawFilledCircle(screen, playerX, playerY, float32(radius), r2d.PlayerColor, false)
	for _, ray := range rays {
		x1 := r2d.translateX(screen, r2d.player.Coord.X)
		y1 := r2d.translateY(screen, r2d.player.Coord.Y)
		x2 := r2d.translateX(screen, ray.Coord.X)
		y2 := r2d.translateY(screen, ray.Coord.Y)
		vector.StrokeLine(screen, x1, y1, x2, y2, 1, r2d.PlayerColor, false)
	}
	directionRayX := r2d.translateX(screen, r2d.player.Coord.X+math.Cos(r2d.player.Angle)*radius*10)
	directionRayY := r2d.translateY(screen, r2d.player.Coord.Y+math.Sin(r2d.player.Angle)*radius*10)
	vector.StrokeLine(screen, playerX, playerY, directionRayX, directionRayY, 1, r2d.DirectionColor, false)
}
