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
	WORLD_WIDTH      = 1200.
	WORLD_HEIGHT     = 1200.
	BLOCK_SIZE       = 40.
	BLOCKS_X     int = WORLD_WIDTH / BLOCK_SIZE
	BLOCKS_Y     int = WORLD_HEIGHT / BLOCK_SIZE
	FOV              = 60
	NO_OF_RAYS       = (60 * 4) + 1
	DEG_BOUNDS       = (NO_OF_RAYS - 1) / 2
	DEG_PER_RAY      = FOV / (NO_OF_RAYS - 1.)
)

type Game struct {
	player           *Player
	mab              [][]raycasting.WallType
	represntation    int
	cursorX, cursorY int
	yMod             float32
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

	// calculate mouse deltas
	newCursorX, newCursorY := ebiten.CursorPosition()
	deltaX := g.cursorX - newCursorX
	deltaY := g.cursorY - newCursorY

	// Update y, to look up or down
	if deltaY != 0 && g.cursorY != 0 {
		g.yMod = g.yMod + float32(deltaY)
	}

	// update mouse position
	g.cursorX, g.cursorY = newCursorX, newCursorY

	// look left or right with mouse
	xMultiplier := 1.
	if deltaX != 0 {
		xMultiplier += math.Abs(float64(deltaX)) / 150.
	}
	if deltaX > 0 {
		g.player.IncreaseAngle(-0.05 * xMultiplier)
	} else if deltaX < 0 {
		g.player.IncreaseAngle(0.05 * xMultiplier)
	}

	// move forward or backwards with keyboard
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		ray, err := raycasting.CastRay(*g.player.coordinate, g.player.Angle, BLOCK_SIZE, g.mab)
		if err == nil && ray.C.DistanceTo(*g.player.coordinate) > BLOCK_SIZE {
			g.player.Move(1)
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		ray, err := raycasting.CastRay(*g.player.coordinate, raycasting.NormalizeAngle(g.player.Angle+raycasting.PI), BLOCK_SIZE, g.mab)
		if err == nil && ray.C.DistanceTo(*g.player.coordinate) > BLOCK_SIZE {
			g.player.Move(-1)
		}
	}

	// strafe
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		angle := -raycasting.PI_HALF
		ray, err := raycasting.CastRay(*g.player.coordinate, raycasting.NormalizeAngle(g.player.Angle+angle), BLOCK_SIZE, g.mab)
		if err == nil && ray.C.DistanceTo(*g.player.coordinate) > BLOCK_SIZE {
			g.player.MoveWithAngle(1, angle)
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		angle := raycasting.PI_HALF
		ray, err := raycasting.CastRay(*g.player.coordinate, raycasting.NormalizeAngle(g.player.Angle+angle), BLOCK_SIZE, g.mab)
		if err == nil && ray.C.DistanceTo(*g.player.coordinate) > BLOCK_SIZE {
			g.player.MoveWithAngle(1, angle)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	coords := make([]raycasting.Coordinate, NO_OF_RAYS)
	rayDistances := make([]float32, NO_OF_RAYS)
	directions := make([]raycasting.Direction, NO_OF_RAYS)

	for i := -DEG_BOUNDS; i <= DEG_BOUNDS; i++ {
		rayAngle := raycasting.NormalizeAngle(g.player.Angle + (float64(i) * DEG_PER_RAY * raycasting.DEG_TO_RAD))
		ray, err := raycasting.CastRay(*g.player.coordinate, rayAngle, BLOCK_SIZE, g.mab)
		if err != nil {
			continue
		}
		coords[i+DEG_BOUNDS] = ray.C
		noFish := math.Cos(raycasting.NormalizeAngle(rayAngle - g.player.Angle))
		rayDistances[i+DEG_BOUNDS] = float32(ray.C.DistanceTo(*g.player.coordinate) * noFish)
		directions[i+DEG_BOUNDS] = ray.D
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
	ebiten.SetCursorMode(ebiten.CursorModeCaptured)

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
		mab:           mab,
		represntation: 2,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
