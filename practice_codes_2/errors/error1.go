package main

import (
	"errors"
	"fmt"
	"log"
)

func divide(x, y int) (val int, err error) {
	if y == 0 {
		return 0, errors.New("can not divide by 0")
	} else {
		return (x / y), nil
	}
}
func main_error1() {
	//success case
	if res, err := divide(3, 6); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("The answer is", res)
	}

	//failure case
	if res, err := divide(3, 0); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("The answer is", res)
	}

}
