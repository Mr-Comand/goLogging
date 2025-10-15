package logging

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestLoggerLevels(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(log.New(&buf, "", 0), INFO)

	// Test INFO level
	logger.Info("test info")
	if !strings.Contains(buf.String(), "test info") {
		t.Errorf("Expected 'test info' in output, got: %s", buf.String())
	}

	buf.Reset()
	logger.Debug("test debug") // Should not print at INFO level
	if strings.Contains(buf.String(), "test debug") {
		t.Errorf("Debug message should not print at INFO level, got: %s", buf.String())
	}
}

func TestLoggerSetLevel(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(log.New(&buf, "", 0), NONE)

	logger.SetLogLevel(DEBUG)
	if logger.GetLogLevel() != DEBUG {
		t.Errorf("Expected log level DEBUG, got %v", logger.GetLogLevel())
	}

	logger.SetLogLevel(INFO)
	if logger.GetLogLevel() != INFO {
		t.Errorf("Expected log level INFO, got %v", logger.GetLogLevel())
	}
}

func TestSystemModuleLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(log.New(&buf, "", 0), DEBUG)

	moduleLogger := logger.NewSystemModuleLogger("TestModule", Blue, Green)
	moduleLogger.Info("module test")

	output := buf.String()
	if !strings.Contains(output, "[TestModule]") {
		t.Errorf("Expected '[TestModule]' in output, got: %s", output)
	}
	if !strings.Contains(output, "module test") {
		t.Errorf("Expected 'module test' in output, got: %s", output)
	}
}

func TestFormattedLogging(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(log.New(&buf, "", 0), INFO)

	logger.InfoF("User %s logged in", "Alice")
	output := buf.String()
	if !strings.Contains(output, "User Alice logged in") {
		t.Errorf("Expected formatted message, got: %s", output)
	}
}

func TestDisableTextModifier(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(log.New(&buf, "", 0), INFO)
	logger.DisableTextModifier = true

	logger.Info("test message")
	output := buf.String()
	// Should not contain ANSI color codes
	if strings.Contains(output, "\033[") {
		t.Errorf("ANSI codes should be disabled, got: %s", output)
	}
	// Should still contain the log level and message
	if !strings.Contains(output, "[INFO]") || !strings.Contains(output, "test message") {
		t.Errorf("Should contain log level and message, got: %s", output)
	}
	// The output should be exactly what we expect without colors
	expected := "[INFO]\t[General]\ttest message\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}

// Helper function to create a new logger for testing
func NewLogger(logLogger *log.Logger, level LogLevel) *Logger {
	return newLogger(logLogger, level)
}
