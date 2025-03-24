package main

import (
	"fmt"
	"sync"
	"time"
)

// will be called based on the order of the init functions

// func init() {
// 	fmt.Println("Hi")
// }

// func init() {
// 	fmt.Println("Hello")
// }

type RateLimiter struct {
	rate      int           // tokens added per interval
	bucket    int           // bucket size (max tokens)
	tokens    int           // current tokens
	interval  time.Duration // interval between token refills
	lastCheck time.Time     // last time tokens were refilled
	mu        sync.Mutex
}

func NewRateLimiter(rate int, bucket int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		rate:      rate,
		bucket:    bucket,
		tokens:    bucket,
		interval:  interval,
		lastCheck: time.Now(),
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastCheck)

	// Add tokens based on the time elapsed
	addedTokens := int(elapsed/rl.interval) * rl.rate

	if addedTokens > 0 {
		rl.tokens += addedTokens
		if rl.tokens > rl.bucket {
			rl.tokens = rl.bucket
		}
		rl.lastCheck = now
	}

	//rl.tokens = 0

	// Check if we have enough tokens
	if rl.tokens > 0 {
		rl.tokens--
		return true
	}

	return false
}

func main_1() {
	rl := NewRateLimiter(1, 5, time.Second)

	for i := 0; i < 10; i++ {
		if rl.Allow() {
			fmt.Println("Request", i+1, "allowed")
		} else {
			fmt.Println("Request", i+1, "denied")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
