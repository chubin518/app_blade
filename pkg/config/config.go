package config

import (
	"app_blade/pkg/profile"
	"os"
	"sync/atomic"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

var (
	ProviderSet = wire.NewSet(New)
	instance    atomic.Value
)

func init() {
	p, err := New()
	if err != nil {
		panic(err)
	}
	SetDefault(p)
}

func SetDefault(p Provider) {
	instance.Store(p)
}

func Default() Provider {
	data := instance.Load()
	if data != nil {
		return data.(Provider)
	}
	return nil
}

func New(options ...Option) (Provider, error) {
	cfg := &appConfig{
		path:   "config",
		name:   profile.Get(),
		suffix: "yaml",
	}

	for _, apply := range options {
		apply(cfg)
	}

	if _, err := os.Stat(cfg.path); os.IsNotExist(err) {
		os.MkdirAll(cfg.path, os.ModePerm)
	}

	v := viper.New()
	v.AutomaticEnv()

	v.AddConfigPath(cfg.path)
	v.SetConfigName(cfg.name)
	v.SetConfigType(cfg.suffix)

	err := v.MergeInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			v.SetDefault("logging.level", "info")
			if err := v.SafeWriteConfig(); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	cfg.cfg = v

	return cfg, nil
}
