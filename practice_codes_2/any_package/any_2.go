package main

import (
	"fmt"
)

func printValue(value any) {
	fmt.Println(value)
	fmt.Printf("%T\n", value)
}

func main() {
	printValue(42)
	printValue("hello")
	printValue(true)
}
