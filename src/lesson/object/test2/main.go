package main

import "fmt"

func main() {
	var i int
	i = 5
	p := &i

	//fmt.Printf("type=%T, address=%p\n", i, i)
	//fmt.Printf("type=%T, address=%p\n", &i, &i)
	//fmt.Printf("type=%T, address=%p\n", p, p)
	//fmt.Println(*p)

	q := *p + 5
	//fmt.Printf("type=%T, address=%p, int=%d\n", i, i, i)
	//fmt.Println(*p)
	//fmt.Println(p)
	//fmt.Println(i)
	//fmt.Println(*p)
	fmt.Println(q)
	fmt.Println(*p)

	//p := &i

}
