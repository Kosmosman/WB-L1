package point

import "math"

type Point struct {
	x float64
	y float64
}

func Distance(p, other Point) float64 {
	divX := p.x - other.x
	divY := p.y - other.y
	return math.Sqrt(divX*divX + divY*divY)
}

func NewPoint(x float64, y float64) Point {
	return Point{x: x, y: y}
}
