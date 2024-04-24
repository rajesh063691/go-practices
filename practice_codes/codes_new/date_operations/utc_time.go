package main

import (
	"fmt"
	"time"
)

func main_utc() {

	utc := time.Now().UTC()
	fmt.Println(utc)
}
