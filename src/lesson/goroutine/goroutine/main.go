package main

import (
	"log"
	"time"
)

/* 6病かかる処理
func main() {
	log.Print("started.")

	// 1病かかるコマンド
	log.Print("sleep1 started.")
	time.Sleep(1 * time.Second)
	log.Print("sleep1 finished.")

	// 2秒かかるコマンド
	log.Print("sleep2 started.")
	time.Sleep(2 * time.Second)
	log.Print("sleep2 finished.")

	// 3病かかるコマンド
	log.Print("sleep3 started.")
	time.Sleep(3 * time.Second)
	log.Print("sleep3 finished.")

	log.Print("all finished.")
}

*/

// チャネルを１つにまとめ、3回finishedカウントさせて終了
func main() {
	log.Print("started.")

	// チャネル
	finished := make(chan bool)

	go func() {
		log.Print("sleep1 started.")
		time.Sleep(1 * time.Second)
		log.Print("sleep1 finished.")
		finished <- true
	}()

	go func() {
		log.Print("sleep2 started.")
		time.Sleep(2 * time.Second)
		log.Print("sleep2 finished.")
		finished <- true
	}()

	go func() {
		log.Print("sleep3 started.")
		time.Sleep(3 * time.Second)
		log.Print("sleep3 finished.")
		finished <- true
	}()

	for i := 0; i < 3; i++ {
		<-finished
	}

	log.Print("all finished.")
}
