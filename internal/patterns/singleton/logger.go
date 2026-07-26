package singleton

import (
	"GourmetOS/pkg/logger"
	"sync"
)

type LoggerManager struct {
	logger *logger.Logger
}

var (
	loggerInstance *LoggerManager
	loggerOnce     sync.Once
)

func GetLoggerManager() *LoggerManager {
	loggerOnce.Do(func() {
		loggerInstance = &LoggerManager{
			logger: logger.NewLogger("info"),
		}
	})

	return loggerInstance
}

func (m *LoggerManager) GetLogger() *logger.Logger {
	return m.logger
}

func (m *LoggerManager) Info(msg string, fields ...interface{}) {
	m.logger.Info(msg, fields...)
}

func (m *LoggerManager) Error(msg string, fields ...interface{}) {
	m.logger.Error(msg, fields...)
}

func (m *LoggerManager) Debug(msg string, fields ...interface{}) {
	m.logger.Debug(msg, fields...)
}

func (m *LoggerManager) Warn(msg string, fields ...interface{}) {
	m.logger.Warn(msg, fields...)
}
