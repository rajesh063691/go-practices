package main

import "fmt"

func defBar() {
	fmt.Println("defBar() started")

	if r := recover(); r != nil {
		fmt.Println("WOHA! Program is panicking with value", r)
	}

	fmt.Println("defBar() done")
}

func defFoo6() {
	fmt.Println("defFoo() started")

	defer defBar() // defer call

	fmt.Println("defFoo() done")
}

func normMain6() {

	fmt.Println("normMain() started")

	defer defFoo() // defer defFoo call

	panic("HELP") // panic here

	fmt.Println("normMain() done")
}

func defMain6() {

	fmt.Println("defMain() started")

	fmt.Println("defMain() done")
}

func main_p6() {
	fmt.Println("main() started")

	defer defMain() // defer call

	normMain() // normal call

	fmt.Println("main() done")
}
