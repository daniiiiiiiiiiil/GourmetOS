package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

var (
	instance *Logger
)

func NewLogger(level string) *Logger {
	if instance != nil {
		return instance
	}

	config := zap.NewProductionConfig()
	config.Level = getLevel(level)

	if os.Getenv("LOG_FORMAT") == "console" {
		config.Encoding = "console"
	}

	logger, err := config.Build()
	if err != nil {
		panic("failed to create logger: " + err.Error())
	}

	instance = &Logger{Logger: logger}
	return instance
}

func GetLogger() *Logger {
	if instance == nil {
		return NewLogger("info")
	}
	return instance
}

func getLevel(level string) zap.AtomicLevel {
	switch level {
	case "debug":
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "info":
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		return zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	default:
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
}

func convertFields(fields ...interface{}) []zap.Field {
	if len(fields)%2 != 0 {
		return []zap.Field{zap.String("error", "invalid number of fields")}
	}

	zapFields := make([]zap.Field, 0, len(fields)/2)
	for i := 0; i < len(fields); i += 2 {
		key, ok := fields[i].(string)
		if !ok {
			continue
		}
		zapFields = append(zapFields, zap.Any(key, fields[i+1]))
	}
	return zapFields
}

func (l *Logger) Info(msg string, fields ...interface{}) {
	l.Logger.Info(msg, convertFields(fields...)...)
}

func (l *Logger) Error(msg string, fields ...interface{}) {
	l.Logger.Error(msg, convertFields(fields...)...)
}

func (l *Logger) Debug(msg string, fields ...interface{}) {
	l.Logger.Debug(msg, convertFields(fields...)...)
}

func (l *Logger) Warn(msg string, fields ...interface{}) {
	l.Logger.Warn(msg, convertFields(fields...)...)
}

func (l *Logger) Fatal(msg string, fields ...interface{}) {
	l.Logger.Fatal(msg, convertFields(fields...)...)
}

func (l *Logger) With(fields ...interface{}) *Logger {
	return &Logger{
		Logger: l.Logger.With(convertFields(fields...)...),
	}
}
