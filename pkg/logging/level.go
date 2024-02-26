package logging

import "go.uber.org/zap/zapcore"

const (
	DebugLevel = Level(1)
	InfoLevel  = Level(2)
	WarnLevel  = Level(3)
	ErrorLevel = Level(4)
	FatalLevel = Level(5)
	PanicLevel = Level(6)
)

// Level 日志输出级别。
type Level uint16

func (lvl Level) String() string {
	switch lvl {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	case PanicLevel:
		return "panic"
	}
	return ""
}

func (lvl Level) toZapLevel() zapcore.Level {
	switch lvl {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case FatalLevel:
		return zapcore.FatalLevel
	case PanicLevel:
		return zapcore.PanicLevel
	}
	return zapcore.InfoLevel
}

func ParseLevel(text string) Level {
	switch text {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	case "panic":
		return PanicLevel
	}
	return InfoLevel
}
