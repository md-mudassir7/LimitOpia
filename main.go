package main

import (
	"log"
	"net/http"

	"github.com/md-mudassir7/LimitOpia/algorithms"
	"github.com/md-mudassir7/LimitOpia/api/handlers"
	"github.com/md-mudassir7/LimitOpia/api/middleware"
)

func main() {
	algorithm := middleware.NewRateLimiterMiddleware(algorithms.NewTokenBucket(5, 1))

	// Example: Sliding Window Algorithm
	// algorithm := middleware.NewRateLimiterMiddleware(algorithms.NewSlidingWindow(10, time.Second))

	// Example: Leaky Bucket Algorithm
	// algorithm := middleware.NewRateLimiterMiddleware(algorithms.NewLeakyBucket(10, 5)) // 10 capacity, 5 leaks per second

	// Setup HTTP server
	mux := http.NewServeMux()
	mux.Handle("/example", algorithm(http.HandlerFunc(handlers.ExampleRouter)))

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
