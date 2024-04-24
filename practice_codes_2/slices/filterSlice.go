package main

import "fmt"

type User struct {
	name       string
	occupation string
	country    string
}

func IsProgrammer(user User) bool {
	return (user.occupation == "programmer") && (user.country == "Canada")
}

func main_filterSlice() {
	users := []User{
		{"John Doe", "gardener", "USA"},
		{"Roger Roe", "driver", "UK"},
		{"Paul Smith", "programmer", "Canada"},
		{"Lucia Mala", "teacher", "Slovakia"},
		{"Patrick Connor", "shopkeeper", "USA"},
		{"Tim Welson", "programmer", "Canada"},
		{"Tomas Smutny", "programmer", "Slovakia"},
	}
	onlyProgrammers := []User{}

	for _, user := range users {
		if IsProgrammer(user) {
			onlyProgrammers = append(onlyProgrammers, user)
		}
	}

	fmt.Println("Struct with only Programmers and from Country Canada")
	for _, programmer := range onlyProgrammers {
		fmt.Println(programmer)
	}

}
