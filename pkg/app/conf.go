package app

import (
	"app_blade/buildinfo"
	"time"
)

type Config struct {
	Name            string        `mapstructure:"name"`
	StartTimeout    time.Duration `mapstructure:"start_timeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
}

func DefaultConfig() *Config {
	return &Config{
		Name:            buildinfo.Name,
		StartTimeout:    15 * time.Second,
		ShutdownTimeout: 15 * time.Second,
	}
}
