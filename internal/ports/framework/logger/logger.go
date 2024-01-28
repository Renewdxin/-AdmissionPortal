package logger

import (
	"go.uber.org/zap"
)

// LogPort interface defines the contract for a logger
type LogPort interface {
	Log(level int, msg string, fields ...zap.Field)
	Logf(level int, format string, args ...interface{})
	SugarLogger() *zap.SugaredLogger
}
