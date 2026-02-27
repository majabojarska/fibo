package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	router := setupRouter()
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/api/v1/fibonacci/0", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
