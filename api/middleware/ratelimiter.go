package middleware

import (
	"net/http"
)

type RateLimiterAlgorithm interface {
	Allow() bool
}

func NewRateLimiterMiddleware(algorithm RateLimiterAlgorithm) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !algorithm.Allow() {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
