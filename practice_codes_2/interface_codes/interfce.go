package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main_() {
	// no need to create derived type for interface
	var s Shape
	s = Rect{5, 4}
	fmt.Println("Area=", s.Area())
	//fmt.Println("Perimeter=", s.Perimeter())
}
