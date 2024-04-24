package main

import "fmt"

func defFoo1() {
	fmt.Println("defFoo() started")

	if r := recover(); r != nil {
		fmt.Println("Whoa! Program is panicking with value:", r)
		//return
	}

	fmt.Println("defFoo() done")
}

func normMain1() {

	fmt.Println("normMain1() started")

	defer defFoo1() // defer defFoo call

	panic("NEED HELP") // panic here

	fmt.Println("normMain() done")
}

func main_4() {
	fmt.Println("main() started")

	normMain1() // normal call

	fmt.Println("main() done")
}
