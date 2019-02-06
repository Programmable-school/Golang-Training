package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	// 値として、pに代入
	p := Person{
		Name: "Stieve",
		Age:  20,
	}

	fmt.Printf("first p :%+v\n", p)

	p2 := p
	p2.Name = "Mike"
	p2.Age = 21
	// pではなく
	fmt.Printf("p2で書き換えを行ったはずのp :%+v\n", p)

	p3 := &p
	p3.Name = "二郎"
	p3.Age = 21

	fmt.Printf("p3で二郎に書き換えを行なったp :%+v\n", p3)

}
