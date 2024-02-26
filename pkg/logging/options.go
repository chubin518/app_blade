package logging

import "app_blade/pkg/config"

type Option func(*logger)

func WithLevel(level string) Option {
	return func(l *logger) {
		l.Config.Level = level
	}
}

func WithPath(path string) Option {
	return func(l *logger) {
		l.Path = path
	}
}

func WithMaxSize(maxSize int) Option {
	return func(l *logger) {
		l.MaxAge = maxSize
	}
}

func WithMaxBackups(maxBackups int) Option {
	return func(l *logger) {
		l.MaxBackups = maxBackups
	}
}

func WithMaxAge(maxAge int) Option {
	return func(l *logger) {
		l.MaxAge = maxAge
	}
}

func WithStdout(stdout bool) Option {
	return func(l *logger) {
		l.Stdout = stdout
	}
}

func NewOptions(conf config.Provider) ([]Option, error) {
	cfg := DefaultConfig()
	if err := conf.UnmarshalKey("logging", cfg); err != nil {
		return nil, err
	}
	return []Option{
		WithLevel(cfg.Level),
		WithMaxAge(cfg.MaxAge),
		WithMaxBackups(cfg.MaxBackups),
		WithMaxSize(cfg.MaxSize),
		WithPath(cfg.Path),
		WithStdout(cfg.Stdout),
	}, nil
}
