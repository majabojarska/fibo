package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/majabojarska/fibo/controller"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const HTTP_PORT_DEFAULT = 8080

type GetFibonacciRequest struct {
	Count int `json:"count"`
}

//	@title			Fibonacci API
//	@version		1.0
//	@description	This is a streaming Fibonacci API

//	@contact.name	Maja Bojarska
//	@contact.url	https://github.com/majabojarska/fibo/issues/new

// @license.name	MIT License
// @license.url	https://github.com/majabojarska/fibo?tab=MIT-1-ov-file#readme

//	@host		localhost:8080
//	@BasePath	/api/v1

func main() {
	router := gin.Default()

	ctrl := controller.NewController()

	groupV1 := router.Group("/api/v1")
	{
		groupFibonacci := groupV1.Group("")
		{
			groupFibonacci.GET("fibonacci", func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "Hello there! - Obi-Wan Kenobi")
			})
		}
	}
	fmt.Printf("%+v", ctrl)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
