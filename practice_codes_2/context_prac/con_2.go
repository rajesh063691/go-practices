package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// ques:: have 3 endpoints - write code to send the requests parallely and get the response using channels

func FetchFromUrls(ctx context.Context, wg *sync.WaitGroup, url string, ch chan string) {
	defer wg.Done()
	// check if context is cancelled
	if !isContextAlive(ctx) {
		ch <- "err :: context deadline exceed, can not call endpoints"
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		ch <- "err :: " + err.Error()
		return
	}

	// call the endpoints
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- "err :: " + err.Error()
		return
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		ch <- "err :: " + err.Error()
		return
	}

	ch <- string(resBody)
}

func isContextAlive(ctx context.Context) (ok bool) {
	select {
	case <-ctx.Done():
		{
			fmt.Println("ctx err :: ", ctx.Err().Error())
			return true
		}
	case <-time.After(10 * time.Second):
		{
			fmt.Println("ctx timeout err :: ", ctx.Err().Error())
			return true
		}
	}
}

func mai_2() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/users",
	}

	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// create channel of array size
	ch := make(chan string, len(urls))
	start := time.Now()
	for _, url := range urls {
		wg.Add(1)
		go FetchFromUrls(ctx, wg, url, ch)
	}

	fmt.Println("waiting for all goroutines to return...")
	wg.Wait()
	close(ch)

	// read from channel
	for res := range ch {
		fmt.Println(res)
	}

	fmt.Printf("Time taken since start in seconds :: %v \n", int(time.Since(start).Seconds()))

	fmt.Println("All task Done.")
}
