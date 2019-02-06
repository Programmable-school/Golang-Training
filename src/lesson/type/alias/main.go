package main

import "fmt"

// int型宣言をエイリアス、他の変数にもint型を付与可能
type Myint int

// 一見すると、意味がないように見えるが複雑な型宣言に有用
type (
	StringPair [2]string
)

func main() {
	var n1 int = 1
	var n2 Myint = 5

	fmt.Println(n1)
	fmt.Println(n2)

	str1 := StringPair{"Japan", "Samurai"}
	str2 := StringPair{"America", "Cawboy"}

	fmt.Println(str1)
	fmt.Println(str2)

}
