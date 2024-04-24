package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main execution started")
	// create goroutine
	go func() {
		fmt.Println("Hello World!")
	}()
	// schedule another goroutine
	time.Sleep(10 * time.Millisecond)
	fmt.Println("main execution stopped")
}
