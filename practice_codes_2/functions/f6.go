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
	//fmt.Printf("%s", "Hello")
	r := f(x)
	return r
}

func main_3() {
	r1 := apply(3, inc)
	r2 := apply(2, dec)
	fmt.Println(r1)
	fmt.Println(r2)
}
