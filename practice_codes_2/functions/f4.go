package main

import "fmt"

//this is assigning anonymous function to global variable
//var add = func(a int, b int) int { return a + b }
var add = Addition

// now, we will create a function with name Addition and assign it to global variable add
func Addition(a int, b int) int {
	return a + b
}
func main_f4() {
	fmt.Println("5+3 =", add(5, 3))
}
