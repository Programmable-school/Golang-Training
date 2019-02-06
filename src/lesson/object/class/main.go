package main

import (
	"log"
)

type Money struct {
	amount   uint
	currency string
}

func main() {

	var money *Money
	money = &Money{
		amount:   120,
		currency: "yen",
	}
	log.Printf("money :%+v\n", money)
	log.Printf("変数に格納されているアドレス :%p", money)
}
