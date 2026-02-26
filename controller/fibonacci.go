package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/majabojarska/fibo/internal/fibonacci"
)

// GetFibonacci godoc
//
//	@Summary		Get Fibonacci sequence
//	@Description	Generates a finite length Fibonacci sequence
//	@ID				get-fibonacci
//	@Tags			fibonacci
//	@Param			count	path		int		true	"Desired sequence size"
//	@Success		200		{object}	{array}	int
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/fibonacci/{count} [get]
func (c *Controller) GetFibonacci(ctx *gin.Context) {
	writer := ctx.Writer

	countRaw := ctx.Param("count")
	requestedCount, parseErr := strconv.Atoi(countRaw)
	if parseErr != nil {
		if err := ctx.AbortWithError(http.StatusBadRequest, parseErr); err != nil {
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
