package logger

import (
	"fmt"
	stdlog "log"
)

type ILogger interface {
	Info(msg ...interface{})
	Warn(msg ...interface{})
	Error(msg ...interface{})
}

type logger struct {
}

var Logger *logger

// Init NewLogger creates a basic logger that wraps the core log library.
func Init() ILogger {
	Logger = &logger{}
	return Logger
}

// Info log message
func (b *logger) Info(msg ...interface{}) {
	stdlog.Printf("[INFO] %s", fmt.Sprint(msg...))
}

// Warn log message
func (b *logger) Warn(msg ...interface{}) {
	stdlog.Printf("[WARN] %s", fmt.Sprint(msg...))
}

// Error log message
func (b *logger) Error(msg ...interface{}) {
	stdlog.Printf("[ERROR] %s", fmt.Sprint(msg...))
}

func (b *logger) Fatal(msg ...interface{}) {
	stdlog.Printf("[FATAL] %s", fmt.Sprint(msg...))
}
