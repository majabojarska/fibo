package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetFibonacci(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	tests := []struct {
		name           string
		method         string
		url            string
		wantStatusCode int
		wantContains   string
		wantHeaders    gin.H
	}{
		{
			name:           "Valid fibonacci sequence request (empty)",
			method:         http.MethodGet,
			url:            "/api/v1/fibonacci/0",
			wantStatusCode: http.StatusOK,
			wantContains:   "[]",
			wantHeaders: gin.H{
				http.CanonicalHeaderKey("Content-Type"):      "application/json",
				http.CanonicalHeaderKey("Transfer-Encoding"): "chunked",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request

			req, _ = http.NewRequest(tt.method, tt.url, nil)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.wantStatusCode, recorder.Code)
			assert.Contains(t, recorder.Body.String(), tt.wantContains)
		})
	}
}
