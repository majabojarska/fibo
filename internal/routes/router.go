package routes

import (
	"time"

	ginZap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/majabojarska/fibo/internal/middleware"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func setupMiddlewares(router *gin.Engine, logger *zap.Logger) {
	if viper.GetBool("metrics.enabled") {
		// Must happen before API route setup
		middleware.SetupPromMiddleware(router, viper.GetString("metrics.addr"), viper.GetString("metrics.path"))
	}

	// Zap integration
	router.Use(ginZap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginZap.RecoveryWithZap(logger, true))
}

func setupRoutes(router *gin.Engine) {
	groupV1 := router.Group("/api/v1")
	{
		groupFibonacci := groupV1.Group("/fibonacci")
		{
			groupFibonacci.GET(":count", GetFibonacci)
		}
	}

	router.GET("/readyz", GetReadyz)
	router.GET("/livez", GetLivez)

	if viper.GetBool("docs.enabled") {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func SetupRouter(logger *zap.Logger) *gin.Engine {
	router := gin.Default()

	setupMiddlewares(router, logger)
	setupRoutes(router)

	return router
}
