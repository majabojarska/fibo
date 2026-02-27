package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/majabojarska/fibo/controller"

	_ "github.com/majabojarska/fibo/docs" // Swaggo requires this to be imported
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

//	@title			Fibo
//	@version		1.0
//	@description	Fibonacci sequence REST API

//	@contact.name	Maja Bojarska
//	@contact.url	https://github.com/majabojarska/fibo
//	@contact.email	majabojarska98@gmail.com

//	@license.name	MIT License
//	@license.url	https://github.com/majabojarska/fibo?tab=MIT-1-ov-file#readme

//	@host		localhost:8080
//	@BasePath	/

const (
	apiListenAddrDefault     = ":8080"
	metricsListenAddrDefault = ":9090"
	metricsPathDefault       = "/metrics"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	ctrl := controller.NewController()

	groupV1 := router.Group("/api/v1")
	{
		groupFibonacci := groupV1.Group("fibonacci")
		{
			groupFibonacci.GET(":count", ctrl.GetFibonacci)
		}
	}

	router.GET("/readyz", ctrl.GetReadyz)
	router.GET("/livez", ctrl.GetLivez)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	setupPromMiddleware(router, metricsPathDefault)

	return router
}

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

func setupPromMiddleware(router *gin.Engine, metricsPath string) {
	prom := ginprometheus.NewWithConfig(ginprometheus.Config{
		Subsystem: "gin",
	})
	prom.MetricsPath = metricsPathDefault
	prom.SetListenAddress(":9090")
	prom.ReqCntURLLabelMappingFn = promUrlRelabel
	prom.Use(router)
}

func main() {
	router := setupRouter()

	err := router.Run(apiListenAddrDefault)
	if err != nil {
		panic(err)
	}
}
