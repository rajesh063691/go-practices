package main

import "fmt"

func main_milli() {
	milli := int64(85954000)
	yr, mn, d, h, m, s := MilliConversion(milli)
	fmt.Printf("year=%d,month=%d,day=%d, Hour=%d,Minutes=%d,Seconds=%d\n", yr, mn, d, h, m, s)
}

func MilliConversion(milliseconds int64) (year, month, day, hr, min, sec int) {
	hr = 0
	min = 0
	sec = 0
	day = 0
	month = 0
	year = 0

	for milliseconds >= 1000 {
		milliseconds = (milliseconds - 1000)
		sec = sec + 1
		if sec >= 60 {
			min = min + 1
			sec = sec - 60
		}
		if min >= 60 {
			hr = hr + 1
			min = min - 60
		}
		if hr >= 24 {
			hr = (hr - 24)
			day = day + 1
		}
		if day >= 30 {
			day = day - 30
			month = month + 1
		}
		if month >= 12 {
			month = month - 12
			year = year + 1
		}
	}
	return
}
