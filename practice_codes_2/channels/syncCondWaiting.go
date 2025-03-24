package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(cond *sync.Cond, ready *bool) {
	cond.L.Lock()

	if !*ready {
		fmt.Println("Worker waiting!")
		cond.Wait() // Wait for the condition to be signaled
	}

	cond.L.Unlock()
	fmt.Println("Worker resumed!")
}

func main_cond() {
	var mu sync.Mutex
	ready := false

	cond := sync.NewCond(&mu)

	// Start a goroutine
	go worker(cond, &ready)

	time.Sleep(5 * time.Second)

	// Simulate some work
	fmt.Println("Preparing to signal worker...")
	cond.L.Lock()
	ready = true
	cond.L.Unlock()
	cond.Signal()

}
