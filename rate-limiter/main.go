package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the structure of the configuration file.
type Config struct {
	Algorithm   string `yaml:"algorithm"`
	TokenBucket struct {
		Capacity   int `yaml:"capacity"`
		RefillRate int `yaml:"refill_rate"`
	} `yaml:"token_bucket"`
}

// RateLimiter is an interface for any rate-limiting algorithm.
type RateLimiter interface {
	Allow(clientID string) bool
}

// RateLimiterMiddleware applies rate limiting to a handler.
func RateLimiterMiddleware(limiter RateLimiter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientID := r.RemoteAddr // Use client IP as the unique identifier
		if !limiter.Allow(clientID) {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next(w, r)
	}
}

// SimpleHandler handles requests to the main endpoint.
func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request successful!")
}

func main() {
	var rateLimiter RateLimiter
	// Load configuration from file
	config, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Validate algorithm
	switch config.Algorithm {
	case "TokenBucket":
		tbConfig := config.TokenBucket
		rateLimiter = NewTokenBucket(tbConfig.Capacity, tbConfig.RefillRate)
	default:
		fmt.Errorf("unsupported rate limiter algorithm: %s", config.Algorithm)
		return
	}

	// Set up the HTTP server with the rate-limiting middleware
	http.Handle("/", RateLimiterMiddleware(rateLimiter, SimpleHandler))

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// LoadConfig loads the YAML configuration file.
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
