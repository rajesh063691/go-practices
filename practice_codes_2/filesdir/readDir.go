package main

import (
	"fmt"
	"os"
)

// temprory directory

const tmpDir = "/Users/rajesh/Desktop/tmpDir"

func main_readDir() {
	files, _ := os.ReadDir(tmpDir)
	//fmt.Printf("Hello")
	fmt.Println("<===Files Informations===>")
	for _, file := range files {
		// fmt.Printf("Name: %v, Size:%v kv, Mode:%v, IsDir:%v ",
		// 	file.Name(),
		// 	//file.Size(),
		// 	//file.Mode(),
		// 	file.IsDir(),
		// )

		_ = file

		fmt.Println()
	}
}
