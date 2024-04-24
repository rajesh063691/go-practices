package main

import (
	"errors"
	"fmt"
	"os"
)

func main_2() {
	file, err := os.Open("non-existing.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file.Name() + "opened succesfully")
	}

	sampleErr := errors.New("error occured")
	fmt.Println(sampleErr)

	sampleErr = fmt.Errorf("Err is: %s", "database connection issue")
	fmt.Println(sampleErr)
}
