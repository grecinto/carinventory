package logger

// LogLevel enumerates supported log levels
type LogLevel int
const(
	Debug LogLevel = iota
	Info
	Warning
	Error
	Fatal
)

// LoggerFunc defines a function for implementing logging
type LoggerFunc func(logMessage string, logLevel LogLevel)
