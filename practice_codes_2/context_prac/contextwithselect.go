package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Work in progress...")
		case <-ctx.Done():
			// If the context deadline is exceeded or canceled, this case will be triggered
			fmt.Println("Context deadline exceeded:", ctx.Err())
			return
		}
	}
}

func main_se() {
	// Create a context with a 5-second deadline
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel() // Ensure the context is canceled when we are done

	fmt.Println("Starting work...")

	// Start the work in a goroutine
	go doWork(ctx)

	// Wait for 7 seconds to observe what happens
	time.Sleep(7 * time.Second)

	fmt.Println("Main function finished.")
}
