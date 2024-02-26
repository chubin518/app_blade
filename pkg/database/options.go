package database

import (
	"app_blade/pkg/config"
	"errors"
	"time"
)

type Option func(*Config)

func WithPrimary(primary string) Option {
	return func(cfg *Config) {
		cfg.Primary = primary
	}
}

func WithMaxIdleConns(n int) Option {
	return func(cfg *Config) {
		cfg.MaxIdleConns = n
	}
}

func WithMaxOpenConns(n int) Option {
	return func(cfg *Config) {
		cfg.MaxOpenConns = n
	}
}

func WithMaxIdleTime(d time.Duration) Option {
	return func(cfg *Config) {
		cfg.MaxIdleTime = d
	}
}

func WithMaxLifeTime(d time.Duration) Option {
	return func(cfg *Config) {
		cfg.MaxLifeTime = d
	}
}

func WithDataSource(ds map[string]DataSource) Option {
	return func(cfg *Config) {
		cfg.DataSource = ds
	}
}

func NewOptions(conf config.Provider) ([]Option, error) {
	cfg := DefaultConfig()
	if err := conf.UnmarshalKey("database", cfg); err != nil {
		return nil, err
	}

	if len(cfg.DataSource) == 0 {
		return nil, errors.New("database config must have at least one data source")
	}

	return []Option{
		WithPrimary(cfg.Primary),
		WithMaxIdleConns(cfg.MaxIdleConns),
		WithMaxOpenConns(cfg.MaxOpenConns),
		WithMaxLifeTime(cfg.MaxLifeTime),
		WithMaxIdleTime(cfg.MaxIdleTime),
		WithDataSource(cfg.DataSource),
	}, nil
}
