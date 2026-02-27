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
		wantHeaders    http.Header
	}{
		{
			name:           "GET fibonacci valid sequence request (empty)",
			method:         http.MethodGet,
			url:            "/api/v1/fibonacci/0",
			wantStatusCode: http.StatusOK,
			wantContains:   "[]",
			wantHeaders:    http.Header(http.Header{"Content-Type": []string{"application/json"}, "Transfer-Encoding": []string{"chunked"}}),
		},
		{
			name:           "GET fibonacci valid sequence request (1 item)",
			method:         http.MethodGet,
			url:            "/api/v1/fibonacci/1",
			wantStatusCode: http.StatusOK,
			wantContains:   "[0]",
			wantHeaders:    http.Header(http.Header{"Content-Type": []string{"application/json"}, "Transfer-Encoding": []string{"chunked"}}),
		},
		{
			name:           "GET fibonacci valid sequence request (2 items)",
			method:         http.MethodGet,
			url:            "/api/v1/fibonacci/2",
			wantStatusCode: http.StatusOK,
			wantContains:   "[0,1]",
			wantHeaders:    http.Header(http.Header{"Content-Type": []string{"application/json"}, "Transfer-Encoding": []string{"chunked"}}),
		},
		{
			name:           "GET fibonacci invalid seq size (-1)",
			method:         http.MethodGet,
			url:            "/api/v1/fibonacci/-1",
			wantStatusCode: http.StatusBadRequest,
			wantContains:   "{}",
			wantHeaders:    http.Header(http.Header{"Content-Type": []string{"application/json"}, "Transfer-Encoding": []string{"chunked"}}),
		},
		{
			name:           "GET fibonacci invalid seq size (-2)",
			method:         http.MethodGet,
			url:            "/api/v1/fibonacci/-2",
			wantStatusCode: http.StatusBadRequest,
			wantContains:   "{}",
			wantHeaders:    http.Header(http.Header{"Content-Type": []string{"application/json"}, "Transfer-Encoding": []string{"chunked"}}),
		},
		{
			name:           "GET fibonacci invalid seq size (non-numeric)",
			method:         http.MethodGet,
			url:            "/api/v1/fibonacci/BOOM",
			wantStatusCode: http.StatusBadRequest,
			wantContains:   "{}",
			wantHeaders:    http.Header(http.Header{"Content-Type": []string{"application/json"}, "Transfer-Encoding": []string{"chunked"}}),
		},
		{
			name:           "GET fibonacci without path param",
			method:         http.MethodGet,
			url:            "/api/v1/fibonacci/",
			wantStatusCode: http.StatusNotFound,
			wantContains:   "404 page not found",
			wantHeaders:    http.Header{"Content-Type": []string{"text/plain"}},
		},
		{
			name:           "POST fibonacci with valid path param",
			method:         http.MethodGet,
			url:            "/api/v1/fibonacci",
			wantStatusCode: http.StatusNotFound,
			wantContains:   "404 page not found",
			wantHeaders:    http.Header{"Content-Type": []string{"text/plain"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request

			req, _ = http.NewRequest(tt.method, tt.url, nil)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.wantStatusCode, recorder.Code)
			assert.Equal(t, tt.wantHeaders, recorder.Result().Header)
			assert.Equal(t, tt.wantContains, recorder.Body.String())
		})
	}
}
