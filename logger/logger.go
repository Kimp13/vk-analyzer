package logger

import (
	"fmt"
)

type Logger struct {
	errorPrefix string
	infoPrefix  string
}

func (l *Logger) SetErrorPrefix(prefix string) {
	l.errorPrefix = prefix
}

func (l *Logger) SetInfoPrefix(prefix string) {
	l.infoPrefix = prefix
}

func (l *Logger) Error(args ...interface{}) {
	fmt.Print(l.errorPrefix)
	fmt.Print(args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	fmt.Print(l.errorPrefix)
	fmt.Println(args...)
}

func (l *Logger) Errorf(message string, args ...interface{}) {
	fmt.Printf(
		"%s%s",
		l.errorPrefix,
		fmt.Sprintf(message, args...),
	)
}

func (l *Logger) Info(args ...interface{}) {
	fmt.Print(l.infoPrefix)
	fmt.Print(args...)
}

func (l *Logger) Infoln(args ...interface{}) {
	fmt.Print(l.infoPrefix)
	fmt.Println(args...)
}

func (l *Logger) Infof(message string, args ...interface{}) {
	fmt.Printf(
		"%s%s",
		l.infoPrefix,
		fmt.Sprintf(message, args...),
	)
}

func New() *Logger {
	return &Logger{
		"",
		"",
	}
}
