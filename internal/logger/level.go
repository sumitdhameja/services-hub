package logger

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

// These are the different logging levels.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the message passed to Debug, Info, ...
	PanicLevel log.Level = iota
	// FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
)

// LevelFromString convert string to Log Level. Default is Debug.
func LevelFromString(level string) log.Level {
	l := InfoLevel
	switch strings.ToLower(level) {
	case "panic":
		l = PanicLevel
	case "fatal":
		l = FatalLevel
	case "error":
		l = ErrorLevel
	case "warn":
		l = WarnLevel
	case "info":
		l = InfoLevel
	case "debug":
		l = DebugLevel
	default:
		l = InfoLevel
	}
	return l
}
