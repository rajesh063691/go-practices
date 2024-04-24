package main

import "fmt"

func main() {
	var arr [3]string
	arr[0] = "Rajesh"
	arr[1] = "Kumar"
	arr[2] = "Yadav"

	fmt.Println(arr)

	arr1 := [3]string{"Rajesh", "Kumar", "Yadav"}
	fmt.Println(arr1)

	myarr := [...]string{"Rajesh", "Kumar", "Yadav", "testing"}
	fmt.Println(myarr)
	fmt.Println(len(myarr))

}
