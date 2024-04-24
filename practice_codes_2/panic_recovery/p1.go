package main

import "fmt"

// access an element of a slice by index
func accessElement1(a []int, index int) int {
	fmt.Println("Before Panic:")

	if index < len(a) {
		return a[index]
	} else {
		panic("Index is out of bound")
	}
	fmt.Println("After Panic:")
	return 0
}

func main_p1() {
	a := []int{1, 2, 3}

	// access 3rd element
	fmt.Println(accessElement(a, 2))

	// access 4th element (out of bound)
	fmt.Println(accessElement(a, 3))

}
