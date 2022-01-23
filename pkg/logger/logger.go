package logger

import (
	"bytes"
	"io"
	"sync"
)
// Need to work for improvement // Not using !
// Define log levels
type Level int

// All possibale log levels [1 to 5]
const (
	LevelDebug Level = iota
	LevelInfo
	LevelError
	LevelProd
	LevelSecure
)

const defaultLogLevel = LevelDebug

type Logger struct {
	mu     sync.Mutex   // guards
	prefix string       // prefix
	Level  Level        // Log level
	w      io.Writer    // write
	buf    bytes.Buffer // internal buffer
}

// Creates a new logger
func New(w io.Writer, prefix string) *Logger {
	return &Logger{w: w, prefix: prefix, Level: defaultLogLevel}
}

var Console = New(os.Stderr, prefix: "CHI")


// LevelDebug
func (l *Logger) Debug(v ... interface{}) {
	l.w.WriteString(LevelDebug, fmt.Sprintln(v...))
}
//LevelInfo
func (l *Logger) Info(v ... interface{}) {
	if LevelInfo < l.Level {
		return
	}
	l.w.WriteString(LevelInfo, fmt.Sprintln(v...))
}
//LevelError
func (l *Logger) Error(v ... interface{}) {
	if LevelError < l.Level {
		return
	}
	l.w.WriteString(LevelError, fmt.Sprintln(v...))
}
//LevelProd
func (l *Logger) Prod(v ... interface{}) {
	if LevelProd < l.Level {
		return
	}
	l.w.WriteString(LevelProd, fmt.Sprintln(v...))
}
//LevelSecure
func (l *Logger) Secure(v ... interface{}) {
	if LevelSecure < l.Level {
		return
	}
	l.w.WriteString(LevelSecure, fmt.Sprintln(v...))
}
// WriteEntry writes the msg based on selected level
func (l *Logger) WriteEntry(lvl Level, msg string) error{
	l.w.write([]byte(msg))
	return nil
}