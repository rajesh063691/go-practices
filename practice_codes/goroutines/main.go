package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Go Routines Tutorial")

	startTime := time.Now()

	links := []string{
		"https://www.google.com/",
		"https://www.amazon.in/",
		"https://www.flipkart.com/",
		"https://go.dev/dl/",
	}

	CallUrls(links)

	fmt.Printf("Time taken to call the above apis in seconds: %f \n", time.Since(startTime).Seconds())
}

func CallUrls(urls []string) {
	for _, url := range urls {
		CallURL(url)
	}
}

func CallURL(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("We could not connect to url: ", url)
	} else {
		fmt.Println("Successfully connected to url: ", url)
	}
}
