package logger

import "testing"

func TestFileLogger(t *testing.T) {
	logger := NewFileLogger(LogLevelDebug, "e:/logs/", "test")
	logger.Debug("user id[%d] is come from china.", 323232)
	logger.Warn("test warn log.")
	logger.Fatal("test fatal log.")
	logger.Close()
}

