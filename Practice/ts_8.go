package main

import (
	"fmt"
	"strconv"
	"time"
)

func main_8() {

	today := time.Now()
	//cutDate := today.Add(-24 * time.Hour)
	cutDate := today.AddDate(0, +1, 0)
	defaultDateTime := time.Time{}

	// Using time.Before() method
	g1 := today.After(defaultDateTime)
	fmt.Println("today after defaultDateTime:", g1)

	// Using time.After() method
	g2 := today.Before(cutDate)
	fmt.Println("today after cutDate:", g2)

	fmt.Printf("defaultDateTime = %s\n", defaultDateTime)
	fmt.Printf("Today = %s\n", today)
	fmt.Printf("cutDate = %s\n", cutDate)

	if today.After(defaultDateTime) && today.Before(cutDate) {
		fmt.Println("Should Ignore...")
	}

	str1 := "raj"
	str2 := "rajesh"

	fmt.Printf("str1=%s\n", str1[len(str2)-5:])
	fmt.Printf("str2=%s\n", str1[len(str1)-5:])

	val := 123.456789
	fl, err := strconv.ParseFloat(fmt.Sprintf("%.2f", val), 2)
	if err != nil {
		fmt.Println(fl)
	}

}
