package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Api     ApiConfig
	Docs    DocsConfig
	Logging LoggingConfig
	Metrics MetricsConfig
	Pprof   PprofConfig
}

type ApiConfig struct {
	Addr string
}

type DocsConfig struct {
	Enabled bool
}

type LoggingConfig struct {
	// https://pkg.go.dev/go.uber.org/zap#AtomicLevel.UnmarshalText
	Level string
}

type MetricsConfig struct {
	Enabled bool
	Addr    string
	Path    string
}

type PprofConfig struct {
	Enabled bool
}

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
