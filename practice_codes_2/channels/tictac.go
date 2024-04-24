package main

import (
	"fmt"
)

func tic(str chan<- string) {
	str <- "tic"
	close(str)
}

func tac(str <-chan string) string {
	return <-str
	//close(str)
}

func main_tic() {

	ch1 := make(chan<- string)
	ch2 := make(<-chan string)
	//i := 0
	//for i <= 10 {
	go tic(ch1)
	go tac(ch2)
	//}

	for str := range ch2 {
		fmt.Println(str)
	}

}
