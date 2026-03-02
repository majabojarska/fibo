package config

import (
	"strings"
	"time"

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
	Addr         string
	RootURL      string        `mapstructure:"root_url"`
	AllowOrigins []string      `mapstructure:"allow_origins"`
	EventDelay   time.Duration `mapstructure:"event_delay"`
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

func setDefaults() {
	viper.SetDefault("api.addr", ":8080")
	viper.SetDefault("api.root_url", "http://localhost:8080")
	viper.SetDefault("api.allow_origins", []string{"http://localhost"})
	viper.SetDefault("api.event_delay", nil)
	viper.SetDefault("docs.enabled", true)
	viper.SetDefault("logging.level", "info")
	viper.SetDefault("metrics.enabled", true)
	viper.SetDefault("metrics.addr", ":9091")
	viper.SetDefault("metrics.path", "/metrics")
	viper.SetDefault("debug.enabled", false)
}

func parseConfig() (*Config, error) {
	var config Config

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(configName) // TODO: Make this path configurable with a flag
	viper.SetConfigType(configType)

	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	setDefaults()

	// Load config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Does not exist, but we've got defaults for everything so it's fine
			return nil, err
		}
	}

	// Parse
	config, err := parseConfig()
	if err != nil {
		return nil, err
	}
	return config, nil
}
