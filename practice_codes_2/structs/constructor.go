package main

import "fmt"

type User struct {
	name       string
	occupation string
	age        int
}

// setting constructors manually
func setUser(name, occupation string, age int) (user *User) {
	user = &User{
		name:       name,
		occupation: occupation,
		age:        age,
	}
	return user
}

func main_co() {
	u := setUser("Rajesh Kumar Yadav", "Software Engineer", 30)
	fmt.Printf("%s is %d years old and he is a %s.\n", u.name, u.age, u.occupation)
}
