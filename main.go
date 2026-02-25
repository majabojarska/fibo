package main

import (
	"fmt"
	"iter"
)

// fibonacci returns an iter.Seq2[int, int], which generates
// an infinite Fibonacci sequence, starting at 0.
func fibonacci() iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		idx := 0
		left, right := 0, 1

		for {
			if !yield(idx, left) {
				return
			}
			idx++
			left, right = right, left+right
		}
	}
}

func main() {
	desiredCount := 10

	for idx, fibVal := range fibonacci() {
		if idx >= desiredCount {
			break
		}
		fmt.Println(idx, fibVal)
	}
}
