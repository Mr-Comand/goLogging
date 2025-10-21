package errorhandling

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/Mr-Comand/goLogging/logging"
)

type HtmlError struct {
	UserMessage string `json:"userMessage"`
	DevMessage  string `json:"devMessage"`
	TraceId     string `json:"TraceId"`
}

// Implement the Error() method for CustomError to satisfy the error interface
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.LogMessage)
}
func (ep *CustomErrorPreset) Error() string {
	return fmt.Sprintf("ErrorPreset %d: %s", ep.Code, ep.LogMessage)
}
func (e *CustomError) HTML() (HtmlError, int) {
	return HtmlError{
		UserMessage: e.UserMessage,
		DevMessage:  e.DevMessage,
		TraceId:     e.TraceId,
	}, e.HttpCode
}
func (e *CustomError) Log() *CustomError {
	var sml logging.LoggerInterface

	if e.Source.SML != nil {
		sml = e.Source.SML
	} else {
		sml = std.logger.GetSystemModule(e.Source.Name)
		if sml == nil || reflect.ValueOf(sml).IsNil() {
			sml = std.logger
		}
	}

	switch e.Level {
	case ErrorWARN:
		sml.Warn(fmt.Sprintf("{trc-%s}\t%s", e.TraceId, e.LogMessage))
	case ErrorWrongUsage:
		sml.Debug(fmt.Sprintf("{trc-%s}\t%s", e.TraceId, e.LogMessage))
	case ErrorMedium:
		sml.Error(fmt.Sprintf("{trc-%s}\t%s", e.TraceId, e.LogMessage))
	case ErrorFail:
		sml.Fail(fmt.Sprintf("{trc-%s}\t%s", e.TraceId, e.LogMessage))
	default:
		sml.Error(fmt.Sprintf("{trc-%s}\t%s", e.TraceId, e.LogMessage))
	}
	return e
}
func (e *CustomError) HandelWeb(w http.ResponseWriter, r *http.Request) bool {
	if !e.ContinueExecution {
		e.Log()

		html, HttpCode := e.HTML()
		w.Header().Set("X-Trace-ID", e.TraceId) // add Trace ID to response header
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(HttpCode)
		_ = json.NewEncoder(w).Encode(html)
		return false
	}
	w.Header().Set("X-Trace-ID", e.TraceId) // add Trace ID to response header
	return true
}
func (e *CustomError) HandelWebExit(w http.ResponseWriter, r *http.Request) *CustomError {
	e.Log()

	html, HttpCode := e.HTML()
	w.Header().Set("X-Error-Trace-ID", e.TraceId) // add Trace ID to response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(HttpCode)
	_ = json.NewEncoder(w).Encode(html)
	return e

}
