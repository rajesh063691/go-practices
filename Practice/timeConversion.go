package main

import "fmt"

func main_t() {

	duration := int64(1)
	unit := "nanotetet"

	n := timeConversion(unit, duration)
	fmt.Printf("Time in Nano:=%d \n", n)
}

func timeConversion(unit string, duration int64) (nanoDuration int64) {
	switch unit {
	case "nano":
		nanoDuration = duration
	case "micro":
		nanoDuration = duration * 1000
	case "milli":
		nanoDuration = duration * 1000 * 1000
	case "second":
		nanoDuration = duration * 1000 * 1000 * 1000
	case "minute":
		nanoDuration = duration * 60 * 1000 * 1000 * 1000
	case "hour":
		nanoDuration = duration * 60 * 60 * 1000 * 1000 * 1000
	}
	return
}
