package main

import (
	"fmt"
	"time"
)

func main__() {
	go spinner(100 * time.Millisecond)
	const n = 5
	fibN := fib(n) // slow
	fmt.Println()
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/-` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
