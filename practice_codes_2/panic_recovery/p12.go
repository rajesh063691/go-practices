package main

import (
	"fmt"
	"time"
)

func main_12() {
	a := []string{"a", "b"}
	checkAndPrintWithRecover(a, 2)
	time.Sleep(time.Second)
	fmt.Println("Exiting normally")
}
func checkAndPrintWithRecover(a []string, index int) {
	go checkAndPrint(a, 2)
}
func checkAndPrint(a []string, index int) {
	defer handleOutOfBounds()
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	fmt.Println(a[index])
}
func handleOutOfBounds2() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}
