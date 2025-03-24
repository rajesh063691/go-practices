package main

import (
	"fmt"
	"time"
)

func PingPong2(ch chan string) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			ch <- "ping"
		} else {
			ch <- "pong"
		}
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func main_p() {
	ch := make(chan string)

	go PingPong2(ch)

	for msg := range ch {
		println(msg)
	}

	fmt.Println("Done")
}
