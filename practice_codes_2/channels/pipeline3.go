package main

import "fmt"

func main_pip3() {
	fmt.Println(mirroredQuery())

}

func mirroredQuery() []string {
	responses := make(chan string, 3)
	go func() {
		responses <- request("asia.gopl.io")
	}()
	go func() {
		responses <- request("europe.gopl.io")
	}()
	go func() {
		responses <- request("americas.gopl.io")
	}()

	res := make([]string, 0, 3)
	for i := 0; i < 3; i++ {
		res = append(res, <-responses)
	}
	close(responses)

	return res // return the quickest response
}

func request(hostname string) (response string) {
	return hostname
}
