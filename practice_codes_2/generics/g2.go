package main

import "fmt"

type Gen interface {
	int | int16 | int32 | float32 | float64
}

func getData[G Gen](r G) {
	var c G

	c = 3 * 7 * r
	fmt.Println("The circumference is: ", c)
	return
}

func main() {
	getData(10.5)
}
