package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Extra any
}

func main() {
	p := &Person{
		Name: "John",
		Age:  30,
	}

	p.Extra = "Person Eats a lot"

	fmt.Printf("%+v\n", p)
	fmt.Printf("%T\n", p.Extra)

	p.Extra = 42
	fmt.Printf("%T\n", p.Extra)
}
