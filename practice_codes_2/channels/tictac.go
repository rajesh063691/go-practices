package main

import "fmt"

func tic(tic chan<- string) {
	//time.Sleep(1 * time.Second)
	tic <- "tic tac"
	close(tic)
}

func tac(str chan<- string, str2 <-chan string) {
	value := <-str2
	str <- value
	close(str)
}

func main_t() {

	sendTo := make(chan string, 1)
	receiveFrom := make(chan string, 1)
	//i := 0
	//for i <= 10 {
	go tic(sendTo)
	go tac(sendTo, receiveFrom)
	//}

	//fmt.Println(<-sendTo)

	for str := range sendTo {
		fmt.Println(str)
	}

}
