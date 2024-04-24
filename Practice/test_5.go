package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main_5() {

	host := "https://abc.github.com///"
	host = strings.TrimSuffix(host, "/")
	fmt.Println("host=", host)

	defaultTime := time.Time{}
	fmt.Printf("%v\n", defaultTime)

	strTimeStamp := "2022-09-22T10:55:46.000+00:00"
	timeFormatTimeStamp, err := time.Parse("2006-01-02T15:04:05.000+00:00", strTimeStamp)
	log.Printf("Time Parsing error=%v\n", err)
	fmt.Printf("Time=%v\n", timeFormatTimeStamp)

	val := strings.Split("2022-09-22T07:49:22.000+00:00", "+")[0]
	fmt.Printf("Val=%v\n", val)

	curreTime := time.Now()
	fmt.Printf("curtime=%v\n", curreTime.String())

	current_time := time.Now()

	// individual elements of time can
	// also be called to print accordingly
	newTime := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		current_time.Year(), current_time.Month(), current_time.Day(),
		current_time.Hour(), current_time.Minute(), current_time.Second())

	fmt.Printf("newTime=%s\n", newTime)
	fmt.Printf("currentTime=%s\n", current_time)

}
