package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to fiels in golang")
	content := "This is Rajesh Kumar Yadav"
	file, err := os.Create("./mylocalfile.txt")
	if err != nil {
		panic(err)
	}
	len, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is: ", len)
	readFile("./mylocalfile.txt")
	defer file.Close()
}

func readFile(filename string) {
	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text Data inside file is: \n", byteData)
}
