package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(other *Point) float64 {
	xDiff := p.x - other.x
	yDiff := p.y - other.y

	return math.Sqrt(xDiff*xDiff + yDiff*yDiff)
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 0)
	p3 := NewPoint(0, 4)

	// 3(p2)
	// 2
	// 1
	// 0(p1)	1	2	3	4(p3)

	fmt.Printf("bw p1 and p2: %f\n", p1.Distance(p2))
	// expected: 3
	// sqrt( (0-3)^2 + (0-0)^2) = sqrt(3^2) = 3
	fmt.Printf("bw p1 and p3: %f\n", p1.Distance(p3))
	// expected: 4
	// sqrt( (0-0)^2 + (0-4)^2) = sqrt(4^2) = 4
	fmt.Printf("bw p2 and p3: %f\n", p2.Distance(p3))
	// expected: 5
	// sqrt( (3-0)^2 + (0-4)^2) = sqrt(3^2 + 4^2) = sqrt(9 + 16) = sqrt(25) = 5
}
