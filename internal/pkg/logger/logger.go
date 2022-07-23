// Package logger is used for logging. The default one pushes structured (JSON) logs.
package logger

import (
	"go-api/internal/pkg/env"
	"os"

	"github.com/sirupsen/logrus"
)

var std = New(
	env.GetEnvDefault("APP_NAME", "go-api"),
	env.GetEnvDefault("APP_VERSION", "Unknown Version"),
)

// Logger interface defines all the logging methods to be implemented
type Logger interface {
	Info(payload ...interface{}) error
	Warn(payload ...interface{}) error
	Error(payload ...interface{}) error
	Fatal(payload ...interface{}) error
	InfoWithExtra(extraTags map[string]interface{}, payload ...interface{}) error
	ErrorWithExtra(extraTags map[string]interface{}, payload ...interface{}) error
}

// Info is for logging items with severity 'info'
func Info(payload ...interface{}) {
	std.Info(payload...)
}

// Warn is for logging items with severity 'Warn'
func Warn(payload ...interface{}) {
	std.Warn(payload...)
}

// Error is for logging items with severity 'Error'
func Error(payload ...interface{}) {
	std.Error(payload...)
}

// Fatal is for logging items with severity 'Fatal'
func Fatal(payload ...interface{}) {
	std.Fatal(payload...)
}

func InfoWithExtra(extraTags map[string]interface{}, payload ...interface{}) {
	std.WithFields(extraTags).Info(payload...)
}

func ErrorWithExtra(extraTags map[string]interface{}, payload ...interface{}) {
	std.WithFields(extraTags).Error(payload...)
}

// New returns a new instance of Logrus logger
func New(appname string, appversion string) *logrus.Entry {
	log := logrus.New()
	initLogger(log)
	
	return log.WithFields(logrus.Fields {
		"appName": appname,
		"appVersion": appversion,
	})
}

func initLogger(log *logrus.Logger) {
	log.Out = os.Stdout

	if GO_ENV := env.GetEnvDefault("GO_ENV", "development"); GO_ENV != "development" {
		log.Formatter = &logrus.JSONFormatter{}

		// file, err := os.OpenFile("./log/" + GO_ENV + ".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

		// if err == nil {
		// 	log.Out = file
		// 	log.Formatter = &logrus.JSONFormatter{}
		// } else {
		// 	log.Info("Failed to log to file, using default stdout")
		// }
	}
}
