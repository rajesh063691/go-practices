package main

import (
	"fmt"
	"math"
)

// Maximum Sum of a Subarray of Size k
func findMaxSumInSubArray(arr []int, size int) (maxSum int) {
	if len(arr) == 0 || size == 0 {
		return maxSum
	}

	window_sum := 0
	start := 0
	for end := 0; end < len(arr)-1; end++ {
		window_sum += arr[end]
		// check if window is greater than given size
		if end >= size-1 {
			// we got one window
			maxSum = (int)(math.Max(float64(maxSum), float64(window_sum)))
			window_sum -= arr[start]
			start++
		}
	}

	return maxSum
}

// Find the length of the smallest subarray with a sum ≥ k
func findMinSumInSubArray(arr []int, sum int) (minLength int) {
	if len(arr) == 0 {
		return minLength
	}

	minLength = len(arr)
	window_sum := 0
	start := 0
	for end := 0; end < len(arr); end++ {
		window_sum += arr[end]
		for window_sum >= sum {
			minLength = int(math.Min(float64(minLength), float64(end-start+1)))
			window_sum -= arr[start]
			start++
		}
	}

	return minLength
}

func main() {
	arr := []int{2, 1, 5, 1, 3, 2}
	fmt.Println(findMaxSumInSubArray(arr, 3))
	arr = []int{2, 3, 1, 2, 4, 3}
	fmt.Println(findMinSumInSubArray(arr, 7))
}
