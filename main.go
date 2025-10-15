package main

import (
	"log"
	"os"

	"github.com/Mr-Comand/logging/logging"
)

func main() {
	// Create a custom logger using the standard log package
	customLogger := log.New(os.Stdout, "", log.LstdFlags)

	// Use the default logger and set the custom logger
	logger := logging.Default()
	logger.SetLogger(customLogger)
	logger.SetLogLevel(logging.INFO)

	// Log messages at different levels
	logger.Debug("This is a debug message") // Won't print since level is INFO
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
	logger.Fail("This is a fail message")

	// Create a system module logger
	moduleLogger := logger.NewSystemModuleLogger("Database", logging.Blue, logging.Yellow)
	moduleLogger.Info("Connecting to database")
	moduleLogger.Warn("Connection slow")
	moduleLogger.Error("Connection failed")

	// Use formatted logging
	logger.InfoF("User %s logged in at %s", "Alice", "2023-10-01")
	moduleLogger.DebugF("Query executed in %d ms", 150)

	// Demonstrate changing log level
	logger.SetLogLevel(logging.DEBUG)
	logger.Debug("Now debug messages will print")

	// Use the default logger
	defaultLogger := logging.Default()
	defaultLogger.Info("Using the default logger")

	logger.DisableTextModifier = true
	// Log messages at different levels
	logger.Debug("This is a debug message without formatting") // Won't print since level is INFO
	logger.Info("This is an info message without formatting")
	logger.Warn("This is a warning message without formatting")
	logger.Error("This is an error message without formatting")
	logger.Fail("This is a fail message without formatting")
}
