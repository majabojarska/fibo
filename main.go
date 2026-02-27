package main

import (
	"github.com/gin-gonic/gin"
	"github.com/majabojarska/fibo/controller"

	_ "github.com/majabojarska/fibo/docs" // Swaggo requires this to be imported
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	apiListenAddrDefault = ":8080"
	metricsPathDefault   = "/metrics"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	SetupPromMiddleware(router, metricsPathDefault) // Must be before route setup

	ctrl := controller.NewController()

	groupV1 := router.Group("/api/v1")
	{
		groupFibonacci := groupV1.Group("/fibonacci")
		{
			groupFibonacci.GET(":count", ctrl.GetFibonacci)
		}
	}

	router.GET("/readyz", ctrl.GetReadyz)
	router.GET("/livez", ctrl.GetLivez)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func main() {
	router := setupRouter()

	err := router.Run(apiListenAddrDefault)
	if err != nil {
		panic(err)
	}
}
