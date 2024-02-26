package app

import (
	"app_blade/pkg/config"
	"app_blade/pkg/logging"
	"time"
)

type Option func(*App)

func WithConfig(cfg config.Provider) Option {
	return func(app *App) {
		app.conf = cfg
	}
}

func WithLogger(logger logging.Provider) Option {
	return func(app *App) {
		app.logger = logger
	}
}

func WithName(name string) Option {
	return func(app *App) {
		app.Name = name
	}
}

func WithStartTimeout(timeout time.Duration) Option {
	return func(app *App) {
		app.StartTimeout = timeout
	}
}

func WithStopTimeout(timeout time.Duration) Option {
	return func(app *App) {
		app.ShutdownTimeout = timeout
	}
}

func WithRunner(runners ...AppRunner) Option {
	return func(app *App) {
		app.runners = runners
	}
}

func NewOptions(conf config.Provider, logger logging.Provider, initializeRunner InitializeRunner) ([]Option, error) {
	cfg := DefaultConfig()
	if err := conf.UnmarshalKey("server", cfg); err != nil {
		return nil, err
	}

	runners := make([]AppRunner, 0)
	runners = append(runners, &AppDestroy{logger: logger})
	runners = append(runners, initializeRunner()...)

	logging.SetDefault(logger)

	return []Option{
		WithConfig(conf),
		WithLogger(logger),
		WithName(cfg.Name),
		WithStartTimeout(cfg.StartTimeout),
		WithStopTimeout(cfg.ShutdownTimeout),
		WithRunner(runners...),
	}, nil
}
