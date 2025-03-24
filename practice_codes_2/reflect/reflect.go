// Example program to show difference between
// Type and Kind and to demonstrate use of
// Methods provided by Go reflect Package
package main

import (
	"fmt"
	"reflect"
)

type details struct {
	fname   string
	lname   string
	age     int
	balance float64
}
type myType string

func showDetails(i, j interface{}) {
	t1 := reflect.TypeOf(i)
	k1 := t1.Kind()
	t2 := reflect.TypeOf(j)
	k2 := t2.Kind()
	fmt.Println("Type of first interface:", t1)
	fmt.Println("Kind of first interface:", k1)
	fmt.Println("Type of second interface:", t2)
	fmt.Println("Kind of second interface:", k2)
	fmt.Println("Value of first interface:", reflect.ValueOf(i))

	fmt.Println("The values in the first argument are :")
	if reflect.ValueOf(i).Kind() == reflect.Struct {
		value := reflect.ValueOf(i)
		numberOfFields := value.NumField()
		for i := 0; i < numberOfFields; i++ {
			fmt.Printf("%d.Type:%T || Value:%#v\n",
				(i + 1), value.Field(i), value.Field(i))

			fmt.Println("Kind is ", value.Field(i).Kind())
		}
	}
	value := reflect.ValueOf(j)
	fmt.Printf("The Value passed in "+
		"second parameter is %#v", value)
}

func main() {
	iD := myType("12345678")
	person := details{
		fname:   "Go",
		lname:   "Geek",
		age:     32,
		balance: 33000.203,
	}

	showDetails(person, iD)
}
