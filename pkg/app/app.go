package app

import (
	"app_blade/pkg/config"
	"app_blade/pkg/logging"
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/wire"
)

type App struct {
	*Config
	conf   config.Provider
	logger logging.Provider

	runners []AppRunner

	exitChan chan struct{}
}

// Run runs the application.
func (app *App) Run() error {
	startCtx, cancel := context.WithTimeout(context.Background(), app.StartTimeout)
	defer cancel()
	if err := app.start(startCtx); err != nil {
		return err
	}
	app.logger.Infof("application started")

	<-app.exitChan

	stopCtx, cancel := context.WithTimeout(context.Background(), app.ShutdownTimeout)
	defer cancel()
	if err := app.stop(stopCtx); err != nil {
		return err
	}
	app.logger.Infof("application stopped")
	return nil
}

// start starts the application.
func (app *App) start(ctx context.Context) error {
	// 响应控制台 Ctrl+C 或 kill 命令。
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

		sig := <-ch

		// 关闭执行器
		app.logger.Infof("application will exit signal: %v", sig)

		select {
		case <-app.exitChan:
			// chan 已关闭，无需再次关闭
		default:
			close(app.exitChan)
		}
	}()

	return withTimeout(ctx, func(ctx context.Context) error {
		for _, runner := range app.runners {
			if err := ctx.Err(); err != nil {
				return err
			}
			if err := runner.OnStart(ctx); err != nil {
				return err
			}
		}
		return nil
	})
}

// stop stops the application.
func (app *App) stop(ctx context.Context) error {
	return withTimeout(ctx, func(ctx context.Context) error {
		var errs []error
		for num := len(app.runners); num > 0; num-- {
			if err := ctx.Err(); err != nil {
				return err
			}
			runner := app.runners[num-1]
			if err := runner.OnShutdown(ctx); err != nil {
				// For best-effort cleanup, keep going after errors.
				errs = append(errs, err)
			}
		}
		return errors.Join(errs...)
	})
}

func New(options ...Option) *App {
	app := &App{
		conf:     config.Default(),
		logger:   logging.Default(),
		Config:   DefaultConfig(),
		runners:  make([]AppRunner, 0),
		exitChan: make(chan struct{}),
	}

	for _, apply := range options {
		apply(app)
	}

	return app
}

var ProviderSet = wire.NewSet(NewOptions, New)
