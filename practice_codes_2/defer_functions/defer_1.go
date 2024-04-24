package main

import "fmt"

func main() {

	fmt.Println("start")
	for i := 1; i <= 5; i++ {
		if i == 1 || i == 2 || i == 3 {
			defer fmt.Println(i)
		}

	}
	fmt.Println("end")
}
