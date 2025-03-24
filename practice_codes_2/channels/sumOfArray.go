package main

import (
	"fmt"
	"runtime"
	"sync"
)

func sumRow(row []int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()

	sum := 0
	for _, val := range row {
		sum += val
	}
	resultChan <- sum
}

func main_sumArray() {
	// Define a 2D array
	array := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{9, 10, 11, 12},
		{9, 10, 11, 12},
		{9, 10, 11, 12},
		{9, 10, 11, 12},
		{9, 10, 11, 12},
		{9, 10, 11, 12},
		{9, 10, 11, 12},
	}

	var wg sync.WaitGroup
	resultChan := make(chan int, len(array)) // Buffer to avoid blocking
	fmt.Printf("current number of Go Routines :: %d \n", runtime.NumGoroutine())
	// Launch goroutines to sum each row
	for _, row := range array {
		wg.Add(1)
		go sumRow(row, &wg, resultChan)
	}

	fmt.Printf("current number of Go Routines :: %d \n", runtime.NumGoroutine())

	// Wait for all goroutines to complete
	wg.Wait()
	close(resultChan)

	// Aggregate the total sum
	totalSum := 0
	for sum := range resultChan {
		totalSum += sum
	}

	fmt.Printf("current number of Go Routines :: %d \n", runtime.NumGoroutine())

	fmt.Printf("Total sum of 2D array: %d\n", totalSum)
}
