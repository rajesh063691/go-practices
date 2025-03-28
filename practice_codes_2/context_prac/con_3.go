package main

import (
	"context"
	"fmt"
	"time"
)

func main_3() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go performTask(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("Task timed out")
	}
}

func performTask(ctx context.Context) {
	_ = ctx
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Task completed successfully")
	}
}
