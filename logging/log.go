package logging

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// LogLevel represents the log level.
type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

// SetLogLevel sets the global log level for zerolog.
func SetLogLevel(level LogLevel) {
	switch level {
	case Debug:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case Info:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case Warning:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case Error:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}
}

// ColoredConsoleWriter creates a console writer with colored output.
func ColoredConsoleWriter() zerolog.ConsoleWriter {
	return zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05", NoColor: false}
}

// InitLogger initializes the logger with the specified log level and output.
func InitLogger(level LogLevel) {
	SetLogLevel(level)
	log.Logger = log.Output(ColoredConsoleWriter())
}

// Log prints a log message with the specified level.
func Log(level LogLevel, message string) {
	switch level {
	case Debug:
		log.Debug().Msg(message)
	case Info:
		log.Info().Msg(message)
	case Warning:
		log.Warn().Msg(message)
	case Error:
		log.Error().Msg(message)
	}
}
