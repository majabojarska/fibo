package routes

import (
	"net/url"
	"time"

	"github.com/gin-contrib/cors"
	ginZap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerDocs "github.com/majabojarska/fibo/docs" // Swaggo requires this to be imported
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

	// CORS, of course
	router.Use(cors.New(cors.Config{
		AllowOrigins: config.Api.AllowOrigins,
	}))
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

func setupRoutes(router *gin.Engine, logger *zap.Logger, config *fiboConfig.Config) error {
	groupV1 := router.Group("/api/v1")
	{
		groupFibonacci := groupV1.Group("/fibonacci")
		{
			groupFibonacci.Any(":count/stream", GetFibonacci)
		}
	}

	router.GET("/readyz", GetReadyz)
	router.GET("/livez", GetLivez)

	if config.Docs.Enabled {
		logger.Info("Docs are enabled, registering Swagger route.")
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return nil
}

func setupSwagger(config *fiboConfig.Config) error {
	parsedURL, err := url.Parse(config.Api.RootURL)
	if err != nil {
		return err
	}

	swaggerDocs.SwaggerInfo.Host = parsedURL.Host
	swaggerDocs.SwaggerInfo.Schemes = []string{parsedURL.Scheme}

	return nil
}

func SetupRouter(logger *zap.Logger, config *fiboConfig.Config) (*gin.Engine, error) {
	router := gin.Default()

	setupMiddlewares(router, logger, config)

	if err := setupSwagger(config); err != nil {
		return nil, err
	}

	if err := setupRoutes(router, logger, config); err != nil {
		return nil, err
	}

	return router, nil
}
