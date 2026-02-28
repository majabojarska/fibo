package main

import (
	"time"

	ginZap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	config "github.com/majabojarska/fibo/config"
	ctrl "github.com/majabojarska/fibo/controller"
	_ "github.com/majabojarska/fibo/docs" // Swaggo requires this to be imported
	middleware "github.com/majabojarska/fibo/internal/middleware"
	"github.com/spf13/viper"
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
	if viper.GetBool("docs.enabled") {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func setupRouter(logger *zap.Logger) *gin.Engine {
	router := gin.Default()

	setupMiddlewares(router, logger)
	setupRoutes(router)

	return router
}

func main() {
	config.LoadConfig()

	var logger *zap.Logger

	if viper.GetBool("debug") {
		gin.SetMode(gin.DebugMode)
		logger = zap.Must(zap.NewDevelopment())
	} else {
		gin.SetMode(gin.ReleaseMode)
		logger = zap.Must(zap.NewProduction())
	}
	defer logger.Sync() // nolint:errcheck

	router := setupRouter(logger)
	err := router.Run(viper.GetString("api.addr"))
	if err != nil {
		panic(err)
	}
}
