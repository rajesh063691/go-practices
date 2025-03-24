package main

import (
	"time"
)

func main_ping_pong() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- "ping"
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			ch2 <- "pong"
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		select {
		case msg1 := <-ch1:
			println(msg1)
		case msg2 := <-ch2:
			println(msg2)
		}
	}

	// for {
	// 	fmt.Println(<-ch1)
	// 	fmt.Println(<-ch2)
	// }
}
