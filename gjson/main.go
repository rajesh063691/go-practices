package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	// json := `{
	//     "name": "John",
	//     "age": 30,
	//     "address": {
	//         "city": "New York",
	//         "zipcode": "10001"
	//     },
	//     "hobbies": ["reading", "gaming"]
	// }`
	json := `{
		"company": {
			"name": "TechCorp",
			"employees": [
				{"name": "Alice", "age": 30},
				{"name": "Bob", "age": 25}
			]
		}
	}`

	companyName := gjson.Get(json, "company.name")
	employeeName := gjson.Get(json, "company.employees.1.name")

	fmt.Println("Company Name:", companyName.String())          // Output: TechCorp
	fmt.Println("Second Employee Name:", employeeName.String()) // Output: Bob

	// Simple key access
	name := gjson.Get(json, "name")
	fmt.Println("Name:", name.String())

	// Nested key access
	city := gjson.Get(json, "address.city")
	fmt.Println("City:", city.String())

	// Array access
	firstHobby := gjson.Get(json, "hobbies.0")
	fmt.Println("First Hobby:", firstHobby.String())
}
