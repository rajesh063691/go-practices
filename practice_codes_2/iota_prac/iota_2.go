package main

import "fmt"

type Size uint8

// const (
// 	small Size = iota
// 	Medium
// 	Low
// 	High
// )

// without iota
const (
	Small  Size = 0
	Medium Size = 1
	Low    Size = 2
	High   Size = 3
)

func main_2() {
	fmt.Println(Small)
	fmt.Println(Medium)
	fmt.Println(Low)
	fmt.Println(High)
}
