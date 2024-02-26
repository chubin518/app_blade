package web

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
)

type Config struct {
	Host           string        `mapstructure:"host"`
	Port           int           `mapstructure:"port"`
	ReadTimeout    time.Duration `mapstructure:"read_timeout"`
	WriteTimeout   time.Duration `mapstructure:"write_timeout"`
	MaxHeaderBytes int           `mapstructure:"max_header_bytes"`
	EnableDump     bool          `mapstructure:"enable_dump"`
	Cors           *CorsConfig   `mapstructure:"cors"`
	IgnorePatterns []string      `mapstructure:"ignore_patterns"`
}

func (cfg *Config) Addr() string {
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}

type CorsConfig struct {
	AllowMethods     []string      `mapstructure:"allow_methods"`
	AllowHeaders     []string      `mapstructure:"allow_headers"`
	AllowOrigins     []string      `mapstructure:"allow_origins"`
	AllowCredentials bool          `mapstructure:"allow_credentials"`
	MaxAge           time.Duration `mapstructure:"max_age"`
}

func DefaultConfig() *Config {
	corsCfg := cors.DefaultConfig()
	return &Config{
		Host:           "0.0.0.0",
		Port:           8080,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
		EnableDump:     false,
		IgnorePatterns: make([]string, 0),
		Cors: &CorsConfig{
			AllowMethods:     corsCfg.AllowMethods,
			AllowHeaders:     corsCfg.AllowHeaders,
			AllowOrigins:     corsCfg.AllowOrigins,
			AllowCredentials: corsCfg.AllowCredentials,
			MaxAge:           corsCfg.MaxAge,
		},
	}
}
