package fibonacci_test

import (
	"fmt"
	"reflect"
	"slices"
	"testing"

	"github.com/majabojarska/fibo/internal/fibonacci"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		wantCount int
		wantSlice []int
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

			if !reflect.DeepEqual(got, tt.wantSlice) {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.wantSlice)
			}
		})
	}
}
