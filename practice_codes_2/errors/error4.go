package main

import (
	"errors"
	"fmt"
)

type errorOne struct{}

func (e errorOne) Error() string {
	return "Error One happended"
}
func main() {
	var err1 errorOne
	err2 := do()
	if err1 == err2 {
		fmt.Println("Equality Operator: Both errors are equal")
	}
	if errors.Is(err1, err2) {
		fmt.Println("Is function: Both errors are equal")
	}
}
func do() error {
	return errorOne{}
}
