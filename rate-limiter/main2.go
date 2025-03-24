package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiterr struct {
	bucket       int
	token        int
	refillRate   int // 10 token every 10 seconds
	lastRefilled time.Time
	mu           sync.Mutex
}

func NewRateLimiterr(bucketSize int, refil int) *RateLimiterr {
	return &RateLimiterr{
		bucket:       bucketSize,
		token:        bucketSize,
		refillRate:   refil,
		lastRefilled: time.Now(),
	}
}

func (r *RateLimiterr) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	// check the time if needs to be refilled the bucket with new token
	now := time.Now()
	eclapsed := now.Sub(r.lastRefilled).Seconds()

	// check if eclapsed time > 60 seconds then refill the bucket with 10 new tokens
	//if eclapsed > 10 {
	r.refilBucket(eclapsed)
	//}

	// check if sufficient token is left or not if yes then consume one
	if r.token > 0 {
		r.token--
		return true
	}
	return false
}

func (r *RateLimiterr) refilBucket(eclapsed float64) {
	if eclapsed > float64(r.refillRate) {
		fmt.Printf("refilling bucket with new token \n")
		toketRefill := eclapsed / float64(r.refillRate) * float64(r.refillRate)
		fmt.Println("new token added :: ", toketRefill)
		if toketRefill > 0 {
			r.token += int(toketRefill)
			if r.token > r.bucket {
				r.token = r.bucket
			}
			r.lastRefilled = time.Now()
		}
	}

}

func main2() {
	r := NewRateLimiterr(8, 10)
	fmt.Printf("Initial Bucket size=%d and token=%d \n", r.bucket, r.token)
	counter := 1
	for {
		if !r.Allow() {
			fmt.Printf("bucket is empty, wait for refill... (bucket size:%d) refill time :: every 10 seconds \n", r.bucket)
		} else {
			fmt.Printf("token # %d is consumed, remaining token in the bucket :: %d \n", counter, r.token)
		}
		counter++
		time.Sleep(1 * time.Second)
		if counter > 30 {
			fmt.Println("60 iteration is done, breaking the loop")
			break
		}

	}

}
