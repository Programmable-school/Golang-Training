package main

import "fmt"

type Point struct {
	X, Y int
}

func swap(p Point) {
	/* フィールドX, Yの値を入れ替える */
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

func swapp(p *Point) {
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

func main() {

	p := Point{X: 1, Y: 2}
	swap(p)

	fmt.Println(p.X)
	fmt.Println(p.Y)

	swapp(&p)
	fmt.Println(p.X)
	fmt.Println(p.Y)

}
