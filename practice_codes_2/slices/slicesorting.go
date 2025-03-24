package main

import (
	"fmt"
	"slices"
	"sort"
)

func main_sort() {
	words := []string{"falcon", "bold", "bear", "sky", "cloud", "ocean"}
	vals := []int{4, 2, 1, 5, 6, 8, 0, -3}

	// both ways is fine and works simillary
	slices.Sort(words)
	slices.Sort(vals)

	sort.Strings(words)
	sort.Ints(vals)

	fmt.Println(words)
	fmt.Println(vals)

	sort.Slice(words, func(i, j int) bool {
		return i > j
	})

	// sort the slice using logic
}
