package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	tb := NewTokenBucket(10, 1)

	// Test initial capacity
	for i := 0; i < 10; i++ {
		if !tb.Allow() {
			t.Errorf("Expected request to be allowed")
		}
	}

	// Test exceeding capacity
	if tb.Allow() {
		t.Errorf("Expected request to be denied")
	}

	// Test refill
	time.Sleep(1 * time.Second)
	if !tb.Allow() {
		t.Errorf("Expected request to be allowed after refill")
	}
}

func TestTokenBucketConcurrent(t *testing.T) {
	tb := NewTokenBucket(10, 1)
	var wg sync.WaitGroup

	// Test concurrent access
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("tb.Allow()", tb.Allow(), "i", i)
		}(i)
	}
	wg.Wait()

	// Test refill after concurrent access
	time.Sleep(1 * time.Second)
	if !tb.Allow() {
		t.Errorf("Expected request to be allowed after refill")
	}
}

func TestTokenBucketNegative(t *testing.T) {
	tb := NewTokenBucket(0, 1)

	// Test zero capacity
	if tb.Allow() {
		t.Errorf("Expected request to be denied with zero capacity")
	}
}
