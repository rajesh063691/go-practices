package main

import "fmt"

func main_1() {

	countries := map[string]string{
		"sk": "Slovakia",
		"ru": "Russia",
		"de": "Germany",
		"no": "Norway",
	}

	sort.countries()

	fmt.Println(countries)

	delete(countries, "ru")

	fmt.Println(countries)
}
