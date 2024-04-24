package main

import (
	"fmt"
	"time"
)

var (
	DateLayout           = "2006-01-02"
	DateTimeLayout       = "2006-01-02 15:04:05"
	TDateTimeLayout      = "2006-01-02T15:04:05"
	DateTimeOffsetLayout = "2006-01-02T15:04:05-0700"
)

func main() {

	now := time.Now()

	fmt.Println("Custom Format:", now.Format(DateLayout))

	fmt.Println("Time: ", now.Format("15:04:05"))
	fmt.Println("Date:", now.Format("Jan 2, 2006"))
	fmt.Println("Timestamp:", now.Format(time.Stamp))
	fmt.Println("ANSIC:", now.Format(time.ANSIC))
	fmt.Println("UnixDate:", now.Format(time.UnixDate))
	fmt.Println("Kitchen:", now.Format(time.Kitchen))
}
