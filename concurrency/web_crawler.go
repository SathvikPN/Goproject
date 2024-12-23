// Exercise: https://go.dev/tour/concurrency/10
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	var str_map = make(map[string]bool)
	var mu sync.Mutex
	var wg sync.WaitGroup

	var crawler func(string, int)
	crawler = func(url string, depth int) {
		defer wg.Done() // decrement waitgroup count when current goroutine finishes.

		if depth <= 0 {
			return
		}

		// A mutex does not lock specific data structures or variables directly;
		// instead, it locks the code that accesses those resources
		// any code between a call to Lock() and Unlock()
		// on a mutex is protected from concurrent access by other goroutines
		mu.Lock()
		if _, ok := str_map[url]; ok {
			mu.Unlock()
			return
		} else {
			str_map[url] = true
			mu.Unlock()
		}

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q %q\n", url, body, urls)

		for _, u := range urls {
			wg.Add(1) // add count to waitGroup for this goroutine
			go crawler(u, depth-1)
		}
	}
	wg.Add(1)
	crawler(url, depth)
	wg.Wait() // blocks main function until wg count reaches zero
}

// ------------------------------------------------------------------------------
// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
