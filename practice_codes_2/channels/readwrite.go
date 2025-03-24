package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	wg *sync.WaitGroup
}

func (d *data) writeData(ch chan<- int) {
	defer d.wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("Writing %d to channel...\n", i)
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch) // Close the channel when done writing
}

func (d *data) readData(ch <-chan int) {
	defer d.wg.Done()
	for val := range ch {
		fmt.Printf("Reading %d from channel...\n", val)
		time.Sleep(1 * time.Second)
	}
}

func main_read() {

	d := &data{wg: &sync.WaitGroup{}}

	// Create a channel for communication
	ch := make(chan int, 5)

	// Start a goroutine to write data
	d.wg.Add(1)
	go d.writeData(ch)

	// Start a goroutine to read data
	d.wg.Add(1)
	go d.readData(ch) // This runs in the main goroutine, blocking until the channel is closed
	d.wg.Wait()
	fmt.Println("All data processed.")
}
