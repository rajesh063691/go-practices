package main

import "fmt"

type Employee1 struct {
	name   string
	salary int
}

func (e *Employee1) changeName(newName string) {
	e.name = newName
}

func main_m3() {
	e := &Employee1{
		name:   "Ross Geller",
		salary: 1200,
	}

	// e before name change
	fmt.Println("e before name change =", e)

	// change name
	e.changeName("Monica Geller")

	// e after name change
	fmt.Println("e after name change =", e)
}
