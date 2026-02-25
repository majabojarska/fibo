package main

import (
	"fmt"

	"github.com/majabojarska/fibo/internal/fibonacci"
)

func main() {
	desiredCount := 10

	idx := 0
	for fibVal := range fibonacci.Fibonacci() {
		if idx >= desiredCount {
			break
		}
		fmt.Println(idx, fibVal)
		idx++
	}
}
