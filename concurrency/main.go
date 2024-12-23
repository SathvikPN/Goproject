package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func main() {
	// equivalent binary trees
	t1 := tree.New(1)
	t2 := tree.New(2)
	tx := tree.New(1)
	fmt.Println("Same(t1, tx)", Same(t1, tx))
	fmt.Println("Same(t1, t2)", Same(t1, t2))

	// web crawler
	Crawl("https://golang.org/", 4, fetcher)
}
