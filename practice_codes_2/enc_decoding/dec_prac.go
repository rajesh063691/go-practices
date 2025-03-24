package main

import (
	"encoding/json"
	"fmt"
)

// Student declares `Student` structure
type Student struct {
	FirstName, lastName string
	Email               string `json:"email,omitempty"`
	Age                 int
	HeightInMeters      float64
}

func main_unmarshal() {

	// some JSON data
	data := []byte(`
	{
		"FirstName": "John",
		"lastName": "Doe",
		"Age": 21,
		"HeightInMeters": 175,
		"Username": "johndoe91"
	}`)

	// create a data container
	//var john Student
	john := Student{}

	// unmarshal `data`
	fmt.Printf("Error: %v\n", json.Unmarshal(data, &john))

	// print `john` struct
	fmt.Printf("%+v\n", john)
}
