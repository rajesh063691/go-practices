package main

import "fmt"

// func uniqueVals(val []int) (uval []int) {
// 	maps := make(map[int]bool)
// 	uval = []int{}

// 	for _, ele := range val {
// 		if _, ok := maps[ele]; !ok {
// 			maps[ele] = true
// 			uval = append(uval, ele)
// 		}
// 	}
// 	return uval
// }

func main() {

	vals := []int{1, 2, 3, 3, 4, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 11, 11}
	uvals := removeDuplicates(vals)
	fmt.Printf("Original Slice: %v \n", vals)
	fmt.Printf("Unique Slice: %v \n", uvals)

	var myslice []int
	myslice = append(myslice, vals[5:len(vals)-1]...)

	fmt.Printf("New Slice: %v \n", myslice)

}

func removeDuplicates(val []int) (uval []int) {
	maps := make(map[int]bool)

	// for _, v := range val {
	// 	if _, ok := maps[v]; !ok {
	// 		maps[v] = v
	// 		uval = append(uval, v)
	// 	}
	// }

	for _, v := range val {
		if _, ok := maps[v]; !ok {
			maps[v] = true
			uval = append(uval, v)
		}
	}

	return uval
}
