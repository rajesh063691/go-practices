package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimitter struct {
	Bucket       int // max token
	Token        int // current
	Rate         int // 10 token per interval - eg. 10 tokens per 60 seconds
	Interval     int
	LastRefilled time.Time  // latest refilled
	mu           sync.Mutex // for race condition
}

func NewRateLimitter(rate int, interval int) *RateLimitter {
	return &RateLimitter{
		Bucket:       rate,
		Token:        rate,
		Interval:     interval,
		Rate:         rate,
		LastRefilled: time.Now(),
	}
}

func (r *RateLimitter) apply() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	timeEclapsed := int(time.Since(r.LastRefilled).Seconds())
	newTokenAdded := timeEclapsed / r.Interval * r.Rate

	fmt.Printf("timeEclapsed:: %d \n", int(timeEclapsed))
	fmt.Printf("newTokenAdded:: %d \n", newTokenAdded)
	fmt.Printf("current token:: %d \n", r.Token)
	//fmt.Printf("Interval Provided:: %d \n", int(r.Interval.Seconds()))

	if newTokenAdded > 0 {
		r.Token += newTokenAdded
		if r.Token > r.Bucket {
			r.Bucket = r.Token
		}
		r.LastRefilled = time.Now()
	}

	if r.Token > 0 {
		r.Token--
		return true
	}

	return false
}

func main() {

	req := NewRateLimitter(10, 15)

	for i := 1; i <= 75; i++ {
		if req.apply() {
			fmt.Printf("Request no %d is processed \n", i)
		} else {
			fmt.Printf("Request no %d is rejected, waiting for token refill \n", i)
		}

		time.Sleep(time.Second)
	}

}
