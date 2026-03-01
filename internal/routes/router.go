package routes

import (
	"time"

	ginZap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/majabojarska/fibo/docs" // Swaggo requires this to be imported
	fiboConfig "github.com/majabojarska/fibo/internal/config"
	"github.com/majabojarska/fibo/internal/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func setupMiddlewares(router *gin.Engine, logger *zap.Logger, config *fiboConfig.Config) {
	if config.Metrics.Enabled {
		// Must happen before API route setup
		logger.Info("Metrics are enabled, setting up Prometheus middleware.")
		middleware.SetupPromMiddleware(router, config.Metrics.Addr, config.Metrics.Path)
	}

	// Zap integration
	router.Use(ginZap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginZap.RecoveryWithZap(logger, true))
}

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

func setupRoutes(router *gin.Engine, logger *zap.Logger, config *fiboConfig.Config) {
	groupV1 := router.Group("/api/v1")
	{
		groupFibonacci := groupV1.Group("/fibonacci")
		{
			groupFibonacci.GET(":count", GetFibonacci)
		}
	}

	router.GET("/readyz", GetReadyz)
	router.GET("/livez", GetLivez)

	if config.Docs.Enabled {
		logger.Info("Docs are enabled, registering Swagger route.")
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func SetupRouter(logger *zap.Logger, config *fiboConfig.Config) *gin.Engine {
	router := gin.Default()

	setupMiddlewares(router, logger, config)
	setupRoutes(router, logger, config)

	return router
}
