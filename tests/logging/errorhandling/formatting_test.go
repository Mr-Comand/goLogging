package errorhandling_test

import (
	"testing"

	"github.com/Mr-Comand/goLogging/logging/errorhandling"
	"github.com/stretchr/testify/assert"
)

func newPresetError(devMsg, userMsg, logMsg string) *errorhandling.CustomError {
	preset := errorhandling.CustomErrorPreset{
		Code:        500,
		DevMessage:  devMsg,
		UserMessage: userMsg,
		LogMessage:  logMsg,
		Source:      &errorhandling.GenericErrorsSource,
		HttpCode:    500,
	}
	return preset.New()
}

func TestFormatAndHTML(t *testing.T) {
	err := newPresetError(
		"Dev: {field|unknown}",
		"User: {field|unknown}",
		"Log: {field|unknown}",
	)

	values := map[string]string{
		"field": "username",
	}

	// Apply formatting
	err.Format(values)

	// Use public HTML() method
	html, code := err.HTML()

	assert.Equal(t, 500, code)
	assert.Equal(t, "Dev: username", html.DevMessage)
	assert.Equal(t, "User: username", html.UserMessage)
}

func TestFormatMultiplePlaceholdersAndHTML(t *testing.T) {
	err := newPresetError(
		"Dev: {field|unknown} at {location|unknown}",
		"User: {field|unknown} in {location|unknown}",
		"",
	)

	values := map[string]string{
		"field":    "password",
		"location": "form",
	}

	err.Format(values)
	html, _ := err.HTML()

	assert.Equal(t, "Dev: password at form", html.DevMessage)
	assert.Equal(t, "User: password in form", html.UserMessage)
}

func TestEscapedPlaceholders(t *testing.T) {
	err := newPresetError(
		"Literal \\{escaped\\} and {field|default}",
		"Check \\{escaped\\} and {field|default}",
		"",
	)

	values := map[string]string{
		"field": "value",
	}

	err.Format(values)
	html, _ := err.HTML()

	// Escaped braces are preserved in output
	assert.Equal(t, "Literal {escaped} and value", html.DevMessage)
	assert.Equal(t, "Check {escaped} and value", html.UserMessage)
}
func TestFormatWithValuesAndDefaults(t *testing.T) {
	err := newPresetError(
		"Dev: {field|unknown} at {location|nowhere} and {missing}",
		"User: {field|unknown} in {location|nowhere} and {missing}",
		"",
	)

	values := map[string]string{
		"field": "username",
	}

	err.Format(values)
	html, _ := err.HTML()

	// field has value -> replaced
	assert.Equal(t, "Dev: username at nowhere and missing", html.DevMessage)
	assert.Equal(t, "User: username in nowhere and missing", html.UserMessage)
}
func TestFormatWithoutDefault(t *testing.T) {
	err := newPresetError(
		"Dev: {missing}",
		"User: {missing}",
		"",
	)

	values := map[string]string{}

	err.Format(values)
	html, _ := err.HTML()

	// No default, no value -> fallback to placeholder name
	assert.Equal(t, "Dev: missing", html.DevMessage)
	assert.Equal(t, "User: missing", html.UserMessage)
}
func TestFormatWithEscapedDefault(t *testing.T) {
	err := newPresetError(
		"Dev: \\{escaped|default\\} and {field|defaultVal}",
		"User: \\{escaped|default\\} and {field|defaultVal}",
		"",
	)

	values := map[string]string{
		"field": "value",
	}

	err.Format(values)
	html, _ := err.HTML()

	// Escaped braces remain literal, default inside is ignored
	assert.Equal(t, "Dev: {escaped|default} and value", html.DevMessage)
	assert.Equal(t, "User: {escaped|default} and value", html.UserMessage)
}

func TestFormatWithNoValueUsesDefault(t *testing.T) {
	err := newPresetError(
		"Dev: {field|defaultVal}",
		"User: {field|defaultVal}",
		"",
	)

	values := map[string]string{} // no value provided

	err.Format(values)
	html, _ := err.HTML()

	// No value -> uses default
	assert.Equal(t, "Dev: defaultVal", html.DevMessage)
	assert.Equal(t, "User: defaultVal", html.UserMessage)
}
