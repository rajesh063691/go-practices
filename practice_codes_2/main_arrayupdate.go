package main

import (
	"codes/arrayupdate"
	"fmt"
)

func main3() {
	var size, calTime, testCases int
	var arr = make([]int, 0)
	fmt.Scanf("%d", &testCases)
	var outArr = make([]int, testCases)
	fmt.Println()
	//taking inputs
	for i := 0; i < testCases; i++ {
		fmt.Println("Enter size of array and time separated by space")
		fmt.Scanf("%d", &size)
		fmt.Scanf("%d", &calTime)
		arr = make([]int, 0)
		for j := 0; j < size; j++ {
			//increase size dynamically
			arr = append(arr, j)
			fmt.Scanf("%d", &arr[j])
		}

		// storing output in array
		outArr[i] = arrayupdate.CalculateTime(arr, calTime)

	}
	for i := 0; i < len(outArr); i++ {
		fmt.Println(outArr[i])
	}
}
