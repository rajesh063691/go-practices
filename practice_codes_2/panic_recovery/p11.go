package main

import "fmt"

func main_11() {

	a := []string{"a", "b"}
	checkAndPrint1(a, 2)
	fmt.Println("Exiting normally")
}

func checkAndPrint1(a []string, index int) {
	defer handleOutOfBounds()
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	fmt.Println(a[index])
}

func handleOutOfBounds1() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}
