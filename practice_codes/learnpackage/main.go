package main

import (
	"fmt"
	"learnpackage/simpleinterest" //importing custom packages
	"log"
)

var p, r, t = -5000.0, 10.0, 1.0

/*
* init function to check if p, r and t are greater than zero
 */
func init() {
	println("Main package initialized")
	if p < 0 {
		log.Fatal("Principal is less than zero")
	}
	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}
	if t < 0 {
		log.Fatal("Duration is less than zero")
	}

	fmt.Println("p==>r===>t", p, r, t)
}

func main() {
	fmt.Println("Simple interest calculation")
	// p := 5000.0
	// r := 10.0
	// t := 2.0
	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}
