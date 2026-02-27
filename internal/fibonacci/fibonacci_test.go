package fibonacci_test

import (
	"fmt"
	"iter"
	"math/big"
	"reflect"
	"slices"
	"testing"

	"github.com/majabojarska/fibo/internal/fibonacci"
)

// TestFibonacciExhaustSeq tests sequence generation correctness (content, length),
// by exhausting all items from the sequence.
func TestFibonacciExhaustSeq(t *testing.T) {
	tests := []struct {
		wantCount int
		wantItems []int // Use int for easier maintenance
	}{
		{
			0,
			nil,
		},
		{
			1,
			[]int{0},
		},
		{
			2,
			[]int{0, 1},
		},
		{
			10,
			[]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Fibonacci %d items", tt.wantCount), func(t *testing.T) {
			fiboIter := fibonacci.Fibonacci(tt.wantCount)
			got := slices.Collect(fiboIter)

			var wantItemsBigInt []*big.Int

			if len(tt.wantItems) > 0 {
				wantItemsBigInt = make([]*big.Int, len(tt.wantItems))
				for idx, item := range tt.wantItems {
					wantItemsBigInt[idx] = big.NewInt(int64(item))
				}
			}

			if !reflect.DeepEqual(got, wantItemsBigInt) {
				t.Errorf("Fibonacci() = %v, want %v", got, wantItemsBigInt)
			}
		})
	}
}

// TestFibonacciStopEarly tests compatibility with range loops and iterator adapters,
// by calling stop early, before the sequence is exhausted.
func TestFibonacciStopEarly(t *testing.T) {
	tests := []struct {
		wantCount      int
		stopAfterCount int
	}{
		{
			0,
			0,
		},
		{
			1,
			0,
		},
		{
			1,
			1,
		},
		{
			2,
			1,
		},
		{
			10,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Want %d items, break range iteration after %d items", tt.wantCount, tt.stopAfterCount), func(t *testing.T) {
			if tt.wantCount < tt.stopAfterCount {
				t.Errorf("wantCount (%d) must be equal to, or greater than, stopAfterCount (%d)", tt.wantCount, tt.stopAfterCount)
			}

			fiboSeq := fibonacci.Fibonacci(tt.wantCount)

			// Convert to a pull iterator to facilitate fetching individual items.
			next, stop := iter.Pull(fiboSeq)

			for pulledCount := range tt.wantCount {
				if pulledCount == tt.stopAfterCount {
					stop()
				}

				_, ok := next()

				if pulledCount < tt.stopAfterCount {
					if !ok {
						t.Errorf("Iterator stopped before stop was called.")
					}
				} else {
					if ok {
						t.Errorf("Iterator didn't stop after stop was called")
					}
				}

			}
		})
	}
}
