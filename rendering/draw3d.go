package rendering

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hvassaa/gaster/player"
	"github.com/hvassaa/gaster/raycasting"
)

type Renderer3D struct {
	TopColor, BottomColor                             color.Color
	BlockSize                                         float64
	WallColors                                        map[raycasting.WallType]color.Color
	Screen                                            *ebiten.Image
	ScreenMid, ColumnWidth, ScreenWidth, ScreenHeight float32
	texture                                           map[uint]Texture
	Player                                            *player.Player
}

func NewRenderer3D(screen *ebiten.Image, player *player.Player, noOfRays int, blockSize float64) *Renderer3D {
	wallColors := make(map[raycasting.WallType]color.Color)
	wallColors[0] = color.RGBA{255, 0, 0, 255}
	wallColors[1] = color.RGBA{155, 0, 0, 255}

	screenHeight := float32(screen.Bounds().Size().Y)
	screenWidth := float32(screen.Bounds().Size().X)

	return &Renderer3D{
		TopColor:     color.RGBA{50, 150, 150, 255},
		BottomColor:  color.RGBA{200, 200, 200, 255},
		BlockSize:    blockSize,
		WallColors:   wallColors,
		Screen:       screen,
		ScreenHeight: screenHeight,
		ScreenWidth:  screenWidth,
		ScreenMid:    screenHeight / 2.,
		ColumnWidth:  screenWidth / float32(noOfRays),
		Player:       player,
		texture: map[uint]Texture{
			1: LoadTexture(CROSS_TEXTURE),
			2: LoadTexture(ASD),
		},
	}
}

func (r3d *Renderer3D) Render(rays []raycasting.Ray) {
	xStart := r3d.Screen.Bounds().Min.X
	// we render walls "half up and down" from this point
	// we initially set it to the middle of the screen
	renderMiddle := r3d.ScreenMid + float32(r3d.Player.HozAngle)*r3d.ScreenHeight*3/180

	for i, ray := range rays {
		columnColor := color.RGBA{0, 0, 0, 255}
		b := r3d.texture[uint(rays[i].Wt)]
		coord := ray.Coord
		var xPosOnBlock float64
		if ray.Dir == raycasting.HORIZONTAL {
			xPosOnBlock = math.Mod(coord.X, r3d.BlockSize)
			columnColor.R = 50
		} else {
			xPosOnBlock = math.Mod(coord.Y, r3d.BlockSize)
		}

		// this avoid fisheye
		noFish := math.Cos(raycasting.NormalizeAngle(r3d.Player.Angle - ray.Ang)) * ray.Dist
		columnHeight := float32((r3d.BlockSize * float64(r3d.ScreenHeight)) / noFish)

		x := float32(xStart) + float32(i)*r3d.ColumnWidth
		yTextureListSize := len(b)
		xTextureListSize := len(b[0])
		xTextureSliceSize := r3d.BlockSize / float64(xTextureListSize)
		xTextureIdx := int(math.Floor(xPosOnBlock / xTextureSliceSize))
		top := renderMiddle - columnHeight/2
		bot := renderMiddle + columnHeight/2
		vertSlice := columnHeight / float32(yTextureListSize)

		// draw top and bottom colors
		vector.StrokeLine(r3d.Screen, x, bot, x, r3d.ScreenHeight, r3d.ColumnWidth, r3d.BottomColor, false)
		vector.StrokeLine(r3d.Screen, x, top, x, 0, r3d.ColumnWidth, r3d.TopColor, false)

		for j := 0; j < yTextureListSize; j++ {
			fj := float64(j)
			var y1 float32 = top + vertSlice*float32(fj)
			var y2 float32 = y1 + vertSlice
			yTextureIdx := j
			angle := ray.Ang
			leftWallHit := angle > raycasting.PI_HALF && angle < raycasting.PI_THREE_HALF && ray.Dir == raycasting.VERTICAL
			bottomWallHit := angle < raycasting.PI && ray.Dir == raycasting.HORIZONTAL
			if leftWallHit || bottomWallHit {
				// when the way is left or down, the texture is x-mirrored
				columnColor.B = b[yTextureIdx][xTextureListSize-xTextureIdx-1]
			} else {
				columnColor.B = b[yTextureIdx][xTextureIdx]
			}
			vector.StrokeLine(r3d.Screen, x, y1, x, y2, r3d.ColumnWidth, columnColor, false)
		}

		// if i%4 == 0 {
		// topColor := color.RGBA{0, 255, 0, 255}
		// 	vector.StrokeLine(r3d.Screen, x, bot, x+float32(math.Cos(ray.Ang))*columnHeight, r3d.ScreenHeight, r3d.ColumnWidth, topColor, false)
		// }
	}

	// for i, ray := range rays {
	// 	// this avoid fisheye on "right ahead walls"
	// 	noFish := math.Cos(raycasting.NormalizeAngle(r3d.Player.Angle - ray.Ang))
	// 	columnHeight := float32((r3d.BlockSize * float64(r3d.ScreenHeight)) / (ray.Dist * noFish))
	// 	x := float32(xStart) + float32(i)*r3d.ColumnWidth
	// 	bot := renderMiddle + columnHeight/2
	// 	// if i == 1 || i == len(rays)-1 {
	// 	topColor := color.RGBA{0, 255, 0, 255}
	// 	vector.StrokeCircle(r3d.Screen, x, bot, 10, 2, color.White, false)
	// 	vector.StrokeCircle(r3d.Screen, r3d.ScreenWidth/2, r3d.ScreenHeight, 10, 2, color.White, false)
	// 	vector.StrokeLine(r3d.Screen, x, bot, r3d.ScreenWidth/2, r3d.ScreenHeight, 1, topColor, false)
	// 	vector.StrokeLine(r3d.Screen, r3d.ScreenWidth/2, r3d.ScreenHeight, x, bot, 1, topColor, false)
	// 	vector.StrokeLine(r3d.Screen, x, bot, x, r3d.ScreenHeight, 1, color.Black, false)
	// 	// vector.StrokeLine(r3d.Screen, x, bot, r3d.ScreenWidth/2, r3d.ScreenHeight, 1, topColor, false)
	// 	// }
	// }
}
