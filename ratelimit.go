package main

import (
	"sync"

	"golang.org/x/time/rate"
)

// IPRateLimiter stores rate limiters for different IP addresses
type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

// NewIPRateLimiter creates a new rate limiter for each IP
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	return &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}
}

// GetLimiter returns the rate limiter for the provided IP address
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.RLock()
	limiter, exists := i.ips[ip]
	i.mu.RUnlock()

	if !exists {
		i.mu.Lock()
		limiter = rate.NewLimiter(i.r, i.b)
		i.ips[ip] = limiter
		i.mu.Unlock()
	}

	return limiter
}

// func RateLimiterMiddleware(c *gin.Context, limiter *IPRateLimiter) {
// 	ip := c.ClientIP()
//
// 	singleIPLimiter := limiter.GetLimiter(ip)
// 	if !singleIPLimiter.Allow() {
// 		c.JSON(http.StatusTooManyRequests, gin.H{
// 			"message": "Too many requests. Please try again later.",
// 		})
// 		c.Abort()
// 		return
// 	}
// 	c.Next()
// }
