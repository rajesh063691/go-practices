package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"codes/arrprac"
)

// Write your code here
func mainarrprac() {
	//fmt.Println("Enter length of bracket sequence")
	length := 0
	fmt.Scanf("%d", &length)
	var bracketSeq = make([]byte, length)
	//fmt.Printf("Enter bracket sequence of length =%d", length)
	//fmt.Println()
	for i := 0; i < length; i++ {
		in := bufio.NewReader(os.Stdin)
		byteVal, err := in.ReadByte()
		if err == io.EOF {
			break
		}
		bracketSeq = append(bracketSeq, byteVal)

	}
	outlen := arrprac.BracketSequence(bracketSeq)
	fmt.Println(outlen)
}
