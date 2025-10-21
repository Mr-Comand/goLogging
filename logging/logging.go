package logging

import (
	"fmt"
	"log"
	"strings"
)

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FAIL
	NONE
)

type LogLevel int

type LoggerInterface interface {
	Debug(msg ...string)
	Info(msg ...string)
	Warn(msg ...string)
	Error(msg ...string)
	Fail(msg ...string)
	Println(msg ...string)
	DebugF(format string, v ...any)
	InfoF(format string, v ...any)
	WarnF(format string, v ...any)
	ErrorF(format string, v ...any)
	FailF(format string, v ...any)
	Printf(format string, v ...any)
	SetLogLevel(logLevel LogLevel)
	GetLogLevel() LogLevel
}

// Logger structure
type Logger struct {
	level               LogLevel
	logger              *log.Logger
	systemModules       map[string]*SystemModuleLogger
	DisableTextModifier bool
}

var std *Logger = NewLogger(log.Default(), INFO)

func Default() *Logger { return std }

// New logger constructor
func NewLogger(logLogger *log.Logger, level LogLevel) *Logger {
	return &Logger{
		level:  level,
		logger: logLogger,
	}
}
func (l *Logger) SetLogLevel(logLevel LogLevel) {
	l.level = max(min(logLevel, 5), 0)
}
func (l *Logger) GetLogLevel() LogLevel {
	return l.level
}
func (l *Logger) SetLogger(logLogger *log.Logger) {
	if logLogger != nil {
		l.logger = logLogger
	}
}
func (l *Logger) prefixGenerator(color TextModifier, level string, module *SystemModuleLogger, textColor TextModifier) string {

	// Save the original prefix
	originalPrefix := l.logger.Prefix()

	if module != nil {
		if module.NameColor == "" {
			textColor = Reset
		}
		if textColor == "" {
			if module.TextColor == "" {
				textColor = Reset
			} else {
				textColor = module.TextColor
			}
		}
		if l.DisableTextModifier {
			// Set new prefix with color and level and module
			l.logger.SetPrefix(fmt.Sprintf("[%s]\t[%s]\t", level, module.ModuleName))
		} else {
			// Set new prefix with color and level and module
			l.logger.SetPrefix(fmt.Sprintf("%s[%s]%s\t[%s]%s\t", color, level, module.NameColor, module.ModuleName, textColor))
		}
	} else {
		if textColor == "" {
			textColor = Reset
		}
		if l.DisableTextModifier {
			// Set new prefix with color and level
			l.logger.SetPrefix(fmt.Sprintf("[%s]\t[General]\t", level))
		} else {
			// Set new prefix with color and level
			l.logger.SetPrefix(fmt.Sprintf("%s[%s]%s\t[General]\t", color, level, textColor))
		}
	}
	return originalPrefix
}

// Helper function to log messages with color and level
func (l *Logger) logWithLevel(color TextModifier, level string, module *SystemModuleLogger, textColor TextModifier, msg ...string) {
	if l.logger == nil {
		return
	}
	originalPrefix := l.prefixGenerator(color, level, module, textColor)
	message := strings.Join(msg, " ")
	if !l.DisableTextModifier {
		message += string(Reset)
	}
	// Log the message
	l.logger.Println(message)

	// Restore the original prefix
	l.logger.SetPrefix(originalPrefix)
}
func (l *Logger) logWithLevelF(color TextModifier, level string, module *SystemModuleLogger, textColor TextModifier, format string, v ...any) {
	if l.logger == nil {
		return
	}
	originalPrefix := l.prefixGenerator(color, level, module, textColor)

	message := fmt.Sprintf(format, v...)
	if !l.DisableTextModifier {
		message += string(Reset)
	}
	// Log the message
	l.logger.Println(message)

	// Restore the original prefix
	l.logger.SetPrefix(originalPrefix)
}

// Debug level log with blue color
func (l *Logger) Debug(msg ...string) {
	if l.level <= DEBUG {
		l.logWithLevel(Blue, "DEBUG", nil, "", msg...)
	}
}

// Info level log with green color
func (l *Logger) Info(msg ...string) {
	if l.level <= INFO {
		l.logWithLevel(Green, "INFO", nil, "", msg...)
	}
}

// Warn level log with yellow color
func (l *Logger) Warn(msg ...string) {
	if l.level <= WARN {
		l.logWithLevel(Yellow, "WARN", nil, "", msg...)
	}
}

// Error level log with red color
func (l *Logger) Error(msg ...string) {
	if l.level <= ERROR {
		l.logWithLevel(Red, "ERROR", nil, "", msg...)
	}
}

// Error level log with red color
func (l *Logger) Fail(msg ...string) {
	if l.level <= FAIL {
		l.logWithLevel(Red+MagentaBG, "FAIL", nil, Red+MagentaBG, msg...)
	}
}

// Debug level log with blue color
func (l *Logger) DebugF(format string, v ...any) {
	if l.level <= DEBUG {
		l.logWithLevelF(Blue, "DEBUG", nil, "", format, v...)
	}
}

// Info level log with green color
func (l *Logger) InfoF(format string, v ...any) {
	if l.level <= INFO {
		l.logWithLevelF(Green, "INFO", nil, "", format, v...)
	}
}

// Warn level log with yellow color
func (l *Logger) WarnF(format string, v ...any) {
	if l.level <= WARN {
		l.logWithLevelF(Yellow, "WARN", nil, "", format, v...)
	}
}

// Error level log with red color
func (l *Logger) ErrorF(format string, v ...any) {
	if l.level <= ERROR {
		l.logWithLevelF(Red, "ERROR", nil, "", format, v...)
	}
}

// Error level log with red color
func (l *Logger) FailF(format string, v ...any) {
	if l.level <= FAIL {
		l.logWithLevelF(Red+MagentaBG, "FAIL", nil, Red+MagentaBG, format, v...)
	}
}

func (l *Logger) Printf(format string, v ...any) {
	l.logWithLevelF("", "????", nil, "", format, v...)
}

func (l *Logger) Println(msg ...string) {
	l.Debug(msg...)
}
