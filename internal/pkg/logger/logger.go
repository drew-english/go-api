// Package logger is used for logging. The default one pushes structured (JSON) logs.
package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"time"
)

const (
	// LogTypeInfo is for logging type 'info'
	LogTypeInfo = "info"
	// LogTypeWarn is for logging type 'warn'
	LogTypeWarn = "warn"
	// LogTypeError is for logging type 'error'
	LogTypeError = "error"
	// LogTypeFatal is for logging type 'fatal'
	LogTypeFatal = "fatal"
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

// LogHandler implements Logger
type LogHandler struct {
	Skipstack  int
	appName    string
	appVersion string
}

func (lh *LogHandler) defaultPayload(severity string) map[string]interface{} {
	_, file, line, _ := runtime.Caller(lh.Skipstack)
	return map[string]interface{}{
		"app":        lh.appName,
		"appVersion": lh.appVersion,
		"severity":   severity,
		"line":       fmt.Sprintf("%s:%d", file, line),
		"time":       time.Now().UTC(),
	}
}

func (lh *LogHandler) serialize(severity string, data ...interface{}) (string, error) {
	payload := lh.defaultPayload(severity)
	for idx, value := range data {
		payload[fmt.Sprintf("%d", idx)] = value
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (lh *LogHandler) serializeWithExtra(severity string, extraTags map[string]interface{}, data ...interface{}) (string, error) {
	payload := lh.defaultPayload(severity)

	for k, v := range extraTags {
		payload[k] = v
	}

	for idx, value := range data {
		payload[fmt.Sprintf("%d", idx)] = value
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func handleFatal(severity, line string) {
	switch severity {
	case "fatal":
		{
			fmt.Println(line)
			os.Exit(1)
		}
	}
}

func (lh *LogHandler) log(severity string, payload ...interface{}) error {
	out, err := lh.serialize(severity, payload...)
	if err != nil {
		return err
	}

	handleFatal(severity, out)
	fmt.Println(out)

	return nil
}

func (lh *LogHandler) logWithExtra(severity string, extraTags map[string]interface{}, payload ...interface{}) error {
	out, err := lh.serializeWithExtra(severity, extraTags, payload...)
	if err != nil {
		return err
	}

	handleFatal(severity, out)
	fmt.Println(out)

	return nil
}

// Info is for logging items with severity 'info'
func (lh *LogHandler) Info(payload ...interface{}) error {
	return lh.log(LogTypeInfo, payload...)
}

// Warn is for logging items with severity 'Warn'
func (lh *LogHandler) Warn(payload ...interface{}) error {
	return lh.log(LogTypeWarn, payload...)
}

// Error is for logging items with severity 'Error'
func (lh *LogHandler) Error(payload ...interface{}) error {
	return lh.log(LogTypeError, payload...)
}

// Fatal is for logging items with severity 'Fatal'
func (lh *LogHandler) Fatal(payload ...interface{}) error {
	return lh.log(LogTypeFatal, payload...)
}

func (lh *LogHandler) InfoWithExtra(extraTags map[string]interface{}, payload ...interface{}) error {
	return lh.logWithExtra(LogTypeInfo, extraTags, payload...)
}

func (lh *LogHandler) ErrorWithExtra(extraTags map[string]interface{}, payload ...interface{}) error {
	return lh.logWithExtra(LogTypeError, extraTags, payload...)
}


// New returns a new instance of LogHandler
func New(appname string, appversion string, skipStack uint) *LogHandler {
	if skipStack <= 1 {
		skipStack = 4
	}

	return &LogHandler{
		Skipstack:  int(skipStack),
		appName:    appname,
		appVersion: appversion,
	}
}
