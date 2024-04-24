package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main_2() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	urls := []string{
		"https://jsonplaceholder.typicode.com/users",
		"https://api.example.com/products",
		"https://api.example.com/orders",
	}

	ch1 := make(chan string)

	for _, url := range urls {
		go fetchAPI(ctx, url, ch1)
	}

	for val := range ch1 {
		fmt.Println(val)
	}
}

func fetchAPI(ctx context.Context, url string, ch1 chan string) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		ch1 <- fmt.Sprintf("error fetching data from %s:: and err=%s", url, err.Error())
		return
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		ch1 <- fmt.Sprintf("error fetching data from %s  and err=%s", url, err.Error())
		return
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		ch1 <- fmt.Sprintf("error fetching data from %s  and err=%s", url, err.Error())
		return
	}

	ch1 <- fmt.Sprintf("data of size %d \n and res=%s\n", len(data), string(data))
	close(ch1)
}
