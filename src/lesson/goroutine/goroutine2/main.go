// for文で投入するとなぜか3番目から実行される
package main

import (
	"log"
	"time"
)

func main() {
	log.Print("started.")

	finished := make(chan bool)

	// 配列
	funcs := []func(){
		func() {
			log.Print("sleep1 started.")
			time.Sleep(1 * time.Second)
			log.Print("sleep1 finished.")
			finished <- true
		},
		func() {
			log.Print("sleep2 started.")
			time.Sleep(2 * time.Second)
			log.Print("sleep2 started.")
			finished <- true
		},
		func() {
			log.Print("sleep3 started.")
			time.Sleep(3 * time.Second)
			log.Print("sleep3 finished.")
			finished <- true
		},
	}

	/*
		for range文にて指数を破棄し、func()をsleepに投入しgo sleep()
	*/
	for _, sleep := range funcs {
		go sleep()
	}

	// 終わるまで待つ
	for i := 0; i < len(funcs); i++ {
		<-finished
	}
	log.Print("all finished.")
}
