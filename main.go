package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hvassaa/gaster/raycasting"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	WORLD_WIDTH      = 800.
	WORLD_HEIGHT     = 1200.
	BLOCK_SIZE       = 80.
	BLOCKS_X     int = WORLD_WIDTH / BLOCK_SIZE
	BLOCKS_Y     int = WORLD_HEIGHT / BLOCK_SIZE
)

type Game struct {
	player *Player
	mab    [][]raycasting.WallType
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.IncreaseAngle(-0.05)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.IncreaseAngle(0.05)
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.player.Move(1)
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.Move(-1)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.Draw2DWalls(screen)
	g.Draw2DPlayer(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("angle: %v / %v", g.player.Angle, g.player.Angle/raycasting.PI))
}

func (g *Game) Draw2DPlayer(screen *ebiten.Image) {
	playerColor := color.RGBA{200, 0, 0, 0}
	yellow := color.RGBA{200, 200, 0, 0}
	vector.DrawFilledCircle(screen, float32(g.player.coordinate.X), float32(g.player.coordinate.Y), 10, playerColor, false)
	X := float32(g.player.coordinate.X + math.Cos(g.player.Angle)*50)
	Y := float32(g.player.coordinate.Y + math.Sin(g.player.Angle)*50)
	vector.StrokeLine(screen, float32(g.player.coordinate.X), float32(g.player.coordinate.Y), X, Y, 1, yellow, false)
	for i := -30; i <= 30; i++ {
		rayAngle := raycasting.NormalizeAngle(g.player.Angle + (float64(i) * raycasting.DEG_TO_RAD))
		c, _, _ := raycasting.CastRay(*g.player.coordinate, rayAngle, BLOCK_SIZE, g.mab)
		if !c.IsInvalid() {
			vector.StrokeLine(screen, float32(g.player.coordinate.X), float32(g.player.coordinate.Y), float32(c.X), float32(c.Y), 1, playerColor, false)
		}
	}
}

func (g *Game) Draw2DWalls(screen *ebiten.Image) {
	wallColor := color.RGBA{0, 50, 50, 0}
	for y, yv := range g.mab {
		yp := float32(y * BLOCK_SIZE)
		vector.StrokeLine(screen, 0, yp, float32(WORLD_WIDTH), yp, 1, wallColor, false)
		for x, wallType := range yv {
			xp := float32(x * BLOCK_SIZE)
			if y == 0 {
				vector.StrokeLine(screen, xp, 0, xp, float32(WORLD_HEIGHT), 1, wallColor, false)
			}
			if wallType != 0 {
				vector.DrawFilledRect(screen, xp, yp, BLOCK_SIZE, BLOCK_SIZE, wallColor, false)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WORLD_WIDTH, WORLD_HEIGHT
}

func makeStandardMap() [][]raycasting.WallType {
	m := make([][]raycasting.WallType, BLOCKS_Y)
	for i := range m {
		m[i] = make([]raycasting.WallType, BLOCKS_X)
	}
	for y := 0; y < BLOCKS_Y; y++ {
		for x := 0; x < BLOCKS_X; x++ {
			if x == 0 || x == BLOCKS_X-1 || y == 0 || y == BLOCKS_Y-1 {
				m[y][x] = 1
			}
		}
	}

	return m
}

func main() {
	ebiten.SetWindowSize(int(WORLD_WIDTH), int(WORLD_HEIGHT))
	mab := makeStandardMap()

	mab[1][1] = 1
	mab[1][7] = 1
	mab[2][7] = 1

	game := &Game{
		player: &Player{
			coordinate: &raycasting.Coordinate{
				X: WORLD_WIDTH / 2.,
				Y: WORLD_HEIGHT / 2.,
			},
			Angle: 0,
			Speed: 10.,
		},
		mab: mab,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
