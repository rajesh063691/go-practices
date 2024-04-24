package main

import (
	"fmt"
	"time"
)

func main_Curr() {
	//shows time in local timezone
	now := time.Now()

	fmt.Println("The current datetime is:", now)
}
