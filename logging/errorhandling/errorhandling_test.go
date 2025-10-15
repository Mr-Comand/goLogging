package errorhandling

import (
	"errors"
	"testing"

	"github.com/Mr-Comand/goLogging/logging"
)

func TestNewCustomError(t *testing.T) {
	preset := CustomErrorPreset{
		Code:        404,
		UserMessage: "Not found",
		DevMessage:  "Resource not found",
		LogMessage:  "Resource not found in database",
		Source:      &GenericErrorsSource,
		Level:       ErrorMedium,
		HttpCode:    404,
	}

	customErr := NewCustomError(preset)
	if customErr.Code != 404 {
		t.Errorf("Expected code 404, got %d", customErr.Code)
	}
	if customErr.UserMessage != "Not found" {
		t.Errorf("Expected user message 'Not found', got %s", customErr.UserMessage)
	}
	if customErr.TraceId == "" {
		t.Error("Expected non-empty TraceId")
	}
}

func TestErrorHandlerParse(t *testing.T) {
	logger := logging.Default()
	handler := NewErrorHandler(logger)

	// Test parsing a standard error
	stdErr := errors.New("standard error")
	parsedErr := handler.Parse(stdErr)

	if parsedErr.Code != 500 {
		t.Errorf("Expected code 500 for generic error, got %d", parsedErr.Code)
	}
	if parsedErr.LogMessage != "standard error" {
		t.Errorf("Expected log message 'standard error', got %s", parsedErr.LogMessage)
	}
}

func TestErrorHandlerRegisterSource(t *testing.T) {
	logger := logging.Default()
	handler := NewErrorHandler(logger)

	source := &ErrorSource{
		Name: "TestSource",
		ParseError: func(err error) *CustomError {
			if err.Error() == "test error" {
				return &CustomError{
					CustomErrorPreset: CustomErrorPreset{
						Code:       400,
						LogMessage: "Parsed test error",
						Source:     &GenericErrorsSource,
					},
					TraceId: "test-trace",
				}
			}
			return nil
		},
	}

	handler.RegisterErrorSource(source)

	testErr := errors.New("test error")
	parsedErr := handler.Parse(testErr)

	if parsedErr.Code != 400 {
		t.Errorf("Expected code 400 for parsed error, got %d", parsedErr.Code)
	}
	if parsedErr.LogMessage != "Parsed test error" {
		t.Errorf("Expected log message 'Parsed test error', got %s", parsedErr.LogMessage)
	}
}

func TestCustomErrorIsFromPreset(t *testing.T) {
	preset := CustomErrorPreset{
		Code:   404,
		Source: &GenericErrorsSource,
	}

	customErr := preset.New()

	if !customErr.IsFromPreset(&preset) {
		t.Error("Expected error to be from preset")
	}

	differentPreset := CustomErrorPreset{
		Code:   500,
		Source: &GenericErrorsSource,
	}

	if customErr.IsFromPreset(&differentPreset) {
		t.Error("Expected error not to be from different preset")
	}
}
