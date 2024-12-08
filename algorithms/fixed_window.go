package algorithms

import (
	"log"
	"sync"
	"time"
)

type FixedWindow struct {
	limit      int
	windowSize time.Duration
	count      int
	windowEnd  time.Time
	mu         sync.Mutex
}

func NewFixedWindow(limit int, windowSize time.Duration) *FixedWindow {
	log.Println("Using Fixed Window algorithm for rate limiting", "limit", limit, "windowSize", windowSize)
	return &FixedWindow{
		limit:      limit,
		windowSize: windowSize,
		windowEnd:  time.Now().Add(windowSize),
	}
}

func (fw *FixedWindow) Allow() bool {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	now := time.Now()

	if now.After(fw.windowEnd) {
		fw.windowEnd = now.Add(fw.windowSize)
		fw.count = 0
	}

	if fw.count < fw.limit {
		fw.count++
		return true
	}
	return false
}
