package main

import "fmt"

// sendToChannel sends data to channel
func Ping(pings chan<- string, message string) {
	for i := 0; i < 5; i++ {
		pings <- message + fmt.Sprint("_") + fmt.Sprint(i)
	}

	close(pings)
}

// receives data from channel
func Pong(pings <-chan string, pongs chan<- string) {
	for i := 0; i < 5; i++ {
		msg := <-pings
		pongs <- msg
	}

	close(pongs)
}

// create basic channel functionality
func main_4() {
	ping := make(chan string, 1)
	pong := make(chan string, 1)
	//for i := 0; i < 5; i++ {
	go Ping(ping, "Ping")
	go Pong(ping, pong)
	//}
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(<-pong)
	// }

	for val := range pong {
		fmt.Println(val)
	}
}
