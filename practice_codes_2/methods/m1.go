package main

import "fmt"

type Employee struct {
	FirstName, MiddleName, LastName string
}

// with receiver it becomes method and earlier it was a function
func (e *Employee) fullName() (fullName string) {
	e.MiddleName = "Kumar"
	fullName = e.FirstName + " " + e.MiddleName + " " + e.LastName
	return
}

func main() {
	var e = new(Employee)
	// e := &Employee{
	// 	FirstName: "Rajesh",
	// 	LastName:  "Yadav",
	// }
	fmt.Printf("%+v", e)
	fmt.Println(e)
	fmt.Println(e.fullName())
	fmt.Printf("%+v", e)

}
