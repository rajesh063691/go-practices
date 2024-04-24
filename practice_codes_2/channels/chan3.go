package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch1 := make(chan string, 1)

	for i := 0; i < 10; i++ {
		if isChannelFull(ch1) {
			ch1 = checkCapacity(ch1)
		} else {
			ch1 <- "Rajesh_" + strconv.Itoa(i)
		}
		fmt.Println(<-ch1)
	}

}

func isChannelFull(ch1 chan string) bool {
	if len(ch1) == cap(ch1) {
		return true
	}
	return false
}

func checkCapacity(ch chan string) chan string {
	if len(ch) == cap(ch) {
		return make(chan string, len(ch)+1)
	}
	return ch
}
