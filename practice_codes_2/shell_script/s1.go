package main

import (
	"fmt"
	"os/exec"
)

func main_s1() {

	// find `go` executable path
	goExecPath, err := exec.LookPath("go1")

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Go Executable: ", goExecPath)
	}

}
