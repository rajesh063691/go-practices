package main

import (
	"fmt"
	"time"
)

func getChars2(s string) {
	for _, c := range s {
		fmt.Printf("%c ", c)
	}
	fmt.Println()
}

func getDigits2(s []int) {
	for _, d := range s {
		fmt.Printf("%d ", d)
	}
	fmt.Println()
}

func main_2() {
	fmt.Println("main execution started")

	// getChars goroutine
	go getChars2("Hello")

	// getDigits goroutine
	go getDigits2([]int{1, 2, 3, 4, 5})

	// schedule another goroutine
	time.Sleep(time.Millisecond)

	fmt.Println("\nmain execution stopped")
}
