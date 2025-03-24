package main

import "fmt"

func main_initVal() {

	//this default value is nil
	//var vals []int

	// this default value is [], means empty slice
	vals := []int{}
	fmt.Println(vals)
	if len(vals) == 0 {
		fmt.Printf("slice is empty\n")
	}

	fmt.Printf("slice: %v; len: %d; cap: %d \n", vals, len(vals), cap(vals))

	fmt.Println("---------------------------")

	var vals2 = make([]int, 5)
	fmt.Printf("slice: %v; len: %d; cap: %d \n", vals2, len(vals2), cap(vals2))
}
