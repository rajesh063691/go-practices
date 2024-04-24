package main

import "fmt"

func greet(message string) {
	fmt.Println("greeting: ", message)
}

func main_f1() {
	fmt.Println("Call one")

	defer greet("Greet one")

	fmt.Println("Call two")

	defer greet("Greet two")

	fmt.Println("Call three")

	defer greet("Greet three")
}
