package main

import (
	"fmt"
	"time"
)

type TestTime struct {
	time1 time.Time `json:"time1"`
	time2 time.Time `json:"time2"`
}

func main_3() {
	s := TestTime{}
	duration := int64(s.time1.Sub(s.time2))
	fmt.Println("Time1=", s.time1)
	fmt.Println("Duration=", duration)

}
