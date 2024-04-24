package main

import "fmt"

type Contact5 struct {
	phone, address string
}
type Employee5 struct {
	name    string
	salary  int
	contact Contact5
}

func (e *Employee5) changePhone(newPhone string) {
	e.contact.phone = newPhone
}

func main_m5() {
	e := Employee5{
		name:    "Ross Geller",
		salary:  1200,
		contact: Contact5{"011 8080 8080", "New Delhi, India"},
	}
	// e before phone change
	fmt.Println("e before phone change =", e)
	// change phone
	e.changePhone("011 1010 1222")
	// e after phone change
	fmt.Println("e after phone change =", e)
}
