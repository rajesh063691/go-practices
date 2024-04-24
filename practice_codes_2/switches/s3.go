package main

import "fmt"

func main_s3() {
	switch number := 20; {
	case number <= 5:
		fmt.Println("number is less than or equal to 5")
		fallthrough
	case number > 5:
		fmt.Println("number is greater than 5")
		fallthrough
	case number > 10:
		fmt.Println("number is greater than 10")
		fallthrough
	case number > 15:
		fmt.Println("number is greater than 15")

	}
}
