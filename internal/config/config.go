package config

import (
	"fmt"
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

	fmt.Printf("metrics.addr: %s", viper.GetString("metrics.addr"))
	viper.AutomaticEnv()
	viper.SetEnvPrefix("FIBO")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
