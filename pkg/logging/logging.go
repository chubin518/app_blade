package logging

import (
	"os"
	"sync/atomic"
	"time"

	"github.com/google/wire"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ProviderSet = wire.NewSet(NewOptions, New)
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
	logger := &logger{
		Config: DefaultConfig(),
		fields: make([]zapcore.Field, 0),
	}

	for _, apply := range options {
		apply(logger)
	}

	wss := make([]zapcore.WriteSyncer, 0)
	if logger.Stdout {
		wss = append(wss, zapcore.AddSync(os.Stdout))
	}

	if !(len(logger.Path) == 0 || logger.Path == " ") {
		wss = append(wss, zapcore.AddSync(&lumberjack.Logger{
			Filename:   logger.Path,
			MaxSize:    logger.MaxSize,
			MaxBackups: logger.MaxBackups,
			MaxAge:     logger.MaxAge,
			Compress:   true,
			LocalTime:  true,
		}))
	}

	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:       "timestamp",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		FunctionKey:   zapcore.OmitKey,
		StacktraceKey: zapcore.OmitKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendString(t.Format("2006-01-02 15:04:05.999"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	zlog := zap.New(
		zapcore.NewCore(
			encoder,
			zapcore.NewMultiWriteSyncer(wss...),
			zap.NewAtomicLevelAt(ParseLevel(logger.Config.Level).toZapLevel()),
		),
		zap.AddCaller(),
		zap.AddCallerSkip(2),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	zap.ReplaceGlobals(zlog)

	logger.Logger = zlog

	return logger, nil
}
