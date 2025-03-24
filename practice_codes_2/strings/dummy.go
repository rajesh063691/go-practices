package main

import (
	"fmt"
)

func main() {
	str := "rajesh"

	//strArr := strings.Split(str, "")

	// for i := len(strArr) - 1; i >= 0; i-- {
	// 	fmt.Printf(strArr[i])
	// }

	mods := ""

	for i := 0; i < len(str); i++ {
		mods += fmt.Sprintf("%c", str[i])
		//fmt.Printf("Byte Index: %d, Value: %c\n", i, str[i])
	}

	fmt.Println(mods)
}

// fmt.Println()
// fmt.Println(strings.Join(strArr, ""))
