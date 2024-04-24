package main

import (
	"fmt"
	"sort"
)

func main_sort() {
	words := []string{"falcon", "bold", "bear", "sky", "cloud", "ocean"}
	vals := []int{4, 2, 1, 5, 6, 8, 0, -3}

	sort.Strings(words)
	sort.Ints(vals)

	fmt.Println(words)
	fmt.Println(vals)
}
