package main

import (
	"fmt"
	"strings"
)

type MyString string

func (s MyString) toUpperCase() string {
	normalString := string(s)
	return strings.ToUpper(normalString)
}

func main_m6() {
	str := MyString("Hello World")
	fmt.Println(str.toUpperCase())
}
