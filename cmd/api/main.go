package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	_ "github.com/majabojarska/fibo/docs" // Swaggo requires this to be imported
	config "github.com/majabojarska/fibo/internal/config"
	"github.com/majabojarska/fibo/internal/routes"
	"github.com/spf13/viper"
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

func main() {
	config.LoadConfig()

	var logger *zap.Logger

	isDebug := viper.GetBool("debug")
	if isDebug {
		gin.SetMode(gin.DebugMode)
		logger = zap.Must(zap.NewDevelopment())
	} else {
		gin.SetMode(gin.ReleaseMode)
		logger = zap.Must(zap.NewProduction())
	}
	defer logger.Sync() // nolint:errcheck

	router := routes.SetupRouter(logger)

	if isDebug {
		pprof.Register(router)
	}

	err := router.Run(viper.GetString("api.addr"))
	if err != nil {
		panic(err)
	}
}
