package main

import (
	"fmt"
)

// Function to find the largest sequential substring
func findLargestSequentialSubstring(input string) string {
	if len(input) == 0 {
		return ""
	}

	longestSeq := ""
	currentSeq := string(input[0])

	for i := 1; i < len(input); i++ {
		// Check if the current character is sequential to the previous one
		if input[i] == input[i-1]+1 {
			currentSeq += string(input[i])
		} else {
			// Update longest sequence if needed
			if len(currentSeq) > len(longestSeq) {
				longestSeq = currentSeq
			}
			// Reset current sequence
			currentSeq = string(input[i])
		}
	}

	// Final check in case the longest sequence ends at the last character
	if len(currentSeq) > len(longestSeq) {
		longestSeq = currentSeq
	}

	return longestSeq
}

func main() {
	input := "abcfghijklmabxyzdefghijklmnopqr"
	longestSeq := findLargestSequentialSubstring(input)
	fmt.Printf("The largest sequential substring is: %s\n", longestSeq)
}
