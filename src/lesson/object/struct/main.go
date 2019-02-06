package main

import "fmt"

type Test struct {
	a int
	b int
}

func main() {
	var foo Test
	foo.a = 1
	foo.b = 2

	fmt.Println(foo.a)
	fmt.Println(foo.b)

	bar := Test{a: 3, b: 4}

	fmt.Println(bar.a)
	fmt.Println(bar.b)
}
