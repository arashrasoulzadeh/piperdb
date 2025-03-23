package config

import (
	"github.com/BurntSushi/toml"
)

var instance *Config

func NewConfig(flags ...func(*Config)) Config {
	config := Config{}

	for _, f := range flags {
		f(&config)
	}

	instance = &config

	return config
}

func Current() *Config {
	if instance == nil {
		instance = &Config{}
	}
	return instance
}

func WithInstanceName(name string) func(*Config) {
	return func(c *Config) {
		c.Instance.Name = name
	}
}

func WithInstancePath(path string) func(*Config) {
	return func(c *Config) {
		c.Instance.Path = path
	}
}

func WithServerHost(host string) func(*Config) {
	return func(c *Config) {
		c.Server.Host = host
	}
}

func WithServerPort(port string) func(*Config) {
	return func(c *Config) {
		c.Server.Port = port
	}
}

func WithConfigFile(filePath string) func(*Config) {
	return func(c *Config) {
		if _, err := toml.DecodeFile(filePath, &c); err != nil {
			panic("cannot load config file: " + err.Error())
		}
	}
}
