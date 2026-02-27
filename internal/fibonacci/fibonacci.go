// Package fibonacci implements Fibonacci sequence generation logic.
package fibonacci

import (
	"iter"
	"math/big"
)

// Fibonacci returns an iter.Seq[*big.Int], which generates
// a finite Fibonacci sequence, starting at 0.
func Fibonacci(count int) iter.Seq[*big.Int] {
	return func(yield func(*big.Int) bool) {
		returnCount := 0
		left, right := big.NewInt(0), big.NewInt(1)

		for {
			yieldVal := new(big.Int).Set(left)
			if returnCount >= count || !yield(yieldVal) {
				return
			}
			returnCount++

			left.Add(left, right)
			left, right = right, left
		}
	}
}
