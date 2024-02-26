package app

import (
	"app_blade/pkg/logging"
	"context"
)

// InitializeRunner 初始化 AppRunner
type InitializeRunner func() []AppRunner

type AppRunner interface {
	OnStart(ctx context.Context) error
	OnShutdown(ctx context.Context) error
}

var _ AppRunner = (*AppDestroy)(nil)

type AppDestroy struct {
	logger logging.Provider
}

// OnShutdown implements AppRunner.
func (ad *AppDestroy) OnShutdown(ctx context.Context) error {
	ad.logger.Flush()
	return nil
}

// OnStart implements AppRunner.
func (ad *AppDestroy) OnStart(ctx context.Context) error {
	return nil
}
