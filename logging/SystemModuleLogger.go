package logging

type SystemModuleLogger struct {
	level      *LogLevel
	ModuleName string
	NameColor  TextModifier
	TextColor  TextModifier
	logger     *Logger
}

func (l *Logger) NewSystemModuleLogger(moduleName string, nameColor, textColor TextModifier) *SystemModuleLogger {
	if l.systemModules == nil {
		l.systemModules = make(map[string]*SystemModuleLogger)
	}
	// Check if module already exists
	if existing, ok := l.systemModules[moduleName]; ok {
		return existing
	}

	systemModuleLogger := &SystemModuleLogger{
		level:      &l.level,
		logger:     l,
		ModuleName: moduleName,
		NameColor:  nameColor,
		TextColor:  textColor,
	}
	l.systemModules[moduleName] = systemModuleLogger
	return systemModuleLogger
}

func (l *Logger) GetSystemModule(moduleName string) *SystemModuleLogger {
	if l.systemModules == nil {
		return nil
	}
	return l.systemModules[moduleName]
}

// Set LogLevel of the SystemModuleLogger
// set logLevel to -1 for inherit LogLevel of logger
func (sm *SystemModuleLogger) SetLogLevel(logLevel LogLevel) {
	if logLevel == -1 {
		sm.level = &sm.logger.level
	}
	newLevel := max(min(logLevel, NONE), DEBUG)
	sm.level = &newLevel
}
func (sm *SystemModuleLogger) ResetLogLevel() {
	sm.level = &sm.logger.level
}

func (sm *SystemModuleLogger) GetLogLevel() LogLevel {
	return *sm.level
}

// Debug level log with blue color
func (sm *SystemModuleLogger) Debug(msg ...string) {
	if *sm.level <= DEBUG {
		sm.logger.logWithLevel(Blue, "DEBUG", sm, "", msg...)
	}
}

// Info level log with green color
func (sm *SystemModuleLogger) Info(msg ...string) {
	if *sm.level <= INFO {
		sm.logger.logWithLevel(Green, "INFO", sm, "", msg...)
	}
}

// Warn level log with yellow color
func (sm *SystemModuleLogger) Warn(msg ...string) {
	if *sm.level <= WARN {
		sm.logger.logWithLevel(Yellow, "WARN", sm, "", msg...)
	}
}

// Error level log with red color
func (sm *SystemModuleLogger) Error(msg ...string) {
	if *sm.level <= ERROR {
		sm.logger.logWithLevel(Red, "ERROR", sm, "", msg...)
	}
}

// Error level log with red color
func (sm *SystemModuleLogger) Fail(msg ...string) {
	if *sm.level <= FAIL {
		sm.logger.logWithLevel(Red+MagentaBG, "FAIL", sm, Red+MagentaBG, msg...)
	}
}

// Debug level log with blue color
func (sm *SystemModuleLogger) DebugF(format string, v ...any) {
	if *sm.level <= DEBUG {
		sm.logger.logWithLevelF(Blue, "DEBUG", sm, "", format, v...)
	}
}

// Info level log with green color
func (sm *SystemModuleLogger) InfoF(format string, v ...any) {
	if *sm.level <= INFO {
		sm.logger.logWithLevelF(Green, "INFO", sm, "", format, v...)
	}
}

// Warn level log with yellow color
func (sm *SystemModuleLogger) WarnF(format string, v ...any) {
	if *sm.level <= WARN {
		sm.logger.logWithLevelF(Yellow, "WARN", sm, "", format, v...)
	}
}

// Error level log with red color
func (sm *SystemModuleLogger) ErrorF(format string, v ...any) {
	if *sm.level <= ERROR {
		sm.logger.logWithLevelF(Red, "ERROR", sm, "", format, v...)
	}
}

// Error level log with red color
func (sm *SystemModuleLogger) FailF(format string, v ...any) {
	if *sm.level <= FAIL {
		sm.logger.logWithLevelF(Red+MagentaBG, "FAIL", sm, Red+MagentaBG, format, v...)
	}
}
func (sm *SystemModuleLogger) Printf(format string, v ...any) {
	sm.logger.logWithLevelF("", "????", sm, "", format, v...)
}

func (sm *SystemModuleLogger) Println(msg ...string) {
	sm.Debug(msg...)
}
