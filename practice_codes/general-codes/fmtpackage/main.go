package main

import (
	"fmt"
	"time"
)

func main() {

	//inside %10d means, it says about its size it can occupy
	// like if field Rajesh has %10s So its total width will be 10 strting from left to right i.e    Rajesh
	// and if same has %-10s, then it will be same but in oppsite direction
	//i.eRajesh and 4 space bcoz Rajesh already has 6 letter so remaining space will be 4
	fmt.Printf("%-40s  %-10s  %-10s  %-60s\n", "Repo", "ID", "Status", "Title")
	fmt.Printf("%-40s  %-10s  %-10s  %-60s\n", "=====", "=====", "=====", "=====")

	fmt.Printf("width5: |%-5s|%-6s|\n", "foo", "b")
	fmt.Printf("width6: |%5s|%6s|\n", "foo", "b")
	truncatedTime := time.Now().UTC().Truncate(time.Second)
	fmt.Println("Current Time in UTC: ", time.Now().UTC())
	fmt.Println("Truncated Time in UTC: ", truncatedTime)
}
