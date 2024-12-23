// Exercise: https://go.dev/tour/concurrency/8
package main

import (
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}

	walker(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	// Start walking both trees concurrently
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// Loop to read values from both channels simultaneously
	for {
		// Receive values from both channels
		v1, isOpenCh1 := <-ch1 // isOpenCh1 is false when ch1 is closed
		v2, isOpenCh2 := <-ch2 // isOpenCh2 is false when ch2 is closed

		// Even if one Goroutine is faster,
		// it must wait for the 'Same' function to consume the value before releasing new value from chan.

		// If one channel is closed while the other is still open,
		// it means one tree has more values than the other
		if isOpenCh1 != isOpenCh2 {
			return false
		}

		// If the values differ, the trees are not the same
		if v1 != v2 {
			return false
		}

		// If both channels are closed, no mismatched val
		if !isOpenCh1 && !isOpenCh2 {
			// break
			return true
		}
	}

}
