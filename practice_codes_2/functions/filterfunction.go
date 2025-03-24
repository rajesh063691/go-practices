package main

import "fmt"

type User struct {
	name       string
	occupation string
	married    bool
}

func returnIfMarried(u User) bool {
	return u.married
}

func returnIfTeacher(u User) bool {
	return u.occupation == "teacher"
}

func main_fil() {

	// u1 := User{"John Doe", "gardener", false}
	// u2 := User{"Richard Roe", "driver", true}
	// u3 := User{"Bob Martin", "teacher", true}
	// u4 := User{"Lucy Smith", "accountant", false}
	// u5 := User{"James Brown", "teacher", true}

	// users := []User{u1, u2, u3, u4, u5}

	users := []User{
		{"John Doe", "gardener", false},
		{"Richard Roe", "driver", true},
		{"Bob Martin", "teacher", true},
		{"Lucy Smith", "accountant", false},
		{"James Brown", "teacher", true},
	}

	// annonymous way of calling
	// //this is a annonymous function where definition is at the same place
	// married := filter(users, func(u User) bool {
	// 	if u.married == true {
	// 		return true
	// 	}
	// 	return false
	// })
	// //this is a annonymous function
	// teachers := filter(users, func(u User) bool {

	// 	if u.occupation == "teacher" {
	// 		return true
	// 	}
	// 	return false
	// })

	//normal way of calling
	married := filter(users, returnIfMarried)
	teachers := filter(users, returnIfTeacher)

	fmt.Println("Married:")
	fmt.Printf("%+v\n", married)

	fmt.Println("Teachers:")
	fmt.Printf("%+v\n", teachers)

}

func filter(s []User, f func(User) bool) []User {
	var res []User

	for _, v := range s {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}
