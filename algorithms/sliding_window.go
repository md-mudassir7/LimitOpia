package algorithms

import (
	"log"
	"sync"
	"time"
)

type SlidingWindow struct {
	limit      int
	windowSize time.Duration
	events     []time.Time
	mu         sync.Mutex
}

func NewSlidingWindow(limit int, windowSize time.Duration) *SlidingWindow {
	log.Println("Using Sliding Window algorithm for rate limiting", "limit", limit, "windowSize", windowSize)

	return &SlidingWindow{
		limit:      limit,
		windowSize: windowSize,
		events:     []time.Time{},
	}
}

func (sw *SlidingWindow) Allow() bool {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	now := time.Now()

	validStart := now.Add(-sw.windowSize)
	for len(sw.events) > 0 && sw.events[0].Before(validStart) {
		sw.events = sw.events[1:]
	}

	if len(sw.events) < sw.limit {
		sw.events = append(sw.events, now)
		return true
	}
	return false
}
