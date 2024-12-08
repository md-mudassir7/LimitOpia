# LimitOpia

A flexible, extensible rate-limiting library written in GoLang with support for multiple algorithms, middleware integration, and Redis-based rate limiting.


# Features

**Rate-Limiting Algorithms**

1. Token Bucket

2. Fixed Window

3. Sliding Window

4. Leaky Bucket


**Redis Integration for distributed rate limiting.**

**Middleware Support for easy integration with HTTP servers.**


# Getting Started

**Run the Project**

```
go run main.go
```
The server will start on http://localhost:8080.



**Example API Endpoint**

```
curl http://localhost:8080/example
```

By default, rate limiting is applied using the Token Bucket algorithm with 5 requests per second.


# Usage

**Middleware Integration**


1. Sliding Window
```
algorithm := middleware.NewRateLimiterMiddleware(algorithms.NewSlidingWindow(10, time.Second))
```

2. Leaky Bucket
```
algorithm := middleware.NewRateLimiterMiddleware(algorithms.NewLeakyBucket(10, 5))
```

3. Redis-Based Rate Limiting
```
redisLimiter := redis.NewRedisLimiter(redisClient, "example_key", 10, time.Second)
algorithm := middleware.NewRateLimiterMiddleware(redisLimiter)
```

