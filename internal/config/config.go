package config

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	envPrefix  = "FIBO"
	configName = "fibo"
	configType = "yaml"
)

type Config struct {
	Api     ApiConfig
	Docs    DocsConfig
	Logging LoggingConfig
	Metrics MetricsConfig
	Debug   DebugConfig
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

type DebugConfig struct {
	Enabled bool
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("api.addr", ":8080")
	v.SetDefault("docs.enabled", true)
	v.SetDefault("logging.level", "info")
	v.SetDefault("metrics.enabled", true)
	v.SetDefault("metrics.addr", ":9091")
	v.SetDefault("metrics.path", "/metrics")
	v.SetDefault("debug.enabled", false)
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

	setDefaults(v)

	v.AddConfigPath("$HOME/.appname")
	v.AddConfigPath(".")
	v.SetConfigName(configName) // TODO: Make this path configurable with a flag
	v.SetConfigType(configType)

	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

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
