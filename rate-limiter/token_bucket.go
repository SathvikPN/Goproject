package main

import (
	"sync"
	"time"
)

// TokenBucket implements the Token Bucket rate-limiting algorithm.
type TokenBucket struct {
	capacity   int                     // Max tokens the bucket can hold
	tokens     int                     // Current token count
	refillRate int                     // Tokens added per second
	lastRefill time.Time               // Last refill timestamp
	mu         sync.Mutex              // Protects concurrent access
	buckets    map[string]*TokenBucket // Per-client buckets
}

// NewTokenBucket creates a new TokenBucket instance.
func NewTokenBucket(capacity, refillRate int) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
		buckets:    make(map[string]*TokenBucket),
	}
}

// refill adds tokens to the bucket based on elapsed time.
func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	newTokens := int(elapsed) * tb.refillRate
	if newTokens > 0 {
		tb.tokens = min(tb.capacity, tb.tokens+newTokens)
		tb.lastRefill = now
	}
}

// Allow checks if a request is allowed for the client.
func (tb *TokenBucket) Allow(clientID string) bool {
	return true // TODO
}

// Helper function to get the smaller of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
