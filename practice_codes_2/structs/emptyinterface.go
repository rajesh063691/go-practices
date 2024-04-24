package main

import (
	"fmt"
)

// type test struct {
// 	Id         int
// 	Name       string
// 	Othersinfo interface{}
// }

func testType(a interface{}) (middlename, lastname string) {
	if val, ok := a.(map[string]string); ok {
		middlename = val["middlename"]
		lastname = val["lastname"]
		//fmt.Printf("%v", val)
		//fmt.Println(lastname)
	}
	// switch a := a.(type) {
	// case map[string]string:
	// 	for key, val := range a {
	// 		if key == "middlename" {
	// 			middlename = val
	// 		}
	// 		if key == "lastname" {
	// 			lastname = val
	// 		}
	// 	}
	// }
	return
}

func main() {

	mp := make(map[string]string)
	mp["middlename"] = "Kumar"
	mp["lastname"] = "Yadav"
	// st := test{
	// 	Id:         1,
	// 	Name:       "Rajesh",
	// 	Othersinfo: mp,
	// }

	// fmt.Printf("STRUCT:%v", st)
	// fmt.Println()
	// otherInfo := st.Othersinfo
	// fmt.Println(otherInfo)
	// var b []byte
	// err := json.Unmarshal(b, &otherInfo)
	// _ = err
	// fmt.Println(string(b))
	// for key, val := range string(b) {
	// 	fmt.Printf("Key:%s, Val=%s", key, val)
	// }

	middname, lastname := testType(mp)
	fmt.Println(middname, lastname)

}
