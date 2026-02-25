// Package fibonacci implements Fibonacci sequence generation logic.
package fibonacci

import "iter"

// Fibonacci returns an iter.Seq[int], which generates
// a finite Fibonacci sequence, starting at 0.
func Fibonacci(count int) iter.Seq[int] {
	return func(yield func(int) bool) {
		returnCount := 0
		left, right := 0, 1

		for {
			if returnCount >= count || !yield(left) {
				return
			}
			left, right = right, left+right
			returnCount++
		}
	}
}
