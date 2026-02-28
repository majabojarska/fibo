package main

import (
	"time"

	ginZap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	ctrl "github.com/majabojarska/fibo/controller"
	_ "github.com/majabojarska/fibo/docs" // Swaggo requires this to be imported
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
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

func setupMiddlewares(router *gin.Engine, logger *zap.Logger) {
	SetupPromMiddleware(router, metricsListenAddrDefault, metricsPathDefault) // Must be before route setup

	router.Use(ginZap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginZap.RecoveryWithZap(logger, true))
}

func setupRoutes(router *gin.Engine) {
	ctrl := ctrl.NewController()

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
}

func setupRouter(logger *zap.Logger) *gin.Engine {
	router := gin.Default()

	setupMiddlewares(router, logger)
	setupRoutes(router)

	return router
}

func main() {
	logger, err := zap.NewDevelopment() // TODO: Make this gated by FIBO_DEBUG once viper is in
	if err != nil {
		panic(err)
	}

	router := setupRouter(logger)

	err = router.Run(apiListenAddrDefault)
	if err != nil {
		logger.Sugar().Error(err)
	}
}
