package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

// global app logger
var (
	level  log.Level
	logger *log.Logger
)

func Println(args ...interface{}) {
	logger.Println(args...)
}

// Panic App-level messages
func Panic(args ...interface{}) {
	if level >= PanicLevel {
		logger.Panic(args...)
	}
}

// Fatal App-level messages
func Fatal(args ...interface{}) {
	if level >= FatalLevel {
		logger.Fatal(args...)
	}
}

// Error App-level messages
func Error(args ...interface{}) {
	if level >= ErrorLevel {
		logger.Errorln(args...)
	}
}

// Warn App-level messages
func Warn(args ...interface{}) {
	if level >= WarnLevel {
		logger.Warnln(args...)
	}
}

// Info App-level messages
func Info(args ...interface{}) {
	if level >= InfoLevel {
		logger.Infoln(args...)
	}
}

// Debug App-level messages
func Debug(args ...interface{}) {
	if level >= DebugLevel {
		logger.Debugln(args...)
	}
}

func InitializeWithWriter(l log.Level, out io.Writer) {
	logger = log.New()
	level = l
	logger.SetLevel(l)
	// logger.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	logger.SetOutput(out)

}

// Initialize initializes the logger object
func Initialize(l log.Level) {
	InitializeWithWriter(l, os.Stdout)
}
