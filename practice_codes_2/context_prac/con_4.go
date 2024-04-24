package main

import (
	"context"
	"fmt"
	"time"
)

func main_4() {
	// data := `{
	// 	"Name":"John",
	// 	"Age":30,
	// 	"City":"New York",
	// 	"Country":"USA"
	// }`

	ctx := context.Background()
	ctx = context.WithValue(ctx, "UserID", 123)

	go doSomeTask(ctx)
	time.Sleep(2 * time.Second)
}

func doSomeTask(ctx context.Context) {
	userID := ctx.Value("UserID").(int)
	fmt.Println("User ID:", userID)
}
