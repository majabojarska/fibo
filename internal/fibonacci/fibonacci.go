// Package fibonacci implements Fibonacci sequence generation logic.
package fibonacci

import (
	"iter"
	"math/big"
)

// FibonacciSeq returns an iter.Seq[*big.Int], which generates
// a finite Fibonacci sequence, starting at 0.
func FibonacciSeq(count int) iter.Seq[*big.Int] {
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

// FibonacciSeq2 returns an iter.Seq2[int, *big.Int], which generates
// a finite iterator, where the first value is the index (0-indexed),
// and the second is the fibonacci sequence value.
func FibonacciSeq2(count int) iter.Seq2[int, *big.Int] {
	return func(yield func(int, *big.Int) bool) {
		yieldCount := 0

		for fiboVal := range FibonacciSeq(count) {
			if !yield(yieldCount, fiboVal) {
				return
			}
			yieldCount++
		}
	}
}
