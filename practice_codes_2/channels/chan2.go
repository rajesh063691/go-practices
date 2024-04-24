package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(10)
}

func main_2() {
	fmt.Println("Start main")
	dataChan := make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for d := range dataChan {
		fmt.Printf("%d\n", d)
	}
	fmt.Println("Ends main")
}
