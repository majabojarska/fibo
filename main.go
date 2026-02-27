package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/majabojarska/fibo/controller"

	_ "github.com/majabojarska/fibo/docs" // Swaggo requires this to be imported
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginProm "github.com/zsais/go-gin-prometheus"
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

const httpPortDefault = 8080

func setupRouter() *gin.Engine {
	router := gin.Default()

	ctrl := controller.NewController()

	groupV1 := router.Group("/api/v1")
	{
		groupFibonacci := groupV1.Group("")
		{
			groupFibonacci.GET("fibonacci/:count", ctrl.GetFibonacci)
		}
	}

	router.GET("/readyz", ctrl.GetReadyz)
	router.GET("/livez", ctrl.GetLivez)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Prometheus
	prom := ginProm.NewWithConfig(ginProm.Config{
		Subsystem: "gin",
	})
	prom.Use(router)

	return router
}

func main() {
	router := setupRouter()

	err := router.Run(fmt.Sprintf(":%d", httpPortDefault))
	if err != nil {
		panic(err)
	}
}
