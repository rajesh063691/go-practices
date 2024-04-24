package main

import "fmt"

func foo1() {
	fmt.Println("foo() executed")

	panic("panic from foo()")

	fmt.Println("foo() done")
}

func main_p2() {
	fmt.Println("main() started")

	foo() // call foo function

	fmt.Println("main() done")
}
