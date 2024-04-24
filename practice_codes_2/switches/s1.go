package main

import "fmt"

func main_s1() {
	finger := 7

	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5, 6, 7:
		{
			fmt.Println("Little")
		}

	default:
		{
			fmt.Println("Finger not found")
		}
	}
}
