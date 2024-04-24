package main

import (
	"codes_new/model"
	"fmt"
)

func main() {

	u := model.User{Name: "John Doe", Occupation: "gardener"}

	fmt.Printf("%s is a %s\n", u.Name, u.Occupation)

	//here address is not accessible bcoz its first letter Starts with small letter
	//u1 := model.Address{}

	//fmt.Printf("%s is a %s\n", u1.City, u1.City)
}
