package main

import (
	"fmt"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	fiboConfig "github.com/majabojarska/fibo/internal/config"
	"github.com/majabojarska/fibo/internal/routes"
	"go.uber.org/zap"
)

func main() {
	config, err := fiboConfig.LoadConfig()
	if err != nil {
		panic(err)
	}
	fmt.Printf("config: %+v", config)

	var logger *zap.Logger
	if config.Debug.Enabled {
		gin.SetMode(gin.DebugMode)
		logger = zap.Must(zap.NewDevelopment())
	} else {
		gin.SetMode(gin.ReleaseMode)
		logger = zap.Must(zap.NewProduction())
	}
	defer logger.Sync() // nolint:errcheck

	router := routes.SetupRouter(logger, config)

	if config.Debug.Enabled {
		pprof.Register(router)
	}

	err = router.Run(config.Api.Addr)
	if err != nil {
		panic(err)
	}
}
