# Go Logging Library

A comprehensive, colored logging library for Go with advanced features including system module loggers, integrated error handling, and customizable output.

## Features

- **Colored Output**: Automatic color coding for different log levels
- **System Module Loggers**: Create dedicated loggers for different parts of your application
- **Error Handling Integration**: Built-in error handling with customizable error presets
- **Multiple Log Levels**: DEBUG, INFO, WARN, ERROR, FAIL, NONE
- **Formatted Logging**: Support for printf-style formatted messages
- **Color Disable Option**: Can disable ANSI colors for plain text output
- **Thread-Safe**: Uses standard Go log package for thread safety

## Installation

```bash
go get github.com/Mr-Comand/logging
```

## Quick Start

```go
package main

import (
    "github.com/Mr-Comand/logging"
)

func main() {
    // Basic logging
    logging.Info("Application started")
    logging.Debug("Debug information")
    logging.Warn("Warning message")
    logging.Error("Error occurred")
    logging.Fail("Critical failure")

    // Formatted logging
    logging.InfoF("User %s logged in at %s", "alice", "2023-10-01")

    // Create a system module logger
    dbLogger := logging.Default().NewSystemModuleLogger("Database", logging.Blue, logging.Cyan)
    dbLogger.Info("Connected to database")
    dbLogger.Error("Connection failed")

    // Disable colors if needed
    logging.Default().DisableTextModifier = true
    logging.Info("This will be plain text")
}
```

## API Reference

### Basic Logging Functions

- `logging.Debug(msg ...string)` - Debug level logging (blue)
- `logging.Info(msg ...string)` - Info level logging (green)
- `logging.Warn(msg ...string)` - Warning level logging (yellow)
- `logging.Error(msg ...string)` - Error level logging (red)
- `logging.Fail(msg ...string)` - Fail level logging (red with magenta background)

### Formatted Logging

- `logging.DebugF(format string, v ...any)`
- `logging.InfoF(format string, v ...any)`
- `logging.WarnF(format string, v ...any)`
- `logging.ErrorF(format string, v ...any)`
- `logging.FailF(format string, v ...any)`

### Logger Management

```go
logger := logging.Default()

// Set log level
logger.SetLogLevel(logging.DEBUG)

// Create system module logger
moduleLogger := logger.NewSystemModuleLogger("ModuleName", logging.Blue, logging.Green)

// Disable colors
logger.DisableTextModifier = true
```

### Error Handling

```go
import "github.com/Mr-Comand/logging/logging/errorhandling"

// Register error sources
errorhandling.RegisterErrorSource(&errorhandling.ErrorSource{
    Name: "Database",
    ParseError: func(err error) *errorhandling.CustomError {
        // Custom error parsing logic
        return nil
    },
})

// Parse errors
customErr := errorhandling.Parse(someError)
if customErr != nil {
    customErr.Log()
}

// Web error handling
func handleError(w http.ResponseWriter, r *http.Request, err error) {
    customErr := errorhandling.Parse(err)
    if customErr != nil {
        // Continue execution if error allows
        if customErr.HandelWeb(w, r) {
            // Handle error but continue
            return
        }
        // Stop execution for critical errors
        return
    }
    // Handle unparsed errors
    http.Error(w, "Internal Server Error", 500)
}

// Exit immediately on critical errors
func handleCriticalError(w http.ResponseWriter, r *http.Request, err error) {
    customErr := errorhandling.Parse(err)
    if customErr != nil {
        customErr.HandelWebExit(w, r)
        return
    }
    http.Error(w, "Internal Server Error", 500)
}
```

## Log Levels

- `DEBUG` (0) - Detailed debug information
- `INFO` (1) - General information
- `WARN` (2) - Warning messages
- `ERROR` (3) - Error messages
- `FAIL` (4) - Critical failures
- `NONE` (5) - No logging

## Colors

The library supports ANSI color codes:

- `Black`, `Red`, `Green`, `Yellow`, `Blue`, `Magenta`, `Cyan`, `White`
- Bright variants: `BrightRed`, `BrightGreen`, etc.
- Background colors: `RedBG`, `GreenBG`, etc.

## Error Handling

The error handling package provides:

- Custom error presets with metadata
- Automatic trace ID generation
- HTTP status code mapping
- Structured error logging
- Error source registration
- Web response handling with `HandelWeb` and `HandelWebExit`

## Contributing

Contributions are welcome! Please ensure:

- All tests pass: `go test ./...`
- Code is properly formatted: `go fmt ./...`
- No linting issues: `go vet ./...`

## License

MIT License - see LICENSE file for details.
