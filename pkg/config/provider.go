package config

import (
	"time"

	"github.com/spf13/viper"
)

var _ Provider = (*appConfig)(nil)

type Provider interface {
	GetString(key string) string
	GetInt(key string) int
	GetInt64(key string) int64
	GetFloat64(key string) float64
	GetBool(key string) bool
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	GetStringMap(key string) map[string]any
	GetStringSlice(key string) []string
	Get(key string) any
	Set(Key string, value any)
	IsSet(key string) bool
	UnmarshalKey(key string, rawVal any) error
}

type appConfig struct {
	path   string
	name   string
	suffix string
	cfg    *viper.Viper
}

// UnmarshalKey implements Provider.
func (ac *appConfig) UnmarshalKey(key string, rawVal any) error {
	return ac.cfg.UnmarshalKey(key, rawVal)
}

// Get implements Provider.
func (ac *appConfig) Get(key string) any {
	return ac.cfg.Get(key)
}

// GetBool implements Provider.
func (ac *appConfig) GetBool(key string) bool {
	return ac.cfg.GetBool(key)
}

// GetDuration implements Provider.
func (ac *appConfig) GetDuration(key string) time.Duration {
	return ac.cfg.GetDuration(key)
}

// GetFloat64 implements Provider.
func (ac *appConfig) GetFloat64(key string) float64 {
	return ac.cfg.GetFloat64(key)
}

// GetInt implements Provider.
func (ac *appConfig) GetInt(key string) int {
	return ac.cfg.GetInt(key)
}

// GetInt64 implements Provider.
func (ac *appConfig) GetInt64(key string) int64 {
	return ac.cfg.GetInt64(key)
}

// GetString implements Provider.
func (ac *appConfig) GetString(key string) string {
	return ac.cfg.GetString(key)
}

// GetStringMap implements Provider.
func (ac *appConfig) GetStringMap(key string) map[string]any {
	return ac.cfg.GetStringMap(key)
}

// GetStringSlice implements Provider.
func (ac *appConfig) GetStringSlice(key string) []string {
	return ac.cfg.GetStringSlice(key)
}

// GetTime implements Provider.
func (ac *appConfig) GetTime(key string) time.Time {
	return ac.cfg.GetTime(key)
}

// IsSet implements Provider.
func (ac *appConfig) IsSet(key string) bool {
	return ac.cfg.IsSet(key)
}

// Set implements Provider.
func (ac *appConfig) Set(Key string, value any) {
	ac.cfg.Set(Key, value)
}
