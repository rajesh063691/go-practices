package main

import (
	"fmt"
	"time"
)

func main() {

	t1 := time.Date(2021, time.November, 10, 23, 0, 0, 0, time.UTC)
	t2 := time.Date(2021, time.November, 10, 24, 1, 1, 0, time.UTC)

	elapsed := t2.Sub(t1)

	fmt.Println(elapsed)
}
