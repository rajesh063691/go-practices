package main

import "fmt"

// it only sends to channel
func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

// it recives and sends
func squarer(out <-chan int, in chan<- int) {
	for i := range out {
		in <- i * i
	}
	close(in)
}

func print(sq chan int) {
	for val := range sq {
		fmt.Println(val)
	}
}

func main_pip1() {

	workerCount := uint64(2)
	fmt.Println(workerCount)
	co := make(chan int)
	sq := make(chan int)

	go counter(co)
	go squarer(co, sq)
	print(sq)
}
