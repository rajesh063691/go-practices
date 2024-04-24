package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Printf("URL Parsing practice!\n")
	urlSlice := []string{"http://www.google.com", "www.google.com", "google.com", "https://google.com?type=new"}

	for _, url := range urlSlice {
		valid := IsValidUrl(url)
		if valid {
			fmt.Printf("Url: %s is Valid\n", url)
		} else {

			fmt.Printf("Url: %s is InValid\n", url)
		}
	}
}

func IsValidUrl(str string) bool {
	u, err := url.Parse(str)
	fmt.Printf("URL struct: %+v\n", u)
	fmt.Printf("Schema: %s\n", u.Scheme)
	fmt.Printf("Host: %s\n", u.Host)
	return err == nil && u.Scheme != "" && u.Host != ""
}
