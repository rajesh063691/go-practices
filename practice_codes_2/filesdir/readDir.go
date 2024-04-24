package main

import (
	"fmt"
	"io/ioutil"
)

// temprory directory

const tmpDir = "/Users/rajesh/Desktop/tmpDir"

func main_readDir() {
	files, _ := ioutil.ReadDir(tmpDir)
	//fmt.Printf("Hello")
	fmt.Println("<===Files Informations===>")
	for _, file := range files {
		fmt.Printf("Name: %v, Size:%v kv, Mode:%v, IsDir:%v ",
			file.Name(),
			file.Size(),
			file.Mode(),
			file.IsDir(),
		)

		fmt.Println()
	}
}
