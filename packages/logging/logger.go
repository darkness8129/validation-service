package logging

// Logger provides a logic for logging throughout the app.
// Fatal method causes the app to exit with "1" status code.
type Logger interface {
	Named(name string) Logger
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
}
