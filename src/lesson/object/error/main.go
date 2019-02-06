package main

import (
	"errors"
	"fmt"
	"os"
)

func calc(atai1 int, atai2 int) (ans int, err error) {
	if atai1 == 0 || atai2 == 0 {
		return 0, errors.New("不正な数値が設定されてます")
	}
	ans = atai1 + atai2
	return ans, nil
}

func main() {
	ans, err := calc(0, 20)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー:%d", err)
		os.Exit(1)
	}
	fmt.Println(ans)
}
