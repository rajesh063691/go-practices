package main

import (
	"fmt"
)

func main_s2() {
	letter := "a"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Letter is a Vowel !")
	default:
		fmt.Println("Letter is not a Vowel")
	}
}
