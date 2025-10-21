package errorhandling

import (
	"slices"

	"github.com/Mr-Comand/goLogging/logging"
)

var std *ErrorHandler = NewErrorHandler(logging.Default())

type ErrorHandler struct {
	logger       *logging.Logger
	sml          *logging.SystemModuleLogger
	errorSources []*ErrorSource
}

func NewErrorHandler(logger *logging.Logger) *ErrorHandler {
	sml := logger.NewSystemModuleLogger("Error Handler", "", "")
	errorHandler := &ErrorHandler{
		logger:       logger,
		errorSources: []*ErrorSource{},
		sml:          sml,
	}
	return errorHandler
}

func Default() *ErrorHandler {
	return std
}

func (hand *ErrorHandler) UpdateLogger(logger *logging.Logger) {
	if logger != nil {
		sml := logger.NewSystemModuleLogger("Error Handler", "", "")
		hand.logger = logger
		hand.sml = sml
		hand.sml.Info("Switched to this Logger.")
	}
}

// Registers one or more error sources to the handler.
func (hand *ErrorHandler) RegisterErrorSource(sources ...*ErrorSource) *ErrorHandler {
	for _, src := range sources {
		// Avoid duplicates
		exists := slices.Contains(hand.errorSources, src)
		if !exists {
			if src.SML == nil {
				src.SML = hand.logger.NewSystemModuleLogger(src.Name, "", "")
			}
			hand.errorSources = append(hand.errorSources, src)
			if hand.sml != nil {
				hand.sml.Info("Registered error source: " + src.Name)
			}
		} else {
			hand.sml.Warn("ErrorSource already registered.")
		}
	}
	return hand
}

// Unregisters a specific error source.
func (hand *ErrorHandler) UnregisterErrorSource(source *ErrorSource) *ErrorHandler {
	for i, existing := range hand.errorSources {
		if existing == source {
			hand.errorSources = append(hand.errorSources[:i], hand.errorSources[i+1:]...)
			if hand.sml != nil {
				hand.sml.Info("Unregistered error source: " + source.Name)
			}
			break
		}
	}
	return hand
}

// Attempts to parse an error using registered error sources.
// Returns a CustomError if recognized, otherwise wraps it in a generic CustomError.
func (hand *ErrorHandler) Parse(err error) *CustomError {
	// If the error is already a CustomError, return it directly
	if cErr, ok := err.(*CustomError); ok {
		return cErr
	}

	// Otherwise, try all registered error sources
	for _, errorSource := range hand.errorSources {
		if errorSource.ParseError == nil {
			hand.sml.DebugF("errorSource.ParseError of %s is nil-Pointer", errorSource.Name)
			continue
		}
		if cError := errorSource.ParseError(err); cError != nil {
			return cError
		}
	}

	// Fallback: wrap in a new CustomError
	return NewError(err)
}
func (hand *ErrorHandler) ParseWith(err error, ErrorSources ...ErrorSource) *CustomError {
	// If the error is already a CustomError, return it directly
	if cErr, ok := err.(*CustomError); ok {
		return cErr
	}

	// Otherwise, try all registered error sources
	for _, errorSource := range ErrorSources {
		if cError := errorSource.ParseError(err); cError != nil {
			return cError
		}
	}

	// Fallback: wrap in a new CustomError
	return NewError(err)
}
func (hand *ErrorHandler) GetErrorSources(ErrorSourceNames ...string) []*ErrorSource {
	if hand == nil || hand.errorSources == nil {
		return nil
	}

	var result []*ErrorSource
	nameSet := make(map[string]struct{}, len(ErrorSourceNames))

	// Build a set for quick lookup
	for _, n := range ErrorSourceNames {
		nameSet[n] = struct{}{}
	}

	for _, src := range hand.errorSources {
		if _, ok := nameSet[src.Name]; ok {
			result = append(result, src)
		}
	}

	return result
}
