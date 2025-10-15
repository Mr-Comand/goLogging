package logging

// Debug level log with blue color
func Debug(msg ...string) {
	std.Debug(msg...)
}

// Info level log with green color
func Info(msg ...string) {
	std.Info(msg...)
}

// Warn level log with yellow color
func Warn(msg ...string) {
	std.Warn(msg...)
}

// Error level log with red color
func Error(msg ...string) {
	std.Error(msg...)
}

// Error level log with red color
func Fail(msg ...string) {
	std.Fail(msg...)
}

// Debug level log with blue color
func DebugF(format string, v ...any) {
	std.DebugF(format, v...)
}

// Info level log with green color
func InfoF(format string, v ...any) {
	std.InfoF(format, v...)
}

// Warn level log with yellow color
func WarnF(format string, v ...any) {
	std.WarnF(format, v...)
}

// Error level log with red color
func ErrorF(format string, v ...any) {
	std.ErrorF(format, v...)
}

// Error level log with red color
func FailF(format string, v ...any) {
	std.FailF(format, v...)
}
