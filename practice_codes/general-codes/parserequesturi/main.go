package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Printf("URL Parsing practice!\n")
	u, err := url.ParseRequestURI("test:test")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", u)
}
func IsValidUrl(str string) bool {
	u, err := url.Parse(str)
	fmt.Printf("URL struct: %+v\n", u)
	fmt.Printf("Schema: %s\n", u.Scheme)
	fmt.Printf("Host: %s\n", u.Host)
	return err == nil && u.Scheme != "" && u.Host != ""
}
