package main

import "fmt"

func defFoo4() {
	fmt.Println("defFoo() started")

	if r := recover(); r != nil {
		fmt.Println("WOHA! Program is panicking with value", r)
	}

	fmt.Println("defFoo() done")
}

func normMain4() {

	fmt.Println("normMain() started")

	defer defFoo() // defer defFoo call

	panic("HELP") // panic here

	fmt.Println("normMain() done")
}

func defMain5() {

	fmt.Println("defMain() started")

	fmt.Println("defMain() done")
}

func main_p5() {
	fmt.Println("main() started")

	defer defMain() // defer call

	normMain() // normal call

	fmt.Println("main() done")
}
