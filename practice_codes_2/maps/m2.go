package main

import (
	"fmt"
	"maps"
)

func main() {
	numOfCountries := map[string]int{
		"africa": 54,
		"europe": 44,
	}

	numOfCountries1 := map[string]int{
		"africa": 54,
		"europe": 44,
	}
	if country, ok := numOfCountries["africa1"]; ok {
		println("Number of countries in Africa is", country)
	}
	// clear(numOfCountries)
	// delete(numOfCountries, "europe")

	fmt.Println(numOfCountries)
	fmt.Println(maps.Equal(numOfCountries, numOfCountries1))
}
