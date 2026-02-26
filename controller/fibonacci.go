package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/majabojarska/fibo/internal/fibonacci"
)

type FibonacciRequest struct {
	Count int `json:"count" example:"5" binding:"required,min=0"`
}

func (c *Controller) GetFibonacci(ctx *gin.Context) {
	writer := ctx.Writer

	countRaw := ctx.Param("count")
	requestedCount, err := strconv.Atoi(countRaw)
	if err != nil {
		if err := ctx.AbortWithError(http.StatusBadRequest, err); err != nil {
			panic(err)
		}
		return
	}

	if requestedCount < 0 {
		if err := ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("count must be equal to or greater than 0")); err != nil {
			panic(err)
		}
		return
	}

	header := writer.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Transfer-Encoding", "chunked")
	writer.WriteHeader(http.StatusOK)

	if _, err := writer.WriteString("["); err != nil {
		panic(err)
	}
	writer.(http.Flusher).Flush()

	sentCount := 0
	for fiboVal := range fibonacci.Fibonacci(requestedCount) {
		// Sequence item
		if _, err := writer.WriteString(strconv.Itoa(fiboVal)); err != nil {
			panic(err)
		}

		sentCount++

		// Omit delimiter after the last seq item
		if sentCount < requestedCount {
			if _, err := writer.WriteString(","); err != nil {
				panic(err)
			}
		}

		writer.(http.Flusher).Flush()
	}

	if _, err := writer.WriteString("]"); err != nil {
		panic(err)
	}
}
