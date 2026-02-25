package main

import (
	"fmt"

	"github.com/majabojarska/fibo/internal/fibonacci"
)

func main() {
	for fibVal := range fibonacci.Fibonacci(10) {
		fmt.Println(fibVal)
	}
}
