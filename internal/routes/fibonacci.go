package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/majabojarska/fibo/internal/fibonacci"
)

const fiboEvent = "fibonacci"

type GetFibonacciPathParams struct {
	Count int `uri:"count" binding:"min=0"`
}

type FiboEvent struct {
	Id    int      `json:"id"`
	Event string   `json:"event,omitempty"`
	Error string   `json:"error,omitempty"`
	Data  FiboData `json:"data,omitempty"`
}

type FiboData struct {
	Ordinal int    `json:"ordinal"`
	Value   string `json:"value"`
}

// GetFibonacci godoc
//
//	@Summary		Get Fibonacci sequence
//	@Description	Generates a finite length Fibonacci sequence
//	@ID				get-fibonacci
//	@Tags			fibonacci
//	@Param			count	path	GetFibonacciPathParams	true	"Desired sequence size"
//	@Produce		event-stream
//	@Router			/api/v1/fibonacci/{count} [get]
func GetFibonacci(ctx *gin.Context) {
	writer := ctx.Writer
	header := writer.Header()
	header.Set("Content-Type", "text/event-stream")
	header.Set("Connection", "keep-alive")
	header.Set("Cache-Control", "no-cache")
	header.Set("Transfer-Encoding", "chunked")
	header.Set("X-Accel-Buffering", "no")

	var pathParams GetFibonacciPathParams
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		// Will write status code header
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	if err := writeFibo(writer, pathParams.Count); err != nil {
		panic(err)
	}
}

func writeFibo(writer gin.ResponseWriter, wantCount int) error {
	for iterIdx, fiboVal := range fibonacci.FibonacciSeq2(wantCount) {
		event := FiboEvent{
			Id:    iterIdx,
			Event: fiboEvent,
			Data: FiboData{
				Ordinal: iterIdx + 1, // Start ordinals from 1 in the responses
				Value:   fiboVal.String(),
			},
		}

		jsonLine, err := json.Marshal(event)
		if err != nil {
			return err
		}

		jsonLine = append(jsonLine, []byte("\n\n")...) // Event terminator
		if _, err := writer.Write(jsonLine); err != nil {
			return err
		}

		if flusher := writer.(http.Flusher); flusher != nil {
			flusher.Flush()
		}
	}

	return nil
}

// // writeFibo writes the Fibonacci sequence into the provided writer in a JSON array format.
// // Writes in chunks, flushing after brackets and each element (and delimiter).
// func writeFiboOld(writer gin.ResponseWriter, wantCount int) error {
// 	if _, err := writer.WriteString("["); err != nil {
// 		return err
// 	}
//
// 	sentCount := 0
// 	for fiboVal := range fibonacci.FibonacciSeq(wantCount) {
// 		// Sequence item
//
// 		var buf bytes.Buffer
//
// 		buf.WriteRune('"')
// 		buf.WriteString(fiboVal.String())
// 		buf.WriteRune('"')
//
// 		if _, err := writer.WriteString(buf.String()); err != nil {
// 			return err
// 		}
//
// 		sentCount++
//
// 		// Omit delimiter after the last seq item
// 		if sentCount < wantCount {
// 			if _, err := writer.WriteString(","); err != nil {
// 				return err
// 			}
// 		}
//
// 		writer.(http.Flusher).Flush()
// 	}
//
// 	if _, err := writer.WriteString("]"); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
