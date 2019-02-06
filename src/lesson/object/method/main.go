package main

import (
	"fmt"
	"log"
)

type Money struct {
	amount   uint
	currency string
}

func (this *Money) Format() string {
	return fmt.Sprintf("%d %s", this.amount, this.currency)
}

func (this *Money) Add(that *Money) {
	this.amount += that.amount
}

func main() {
	money := &Money{120, "yen"}
	log.Print(money.Format())
	money.Add(&Money{180, "yen"})
	log.Print(money.Format())
}
