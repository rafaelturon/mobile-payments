package structuredLog

import (
	stdlog "log"
	"os"
	"strings"
	"time"

	kitlog "github.com/go-kit/kit/log"
)

var (
	logger      = kitlog.NewLogfmtLogger(os.Stdout)
	logLevel    = "debug"
	logFileName = "decred-pi-wallet.log"
)

func init() {
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)
	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger))
}

func setupLog() {
	lf, err := NewLogFile(logFileName, os.Stderr)
	if err != nil {
		stdlog.Fatal("Unable to create log file: ", err)
	}
	stdlog.SetOutput(lf)
	// rotate log every 30 seconds.
	rotateLogSignal := time.Tick(30 * time.Second)
	go func() {
		for {
			<-rotateLogSignal
			if err := lf.Rotate(); err != nil {
				stdlog.Fatal("Unable to rotate log: ", err)
			}
		}
	}()
}

func getLevel(currentLevel string) int {
	switch currentLevel {
	case "debug":
		return 0
	case "info":
		return 1
	case "warn":
		return 2
	case "error":
		return 3
	}

	return 0
}

func canLog(selectedLevel string, initLevel string) bool {
	if getLevel(selectedLevel) >= getLevel(initLevel) {
		return true
	}

	return false
}

// Setup of logging
// Usage: structuredLog.Setup("debug", "log-test.log")
// structuredLog.Debug("Debugging...")
func Setup(level, fileName string) {
	logLevel = level
	logFileName = fileName
	//setupLog()
}

// Debug designates fine-grained informational events
// that are most useful to debug an application.
func Debug(message string) {
	logMessage(message, "debug")
}

// Info designates informational messages
// that highlight the progress of the application at coarse-grained level.
func Info(message string) {
	logMessage(message, "info")
}

// Warn designates potentially harmful situations.
func Warn(message string) {
	logMessage(message, "warn")
}

// Error designates error events
// that might still allow the application to continue running
func Error(message string) {
	logMessage(message, "error")
}

func logMessage(logMessage string, level string) {
	logger = kitlog.With(logger, "caller", kitlog.Caller(5))
	if canLog(level, logLevel) {
		logger.Log("level", level, "msg", logMessage, "minLog", logLevel)
		stdlog.Printf("[%s] %s", strings.ToUpper(level), logMessage)
	}
	// Output:
	// ts=2018-01-08T00:55:42.190345655Z caller=dcrd.go:279 level=info msg=logMessage
}
