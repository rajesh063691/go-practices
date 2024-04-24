package main

import (
	"encoding/json"
	"fmt"
)

func main_1() {
	data := `{
		"Name":"John",
		"Age":30,
		"City":"New York",
		"Country":"USA"
	}`

	var obj any
	err := json.Unmarshal([]byte(data), &obj)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	m := obj.(map[string]any)
	fmt.Printf("Name: %s\n", m["Name"])
	fmt.Printf("Age: %0.2f\n", m["Age"])
	fmt.Printf("City: %s\n", m["City"])
	fmt.Printf("Country: %s\n", m["Country"])

}
