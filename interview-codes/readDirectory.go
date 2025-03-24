package main

import (
	"fmt"
	"os"
)

// call isDirectoty with base directory
// if null return
// check if directory - append the base dire with / in and cretae whole new directory and call isDirectory
// if not then loop out through all files and append to file slice or map of slice

var dirMap = make(map[string][]string)

func main_d() {

	basedir := "/Users/rajesh/Desktop/dir_rajesh"
	readDirectory(basedir)

	for k, v := range dirMap {
		fmt.Println(k, "=", v)
	}

}

func readDirectory(basedir string) {
	if basedir == "" {
		return
	}

	dir, err := os.ReadDir(basedir)
	if err != nil {
		return
	}

	for _, d := range dir {
		if d.Type().IsDir() {
			newDir := basedir + "/" + d.Name()
			readDirectory(newDir)
		} else {
			sl := []string{}
			sl = append(sl, dirMap[basedir]...)
			dirMap[basedir] = append(sl, d.Name())
		}
	}
}

// write a code to find out all the directories and files in a given directory.
