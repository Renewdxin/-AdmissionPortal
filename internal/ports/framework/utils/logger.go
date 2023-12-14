package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerPorts interface {
	Info(msg string, fields ...zap.Field)
	Infof(msg string, format string, args ...interface{})
	Error(msg string, fields ...zap.Field)
	Errorf(msg string, format string, args ...interface{})
	Warn(msg string, fields ...zap.Field)
	Warnf(msg string, format string, args ...interface{})
	Debug(msg string, fields ...zap.Field)
	Debugf(msg string, format string, args ...interface{})
	Fatal(msg string, fields ...zap.Field)
	Fatalf(msg string, format string, args ...interface{})
	Panic(msg string, fields ...zap.Field)
	Panicf(msg string, format string, args ...interface{})
	//With(fields ...zap.Field) LoggerPorts
	Sync() error
	Level() zapcore.Level
}
