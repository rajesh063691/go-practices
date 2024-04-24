package main

import "fmt"

func add2(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

//1st way of defining
// func calc(a, b int, f func(a1, b1 int) int) int {
// 	v := f(a, b)
// 	return v
// }

// 2nd way, we can define derived type to make it easy to understand
type Calculate func(a, b int) int

func calc(a, b int, c Calculate) int {
	v := c(a, b)
	return v
}

func main_f3() {
	addResult := calc(5, 3, add)
	subResult := calc(5, 3, subtract)
	fmt.Println("5+3 =", addResult)
	fmt.Println("5-3 =", subResult)
}
