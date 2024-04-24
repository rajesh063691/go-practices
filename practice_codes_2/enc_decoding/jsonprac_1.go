package main

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	Username  string
	Followers int
	Grades    map[string]string
}

type Student1 struct {
	FirstName, LastName string
	Age                 int
	Profile             Profile
	Languages           []string
}

func main_1() {
	var st Student1

	st = Student1{
		FirstName: "Rajesh",
		LastName:  "Kumar",
		Age:       30,
		Profile: Profile{
			Username:  "raj",
			Followers: 121,
			Grades:    map[string]string{"Math": "A+", "Science": "A"},
		},
		Languages: []string{"English", "Hindi"},
	}

	stJSON, _ := json.Marshal(st)
	fmt.Println(string(stJSON))
}
