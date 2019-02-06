package main

import "fmt"

// ①引数にint型をとり、int型を返すfunc()をCallbackと命名
type Callback func(i int) int

// Sum()は第2引数に①をとっている
func Sum(ints []int, callback Callback) int {
	var sum int
	for _, i := range ints {
		sum += i
	}
	return callback(sum)
}

func main() {
	n := Sum(
		[]int{1, 2, 3, 4, 5},
		func(i int) int {
			return i * 2
		},
	)
	fmt.Println(n)
}
