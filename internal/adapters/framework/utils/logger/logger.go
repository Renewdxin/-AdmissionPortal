package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// LogLevel is a custom type for log levels
type LogLevel int

const (
	// InfoLevel represents the info log level
	InfoLevel LogLevel = iota
	// DebugLevel represents the debug log level
	DebugLevel
	// WarnLevel represents the warning log level
	WarnLevel
	// ErrorLevel represents the error log level
	ErrorLevel
	// FatalLevel represents the fatal log level
	FatalLevel
)

type LoggerAdapter struct {
	Logger *zap.Logger
}

func NewAdapter(Logger *zap.Logger) *LoggerAdapter {
	return &LoggerAdapter{
		Logger: Logger,
	}
}

var (
	// Logger is the global logger instance
	Logger *zap.Logger
)

// InitLogger initializes the logger with file and console output
func InitLogger(logFilePath string) {
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
	fileOutput, _, _ := zap.Open(logFilePath)

	// Combine console and file outputs
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleDebugging, zap.InfoLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(fileOutput), zap.DebugLevel),
	)

	Logger = zap.New(core, zap.AddCaller())
	defer Logger.Sync()
}

// SugarLogger returns a SugaredLogger from the global logger
func SugarLogger() *zap.SugaredLogger {
	return Logger.Sugar()
}

// Log logs a message with the specified level
func Log(level LogLevel, msg string, fields ...zap.Field) {
	switch level {
	case InfoLevel:
		Logger.Info(msg, fields...)
	case DebugLevel:
		Logger.Debug(msg, fields...)
	case WarnLevel:
		Logger.Warn(msg, fields...)
	case ErrorLevel:
		Logger.Error(msg, fields...)
	case FatalLevel:
		Logger.Fatal(msg, fields...)
	}
}

func (adapter LoggerAdapter) Info(msg string, fields ...zap.Field) {

}

func (adapter LoggerAdapter) Infof(msg string, format string, args ...interface{}) {

}

func (adapter LoggerAdapter) Error(msg string, fields ...zap.Field) {

}

func (adapter LoggerAdapter) Errorf(msg string, format string, args ...interface{}) {

}

func (adapter LoggerAdapter) Warn(msg string, fields ...zap.Field) {

}

func (adapter LoggerAdapter) Warnf(msg string, format string, args ...interface{}) {

}

func (adapter LoggerAdapter) Debug(msg string, fields ...zap.Field) {

}

func (adapter LoggerAdapter) Debugf(msg string, format string, args ...interface{}) {

}

func (adapter LoggerAdapter) Fatal(msg string, fields ...zap.Field) {

}

func (adapter LoggerAdapter) Fatalf(msg string, format string, args ...interface{}) {

}

func (adapter LoggerAdapter) Panic(msg string, fields ...zap.Field) {

}

func (adapter LoggerAdapter) Panicf(msg string, format string, args ...interface{}) {

}

func (adapter LoggerAdapter) Sync() error {
	return nil
}

func (adapter LoggerAdapter) Level() zapcore.Level {

}
