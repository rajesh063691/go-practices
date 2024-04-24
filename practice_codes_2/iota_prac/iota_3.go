package main

import "fmt"

type SizeType int

const (
	small = SizeType(iota)
	medium
	large
	extraLarge
)

func main() {
	var m SizeType = 0
	m.toString()
}
func (s SizeType) toString() {
	switch s {
	case small:
		fmt.Println("Small")
	case medium:
		fmt.Println("Medium")
	case large:
		fmt.Println("Large")
	case extraLarge:
		fmt.Println("Extra Large")
	default:
		fmt.Println("Invalid Size entry")
	}
}
