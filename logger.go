package later

import (
	"fmt"
	"os"
	"time"

	"github.com/hippoai/env"
)

type logger struct {
	Verbose bool
}

// newLogger instanciates
func newLogger(verbose bool) *logger {
	return &logger{
		Verbose: verbose,
	}
}

// newLoggerFromEnv uses environment variable to define the logger
// By default it's verbose
func newLoggerFromEnv() *logger {
	parsed, err := env.Parse(Env_Verbosity)
	if (err != nil) || (parsed[Env_Verbosity] != "0") {
		return newLogger(true)
	}

	return newLogger(false)
}

// ForceLog for important messages
func (logger *logger) ForceLog(pattern string, itfs ...interface{}) {
	fullPattern := fmt.Sprintf("[%s] %s \n",
		time.Now().UTC().Format(time.RFC3339Nano),
		pattern,
	)

	fmt.Fprintf(
		os.Stderr,
		fullPattern,
		itfs...,
	)
}

// Log only if verbosity mode is turned on
func (logger *logger) Log(pattern string, itfs ...interface{}) {
	if !logger.Verbose {
		return
	}

	logger.ForceLog(pattern, itfs...)
}
