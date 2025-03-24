package main

import (
	"design-pattern/singleton"
	"fmt"
)

func main() {
	sin := singleton.GetSingletonInstance()
	fmt.Printf("%v \n", sin)
	sin2 := singleton.GetSingletonInstance()
	fmt.Printf("%v \n", sin2)
}
