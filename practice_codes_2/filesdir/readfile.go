package main

import (
	"fmt"
	"io/ioutil"
)

const tmpDir1 = "/Users/rajesh/Desktop/tmpDir"

func main_readfile() {
	//location of html file
	html, err := ioutil.ReadFile(tmpDir1 + "/page.html")
	if err != nil {
		fmt.Printf("Error in reading file: %v", err.Error())
	}

	fmt.Println("Raw File Data:", html)
	//string represenatation
	fmt.Println("String Data:", string(html))
}
