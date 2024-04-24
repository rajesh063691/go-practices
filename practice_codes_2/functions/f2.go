package main

import "fmt"

func add1(a int, b int) int {
	return a + b
}

func subtract1(a int, b int) int {
	return a - b
}

func main_f2() {
	fmt.Printf("Type of function add is			%T\n", add)
	fmt.Printf("Type of function subtract is		%T\n", subtract)
}
