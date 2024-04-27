package main

import (
	"image"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hvassaa/gaster/player"
	"github.com/hvassaa/gaster/raycasting"
	"github.com/hvassaa/gaster/rendering"
)

const (
	WORLD_WIDTH      = 1200.
	WORLD_HEIGHT     = 1200.
	BLOCK_SIZE       = 40.
	BLOCKS_X     int = WORLD_WIDTH / BLOCK_SIZE
	BLOCKS_Y     int = WORLD_HEIGHT / BLOCK_SIZE
	FOV              = 60
	NO_OF_RAYS       = 61
	DEG_BOUNDS       = (NO_OF_RAYS - 1) / 2
	DEG_PER_RAY      = FOV / (NO_OF_RAYS - 1.)
)

type Game struct {
	player           *player.Player
	mab              [][]raycasting.WallType
	represntation    int
	cursorX, cursorY int
	Paused           bool
	r3d              rendering.Renderer
	r2d              rendering.Renderer
	updateRenders    bool
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		os.Exit(0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		g.Paused = true
		ebiten.SetCursorMode(ebiten.CursorModeVisible)
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.Paused = false
		ebiten.SetCursorMode(ebiten.CursorModeCaptured)
		g.cursorX, g.cursorY = ebiten.CursorPosition()
	}

	if g.Paused {
		return nil
	}
	if ebiten.IsKeyPressed(ebiten.Key1) {
		g.represntation = 0
		g.updateRenders = true
	} else if ebiten.IsKeyPressed(ebiten.Key2) {
		g.represntation = 1
		g.updateRenders = true
	} else if ebiten.IsKeyPressed(ebiten.Key3) {
		g.represntation = 2
		g.updateRenders = true
	}

	// calculate mouse deltas
	newCursorX, newCursorY := ebiten.CursorPosition()
	deltaX := g.cursorX - newCursorX
	deltaY := g.cursorY - newCursorY

	// Update y, to look up or down
	if deltaY != 0 && g.cursorY != 0 {
		g.player.IncreaseHozAngle(float64(deltaY))
	}

	// update mouse position
	g.cursorX, g.cursorY = newCursorX, newCursorY

	// look left or right with mouse
	xMultiplier := 0.
	if deltaX != 0 {
		xMultiplier += float64(deltaX) / -2. * raycasting.DEG_TO_RAD
	}
	if deltaX != 0 {
		g.player.IncreaseAngle(xMultiplier)
	} 

	// move forward or backwards with keyboard
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		ray, err := raycasting.CastRay(*g.player.Coord, g.player.Angle, BLOCK_SIZE, g.mab)
		if err == nil && ray.Coord.DistanceTo(*g.player.Coord) > BLOCK_SIZE/2 {
			g.player.Move(1)
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		ray, err := raycasting.CastRay(*g.player.Coord, raycasting.NormalizeAngle(g.player.Angle+raycasting.PI), BLOCK_SIZE, g.mab)
		if err == nil && ray.Coord.DistanceTo(*g.player.Coord) > BLOCK_SIZE/2 {
			g.player.Move(-1)
		}
	}

	// strafe
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		angle := -raycasting.PI_HALF
		ray, err := raycasting.CastRay(*g.player.Coord, raycasting.NormalizeAngle(g.player.Angle+angle), BLOCK_SIZE, g.mab)
		if err == nil && ray.Coord.DistanceTo(*g.player.Coord) > BLOCK_SIZE/2 {
			g.player.MoveWithAngle(1, angle)
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		angle := raycasting.PI_HALF
		ray, err := raycasting.CastRay(*g.player.Coord, raycasting.NormalizeAngle(g.player.Angle+angle), BLOCK_SIZE, g.mab)
		if err == nil && ray.Coord.DistanceTo(*g.player.Coord) > BLOCK_SIZE/2 {
			g.player.MoveWithAngle(1, angle)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.player.IncreaseAngle(-0.04)
	} else if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.player.IncreaseAngle(0.04)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	coords := make([]raycasting.Coordinate, NO_OF_RAYS)
	rayDistances := make([]float32, NO_OF_RAYS)
	directions := make([]raycasting.Direction, NO_OF_RAYS)
	rays := make([]raycasting.Ray, NO_OF_RAYS)
	for i := -DEG_BOUNDS; i <= DEG_BOUNDS; i++ {
		rayAngle := raycasting.NormalizeAngle(g.player.Angle + (float64(i) * DEG_PER_RAY * raycasting.DEG_TO_RAD))
		ray, err := raycasting.CastRay(*g.player.Coord, rayAngle, BLOCK_SIZE, g.mab)
		if err != nil {
			continue
		}
		coords[i+DEG_BOUNDS] = ray.Coord
		noFish := math.Cos(raycasting.NormalizeAngle(rayAngle - g.player.Angle))
		rayDistances[i+DEG_BOUNDS] = float32(ray.Coord.DistanceTo(*g.player.Coord) * noFish)
		directions[i+DEG_BOUNDS] = ray.Dir
		rays[i+DEG_BOUNDS] = *ray
	}

	if g.represntation == 0 {
		if g.updateRenders {
			width := screen.Bounds().Dx()
			height := screen.Bounds().Dy()
			twoDScreen := screen.SubImage(image.Rect(0, 0, width/2, height)).(*ebiten.Image)
			threeDScreen := screen.SubImage(image.Rect(width/2, 0, width, height)).(*ebiten.Image)
			g.r2d = rendering.NewRenderer2D(twoDScreen, WORLD_WIDTH, WORLD_HEIGHT, BLOCK_SIZE, g.player, g.mab)
			g.r3d = rendering.NewRenderer3D(threeDScreen, g.player, NO_OF_RAYS, BLOCK_SIZE)
			g.updateRenders = false
		}

		g.r3d.Render(rays)
		g.r2d.Render(rays)
	} else if g.represntation == 1 {
		if g.updateRenders {
			g.r3d = rendering.NewRenderer3D(screen, g.player, NO_OF_RAYS, BLOCK_SIZE)
			g.updateRenders = false
		}
		g.r3d.Render(rays)
	} else if g.represntation == 2 {
		if g.updateRenders {
			twoDScreen := screen.SubImage(image.Rect(0, 0, 300, 300)).(*ebiten.Image)
			g.r2d = rendering.NewRenderer2D(twoDScreen, WORLD_WIDTH, WORLD_HEIGHT, BLOCK_SIZE, g.player, g.mab)
			g.r3d = rendering.NewRenderer3D(screen, g.player, NO_OF_RAYS, BLOCK_SIZE)
			g.updateRenders = false
		}

		g.r3d.Render(rays)
		g.r2d.Render(rays)
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
	// initialize some ebiten options
	ebiten.SetWindowSize(1600, 800)
	ebiten.SetCursorMode(ebiten.CursorModeCaptured)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// ebiten.SetFullscreen(true)
	ebiten.SetScreenClearedEveryFrame(false)

	// create some map
	mab := makeStandardMap()
	mab[1][7] = 2
	mab[2][7] = 2
	mab[3][7] = 2
	mab[4][7] = 2
	mab[4][8] = 2
	mab[4][9] = 2
	mab[4][10] = 2
	mab[4][11] = 2
	mab[4][12] = 2
	mab[4][13] = 2
	mab[4][14] = 2
	mab[13][14] = 2
	mab[14][14] = 2
	mab[15][14] = 2
	mab[16][14] = 2
	mab[16][13] = 2
	mab[16][12] = 2
	mab[16][11] = 2

	// create the game struct
	game := &Game{
		player: &player.Player{
			Coord: &raycasting.Coordinate{
				X: WORLD_WIDTH / 2.,
				Y: WORLD_HEIGHT / 2.,
			},
			Angle: 0,
			Speed: 10.,
		},
		mab:           mab,
		represntation: 2,
		updateRenders: true,
	}

	// run the main loop
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
