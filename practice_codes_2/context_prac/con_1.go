package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main_1() {

	// create http request
	//req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/users", nil)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	context.TODO()
	context.Background()
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://jsonplaceholder.typicode.com/users", nil)
	if err != nil {
		fmt.Printf("error fetching users")
	}

	//perform the http request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error processing request")
	}

	defer res.Body.Close()

	//get body from the response
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error reading response body")
	}

	fmt.Printf("data of size %d \n and res=%s\n", len(data), string(data))
}
