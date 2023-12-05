package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"runtime"
	"time"
)

type Level int8

type Fields map[string]interface{}

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields //日志字段
	callers   []string
}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (level Level) String() string {
	switch level {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	default:
		return "panic"
	}
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	return &Logger{
		newLogger: log.New(w, prefix, flag),
	}
}

// clone 存在主要是为了支持链式调用以及保持日志记录器的不可变性
// 保持 Logger 实例的不可变性，更容易进行并发操作，避免竞态条件，有助于避免潜在的并发问题
func (logger *Logger) clone() *Logger {
	return &Logger{
		newLogger: logger.newLogger,
		ctx:       logger.ctx,
		fields:    logger.fields,
		callers:   logger.callers,
	}
}

func (logger *Logger) WithFields(f Fields) *Logger {
	logger1 := logger.clone()
	if logger1.fields == nil {
		logger1.fields = make(Fields)
	}
	for k, v := range f {
		logger1.fields[k] = v
	}
	return logger1
}

func (logger *Logger) WithContext(ctx context.Context) *Logger {
	logger1 := logger.clone()
	logger1.ctx = ctx
	return logger1
}

// WithCaller 记录当前某一层调用栈的信息（程序计数器、文件信息和行号）。
//
// 参数:
//   - skip: 整数，表示要跳过的调用栈层数。
//
// 返回值:
//   - *Logger: 新创建的 Logger 实例，包含了添加调用者信息后的信息。
func (logger *Logger) WithCaller(skip int) *Logger {
	logger1 := logger.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		logger1.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return logger1
}

// WithCallersFrames 方法获取调用者的调用栈信息，并将信息格式化为字符串数组。
//
// 返回值:
//   - *Logger: 新创建的 Logger 实例，包含了添加多个调用者信息后的信息。
func (logger *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	var callers []string
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}
	logger1 := logger.clone()
	logger1.callers = callers
	return logger1
}

func (logger *Logger) WithTrace() *Logger {
	ginCtx, ok := logger.ctx.(*gin.Context)
	if ok {
		return logger.WithFields(Fields{
			"trace_id": ginCtx.MustGet("X-Trace-ID"),
			"span_id":  ginCtx.MustGet("X-Span-ID"),
		})
	}
	return logger
}

func (logger *Logger) WithLevel(level Level) *Logger {
	// Create a new instance of the logger with the desired log level
	newLogger := logger.clone()
	newLogger.fields = make(Fields) // If you want to reset fields when changing the level
	newLogger.fields["level"] = level.String()
	return newLogger
}

// JSONFormat 日志内容格式化
func (logger *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	data := make(Fields, len(logger.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = logger.callers
	if len(logger.fields) > 0 {
		for k, v := range logger.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	return data
}

func (logger *Logger) Output(message string) {
	level := LevelInfo // Set a default level if not explicitly set before
	if value, ok := logger.fields["level"]; ok {
		// Try to extract the log level from the logger fields
		if logLevel, ok := value.(Level); ok {
			level = logLevel
		}
	}
	body, _ := json.Marshal(logger.JSONFormat(level, message))
	content := string(body)

	switch level {
	case LevelDebug, LevelInfo, LevelWarn, LevelError:
		logger.newLogger.Print(content)
	case LevelFatal:
		logger.newLogger.Fatal(content)
	case LevelPanic:
		logger.newLogger.Panic(content)
	default:
		logger.newLogger.Print(content)
	}
}

func (logger *Logger) Info(ctx context.Context, v ...interface{}) {
	logger.WithLevel(LevelInfo).WithContext(ctx).WithTrace().Output(fmt.Sprint(v...))
}

func (logger *Logger) Infof(ctx context.Context, format string, v ...interface{}) {
	logger.WithLevel(LevelInfo).WithContext(ctx).WithTrace().Output(fmt.Sprintf(format, v...))
}

func (logger *Logger) Fatal(ctx context.Context, v ...interface{}) {
	logger.WithLevel(LevelFatal).WithContext(ctx).WithTrace().Output(fmt.Sprint(v...))
}

func (logger *Logger) Fatalf(ctx context.Context, format string, v ...interface{}) {
	logger.WithLevel(LevelFatal).WithContext(ctx).WithTrace().Output(fmt.Sprintf(format, v...))
}

func (logger *Logger) Error(ctx context.Context, v ...interface{}) {
	logger.WithLevel(LevelError).WithContext(ctx).WithTrace().Output(fmt.Sprint(v...))
}

func (logger *Logger) Errorf(ctx context.Context, format string, v ...interface{}) {
	logger.WithLevel(LevelError).WithContext(ctx).WithTrace().Output(fmt.Sprintf(format, v...))
}

func (logger *Logger) Panicf(ctx context.Context, format string, v ...interface{}) {
	logger.WithLevel(LevelPanic).WithContext(ctx).WithTrace().Output(fmt.Sprintf(format, v...))
}
