package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger *zerolog.Logger
}

func New(isDebug bool) *Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "time"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &Logger{logger: &logger}
}
func NewConsole(isDebug bool) *Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &Logger{logger: &logger}
}

// Output duplicates the global logger and sets w as its output.
func (l *Logger) Output(w io.Writer) zerolog.Logger {
	return l.logger.Output(w)
}
func (l *Logger) Level(level zerolog.Level) zerolog.Logger {
	return l.logger.Level(level)
}
func (l *Logger) Debug() *zerolog.Event {
	return l.logger.Debug()
}
func (l *Logger) Info() *zerolog.Event {
	return l.logger.Info()
}
func (l *Logger) Warn() *zerolog.Event {
	return l.logger.Warn()
}

func (l *Logger) Error() *zerolog.Event {
	return l.logger.Error()
}
func (l *Logger) Fatal() *zerolog.Event {
	return l.logger.Fatal()
}
func (l *Logger) Panic() *zerolog.Event {
	return l.logger.Panic()
}
