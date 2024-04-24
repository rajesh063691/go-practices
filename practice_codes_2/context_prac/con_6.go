package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	go performTask2(ctx)

	time.Sleep(6 * time.Second)
}

func performTask2(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Task completed or deadline exceeded:", ctx.Err())
		return
	}
}
