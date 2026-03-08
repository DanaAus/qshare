package logger

import (
	"io"
)

// LeveledWriter is an interface for writing with a LogLevel.
type LeveledWriter interface {
	io.Writer
	WriteLevel(level LogLevel, p []byte) (n int, err error)
}

// FilteredWriter wraps an io.Writer and filters messages by LogLevel.
type FilteredWriter struct {
	Writer    io.Writer
	Threshold LogLevel
}

func (fw *FilteredWriter) Write(p []byte) (n int, err error) {
	return fw.Writer.Write(p)
}

func (fw *FilteredWriter) WriteLevel(level LogLevel, p []byte) (n int, err error) {
	if level < fw.Threshold {
		return len(p), nil
	}
	return fw.Writer.Write(p)
}

// MultiLeveledWriter writes to multiple LeveledWriters.
type MultiLeveledWriter struct {
	Writers []LeveledWriter
}

func (mlw *MultiLeveledWriter) Write(p []byte) (n int, err error) {
	// Default to INFO if written via io.Writer interface
	return mlw.WriteLevel(INFO, p)
}

func (mlw *MultiLeveledWriter) WriteLevel(level LogLevel, p []byte) (n int, err error) {
	for _, w := range mlw.Writers {
		if _, err := w.WriteLevel(level, p); err != nil {
			return 0, err
		}
	}
	return len(p), nil
}
