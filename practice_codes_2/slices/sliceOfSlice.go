package main

import "fmt"

func main_sl() {

	words := [][]string{
		{"sky", "ocean"},
		{"red", "blue"},
		{"C#", "Go"},
	}
	fmt.Println("<======Individual String Vals====>")
	for _, sl := range words {
		for _, str := range sl {
			fmt.Printf(str + " ")
		}
		fmt.Println()
	}

	fmt.Printf("slice: %v; len: %d; cap: %d \n", words, len(words), cap(words))
}
