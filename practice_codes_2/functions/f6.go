package main

import "fmt"

func inc(x int) int {
	x++
	return x
}

func dec(x int) int {
	x--
	return x
}

func apply(x int, f func(int) int) int {
	fmt.Printf("x=%d, funcName=%s, funcType=%T  \n", x, f, f)
	r := f(x)
	return r
}

func main_f6() {
	r1 := apply(3, inc)
	r2 := apply(2, dec)
	fmt.Println(r1)
	fmt.Println(r2)
}
