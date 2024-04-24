package main

import (
	"fmt"
	"time"
)

func main_date() {
	t1 := time.Now().UTC()
	t2 := time.Now().AddDate(0, 0, 1).UTC()

	diff := t2.Sub(t1)
	fmt.Println(diff)
	fmt.Println(diff.Nanoseconds())

}
