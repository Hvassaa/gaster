package raycasting

import (
	"math"
	"testing"
)

func closeTo(a, b float64) bool {
	return math.Abs(a-b) <= 1e-9
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
		if !c.isInf() {
			t.Fatalf("Coordinate should be invalid with angle %v, got: %v", angle, c)
		}
	})

	t.Run("Looking directly left", func(t *testing.T) {
		angle := PI
		c, _, _ := castRayHorizontal(c, angle, blockSize, m)
		if !c.isInf() {
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
		if !c.isInf() {
			t.Fatalf("Coordinate should be invalid with angle %v, got: %v", angle, c)
		}
	})

	t.Run("Looking directly down", func(t *testing.T) {
		angle := PI_THREE_HALF
		c, _, _ := castRayVertical(c, angle, blockSize, m)
		if !c.isInf() {
			t.Fatalf("Coordinate should be invalid with angle %v, got: %v", angle, c)
		}
	})
}

func TestVerticalRay(t *testing.T) {
	c := Coordinate{ 7.5, 7.5}
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
