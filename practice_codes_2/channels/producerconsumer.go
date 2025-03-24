package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 1
	for {
		fmt.Println("produced :: ", i)
		ch <- i
		i++
		time.Sleep(time.Second)

		if i == 11 {
			break
		}
	}
	close(ch)
}

func consume(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range ch {
		fmt.Printf("consumed :: %d\n", msg)
		time.Sleep(time.Second)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 3)

	wg.Add(2)
	go producer(ch, wg)
	go consume(ch, wg)

	wg.Wait()

	fmt.Println("All msg consumed")
}
