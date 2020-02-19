package logger

import (
	"fmt"
)

// Level is linked to verbosity conf
const (
	NONE int = -2
	ERROR int = -1
	NOTICE int = 1
	INFO int = 2
	DEBUG int = 3
)
type Logger struct {
	VerbosityLevel int
}
func NewLogger(verbosity int) Logger {
	logger := Logger{}
	logger.VerbosityLevel = verbosity
	return logger
}
func(this Logger) Log(level int, message string) {
	if(this.VerbosityLevel >= level) {
		fmt.Println(message)
	}
}
func(this Logger) Error(message string) {
	this.Log(ERROR, message)
}
func(this Logger) Info(message string) {
	this.Log(INFO, message)
}
func(this Logger) Notice(message string) {
	this.Log(NOTICE, message)
}
func(this Logger) Debug(message string) {
	this.Log(DEBUG, message)
}
