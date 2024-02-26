package logging

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

var CtxKey = struct{}{}

var _ Provider = (*logger)(nil)

type Provider interface {
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
	Panicf(format string, args ...any)
	DebugfContext(ctx context.Context, format string, args ...any)
	InfofContext(ctx context.Context, format string, args ...any)
	WarnfContext(ctx context.Context, format string, args ...any)
	ErrorfContext(ctx context.Context, format string, args ...any)
	FatalfContext(ctx context.Context, format string, args ...any)
	PanicfContext(ctx context.Context, format string, args ...any)
	Logf(ctx context.Context, level Level, format string, args ...any)
	WithContext(ctx context.Context, attrs map[string]any) context.Context
	FromContext(ctx context.Context) Provider
	AddCallerSkip(skip int) Provider
	Flush()
}

type logger struct {
	*Config
	fields []zap.Field
	*zap.Logger
}

// AddCallerSkip implements Provider.
func (l *logger) AddCallerSkip(skip int) Provider {
	return &logger{
		Config: l.Config,
		Logger: l.Logger.WithOptions(zap.AddCallerSkip(skip)),
		fields: l.fields,
	}
}

// Debugf implements Provider.
func (l *logger) Debugf(format string, args ...any) {
	l.Logf(context.Background(), DebugLevel, format, args...)
}

// DebugfContext implements Provider.
func (l *logger) DebugfContext(ctx context.Context, format string, args ...any) {
	l.Logf(ctx, DebugLevel, format, args...)
}

// Errorf implements Provider.
func (l *logger) Errorf(format string, args ...any) {
	l.Logf(context.Background(), ErrorLevel, format, args...)
}

// ErrorfContext implements Provider.
func (l *logger) ErrorfContext(ctx context.Context, format string, args ...any) {
	l.Logf(ctx, ErrorLevel, format, args...)
}

// Fatalf implements Provider.
func (l *logger) Fatalf(format string, args ...any) {
	l.Logf(context.Background(), FatalLevel, format, args...)
}

// FatalfContext implements Provider.
func (l *logger) FatalfContext(ctx context.Context, format string, args ...any) {
	l.Logf(ctx, FatalLevel, format, args...)
}

// Infof implements Provider.
func (l *logger) Infof(format string, args ...any) {
	l.Logf(context.Background(), InfoLevel, format, args...)
}

// InfofContext implements Provider.
func (l *logger) InfofContext(ctx context.Context, format string, args ...any) {
	l.Logf(ctx, InfoLevel, format, args...)
}

// Panicf implements Provider.
func (l *logger) Panicf(format string, args ...any) {
	l.Logf(context.Background(), PanicLevel, format, args...)
}

// PanicfContext implements Provider.
func (l *logger) PanicfContext(ctx context.Context, format string, args ...any) {
	l.Logf(ctx, PanicLevel, format, args...)
}

// Warnf implements Provider.
func (l *logger) Warnf(format string, args ...any) {
	l.Logf(context.Background(), WarnLevel, format, args...)
}

// WarnfContext implements Provider.
func (l *logger) WarnfContext(ctx context.Context, format string, args ...any) {
	l.Logf(ctx, WarnLevel, format, args...)
}

// Logf implements Provider.
func (l *logger) Logf(ctx context.Context, level Level, format string, args ...any) {
	log := l.FromContext(ctx).(*logger)
	log.Log(level.toZapLevel(), fmt.Sprintf(format, args...), log.fields...)
}

// WithContext implements Provider.
func (l *logger) WithContext(ctx context.Context, attrs map[string]any) context.Context {
	fields := make([]zap.Field, 0, len(attrs))
	for k, v := range attrs {
		fields = append(fields, zap.Any(k, v))
	}
	return context.WithValue(ctx, CtxKey, &logger{
		fields: fields,
		Logger: l.Logger,
		Config: l.Config,
	})
}

// FromContext implements Provider.
func (l *logger) FromContext(ctx context.Context) Provider {
	if ctx == nil {
		return l
	}
	if val, ok := ctx.Value(CtxKey).(*logger); ok {
		return val
	}
	return l
}

// Flush implements Provider.
func (l *logger) Flush() {
	l.Sync()
}
