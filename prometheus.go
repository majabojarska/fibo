package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func promUrlRelabel(ctx *gin.Context) string {
	url := ctx.Request.URL.Path
	for _, param := range ctx.Params {
		if param.Key == "count" {
			url = strings.Replace(url, param.Value, ":count", 1)
			break
		}
	}
	return url
}

func SetupPromMiddleware(router *gin.Engine, path string) {
	prom := ginprometheus.NewWithConfig(ginprometheus.Config{
		Subsystem: "gin",
	})
	prom.ReqCntURLLabelMappingFn = promUrlRelabel
	prom.Use(router)
}
