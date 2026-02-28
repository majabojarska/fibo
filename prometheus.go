package main

import (
	"regexp"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

// promUrlRelabel rewrites the request URLs in order to reduce metrics cardinality
// Path and query parameters are replaced with the respective param names.
func promUrlRelabel(ctx *gin.Context) string {
	url := ctx.Request.URL.Path

	// Hacky, but does the job for now.
	// The goal is to reduce metrics cardinality.
	re := regexp.MustCompile(`^/api/v1/fibonacci/(\d+)$`)
	url = re.ReplaceAllString(url, "/api/v1/fibonacci/:count")

	return url
}

// SetupPromMiddleware sets up a Gin middleware that generates Prometheus metrics
// from each request. Must be called before the monitored routes are registered.
func SetupPromMiddleware(router *gin.Engine, address string, path string) {
	prom := ginprometheus.NewWithConfig(ginprometheus.Config{
		Subsystem:          "gin",
		DisableBodyReading: true,
	})
	prom.ReqCntURLLabelMappingFn = promUrlRelabel

	prom.MetricsPath = path
	prom.SetListenAddress(address)
	prom.Use(router)
}
