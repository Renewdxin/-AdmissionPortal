package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// LogLevel is a custom type for log levels

const (
	// InfoLevel represents the info log level
	InfoLevel int = iota
	// DebugLevel represents the debug log level
	DebugLevel
	// WarnLevel represents the warning log level
	WarnLevel
	// ErrorLevel represents the error log level
	ErrorLevel
	// FatalLevel represents the fatal log level
	FatalLevel
)

// LogAdapter is an implementation of the Logger interface
type LogAdapter struct {
	logger *zap.Logger
}

var Logger *LogAdapter

// NewLogger creates a new Logger instance
func NewLogger() *LogAdapter {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		TimeKey:      "time",
		CallerKey:    "caller",
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	// Configure console output
	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)
	consoleDebugging := zapcore.Lock(zapcore.AddSync(os.Stdout))

	// Configure file output
	fileEncoder := zapcore.NewJSONEncoder(encoderCfg)
	fileOutput, _, err := zap.Open("storage/log.txt")
	if err != nil {
		panic(err)
	}

	// Combine console and file outputs
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleDebugging, zap.InfoLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(fileOutput), zap.DebugLevel),
	)

	// 使用指针接收者赋值给 zl.logger
	logger := zap.New(core, zap.AddCaller())
	return &LogAdapter{logger: logger}
}

// Log logs a message with the specified level
func (zl LogAdapter) Log(level int, msg string, fields ...zap.Field) {
	switch level {
	case InfoLevel:
		zl.logger.Info(msg, fields...)
	case DebugLevel:
		zl.logger.Debug(msg, fields...)
	case WarnLevel:
		zl.logger.Warn(msg, fields...)
	case ErrorLevel:
		zl.logger.Error(msg, fields...)
	case FatalLevel:
		zl.logger.Fatal(msg, fields...)
	}
}

func (zl LogAdapter) Logf(level int, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	zl.Log(level, msg) // You can adjust the log level as needed
}

// SugarLogger returns a SugaredLogger from the global logger
func (zl LogAdapter) SugarLogger() *zap.SugaredLogger {
	return zl.logger.Sugar()
}
