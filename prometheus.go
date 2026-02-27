package main

import (
	"regexp"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func promUrlRelabel(ctx *gin.Context) string {
	url := ctx.Request.URL.Path

	// Hacky, but does the job for now.
	// The goal is to reduce metrics cardinality.
	re := regexp.MustCompile(`^/api/v1/fibonacci/(\d+)$`)
	url = re.ReplaceAllString(url, "/api/v1/fibonacci/:count")

	return url
}

func SetupPromMiddleware(router *gin.Engine, path string) {
	prom := ginprometheus.NewWithConfig(ginprometheus.Config{
		Subsystem: "gin",
	})
	prom.ReqCntURLLabelMappingFn = promUrlRelabel
	prom.Use(router)
}
