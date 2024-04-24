package main

import (
	"bufio"
	"fmt"
	"os"
)

func main_scanner() {
	names := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter Names...")
		scanner.Scan()
		text := scanner.Text()

		if len(text) != 0 {
			names = append(names, text)
		} else {
			break
		}
	}
	fmt.Println(names)
}
