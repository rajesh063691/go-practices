package main

import (
	"fmt"
	"strconv"
)

// func channel1(ch chan int) {
// 	fmt.Println(23 + <-ch)
// }

func mychannel(ch chan string) {
	for i := 0; i < 4; i++ {
		ch <- "Hello_" + strconv.Itoa(i)
	}
	close(ch)
}

// sending string data to channel
func sending(ch chan string) {
	ch <- "Rajesh Kumar"
	ch <- "Yadav"
	close(ch)
}

func main_1() {
	fmt.Println("Start main")
	// ch := make(chan int)
	// go channel1(ch)
	// ch <- 7

	// ch := make(chan string)
	// go mychannel(ch)
	//select {
	// case str := <-ch:
	// 	fmt.Println(str)
	// // default:
	// // 	fmt.Println("no channel found")
	// }

	// for {
	// 	str, ok := <-ch
	// 	if !ok {
	// 		fmt.Println("channel closed :: ", ok)
	// 		break
	// 	}

	// 	fmt.Println("channel open :: ", str, ok)
	// }

	// for str := range ch {
	// 	fmt.Println("receives value from channel :: ", str)
	// }
	// receives data from channel
	// channel1 := make(<-chan bool)
	// // sends data to channel
	// channel2 := make(chan<- bool)

	// fmt.Printf("%T \n", channel1)
	// fmt.Printf("%T \n", channel2)

	// creating bidirectional channel
	ch1 := make(chan string, 1)

	go sending(ch1)

	for str := range ch1 {
		fmt.Println(str)
	}

	fmt.Println("End main")
}
