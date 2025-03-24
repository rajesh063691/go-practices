package main

import (
	"fmt"
	"strings"
)

// func testVariadicFunc(args ...string)(strVal []string, totalArgs int){
// 	return args, len(args)
// }

func testVariadicFunc(args ...interface{}) (str string, totalArgs int) {
	for i, val := range args {
		if strVal, ok := val.(string); ok {
			str += strVal + ", "
		} else {
			fmt.Printf("Mismatched interface type found of Type:%T with Value:%+v\n", args[i], args[i])
			return "", 0
		}

	}
	return strings.TrimSuffix(str, ", "), len(args)
}

func CommonInterface(inter interface{}) {
	if mapVal, ok := inter.(map[string]string); ok {
		for key, val := range mapVal {
			fmt.Printf("Key=%s, Val=%s \n", key, val)
		}
	} else {
		fmt.Printf("Invalid interface type conversion")
	}
}

type NameDetails struct {
	firstName  string
	middleName string
	lastName   string
}

func CommonInterface1(inter interface{}) {

	if st, ok := inter.(NameDetails); ok {
		fmt.Printf("Struct=%+v \n", st)
	} else {
		fmt.Printf("Invalid interface type conversion")
	}
}

func CommonInterface2(inter interface{}) {

	strMap := make(map[string]string)

	if interfaceMap, ok := inter.(map[interface{}]interface{}); ok {

		for key, val := range interfaceMap {
			if strKey, keyOk := key.(string); keyOk {
				if strVal, valOk := val.(string); valOk {
					strMap[strKey] = strVal
				}
			}
		}
		fmt.Printf("String MAP=%+v \n", strMap)

	} else {
		fmt.Printf("Invalid interface type conversion")
	}
}

func main_3() {
	//0th case
	// strArr:=make([]string,10)

	// strArr[0]="Rajesh"
	// strArr[1]="Kumar"
	// strArr[2]="Yadav"

	//str, total:=testVariadicFunc("rajesh", "kumar", "yadav")
	//str, total:=testVariadicFunc("123","2","2234")
	//fmt.Printf("Val=%s and Total Arguments=%d \n",str, total)
	//1st case
	map1 := map[string]string{
		"firstName":  "Rajesh",
		"MiddleName": "Kumar",
		"LastName":   "Yadav",
	}
	CommonInterface(map1)
	//2nd case
	nd := NameDetails{
		firstName:  "Rajesh",
		middleName: "Kumar",
		lastName:   "Yadav",
	}
	CommonInterface1(nd)
	//3rd case
	inMap := map[interface{}]interface{}{
		"firstName":  "Puja",
		"MiddleName": "Kumari",
		"LastName":   "Yadav",
	}

	CommonInterface2(inMap)
}
