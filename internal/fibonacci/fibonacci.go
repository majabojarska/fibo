package fibonacci

import "iter"

// Fibonacci returns an iter.Seq[int], which generates
// an infinite Fibonacci sequence, starting at 0.
func Fibonacci() iter.Seq[int] {
	return func(yield func(int) bool) {
		left, right := 0, 1

		for {
			if !yield(left) {
				return
			}
			left, right = right, left+right
		}
	}
}
