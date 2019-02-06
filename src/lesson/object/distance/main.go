package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) Distance(dp Point) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

func main() {
	p := Point{X: 0, Y: 0}
	dis := p.Distance(Point{X: 1, Y: 1})

	fmt.Println(dis)
}
