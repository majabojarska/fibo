package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	config "github.com/majabojarska/fibo/internal/config"
	"github.com/majabojarska/fibo/internal/routes"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

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
