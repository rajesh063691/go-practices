package main

import (
	"fmt"
)

func main_filter1() {

	vals := []int{-2, 0, 1, 9, 7, -3, -5, 6}

	n := 0
	for _, val := range vals {
		if isPositive(val) {
			vals[n] = val
			n++
		}
	}
	fmt.Println(vals)
	vals = vals[:n]
	fmt.Println(vals)
}

func isPositive(val int) bool {
	if val > 0 {
		return true
	} else {
		return false
	}
}
