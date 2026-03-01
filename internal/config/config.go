package config

import (
	"strings"

	"github.com/spf13/viper"
)

func setDefaults() {
	viper.SetDefault("api.addr", ":8080")
	viper.SetDefault("docs.enabled", true)
	viper.SetDefault("metrics.enabled", true)
	viper.SetDefault("metrics.addr", ":8081")
	viper.SetDefault("metrics.path", "/metrics")
	viper.SetDefault("debug", false)
}

func LoadConfig() {
	setDefaults()

	viper.AutomaticEnv()
	viper.SetEnvPrefix("FIBO")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
