package logger

import (
	"go.uber.org/zap"
)

// LoggerPorts interface defines the contract for a logger
type LoggerPorts interface {
	Init(logFilePath string)
	Log(level int, msg string, fields ...zap.Field)
	SugarLogger() *zap.SugaredLogger
}
