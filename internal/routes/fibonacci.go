package routes

import (
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/majabojarska/fibo/internal/fibonacci"
	"github.com/spf13/viper"
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
//	@Failure		400	{string}	string	"Bad request"
//	@Failure		405	{string}	string	"Method not allowed"
//	@Failure		415	{string}	string	"Unsupported media type"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/api/v1/fibonacci/{count}/stream [get]
func GetFibonacci(ctx *gin.Context) {
	writer := ctx.Writer
	header := writer.Header()

	var pathParams GetFibonacciPathParams
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		// Will write status code header
		ctx.String(http.StatusBadRequest, "Path parameter 'count' must be a non-negative integer.\n")
		return
	}

	if ctx.Request.Method != http.MethodGet {
		header.Set("Allow", "GET")
		ctx.String(http.StatusMethodNotAllowed, "Method not allowed.\n")
		return
	}

	if !slices.Contains(ctx.Request.Header["Accept"], "text/event-stream") {
		ctx.String(http.StatusUnsupportedMediaType, "User agent must accept content type 'text/event-stream'.\n")
		return
	}

	header.Set("Content-Type", "text/event-stream")
	header.Set("Connection", "keep-alive")
	header.Set("Cache-Control", "no-cache")
	header.Set("Transfer-Encoding", "chunked")
	header.Set("X-Accel-Buffering", "no")

	if err := writeFibo(writer, pathParams.Count); err != nil {
		panic(err)
	}
}

func writeFibo(writer gin.ResponseWriter, wantCount int) error {
	eventDelay := viper.GetDuration("api.event_delay")

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
		if eventDelay > 0 {
			time.Sleep(eventDelay)
		}
	}

	return nil
}
