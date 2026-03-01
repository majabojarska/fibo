package config

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	envPrefix  = "FIBO"
	configName = "fibo.yaml"
	configPath = "."
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

	viper.SetDefault("logging.level", "info")

	viper.SetDefault("metrics.enabled", true)
	viper.SetDefault("metrics.addr", ":8081")
	viper.SetDefault("metrics.path", "/metrics")

	viper.SetDefault("pprof.enabled", false)
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var config Config

	err := v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigName(configName) // TODO: Make this path configurable with a flag
	v.AddConfigPath(configPath)

	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	setDefaults()

	// Load config file
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Does not exist, but we've got defaults for everything so it's fine
			return nil, err
		}
	}

	// Parse
	config, err := parseConfig(v)
	if err != nil {
		return nil, err
	}
	return config, nil
}
