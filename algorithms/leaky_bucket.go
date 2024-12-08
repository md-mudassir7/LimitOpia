package algorithms

import (
	"log"
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity   int
	rate       int // Leak rate (tokens per second)
	waterLevel int
	lastLeak   time.Time
	mu         sync.Mutex
}

func NewLeakyBucket(capacity, rate int) *LeakyBucket {
	log.Println("Using Leaky Bucket algorithm for rate limiting", "capacity", capacity, "rate", rate)
	return &LeakyBucket{
		capacity: capacity,
		rate:     rate,
		lastLeak: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(lb.lastLeak).Seconds()

	lb.waterLevel -= int(elapsed * float64(lb.rate))
	if lb.waterLevel < 0 {
		lb.waterLevel = 0
	}
	lb.lastLeak = now

	if lb.waterLevel < lb.capacity {
		lb.waterLevel++
		return true
	}
	return false
}
