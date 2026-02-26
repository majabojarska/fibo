package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/majabojarska/fibo/controller"
	fiboDocs "github.com/majabojarska/fibo/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const httpPortDefault = 8080

func main() {
	fiboDocs.SwaggerInfo.Title = "Fibonacci API"
	fiboDocs.SwaggerInfo.Description = "This is a Fibonacci API"
	fiboDocs.SwaggerInfo.Version = "1.0"
	fiboDocs.SwaggerInfo.BasePath = "/api/v1"
	fiboDocs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()

	ctrl := controller.NewController()

	groupV1 := router.Group("/api/v1")
	{
		groupFibonacci := groupV1.Group("")
		{
			groupFibonacci.GET("fibonacci/:count", ctrl.GetFibonacci)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(fmt.Sprintf(":%d", httpPortDefault))
	if err != nil {
		panic(err)
	}
}
