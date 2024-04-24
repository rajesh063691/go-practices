package main

import (
	"fmt"
	"time"
)

var layout string = "2006-01-02T15:04:05.00"
var layout1 string = "2006-Jan-02 Monday 15:04:05"

func main() {
	// t1 := time.Now().Format(time.RFC3339)
	// t2 := time.Now().Format(layout)
	// fmt.Println(t1)
	// fmt.Println(t2)

	// t1 := time.Now()
	// fmt.Println(t1)

	// t2 := time.Now().Add(2 * time.Hour)
	// sub := t2.Sub(t1)
	// sin := time.Since(t2)
	// fmt.Println(sub)
	// fmt.Println(sin)
	// str := t1.String()
	// time.Parse(layout, str)
	// str = "2020-Jan-29 Wednesday 12:19:25"
	// t, _ := time.Parse(layout1, str)
	// fmt.Println(t)

	t1 := time.Now()
	fmt.Println(t1)
	unixTime := t1.Unix()
	fmt.Println(unixTime)
	t2 := time.Unix(unixTime, 0)
	fmt.Println(t2)

}
