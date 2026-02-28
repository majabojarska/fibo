package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// TestHealthchecks tests the /readyz and /livez GET endpoints.
func TestHealthchecks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Error(err)
	}
	router := SetupRouter(logger)

	tests := []struct {
		name           string
		url            string
		method         string
		wantStatusCode int
		wantContains   string
	}{
		{
			name:           "GET readyz",
			url:            "/readyz",
			method:         http.MethodGet,
			wantStatusCode: http.StatusOK,
			wantContains:   "",
		},
		{
			name:           "POST readyz",
			url:            "/readyz",
			method:         http.MethodPost,
			wantStatusCode: http.StatusNotFound,
			wantContains:   "404 page not found",
		},
		{
			name:           "GET livez",
			method:         http.MethodGet,
			url:            "/livez",
			wantStatusCode: http.StatusOK,
			wantContains:   "",
		},
		{
			name:           "PUT livez",
			method:         http.MethodPut,
			url:            "/livez",
			wantStatusCode: http.StatusNotFound,
			wantContains:   "404 page not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request

			req, _ = http.NewRequest(tt.method, tt.url, nil)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.wantStatusCode, recorder.Code)
			assert.Equal(t, tt.wantContains, recorder.Body.String())
		})
	}
}
