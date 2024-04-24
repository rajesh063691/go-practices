package main

import "fmt"

// access an element of a slice by index
func accessElement8(a []int, index int) int {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Program recovered but nothing to do here!")
		}
	}()
	return a[index]
}

func main_p8() {
	a := []int{1, 2, 3}

	// access 3rd element
	fmt.Println(accessElement(a, 2))

	// access 4th element (out of bound)
	fmt.Println(accessElement(a, 3))
}
