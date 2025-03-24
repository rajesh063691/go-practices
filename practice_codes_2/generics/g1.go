package main

import "fmt"

type Radius interface {
	int | int64 | int8 | float32 | float64
}

// type II
func generic_circumference[R Radius](r R) {
	var c R
	c = 2 * 3 * r
	fmt.Println("The circumference is: ", c)
}

// type I

// func generic_circumference[r int | float32](radius r) {

// 	c := 2 * 3 * radius
// 	fmt.Println("The circumference is: ", c)

// }

func main_1() {
	var radius1 int = 8
	var radius2 float32 = 9.5

	generic_circumference(radius1)
	generic_circumference(radius2)
}
