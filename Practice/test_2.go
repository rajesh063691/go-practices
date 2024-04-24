package main

import (
	"fmt"
	"time"
)

func main_2() {
	ok := isDateValue("3092022")
	println(ok)

}

func isDateValue(strDate string) (ok bool) {
	t, err := time.Parse("02012006", strDate)
	year := t.Year()
	month := t.Month().String()
	day := fmt.Sprintf("%02d", t.Day())
	fmt.Printf("Year=[%v],month=[%v]and day=[%v]", year, month, day)
	return err == nil
}
