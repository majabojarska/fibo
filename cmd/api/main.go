package main

import (
	"github.com/gin-contrib/pprof"
	config "github.com/majabojarska/fibo/internal/config"
	"github.com/majabojarska/fibo/internal/routes"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	logger := zap.New(zapcore.NewCore())

	if isDebug {
		logger = zap.Must(zap.NewDevelopment())
	} else {
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
