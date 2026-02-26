package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/majabojarska/fibo/controller"

	_ "github.com/majabojarska/fibo/docs" // Swaggo requires this to be imported
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const httpPortDefault = 8080

func main() {
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
