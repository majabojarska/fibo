package fibonacci

import "iter"

// Fibonacci returns an iter.Seq2[int, int], which generates
// an infinite Fibonacci sequence, starting at 0.
func Fibonacci() iter.Seq2[int, int] {
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
