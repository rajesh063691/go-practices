package main

import (
	"fmt"
	"strings"
)

func main_1() {
	var str string
	fmt.Println(str)
	str = "Welcome to GeeksforGeeks"

	s := strings.Split(str, "")

	fmt.Println(s)

	j := strings.Join(s, "-")
	fmt.Println(j)

	// if str[0] == str[len(str)-1] {
	// 	fmt.Println("First and last characters are same")
	// } else {
	// 	fmt.Println("First and last characters are not same")
	// }

	// //str[0] = "h"
	// fmt.Println(str)

	for _, val := range str {
		//fmt.Println(string(val))
		fmt.Printf("%s", string(val))
	}
	fmt.Println()
}
