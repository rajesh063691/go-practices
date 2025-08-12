package main

import (
	"fmt"
	"sync"
	"time"
)

// Job represents the task that will be run by the worker
type Job struct {
	ID      int
	Payload string
}

// WorkerPool manages a pool of workers to process jobs
type WorkerPool struct {
	JobQueue    chan Job
	WorkerCount int
	wg          *sync.WaitGroup
}

// NewWorkerPool creates a new WorkerPool
func NewWorkerPool(workerCount int, jobQueueSize int) *WorkerPool {
	return &WorkerPool{
		JobQueue:    make(chan Job, jobQueueSize),
		WorkerCount: workerCount,
		wg:          &sync.WaitGroup{},
	}
}

// Start initializes the worker pool
func (wp *WorkerPool) Start() {
	for i := 1; i <= wp.WorkerCount; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// worker is a function that runs in a goroutine and processes jobs
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for job := range wp.JobQueue {
		fmt.Printf("Worker %d processing job %d with payload: %s\n", id, job.ID, job.Payload)
		time.Sleep(1 * time.Second) // Simulate work
	}
}

// AddJob adds a job to the job queue
func (wp *WorkerPool) AddJob(job Job) {
	wp.JobQueue <- job
}

// Wait waits for all workers to finish
func (wp *WorkerPool) Wait() {
	close(wp.JobQueue)
	wp.wg.Wait()
}

// commenting for better readability
func main() {
	// Create a worker pool with 3 workers and a job queue of size 10
	pool := NewWorkerPool(3, 10)

	// Start the worker pool
	pool.Start()

	// Add jobs to the job queue
	for i := 1; i <= 21; i++ {
		pool.AddJob(Job{ID: i, Payload: fmt.Sprintf("Payload %d", i)})
	}

	// Wait for all jobs to complete
	pool.Wait()

	fmt.Println("All jobs have been processed.")
}
