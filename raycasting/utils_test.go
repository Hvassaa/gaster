package raycasting

import (
	"math"
	"testing"
)


func closeTo(a, b float64) bool {
	return math.Abs(a-b) <= 0.01
}

func TestHorizontalRayAngleInvalidAngles(t *testing.T) {
	c := Coordinate{1, 1}
	blockSize := 1.
	m := make([][]WallType, 20)
	for i := range m {
		m[i] = make([]WallType, 20)
	}

	t.Run("Looking directly right", func(t *testing.T) {
		angle := 0.
		c, _, _ := castRayHorizontal(c, angle, blockSize, m)
		if !c.IsInvalid() {
			t.Fatalf("Coordinate should be invalid with angle %v, got: %v", angle, c)
		}
	})

	t.Run("Looking directly left", func(t *testing.T) {
		angle := PI
		c, _, _ := castRayHorizontal(c, angle, blockSize, m)
		if !c.IsInvalid() {
			t.Fatalf("Coordinate should be invalid with angle %v, got: %v", angle, c)
		}
	})
}

func TestHorizontalRay(t *testing.T) {
	c := Coordinate{7.5, 7.5}
	blockSize := 5.
	// 1 1 1
	// 1 0 1
	// 1 1 1
	m := make([][]WallType, 3)
	for i := range m {
		m[i] = make([]WallType, 3)
		for j := range m[i] {
			m[i][j] = 1
		}
	}
	m[1][1] = 0

	t.Run("Looking directly up", func(t *testing.T) {
		angle := PI_HALF
		expectedX, expectedY := 7.5, 10.
		c, _, _ := castRayHorizontal(c, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking directly down", func(t *testing.T) {
		angle := PI_THREE_HALF
		expectedX, expectedY := 7.5, 5.
		c, _, _ := castRayHorizontal(c, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking left-up", func(t *testing.T) {
		angle := PI_HALF + PI_HALF/2
		expectedX, expectedY := 5., 10.
		c, _, _ := castRayHorizontal(c, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking right-up", func(t *testing.T) {
		angle := PI_HALF - PI_HALF/2
		expectedX, expectedY := 10., 10.
		c, _, _ := castRayHorizontal(c, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking down-left", func(t *testing.T) {
		angle := PI + PI_HALF/2
		expectedX, expectedY := 5., 5.
		c, _, _ := castRayHorizontal(c, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking down-right", func(t *testing.T) {
		angle := PI_THREE_HALF + PI_HALF/2
		expectedX, expectedY := 10., 5.
		c, _, _ := castRayHorizontal(c, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})
}

func TestVerticalRayAngleInvalidAngles(t *testing.T) {
	c := Coordinate{1, 1}
	blockSize := 1.
	m := make([][]WallType, 20)
	for i := range m {
		m[i] = make([]WallType, 20)
	}

	t.Run("Looking directly up", func(t *testing.T) {
		angle := PI_HALF
		c, _, _ := castRayVertical(c, angle, blockSize, m)
		if !c.IsInvalid() {
			t.Fatalf("Coordinate should be invalid with angle %v, got: %v", angle, c)
		}
	})

	t.Run("Looking directly down", func(t *testing.T) {
		angle := PI_THREE_HALF
		c, _, _ := castRayVertical(c, angle, blockSize, m)
		if !c.IsInvalid() {
			t.Fatalf("Coordinate should be invalid with angle %v, got: %v", angle, c)
		}
	})
}

func TestVerticalRay(t *testing.T) {
	c := Coordinate{7.5, 7.5}
	blockSize := 5.
	// 1 1 1
	// 1 0 1
	// 1 1 1
	m := make([][]WallType, 3)
	for i := range m {
		m[i] = make([]WallType, 3)
		for j := range m[i] {
			m[i][j] = 1
		}
	}
	m[1][1] = 0

	t.Run("Looking directly left", func(t *testing.T) {
		angle := PI
		expectedX, expectedY := 5., 7.5
		c, _, _ := castRayVertical(c, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking directly right", func(t *testing.T) {
		angle := 0.
		expectedX, expectedY := 10., 7.5
		c, _, _ := castRayVertical(c, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking left-up", func(t *testing.T) {
		angle := PI_HALF + PI_HALF/2
		expectedX, expectedY := 5., 10.
		c, _, _ := castRayVertical(c, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking right-up", func(t *testing.T) {
		angle := PI_HALF - PI_HALF/2
		expectedX, expectedY := 10., 10.
		c, _, _ := castRayVertical(c, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking down-left", func(t *testing.T) {
		angle := PI + PI_HALF/2
		expectedX, expectedY := 5., 5.
		c, _, _ := castRayVertical(c, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking down-right", func(t *testing.T) {
		angle := PI_THREE_HALF + PI_HALF/2
		expectedX, expectedY := 10., 5.
		c, _, _ := castRayVertical(c, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})
}

func TestAsd(t *testing.T) {
	blockSize := 5.
	// 1 1 1
	// 1 1 1
	m := make([][]WallType, 2)
	for i := range m {
		m[i] = make([]WallType, 3)
		for j := range m[i] {
			m[i][j] = 1
		}
	}

	t.Run("Looking right up low angle", func(t *testing.T) {
		cord := Coordinate{0, 0}
		angle := 21.04 * DEG_TO_RAD
		expectedX, expectedY := 13., 5.
		c, _, _ := castRayHorizontal(cord, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking left up low angle", func(t *testing.T) {
		cord := Coordinate{15, 0}
		angle := 158.96 * DEG_TO_RAD
		expectedX, expectedY := 2., 5.
		c, _, _ := castRayHorizontal(cord, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking down right high angle", func(t *testing.T) {
		cord := Coordinate{0, 9.99999} // if we are exactly on the line we will hit it looking down
		angle := 338.96 * DEG_TO_RAD
		expectedX, expectedY := 13., 5.
		c, _, _ := castRayHorizontal(cord, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking down right high angle", func(t *testing.T) {
		cord := Coordinate{15, 9.99999} // if we are exactly on the line we will hit it looking down
		angle := 199.65 * DEG_TO_RAD
		expectedX, expectedY := 1., 5.
		c, _, _ := castRayHorizontal(cord, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})
}

func TestHorizontalRaySharpFlatAngles(t *testing.T) {
	blockSize := 5.
	// 1 1 1
	// 1 1 1
	m := make([][]WallType, 2)
	for i := range m {
		m[i] = make([]WallType, 3)
		for j := range m[i] {
			m[i][j] = 1
		}
	}

	t.Run("Looking right up low angle", func(t *testing.T) {
		cord := Coordinate{0, 0}
		angle := 21.04 * DEG_TO_RAD
		expectedX, expectedY := 13., 5.
		c, _, _ := castRayHorizontal(cord, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking left up low angle", func(t *testing.T) {
		cord := Coordinate{15, 0}
		angle := 158.96 * DEG_TO_RAD
		expectedX, expectedY := 2., 5.
		c, _, _ := castRayHorizontal(cord, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking down right high angle", func(t *testing.T) {
		cord := Coordinate{0, 9.99999} // if we are exactly on the line we will hit it looking down
		angle := 338.96 * DEG_TO_RAD
		expectedX, expectedY := 13., 5.
		c, _, _ := castRayHorizontal(cord, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking down right high angle", func(t *testing.T) {
		cord := Coordinate{15, 9.99999} // if we are exactly on the line we will hit it looking down
		angle := 199.65 * DEG_TO_RAD
		expectedX, expectedY := 1., 5.
		c, _, _ := castRayHorizontal(cord, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})
}

func TestVerticalRaySharpFlatAngles(t *testing.T) {
	blockSize := 5.
	// 1 1
	// 1 1
	// 1 1
	m := make([][]WallType, 3)
	for i := range m {
		m[i] = make([]WallType, 2)
		for j := range m[i] {
			m[i][j] = 1
		}
	}

	t.Run("Looking right up high angle", func(t *testing.T) {
		cord := Coordinate{0, 0}
		angle := 67.38 * DEG_TO_RAD
		expectedX, expectedY := 5., 12.
		c, _, _ := castRayVertical(cord, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking left up high angle", func(t *testing.T) {
		cord := Coordinate{9.9999, 0}
		angle := 112.62 * DEG_TO_RAD
		expectedX, expectedY := 5., 12.
		c, _, _ := castRayVertical(cord, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking right down high angle", func(t *testing.T) {
		cord := Coordinate{0, 14.9999}
		angle := 291.04 * DEG_TO_RAD
		expectedX, expectedY := 5., 2.
		c, _, _ := castRayVertical(cord, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking left down high angle", func(t *testing.T) {
		cord := Coordinate{9.9999, 15}
		angle := 248.96 * DEG_TO_RAD
		expectedX, expectedY := 5., 2.
		c, _, _ := castRayVertical(cord, angle, blockSize, m)
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})
}

func TestTheShortestRayWins(t *testing.T) {
	blockSize := 5.
	// 1 1 1 1
	// 1 0 0 1
	// 1 0 0 1
	// 1 1 1 1
	m := make([][]WallType, 4)
	for i := range m {
		m[i] = make([]WallType, 4)
		for j := range m[i] {
			m[i][j] = 1
		}
	}
	m[1][1] = 0
	m[1][2] = 0
	m[2][1] = 0
	m[2][2] = 0

	t.Run("Looking right should hit vertical", func(t *testing.T) {
		cord := Coordinate{8, 6}
		angle := 40.6 * DEG_TO_RAD
		expectedX, expectedY := 15., 12.
		c, d, _ := CastRay(cord, angle, blockSize, m)
		if d != VERTICAL {
			t.Fatalf("Direction should be %v, got %v", VERTICAL.asText(), d.asText())
		}
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})

	t.Run("Looking right should hit horizontal", func(t *testing.T) {
		cord := Coordinate{8, 6}
		angle := 56.31 * DEG_TO_RAD
		expectedX, expectedY := 14., 15.
		c, d, _ := CastRay(cord, angle, blockSize, m)
		if d != HORIZONTAL {
			t.Fatalf("Direction should be %v, got %v", HORIZONTAL.asText(), d.asText())
		}
		if !(closeTo(c.X, expectedX) && closeTo(c.Y, expectedY)) {
			t.Fatalf("Coordinate should be (%v, %v) with %v, got: %g", expectedX, expectedY, angle, c)
		}
	})
}
