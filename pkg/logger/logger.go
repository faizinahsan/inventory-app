package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger wraps zap logger
type Logger struct {
	*zap.Logger
}

// NewLogger creates a new logger instance
func NewLogger(level, format string) (*Logger, error) {
	var config zap.Config

	if format == "json" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Set log level
	switch level {
	case "debug":
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}

	zapLogger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{Logger: zapLogger}, nil
}

// NewDefaultLogger creates a default logger
func NewDefaultLogger() *Logger {
	logger, _ := NewLogger("info", "console")
	return logger
}

// WithField adds a field to the logger
func (l *Logger) WithField(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

// WithFields adds multiple fields to the logger
func (l *Logger) WithFields(fields map[string]interface{}) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return zapFields
}

// GetLogger returns a logger instance based on environment
func GetLogger() *Logger {
	env := os.Getenv("ENV")
	if env == "production" {
		logger, _ := NewLogger("info", "json")
		return logger
	}
	return NewDefaultLogger()
}
