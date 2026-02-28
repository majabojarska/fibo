package config

import (
	"strings"

	"github.com/spf13/viper"
)

func setDefaults() {
	viper.SetDefault("api.addr", ":8080")
	viper.SetDefault("metrics.enabled", true)
	viper.SetDefault("metrics.addr", ":9090")
	viper.SetDefault("metrics.path", "/metrics")
	viper.SetDefault("docs.enabled", true)
	viper.SetDefault("docs.path", "/swagger/*any")
	viper.SetDefault("debug", false)
}

func LoadConfig() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("FIBO")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	setDefaults()
}
