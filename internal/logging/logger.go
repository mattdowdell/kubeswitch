package logging

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

// New creates a new logger with the given verbosity.
func New(verbosity int) *log.Logger {
	return log.NewWithOptions(os.Stderr, log.Options{
		Level:           verbosityToLevel(verbosity),
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})
}

// verbosityToLevel converts a verbosity integer to a log level.
//
// A verbosity of 0 corresponds to an Info level. Any other value corresponds to a Debug level.
func verbosityToLevel(verbosity int) log.Level {
	switch verbosity {
	case 0:
		return log.InfoLevel

	default:
		return log.DebugLevel
	}
}
