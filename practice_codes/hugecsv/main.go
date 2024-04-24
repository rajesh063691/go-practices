package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	// Process start
	start := time.Now()
	// Create CSV file
	w, f := getCSVWriter()
	// Write CSV header
	writeCSVHeader(w)
	// Define data size
	records := 10000

	// ----------------------------------------
	// Uncomment each of functions below to run
	//testSequential(records, w)
	testConcurrencyWaitGroup(records, w)
	// testConcurrencyChannel(records, w)
	// ----------------------------------------

	// Flush buffer to file
	w.Flush()
	// Close the file
	err := f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	// Process finish then print elapsed time
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s to finish", elapsed)
}

func testSequential(records int, w *csv.Writer) {
	allData := make([]data, records)
	for i := 0; i < records; i++ {
		// Wait randomized 0~100ms to simulate API call
		//time.Sleep(getRandomSleepTime(100))
		// Assign dummy data
		allData[i] = getDummyData(i)
		//fmt.Printf("[Write " + allData[i].Name + "] ")
	}
	// Write CSV body
	writeCSVBody(allData, w)
	fmt.Println("Write CSV " + strconv.Itoa(records) + " records is finished")
}

// Note this function shares memory for each goroutine, which violates Go principle
// "Do not communicate by sharing memory; instead, share memory by communicating."
// I just put this as an example
func testConcurrencyWaitGroup(records int, w *csv.Writer) {
	allData := make([]data, records)
	// Define WaitGroup
	var wg sync.WaitGroup
	wg.Add(records)
	for i := 0; i < records; i++ {
		go func(i int) {
			// Mark gotoutine as as finished when data is assigned
			defer wg.Done()
			// Wait randomized 0~100ms to simulate API call
			time.Sleep(getRandomSleepTime(100))
			// Assign dummy data
			allData[i] = getDummyData(i)
			//fmt.Printf("[Write " + allData[i].Name + "] ")
		}(i)
	}
	// Make WaitGroup wait for all goroutines to finish
	wg.Wait()
	// Write CSV body
	writeCSVBody(allData, w)
	fmt.Println("Write CSV " + strconv.Itoa(records) + " records is finished")
}

func testConcurrencyChannel(records int, w *csv.Writer) {
	// Define buffered channel
	ch := make(chan data, records)
	done := make(chan bool)
	// Close channel only if sending is finished
	defer close(ch)
	for i := 0; i < records; i++ {
		go func(i int) {
			// Wait randomized 0~100ms to simulate API call
			time.Sleep(getRandomSleepTime(100))
			// Send data to channel
			ch <- getDummyData(i)
			fmt.Printf("[Write " + getDummyData(i).Name + "] ")
		}(i)
	}
	// Write CSV body
	go writeCSVBodyWithChannel(ch, done, records, w)
	// Notify main goroutine process is finished
	<-done
	fmt.Println("Write CSV " + strconv.Itoa(records) + " records is finished")
}

type data struct {
	ID          string
	Name        string
	Description string
}

func getRandomSleepTime(denom int) time.Duration {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return time.Duration(r.Intn(denom)) * time.Millisecond
}

func getDummyData(num int) data {
	return data{
		ID:          strconv.Itoa(num + 1),
		Name:        "Name" + strconv.Itoa(num+1),
		Description: "Desc",
	}
}

func getCSVWriter() (*csv.Writer, *os.File) {
	f, err := os.Create("./csv/concurrency.csv")
	if err != nil {
		log.Fatalln(err)
	}
	w := csv.NewWriter(f)
	w.UseCRLF = true
	return w, f
}

func writeCSVHeader(w *csv.Writer) {
	err := w.Write([]string{
		"ID",
		"Name",
		"Description",
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func writeCSVBody(allData []data, w *csv.Writer) {
	// Write data from slice to CSV
	for _, data := range allData {
		err := w.Write([]string{
			data.ID,
			data.Name,
			data.Description,
		})
		if err != nil {
			log.Fatalln("File writing failed")
			log.Fatalln(err)
		}
	}
}

func writeCSVBodyWithChannel(ch chan data, done chan bool, records int, w *csv.Writer) {
	// Write data from channel to CSV
	for data := range ch {
		err := w.Write([]string{
			data.ID,
			data.Name,
			data.Description,
		})
		if err != nil {
			log.Fatalln("File writing failed")
			log.Fatalln(err)
		}
		records--
		// Check if all records are processed, if yes then notify channel
		if records == 0 {
			done <- true
		}
	}
}
