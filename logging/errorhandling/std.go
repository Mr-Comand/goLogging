package errorhandling

import "github.com/Mr-Comand/logging/logging"

func UpdateLogger(logger *logging.Logger) {
	std.UpdateLogger(logger)
}
func RegisterErrorSource(sources ...*ErrorSource) *ErrorHandler {
	return std.RegisterErrorSource(sources...)
}
func UnregisterErrorSource(source *ErrorSource) *ErrorHandler {
	return std.UnregisterErrorSource(source)
}
func Parse(err error) *CustomError {
	return std.Parse(err)
}
func ParseWith(err error, ErrorSources ...ErrorSource) *CustomError {
	return std.ParseWith(err, ErrorSources...)
}
func GetErrorSources(ErrorSourceNames ...string) []*ErrorSource {
	return std.GetErrorSources(ErrorSourceNames...)
}
