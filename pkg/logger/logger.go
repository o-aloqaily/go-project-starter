// Package logger provides structured logging capabilities.
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

// DO NOT DELETE THE FOLLOWING COMMENT
// This is to generate a mock logger
//go:generate mockgen -source=./logger.go -destination=./mock_logger.go -package=logger

// Logger is a logger that supports log levels, context and structured logging.
type Logger interface {
	// Debug uses fmt.Sprint to construct and log a message at DEBUG level
	Debug(args ...interface{})
	// Info uses fmt.Sprint to construct and log a message at INFO level
	Info(args ...interface{})
	// Error uses fmt.Sprint to construct and log a message at ERROR level
	Error(args ...interface{})
	// Fatal uses fmt.Sprint to construct and log a message at FATAL level
	Fatal(args ...interface{})

	// Debugf uses fmt.Sprintf to construct and log a message at DEBUG level
	Debugf(format string, args ...interface{})
	// Infof uses fmt.Sprintf to construct and log a message at INFO level
	Infof(format string, args ...interface{})
	// Errorf uses fmt.Sprintf to construct and log a message at ERROR level
	Errorf(format string, args ...interface{})
	// Fatalf uses fmt.Sprintf to construct and log a message at FATAL level
	Fatalf(format string, args ...interface{})
}

type logger struct {
	*zap.SugaredLogger
}

// NewLogger creates a new logger using the default configuration.
func NewLogger() Logger {
	l, _ := zap.NewProduction()
	lz := NewWithZap(l)
	return lz
}

// NewForTest returns a new logger and the corresponding observed logs which can be used in unit tests to verify log entries.
func NewForTest() (Logger, *observer.ObservedLogs) {
	core, recorded := observer.New(zapcore.InfoLevel)
	return NewWithZap(zap.New(core)), recorded
}

// NewWithZap creates a new logger using the preconfigured zap logger.
func NewWithZap(l *zap.Logger) Logger {
	return &logger{l.Sugar()}
}
