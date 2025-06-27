package gg

import (
	"math"

	"golang.org/x/image/math/fixed"
)

type Point struct {
	X, Y float64
}

func (a Point) Fixed() fixed.Point26_6 {
	return fixPoint(a.X, a.Y, false)
}

func (a Point) Distance(b Point) float64 {
	return math.Hypot(a.X-b.X, a.Y-b.Y)
}

func (a Point) Interpolate(b Point, t float64) Point {
	x := a.X + (b.X-a.X)*t
	y := a.Y + (b.Y-a.Y)*t
	return Point{x, y}
}
