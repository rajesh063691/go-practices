package anagram

import (
	"fmt"
	"sort"
	"strings"
)

// input 1 : listen
// input 2 : silent
func IsAnagram(input1 string, input2 string) bool {

	if strings.TrimSpace(input1) == "" || strings.TrimSpace(input2) == "" {
		return false
	}

	input1 = strings.ReplaceAll(input1, " ", "")
	input2 = strings.ReplaceAll(input2, " ", "")

	input1 = strings.ToLower(input1)
	input2 = strings.ToLower(input2)

	if len(input1) != len(input2) {
		return false
	}

	inp1Arr := strings.Split(input1, "")
	inp2Arr := strings.Split(input2, "")

	sort.Slice(inp1Arr, func(i, j int) bool {
		return inp1Arr[i] < inp1Arr[j]
	})

	sort.Slice(inp2Arr, func(i, j int) bool {
		return inp2Arr[i] < inp2Arr[j]
	})

	fmt.Println(inp1Arr)
	fmt.Println(inp2Arr)

	return strings.Join(inp1Arr, "") == strings.Join(inp2Arr, "")
}
