package main

import (
	"fmt"
)

func main() {
	a := []int{5, 6}
	val, err := checkAndGet(a, 2)
	fmt.Printf("Val: %d\n", val)
	fmt.Println("Error: ", err)
}
func checkAndGet(a []int, index int) (value int, err error) {
	value = 10
	defer handleOutOfBounds()
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	value = a[index]
	return value, nil
}
func handleOutOfBounds() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}
