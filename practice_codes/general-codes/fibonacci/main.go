package main

import "fmt"

func main() {
	fmt.Println("Fibonacci numbers")
	Fibo(20)
}
func Fibo(n int) {
	a := 0
	b := 1
	for n > 0 {
		fmt.Printf("%d ", a)
		// need to swap the a and sum of a+b
		temp := b
		b = a + b
		a = temp
		n = n - 1
	}
	fmt.Println()
}
