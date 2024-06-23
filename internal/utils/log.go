package utils

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/fatih/color"
)

const (
	LevelError = iota
	LevelWarning
	LevelInformation
	LevelDebug
)

var Level = LevelDebug
var logger *Logger

type Logger struct {
	level  int
	mu     sync.Mutex
	Writer io.Writer
}

var colors = map[string]func(a ...interface{}) string{
	"warn":  color.New(color.FgYellow).Add(color.Bold).SprintFunc(),
	"panic": color.New(color.BgRed).Add(color.Bold).SprintFunc(),
	"error": color.New(color.FgRed).Add(color.Bold).SprintFunc(),
	"info":  color.New(color.FgCyan).Add(color.Bold).SprintFunc(),
	"debug": color.New(color.FgWhite).Add(color.Bold).SprintFunc(),
}

func (l *Logger) println(prefix, msg string) {
	c := color.New()

	l.mu.Lock()
	defer l.mu.Unlock()

	c.SetWriter(l.Writer)
	c.Printf(
		"%s %s %s \n",
		colors[prefix]("["+prefix+"]"),
		time.Now().Format("2006/01/02 15:04:05"),
		msg,
	)
}

func (l *Logger) Panic(format string, v ...interface{}) {
	if LevelError > l.level {
		return
	}

	msg := fmt.Sprintf(format, v...)
	l.println("panic", msg)
	panic(msg)
}

func (l *Logger) Error(format string, v ...interface{}) {
	if LevelError > l.level {
		return
	}

	l.println("error", fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(format string, v ...interface{}) {
	if LevelWarning > l.level {
		return
	}

	l.println("warn", fmt.Sprintf(format, v...))
}

func (l *Logger) Info(format string, v ...interface{}) {
	if LevelInformation > l.level {
		return
	}

	l.println("info", fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(format string, v ...interface{}) {
	if LevelDebug > l.level {
		return
	}

	l.println("debug", fmt.Sprintf(format, v...))
}

// Returns the logger object.
func Log() *Logger {
	if logger == nil {
		logger = &Logger{
			level: Level,
		}
	} else {
		logger.level = Level
	}
	return logger
}
