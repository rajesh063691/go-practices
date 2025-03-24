package main

import (
	"encoding/json"
	"fmt"
)

// Students ...
// type Students struct {
// 	FirstName, lastName string
// 	Email               string
// 	Age                 int
// 	HeightInMeters      float64
// 	IsMale              bool
// }

type Students map[string]interface{}

func main() {
	// student := Students{
	// 	FirstName: "Rajesh",
	// 	lastName:  "Kumar",
	// 	//Email:          "abc@gmail.com",
	// 	Age:            30,
	// 	HeightInMeters: 167,
	// 	IsMale:         true,
	// }

	student := Students{
		"FirstName": "Rajesh",
		"lastName":  "Kumar",
		//"Email":          "abc@gmail.com",
		"Age":            30,
		"HeightInMeters": 167,
		"IsMale":         true,
	}
	// convert student to json
	stJSON, _ := json.Marshal(student)

	// print json string
	fmt.Println(string(stJSON))

}
