package main

import "fmt"

func main_pipe() {

	naturals := make(chan int)
	squarer := make(chan int)

	// sends natural numbers to naturals chaneel
	go func() {
		for i := 0; i < 10; i++ {
			naturals <- i
		}
		close(naturals)
	}()
	// it take the natural values from naturals and squares it and send to sqaurer
	go func() {
		for x := range naturals {
			squarer <- x * x
		}
		close(squarer)

	}()

	// Printer finally prints the square of naturals
	for val := range squarer {
		fmt.Println(val)
	}

}
