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
func TestFibonacciSeqExhaustSeq(t *testing.T) {
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
			fiboIter := fibonacci.FibonacciSeq(tt.wantCount)
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

func TestFibonacciSeq64bitOverflow(t *testing.T) {
	tests := []struct {
		itemNum       int
		wantValuesStr string
	}{
		{
			93, // Largest one that fits in int64
			"7540113804746346429",
		},
		{
			94, // First one to overflow
			"12200160415121876738",
		},
		{
			100, // Just to be sure, against off-by-ones
			"218922995834555169026",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Fibonacci item %d", tt.itemNum), func(t *testing.T) {
			var last *big.Int
			for last = range fibonacci.FibonacciSeq(tt.itemNum) {
			}

			if last.String() != tt.wantValuesStr {
				t.Errorf("Fibonacci item %d = %s, want %s", tt.itemNum, last.String(), tt.wantValuesStr)
			}
		})
	}
}

// TestFibonacciStopEarly tests compatibility with range loops and iterator adapters,
// by calling stop early, before the sequence is exhausted.
func TestFibonacciSeqStopEarly(t *testing.T) {
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

			fiboSeq := fibonacci.FibonacciSeq(tt.wantCount)

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

func TestFibonacciSeq2ExhaustSeq(t *testing.T) {
	tests := []struct {
		wantCount    int
		wantIndices  []int
		wantFiboVals []int // Use int for easier maintenance
	}{
		{
			0,
			nil,
			nil,
		},
		{
			1,
			[]int{0},
			[]int{0},
		},
		{
			4,
			[]int{0, 1, 2, 3},
			[]int{0, 1, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Fibonacci %d items", tt.wantCount), func(t *testing.T) {
			if tt.wantCount == 0 {
				for range fibonacci.FibonacciSeq2(tt.wantCount) {
					t.Errorf("Iterator should be empty")
				}
			} else {
				testIndex := 0
				for fiboIndex, fiboVal := range fibonacci.FibonacciSeq2(tt.wantCount) {
					if tt.wantIndices[testIndex] != fiboIndex {
						t.Errorf("Index %d = %d, want %s", testIndex, fiboIndex, fiboVal)
					}
					wantFiboVal := big.NewInt(int64(tt.wantFiboVals[testIndex]))
					if wantFiboVal.Cmp(fiboVal) != 0 {
						t.Errorf("Index %d = %v, want %v", testIndex, wantFiboVal, fiboVal)
					}
					testIndex++
				}
			}
		})
	}
}

// TestFibonacciStopEarly tests compatibility with range loops and iterator adapters,
// by calling stop early, before the sequence is exhausted.
func TestFibonacciSeq2StopEarly(t *testing.T) {
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

			loopCount := 0
			for range fibonacci.FibonacciSeq2(tt.wantCount) {
				if loopCount == tt.wantCount {
					break
				}
				loopCount++
			}
			// Should not panic, iterator should stop cleanly.
		})
	}
}
