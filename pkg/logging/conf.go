package logging

type Config struct {
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Stdout     bool   `mapstructure:"stdout"`
	Level      string `mapstructure:"level"`
	Path       string `mapstructure:"path"`
}

func DefaultConfig() *Config {
	return &Config{
		MaxSize:    100,
		MaxBackups: 32,
		MaxAge:     14,
		Stdout:     true,
		Level:      "info",
		Path:       "logs/app.log",
	}
}
