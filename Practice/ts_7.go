package main

import (
	"fmt"
)

func main_7() {
	mapVal := []int{1, 2}
	appVal := []int{1, 2, 3, 5}

	fmt.Println(IfExists(mapVal, appVal))

}

func IfExists(mapVal, appVal []int) (isAppPresentInMap bool) {

	for i, aVal := range appVal {
		isAppPresentInMap = false
		for j, mVal := range mapVal {
			if aVal == mVal {
				isAppPresentInMap = true
				break
			}
			fmt.Println("j=", j)
		}
		if !isAppPresentInMap {
			return
		}
		fmt.Println("i=", i)
	}

	return
}
