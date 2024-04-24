package main

import "fmt"

// access an element of a slice by index
func accessElement(a []int, index int) (v int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Program recoverd successfully!")
			v = a[len(a)-1] // set v to last value
		}
	}()

	v = a[index]
	return
}

func main_9() {
	a := []int{1, 2, 3}

	// access 3rd element
	fmt.Println(accessElement(a, 2))

	// access 4th element (out of bound)
	fmt.Println(accessElement(a, 3))
}
