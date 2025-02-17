package logging

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

// ...
func New(verbosity int) *log.Logger {
	return log.NewWithOptions(os.Stderr, log.Options{
		Level:           verbosityToLevel(verbosity),
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})
}

// ...
func verbosityToLevel(verbosity int) log.Level {
	switch verbosity {
	case 0:
		return log.InfoLevel

	default:
		return log.DebugLevel
	}
}
