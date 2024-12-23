package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	RateLimit struct {
		Algorithm   string `yaml:"algorithm"`
		TokenBucket struct {
			Capacity   int `yaml:"capacity"`
			RefillRate int `yaml:"refill_rate"`
		} `yaml:"token_bucket"`
	} `yaml:"rate_limit"`
}

var rateLimiter RateLimiter

func main() {
	// Load configuration
	configFile, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	var config Config
	decoder := yaml.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		panic(err)
	}

	// Initialize rate limiter based on config
	switch config.RateLimit.Algorithm {
	case "token_bucket":
		rateLimiter = NewTokenBucket(config.RateLimit.TokenBucket.Capacity, config.RateLimit.TokenBucket.RefillRate)
	default:
		panic("Unsupported rate limit algorithm")
	}

	http.HandleFunc("/", rateLimitMiddleware(apiHandler))
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !rateLimiter.Allow() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next(w, r)
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!  ")
}
