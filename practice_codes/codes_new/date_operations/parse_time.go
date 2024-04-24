package main

import (
	"fmt"
	"log"
	"time"
)

func main_parse() {
	vals := []string{"2021-07-28", "2020-11-12", "2019-01-05"}
	for _, val := range vals {
		t, err := time.Parse("2006-01-02", val)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(t)
	}
}
