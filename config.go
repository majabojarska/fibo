package main

type Config struct {
	Api struct {
		Addr string `mapstructure:"addr"`
		Port int    `mapstucture:"port"`
	} `mapstructure:"api"`

	Metrics struct {
		Enable bool   `mapstructure:"enable"`
		Addr   string `mapstructure:"addr"`
		Port   int    `mapstucture:"port"`
	} `mapstructure:"metrics"`

	Docs struct {
		Enable bool   `mapstructure:"enable"`
		Addr   string `mapstructure:"addr"`
		Port   int    `mapstucture:"port"`
	} `mapstructure:"docs"`

	Debug bool `mapstructure:"debug"`
}
