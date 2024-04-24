package main

import "fmt"

// const (
// 	a = iota + 10
// 	b
// 	//tetststs
// 	c
// 	_ // skip value
// 	d
// )

const (
	a = iota + 10
	b = iota + 2
	c = iota + 3
	d = iota + 5
)

func main_1() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
