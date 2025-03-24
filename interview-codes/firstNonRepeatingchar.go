package main

import (
	"fmt"
)

func firstNonRepeatingChar(s string) rune {
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}

	fmt.Println(charCount)

	for _, char := range s {
		if charCount[char] == 1 {
			return char
		}
	}
	return 0 // No non-repeating character
}

func main_r() {
	str := "sssddsse"
	fmt.Printf("%c\n", firstNonRepeatingChar(str)) // Output: "w"
}
