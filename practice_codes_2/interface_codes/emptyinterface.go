package main

import "fmt"

type MyString string

type Rect1 struct {
	width  float64
	height float64
}

func explain(i interface{}) {
	fmt.Printf("value given to explain function is of type '%T' with value %v\n", i, i)
}

func main_4() {
	ms := MyString("Hello World!")
	r := Rect1{5.5, 4.5}
	explain(ms)
	explain(r)

}
