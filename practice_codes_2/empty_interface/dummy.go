package main

import "fmt"

// write a function which reads value from map of interface of interface and put it in the map[ string]string

func readFromInterface(i interface{}) (m map[string]string) {
	// initialize the map
	m = make(map[string]string)

	// read from interface and do type assertion
	if imp, ok := i.(map[interface{}]interface{}); ok {
		for k, v := range imp {
			if key, keyOK := k.(string); keyOK {
				if val, valOK := v.(string); valOK {
					m[key] = val
				}
			}
		}
	}

	return m
}

func main() {
	mp := map[interface{}]interface{}{
		"Name":   "Rajesh Kumar Yadav",
		"Age":    "34",
		"Gender": "Male",
	}

	p := readFromInterface(mp)
	for k, v := range p {
		fmt.Printf("%s :: %s \n", k, v)
	}

}
