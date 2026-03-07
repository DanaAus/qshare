package logger

import (
	"fmt"
	"io"
	"time"
)

// LogLevel defines the severity of a log message.
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// String returns the string representation of the LogLevel.
func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// Logger is the interface for logging messages.
type Logger interface {
	Log(level LogLevel, msg string)
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	WithComponent(component string) Logger
}

// StructuredLogger implements the Logger interface with formatted output.
type StructuredLogger struct {
	Writer    io.Writer
	Component string
	PID       int
}

// WithComponent returns a new logger with the specified component name.
func (l *StructuredLogger) WithComponent(component string) Logger {
	return &StructuredLogger{
		Writer:    l.Writer,
		Component: component,
		PID:       l.PID,
	}
}

// Log writes a message at the specified level.
func (l *StructuredLogger) Log(level LogLevel, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	formatted := fmt.Sprintf("[%s] [%s] [%s] [%d] %s\n", timestamp, level, l.Component, l.PID, msg)
	
	if lw, ok := l.Writer.(LeveledWriter); ok {
		lw.WriteLevel(level, []byte(formatted))
	} else {
		fmt.Fprint(l.Writer, formatted)
	}
}

func (l *StructuredLogger) Debug(msg string) { l.Log(DEBUG, msg) }
func (l *StructuredLogger) Info(msg string)  { l.Log(INFO, msg) }
func (l *StructuredLogger) Warn(msg string)  { l.Log(WARN, msg) }
func (l *StructuredLogger) Error(msg string) { l.Log(ERROR, msg) }
