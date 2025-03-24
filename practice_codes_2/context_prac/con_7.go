package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main_7() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	urls := []string{
		"https://google.com",
		"https://linkedin.com",
		"https://kyndryl.com",
	}

	result := make(chan string)

	for _, url := range urls {
		go getApiDetails(ctx, url, result)
	}

	// read from channels
	for range urls {
		fmt.Println(<-result)
	}

}

func getApiDetails(ctx context.Context, url string, result chan<- string) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		result <- fmt.Sprintf("Error creating request for %s: %s", url, err.Error())
		return
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		result <- fmt.Sprintf("Error processing request for %s: %s", url, err.Error())
		return
	}

	defer res.Body.Close()
	result <- fmt.Sprintf("Response from %s: %d", url, res.StatusCode)

}
