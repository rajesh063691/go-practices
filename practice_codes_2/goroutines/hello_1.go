package main

import (
	"fmt"
	"time"
)

func printHello() {
	time.Sleep(15 * time.Millisecond)
	fmt.Println("Hello World!")
}

func main_1() {
	fmt.Println("main execution started")

	// create goroutine
	go printHello()

	time.Sleep(10 * time.Millisecond)

	fmt.Println("main execution stopped")
}
