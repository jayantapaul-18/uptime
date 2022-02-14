package logger

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/gookit/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	logger *zerolog.Logger
}

func New() *Logger {
	LOGLEVEL, _ := config.Bool("logLevel")
	logLevel := zerolog.InfoLevel
	if LOGLEVEL {
		logLevel = zerolog.DebugLevel
	}
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// zerolog.TimestampFieldName = "time"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"
	zerolog.SetGlobalLevel(logLevel)
	// File Output to /logs/commonlog
	logFile, err := ioutil.TempFile("./logs", "commonlog")
	if err != nil {
		log.Error().Err(err).Msg("there was an error creating a logFile four log")
	}
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	multi := zerolog.MultiLevelWriter(consoleWriter, logFile)
	logger := zerolog.New(multi).With().Timestamp().Logger()
	//logger := zerolog.New(logFile).With().Timestamp().Logger()
	// logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger.Info().Msg("logging configured")
	return &Logger{logger: &logger}
}

// Need to Implement log rotation policy

// func New(isDebug bool) *Logger {
// 	logLevel := zerolog.InfoLevel
// 	if isDebug {
// 		logLevel = zerolog.DebugLevel
// 	}
// 	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
// 	zerolog.TimestampFieldName = "time"
// 	zerolog.LevelFieldName = "level"
// 	zerolog.MessageFieldName = "msg"
// 	zerolog.SetGlobalLevel(logLevel)
// 	// File Output to /logs/commonlog
// 	logFile, err := ioutil.TempFile("./logs", "commonlog")
// 	if err != nil {
// 		log.Error().Err(err).Msg("there was an error creating a logFile four log")
// 	}
// 	logger := zerolog.New(logFile).With().Timestamp().Logger()
// 	// logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
// 	return &Logger{logger: &logger}
// }

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
