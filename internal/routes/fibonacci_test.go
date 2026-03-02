package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	fiboConfig "github.com/majabojarska/fibo/internal/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func wantHeadersFiboResp() http.Header {
	return http.Header{
		"Cache-Control":     []string{"no-cache"},
		"Connection":        []string{"keep-alive"},
		"Content-Type":      []string{"text/event-stream"},
		"Transfer-Encoding": []string{"chunked"},
		"X-Accel-Buffering": []string{"no"},
	}
}

func headersFiboReq() http.Header {
	return http.Header{
		"Accept": []string{"text/event-stream"},
	}
}

func TestGetFibonacci(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Error(err)
	}
	config, err := fiboConfig.LoadConfig()
	if err != nil {
		t.Error(err)
	}
	router, err := SetupRouter(logger, config)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name           string
		method         string
		url            string
		clientHeaders  http.Header
		wantStatusCode int
		wantBody       string
		wantHeaders    http.Header
	}{
		{
			name:           "GET fibonacci valid sequence request (empty)",
			method:         http.MethodGet,
			url:            "/api/v1/fibonacci/0/stream",
			clientHeaders:  headersFiboReq(),
			wantStatusCode: http.StatusOK,
			wantBody:       "",
			wantHeaders:    wantHeadersFiboResp(),
		},
		{
			name:           "GET fibonacci valid sequence request (1 item)",
			method:         http.MethodGet,
			url:            "/api/v1/fibonacci/1/stream",
			clientHeaders:  headersFiboReq(),
			wantStatusCode: http.StatusOK,
			wantBody:       "{\"id\":0,\"event\":\"fibonacci\",\"data\":{\"ordinal\":1,\"value\":\"0\"}}\n\n",
			wantHeaders:    wantHeadersFiboResp(),
		},
		{
			name:           "GET fibonacci valid sequence request (1 item)",
			method:         http.MethodGet,
			url:            "/api/v1/fibonacci/2/stream",
			clientHeaders:  headersFiboReq(),
			wantStatusCode: http.StatusOK,
			wantBody: "{\"id\":0,\"event\":\"fibonacci\",\"data\":{\"ordinal\":1,\"value\":\"0\"}}\n\n" +
				"{\"id\":1,\"event\":\"fibonacci\",\"data\":{\"ordinal\":2,\"value\":\"1\"}}\n\n",
			wantHeaders: wantHeadersFiboResp(),
		},
		{
			name:           "GET fibonacci without accept text/event-stream header",
			method:         http.MethodGet,
			url:            "/api/v1/fibonacci/5/stream",
			wantStatusCode: http.StatusUnsupportedMediaType,
			wantBody:       "User agent must accept content type 'text/event-stream'.\n",
			wantHeaders:    http.Header{"Content-Type": []string{"text/plain; charset=utf-8"}},
		},
		{
			name:           "GET fibonacci invalid seq size (-1)",
			method:         http.MethodGet,
			url:            "/api/v1/fibonacci/-1/stream",
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "Path parameter 'count' must be a non-negative integer.\n",
			wantHeaders:    http.Header{"Content-Type": []string{"text/plain; charset=utf-8"}},
		},
		{
			name:           "GET fibonacci invalid seq size (-2)",
			method:         http.MethodGet,
			url:            "/api/v1/fibonacci/-2/stream",
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "Path parameter 'count' must be a non-negative integer.\n",
			wantHeaders:    http.Header{"Content-Type": []string{"text/plain; charset=utf-8"}},
		},
		{
			name:           "GET fibonacci invalid seq size (non-numeric)",
			method:         http.MethodGet,
			url:            "/api/v1/fibonacci/BOOM/stream",
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "Path parameter 'count' must be a non-negative integer.\n",
			wantHeaders:    http.Header{"Content-Type": []string{"text/plain; charset=utf-8"}},
		},
		{
			name:           "POST fibonacci with valid path param",
			method:         http.MethodPost,
			url:            "/api/v1/fibonacci/5/stream",
			wantStatusCode: http.StatusMethodNotAllowed,
			wantBody:       "Method not allowed.\n",
			wantHeaders:    http.Header{"Allow": []string{"GET"}, "Content-Type": []string{"text/plain; charset=utf-8"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request

			req, _ = http.NewRequest(tt.method, tt.url, nil)
			for hName, hValues := range tt.clientHeaders {
				for _, hVal := range hValues {
					req.Header.Set(hName, hVal)
				}
			}

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.wantStatusCode, recorder.Code)
			assert.Equal(t, tt.wantHeaders, recorder.Result().Header)
			assert.Equal(t, tt.wantBody, recorder.Body.String())
		})
	}
}
