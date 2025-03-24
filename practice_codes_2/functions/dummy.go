package main

import "fmt"

type Cal func(a, b int) int

func A(a, b int) int {
	return a + b
}

func M(a, b int) int {
	return a * b
}

func cals(a, b int, c Cal) int {
	v := c(a, b)
	return v
}

func main() {
	addResult := cals(5, 3, A)
	mulResult := cals(5, 3, M)
	fmt.Println("5+3 =", addResult)
	fmt.Println("5*3 =", mulResult)
}
