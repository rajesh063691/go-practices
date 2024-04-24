package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// temporary directory
const tmpDir2 = "/Users/rajesh/Desktop/tmpDir"

func main_wriet() {

	// welcome message content
	welcomeData := []byte("Welcome to my website.")

	// get `welcome.txt` file path
	path := filepath.Join(tmpDir2, "/welcome.txt")

	// write content to `welcome.txt` file
	err := ioutil.WriteFile(path, welcomeData, 0777)

	// log error
	if err != nil {
		fmt.Println(err)
	}

}
