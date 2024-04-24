package main

import "fmt"

type Contact struct {
	phone, address string
}
type Employee2 struct {
	name    string
	salary  int
	contact Contact
}

func (c *Contact) changePhone(newPhone string) {
	c.phone = newPhone
}

func main_m4() {
	e := Employee2{
		name:   "Ross Geller",
		salary: 1200,
		contact: Contact{
			phone:   "011 8080 8080",
			address: "New Delhi, India",
		},
	}
	// e before phone change
	fmt.Println("e before phone change =", e)
	// change phone
	e.contact.changePhone("011 1010 1222")
	// e after phone change
	fmt.Println("e after phone change =", e)
}
