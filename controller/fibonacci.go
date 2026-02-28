package controller

import (
	"bytes"
	"net/http"

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
//	@Param			count	path	GetFibonacciPathParams	true	"Desired sequence size"
//	@Produce		json
//	@Success		200	{array}		string "Fibonacci sequence items"
//	@Failure		400	{object}	object
//	@Failure		500	{object}	object
//	@Router			/api/v1/fibonacci/{count} [get]
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
		panic(err)
	}
}

// writeFibo writes the Fibonacci sequence into the provided writer in a JSON array format.
// Writes in chunks, flushing after brackets and each element (and delimiter).
func writeFibo(writer gin.ResponseWriter, wantCount int) error {
	if _, err := writer.WriteString("["); err != nil {
		return err
	}

	sentCount := 0
	for fiboVal := range fibonacci.Fibonacci(wantCount) {
		// Sequence item

		var buf bytes.Buffer

		buf.WriteRune('"')
		buf.WriteString(fiboVal.String())
		buf.WriteRune('"')

		if _, err := writer.WriteString(buf.String()); err != nil {
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
