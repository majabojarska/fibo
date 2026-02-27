package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/majabojarska/fibo/internal/fibonacci"
)

type GetFibonacciPathParams struct {
	Count int `uri:"count" binding:"min=0"`
}

// GetFibonacci godoc
//
//	@Summary		Get Fibonacci sequence
//	@Description	Generates a finite length Fibonacci sequence
//	@ID				get-fibonacci
//	@Tags			fibonacci
//	@Param			count	path	CountUri	true	"Desired sequence size"
//	@Produce		json
//	@Success		200	{array}		int
//	@Failure		400	{object}	string
//	@Failure		404	{object}	string
//	@Failure		500	{object}	string
//	@Router			/fibonacci/{count} [get]
func (c *Controller) GetFibonacci(ctx *gin.Context) {
	writer := ctx.Writer
	header := writer.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Transfer-Encoding", "chunked")

	var pathParams GetFibonacciPathParams
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	writer.WriteHeader(http.StatusOK)

	if err := writeFibo(writer, pathParams.Count); err != nil {
		panic(err) // TODO: Log this with Zap
	}
}

// writeFibo writes the Fibonacci sequence into the provided writer in a JSON array format.
// Writes in chunks, flushing after brackets and each element (and delimiter).
func writeFibo(writer gin.ResponseWriter, wantCount int) error {
	if _, err := writer.WriteString("["); err != nil {
		return err
	}
	writer.(http.Flusher).Flush()

	sentCount := 0
	for fiboVal := range fibonacci.Fibonacci(wantCount) {
		// Sequence item
		if _, err := writer.WriteString(strconv.Itoa(fiboVal)); err != nil {
			return err
		}

		sentCount++

		// Omit delimiter after the last seq item
		if sentCount < wantCount {
			if _, err := writer.WriteString(","); err != nil {
				return err
			}
		}

		writer.(http.Flusher).Flush()
	}

	if _, err := writer.WriteString("]"); err != nil {
		return err
	}

	return nil
}
