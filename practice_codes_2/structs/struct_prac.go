package main

import "fmt"

func findNames(i interface{}) (firstName string, middleName string, lastName string) {
	if val, ok := i.(map[string]string); ok {
		firstName = val["f_name"]
		middleName = val["m_name"]
		lastName = val["l_name"]
	}

	return firstName, middleName, lastName
}

func main() {
	mp := make(map[string]string)
	mp["f_name"] = "Rajesh"
	mp["m_name"] = "Kumar"
	mp["l_name"] = "Yadav"

	fName, mName, lName := findNames(mp)
	fmt.Println(fName, mName, lName)

}
