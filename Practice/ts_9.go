package main

import "fmt"

func main() {
	i := 0
	shouldContinue := true
	for shouldContinue {
		shouldContinue = test()

	}
}

func test() bool {
	fmt.Println("Hello")
	return true
}
