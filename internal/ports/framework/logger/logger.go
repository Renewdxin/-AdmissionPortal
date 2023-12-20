package logger

import (
	"go.uber.org/zap"
)

// LoggerPorts interface defines the contract for a logger
type LoggerPorts interface {
	Log(level int, msg string, fields ...zap.Field)
	Logf(level int, format string, args ...interface{})
	SugarLogger() *zap.SugaredLogger
}
