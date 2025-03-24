package main

import (
	"fmt"
	"sync"
)

// doSum - finds sum of each row and put it in the channel
func doSum(arr []int, wg *sync.WaitGroup, sumCh chan int) {
	defer wg.Done()
	sum := 0
	for _, item := range arr {
		sum += item
	}
	sumCh <- sum
}

func main_d() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
		{13, 14, 15},
	}

	wg := &sync.WaitGroup{}
	// create a channle of array length
	ch := make(chan int, len(arr))
	for _, row := range arr {
		wg.Add(1)
		go doSum(row, wg, ch)
	}
	// wait for all goroutines to close
	wg.Wait()
	// signale the channel to close
	close(ch)

	// llop over channel and add the individual row sum
	totalSum := 0
	for sum := range ch {
		totalSum += sum
	}

	fmt.Printf("Total Sum :: %d \n", totalSum)
}
