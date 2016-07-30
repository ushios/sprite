package sprite

import (
	"fmt"
	"math"
)

// Point is only point
type Point struct {
	X, Y int64
}

// String for fmt.Stringer
func (p Point) String() string {
	return fmt.Sprintf("[%d, %d]", p.X, p.Y)
}

// Length is length
type Length struct {
	length float64
}

// Int return length int.
func (l Length) Int() int64 {
	return int64(l.length)
}

// CalcLength get length
func CalcLength(p1, p2 Point) Length {
	x := p1.X - p2.X
	y := p1.Y - p2.Y

	d := math.Sqrt(float64((x * x) + (y * y)))

	return Length{length: d}
}
