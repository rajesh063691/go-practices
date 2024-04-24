package main

import (
	"fmt"
	"time"
)

func main_fun() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launch date: %s\n", t.Local())
}
