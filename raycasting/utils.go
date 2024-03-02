package raycasting

import (
	"math"
)

const (
	PI            = math.Pi
	PI_HALF       = PI / 2
	PI_THREE_HALF = PI_HALF * 3
	PI_TWO        = PI * 2
)

type WallType int
type Direction int

const (
	HORIZONTAL Direction = iota
	VERTICAL
)

const DEG_TO_RAD = 0.0174532925

type Coordinate struct {
	X, Y float64
}

func (c Coordinate) IsInvalid() bool {
	return math.IsNaN(c.X) || math.IsNaN(c.Y)
}

func (c Direction) asText() string {
	switch c {
	case VERTICAL:
		return "VERTICAL"
	case HORIZONTAL:
		return "HORIZONTAL"
	}
	panic("unknown direction")
}

func (c Coordinate) distanceTo(c2 Coordinate) float64 {
	xDiff := math.Abs(c.X - c2.X)
	yDiff := math.Abs(c.Y - c2.Y)
	return math.Sqrt(xDiff*xDiff + yDiff*yDiff)
}

func NormalizeAngle(angle float64) float64 {
	if angle > PI_TWO {
		angle -= PI_TWO
	} else if angle < 0 {
		angle += PI_TWO
	}
	return angle
}

func keepCasting(ix, iy, xOffset, yOffset, blockSize float64, direction Direction, m [][]WallType) (Coordinate, WallType) {
	y_size := len(m)
	x_size := len(m[0])
	for i := 0; i < 100; i++ {
		// translate the position to indices in the map
		xx := int(math.Floor(ix / blockSize))
		yy := int(math.Floor(iy / blockSize))

		// we adjust the index, to look at the "previous"
		// instead of the current if we have a negative offset
		if direction == HORIZONTAL && yOffset < 0 {
			yy -= 1
		}
		if direction == VERTICAL && xOffset < 0 {
			xx -= 1
		}

		if xx < 0 || xx >= x_size || yy < 0 || yy >= y_size {
			return Coordinate{math.NaN(), math.NaN()}, 0
		}

		wallType := m[yy][xx]
		// TODO we might add more walltypes later, then we should switch instead
		if wallType != 0 {
			return Coordinate{ix, iy}, wallType
		}

		// update to check the next wall
		ix += xOffset
		iy += yOffset
	}
	return Coordinate{math.NaN(), math.NaN()}, 0
}

func castRayHorizontal(coordinate Coordinate, angle, blockSize float64, m [][]WallType) (Coordinate, Direction, WallType) {
	// resulting intersection on ix, iy
	var ix, iy float64

	a := -1 / math.Tan(angle)
	x_offset := blockSize * a
	y_offset := blockSize
	iy = math.Floor(coordinate.Y/blockSize) * blockSize
	if angle == 0. || angle == PI {
		return Coordinate{math.NaN(), coordinate.Y}, HORIZONTAL, 0
	} else if angle > PI { // looking down
		// iterate down
		y_offset *= -1
	} else if angle < PI { // looking up
		// we go up a block to look at the top
		iy += blockSize
		x_offset *= -1
	}

	ix = (coordinate.Y-iy)*a + coordinate.X
	coordinate, wallType := keepCasting(ix, iy, x_offset, y_offset, blockSize, HORIZONTAL, m)
	return coordinate, HORIZONTAL, wallType
}

func castRayVertical(coordinate Coordinate, angle, blockSize float64, m [][]WallType) (Coordinate, Direction, WallType) {
	// resulting intersection on ix, iy
	var ix, iy float64

	a := -math.Tan(angle)
	x_offset := blockSize
	y_offset := blockSize * a
	ix = math.Floor(coordinate.X/blockSize) * blockSize
	if angle == PI_HALF || angle == PI_THREE_HALF {
		return Coordinate{coordinate.X, math.NaN()}, HORIZONTAL, 0
	} else if angle < PI_HALF || angle > PI_THREE_HALF { // looking right
		// we go block right to look at the right side
		ix += blockSize
		y_offset *= -1
	} else if angle > PI_HALF { // looking left
		// iterate left
		x_offset *= -1
	}

	iy = (coordinate.X-ix)*a + coordinate.Y
	coordinate, wallType := keepCasting(ix, iy, x_offset, y_offset, blockSize, VERTICAL, m)
	return coordinate, VERTICAL, wallType
}

func CastRay(coordinate Coordinate, angle, blockSize float64, m [][]WallType) (Coordinate, Direction, WallType) {

	// return castRayVertical(coordinate, angle, blockSize, m)

	ch, dh, wh := castRayHorizontal(coordinate, angle, blockSize, m)
	l1 := ch.distanceTo(coordinate)
	cv, dv, wv := castRayVertical(coordinate, angle, blockSize, m)
	l2 := cv.distanceTo(coordinate)
	if ch.IsInvalid() {
		return cv, dv, wv
	}
	if cv.IsInvalid() {
		return ch, dh, wh
	}
	if l1 > l2 {
		return cv, dv, wv
	}
	return ch, dh, wh
}
