package database

import (
	"app_blade/pkg/logging"
	"context"
	"errors"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	_              logger.Interface = (*Logger)(nil)
	gormPackage                     = filepath.Join("gorm.io", "gorm")
	zapgormPackage                  = filepath.Join("app_blade/pkg", "database")
	SlowThreshold                   = 200 * time.Millisecond
)

type Logger struct {
	logger logging.Provider
}

// Error implements logger.Interface.
func (l *Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.Logging(ctx).Errorf(msg, data...)
}

// Info implements logger.Interface.
func (l *Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.Logging(ctx).Infof(msg, data...)
}

// LogMode implements logger.Interface.
func (l *Logger) LogMode(logger.LogLevel) logger.Interface {
	return l
}

// Trace implements logger.Interface.
func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	// SQL执行耗时
	elapsed := time.Since(begin)
	// 获取执行SQL及受影响行数
	sql, rows := fc()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			l.Logging(ctx).Warnf("sql: %s, rowsAffected: %d, elapsed: %s", sql, rows, elapsed)
		} else {
			l.Logging(ctx).Errorf("sql: %s, rowsAffected: %d, elapsed: %s, error: %v", sql, rows, elapsed, err)
		}
		return
	}
	if elapsed > SlowThreshold {
		l.Logging(ctx).Warnf("sql: %s, rowsAffected: %d, elapsed: %s", sql, rows, elapsed)
	} else {
		l.Logging(ctx).Infof("sql: %s, rowsAffected: %d, elapsed: %s", sql, rows, elapsed)
	}
}

// Warn implements logger.Interface.
func (l *Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.Logging(ctx).Warnf(msg, data...)
}

func (l *Logger) Logging(ctx context.Context) logging.Provider {
	log := l.logger.FromContext(ctx)
	for i := 3; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, zapgormPackage):
		default:
			return log.AddCallerSkip(i)
		}
	}
	return log
}

func NewLogger(logger logging.Provider) *Logger {
	return &Logger{logger: logger}
}
