package logger

import (
	"os"
	"sync"
)

var (
	globalLogger Logger
	mu           sync.Mutex
)

// SetGlobalLogger sets the global logger instance.
func SetGlobalLogger(l Logger) {
	mu.Lock()
	defer mu.Unlock()
	globalLogger = l
}

// GetGlobalLogger returns the global logger instance.
func GetGlobalLogger() Logger {
	mu.Lock()
	defer mu.Unlock()
	if globalLogger == nil {
		// Fallback to a basic logger writing to Stdout
		globalLogger = &StructuredLogger{
			Writer:    os.Stdout,
			Component: "default",
			PID:       os.Getpid(),
		}
	}
	return globalLogger
}

// Global functions for convenience
func Info(msg string)  { GetGlobalLogger().Info(msg) }
func Debug(msg string) { GetGlobalLogger().Debug(msg) }
func Warn(msg string)  { GetGlobalLogger().Warn(msg) }
func Error(msg string) { GetGlobalLogger().Error(msg) }

// WithComponent returns a new logger with the specified component name.
func WithComponent(component string) Logger {
	return GetGlobalLogger().WithComponent(component)
}
