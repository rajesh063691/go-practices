package main

import "fmt"

func main_slice1() {

	words := []string{"rajesh", "kumar", "yadav", "puja", "kumari"}
	fmt.Println(words)

	//remove yadav from slice

	s1 := words[:2]
	s2 := words[3:]
	new_words := []string{}
	new_words = append(new_words, s1...)
	new_words = append(new_words, s2...)
	fmt.Println(new_words)

}
