package logger

import (
	"github.com/sirupsen/logrus"
)

// LogrusLog implements logging using logrus API
func LogrusLog(logMessage string, logLevel LogLevel){
	switch logLevel {
	case Debug:
		logrus.Debug(logMessage)
	case Info:
		logrus.Info(logMessage)
	case Warning:
		logrus.Warning(logMessage)
	case Error:
		logrus.Error(logMessage)
	case Fatal:
		logrus.Fatal(logMessage)
	}
}

// SetLogLevel sets logrus global logging to a specific log level
func SetLogLevel(logLevel string){
	ll,err := logrus.ParseLevel(logLevel)
	if err != nil{
		ll = logrus.DebugLevel
	}
	// set global log level
	logrus.SetLevel(ll)
}
