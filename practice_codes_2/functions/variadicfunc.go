package main

import (
	"fmt"
	"strings"
)

func testVariadic(elements ...string) string {
	return strings.Join(elements, "-")
}

func main_vari() {
	var elements = []string{"Rajesh", "Kumar", "Yadav"}
	fmt.Println(testVariadic(elements...))

}
