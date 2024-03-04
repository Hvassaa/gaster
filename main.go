package main

import (
	"image"
	"image/color"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hvassaa/gaster/raycasting"
)

const (
	WORLD_WIDTH      = 800.
	WORLD_HEIGHT     = 800.
	BLOCK_SIZE       = 40.
	BLOCKS_X     int = WORLD_WIDTH / BLOCK_SIZE
	BLOCKS_Y     int = WORLD_HEIGHT / BLOCK_SIZE
)

type Game struct {
	player        *Player
	mab           [][]raycasting.WallType
	represntation int
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	if ebiten.IsKeyPressed(ebiten.Key1) {
		g.represntation = 0
	} else if ebiten.IsKeyPressed(ebiten.Key2) {
		g.represntation = 1
	} else if ebiten.IsKeyPressed(ebiten.Key3) {
		g.represntation = 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.IncreaseAngle(-0.05)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.IncreaseAngle(0.05)
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		directionRay, _, _ := raycasting.CastRay(*g.player.coordinate, g.player.Angle, BLOCK_SIZE, g.mab)
		if directionRay.IsInvalid() || directionRay.DistanceTo(*g.player.coordinate) > BLOCK_SIZE {
			g.player.Move(1)
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		directionRay, _, _ := raycasting.CastRay(*g.player.coordinate, raycasting.NormalizeAngle(g.player.Angle+raycasting.PI), BLOCK_SIZE, g.mab)
		if directionRay.IsInvalid() || directionRay.DistanceTo(*g.player.coordinate) > BLOCK_SIZE {
			g.player.Move(-1)
		}
	}
	return nil
}

// func (g *Game) asd() {
// 	for i := -30; i <= 30; i++ {
// 		rayAngle := raycasting.NormalizeAngle(g.player.Angle + (float64(i) * raycasting.DEG_TO_RAD))
// 		coordinate, _, _ := raycasting.CastRay(*g.player.coordinate, rayAngle, BLOCK_SIZE, g.mab)
// 		if !coordinate.IsInvalid() {
// 		}
// 	}
// }

func (g *Game) Draw(screen *ebiten.Image) {
	coords := make([]raycasting.Coordinate, 61)

	rayDistances := make([]float32, 61)
	directions := make([]raycasting.Direction, 61)

	for i := -30; i <= 30; i++ {
		rayAngle := raycasting.NormalizeAngle(g.player.Angle + (float64(i) * raycasting.DEG_TO_RAD))
		coordinate, direction, _ := raycasting.CastRay(*g.player.coordinate, rayAngle, BLOCK_SIZE, g.mab)
		coords[i+30] = coordinate
		noFish := math.Cos(raycasting.NormalizeAngle(g.player.Angle - rayAngle))
		rayDistances[i+30] = float32(coordinate.DistanceTo(*g.player.coordinate) * noFish)
		directions[i+30] = direction
	}

	if g.represntation == 0 {
		twoDScreen := screen.SubImage(image.Rect(0, 0, WORLD_WIDTH, WORLD_HEIGHT)).(*ebiten.Image)
		maxX := screen.Bounds().Max.X
		threeDScreen := screen.SubImage(image.Rect(WORLD_WIDTH, 0, maxX, WORLD_HEIGHT)).(*ebiten.Image)

		// 2D drawing
		g.Draw2DWalls(twoDScreen)
		g.Draw2DPlayer(twoDScreen, coords)

		// 3D drawing
		g.draw3d(threeDScreen, rayDistances, directions)
	} else if g.represntation == 1 {
		g.draw3d(screen, rayDistances, directions)
	} else if g.represntation == 2 {
		twoDScreen := screen.SubImage(image.Rect(0, 0, 300, 300)).(*ebiten.Image)
		twoDScreen.Fill(color.Black)

		// 3D drawing
		g.draw3d(screen, rayDistances, directions)

		// 2D drawing
		twoDScreen.Fill(color.RGBA{0, 0, 0, 0})
		g.Draw2DWalls(twoDScreen)
		g.Draw2DPlayer(twoDScreen, coords)

	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1600, 800
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
	ebiten.SetWindowSize(1600, 800)
	mab := makeStandardMap()

	mab[1][7] = 1
	mab[2][7] = 1
	mab[3][7] = 1
	mab[4][7] = 1
	mab[4][8] = 1
	mab[4][9] = 1
	mab[4][10] = 1
	mab[4][11] = 1
	mab[4][12] = 1
	mab[4][13] = 1
	mab[4][14] = 1

	mab[13][14] = 1
	mab[14][14] = 1
	mab[15][14] = 1
	mab[16][14] = 1
	mab[16][13] = 1
	mab[16][12] = 1
	mab[16][11] = 1

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
