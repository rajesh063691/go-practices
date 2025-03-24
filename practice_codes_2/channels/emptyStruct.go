package main

import (
	"fmt"
	"sync"
)

func main_empty() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	c := make(chan struct{}, 1)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			c <- struct{}{} // acquire lock
			defer func() {
				<-c // release lock
				wg.Done()
			}()

			fmt.Printf("goroutine %d: hello\n", i)
		}(i)
	}

	//time.Sleep(2 * time.Second)
	fmt.Scanln()
	//wg.Wait()
}
