package errorhandling

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/Mr-Comand/logging/logging"
)

type ErrorLevel int

var presetIDcounter uint = 0

const (
	ErrorWARN ErrorLevel = iota
	ErrorWrongUsage
	ErrorMedium
	ErrorFail
)

// CustomErrorPreset defines a reusable error template with associated metadata.
type CustomErrorPreset struct {
	// PresetID is a unique identifier for this preset. Used for comparisons. if left empty it wil be dynamically created.
	PresetID uint `json:"preset_id"`

	// Code is the internal error code used for logic or identification.
	Code int `json:"code"`

	// UserMessage is a message intended to be shown to end users.
	UserMessage string `json:"user_message"`

	// DevMessage is a message intended for developers or debugging purposes.
	DevMessage string `json:"dev_message"`

	// LogMessage is the message that should be written to logs.
	LogMessage string `json:"log_message"`

	// Source identifies the origin or category of the error.
	Source *ErrorSource `json:"source"`

	// Level indicates the severity or type of the error.
	Level ErrorLevel `json:"level"`

	// HttpCode is the HTTP status code to return if this error is exposed via an API.
	HttpCode int `json:"http_code"`

	// ContinueExecution determines whether program execution should continue after this error.
	ContinueExecution bool `json:"continue_execution"`
}

type ErrorSource struct {
	Name       string
	ParseError func(error) *CustomError
	SML        *logging.SystemModuleLogger
}

func generateTraceId() string {
	length := 16                    // bytes
	bytes := make([]byte, length/2) // half the length because hex encoding doubles the size
	if _, err := rand.Read(bytes); err != nil {
		(&CustomError{CustomErrorPreset: FailedToGenerateTraceId, TraceId: "[TraceId failed]"}).Log()
		return "[TraceId failed]"
	}
	return hex.EncodeToString(bytes)
}

type CustomError struct {
	TraceId string
	CustomErrorPreset
}

func NewCustomError(preset CustomErrorPreset) *CustomError {
	return preset.New()
}

func (preset *CustomErrorPreset) New() *CustomError {
	if preset.PresetID == 0 {
		presetIDcounter++
		preset.PresetID = presetIDcounter
	}
	return &CustomError{CustomErrorPreset: *preset, TraceId: generateTraceId()}
}

func (err *CustomError) IsFromPreset(preset *CustomErrorPreset) bool {
	return err.PresetID == preset.PresetID && err.Source == preset.Source
}
func (err *CustomError) Is(target error) bool {
	preset, ok := target.(*CustomErrorPreset)
	if !ok {
		return false
	}
	return err.IsFromPreset(preset)
}
func NewError(err error) *CustomError {
	customError := CustomError{
		TraceId:           generateTraceId(),
		CustomErrorPreset: GenericInternalServerError,
	}
	customError.LogMessage = err.Error()
	return &customError
}
