package main

import (
	"codes/pairs"
	"fmt"
)

func main1() {
	var eleInArray int
	fmt.Scanf("%d", &eleInArray)

	var actArray = make([]int, eleInArray)
	for i := 0; i < eleInArray; i++ {
		fmt.Scanf("%d", &actArray[i])
	}

	noOfPairs := pairs.FindPossiblePairs(actArray)
	fmt.Println(noOfPairs)
}
