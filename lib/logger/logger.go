package logger

import (
	"fmt"

	"github.com/fatih/color"
)

var errorPrefix string
var infoPrefix string

func SetErrorPrefix(prefix string) {
	errorPrefix = prefix
}

func SetInfoPrefix(prefix string) {
	infoPrefix = prefix
}

func Error(args ...interface{}) {
	fmt.Print(errorPrefix)
	fmt.Print(args...)
}

func Errorln(args ...interface{}) {
	fmt.Print(errorPrefix)
	fmt.Println(args...)
}

func Errorf(message string, args ...interface{}) {
	fmt.Printf(
		"%s%s",
		errorPrefix,
		fmt.Sprintf(message, args...),
	)
}

func Info(args ...interface{}) {
	fmt.Print(infoPrefix)
	fmt.Print(args...)
}

func Infoln(args ...interface{}) {
	fmt.Print(infoPrefix)
	fmt.Println(args...)
}

func Infof(message string, args ...interface{}) {
	fmt.Printf(
		"%s%s",
		infoPrefix,
		fmt.Sprintf(message, args...),
	)
}

var BoldRed = color.New(color.FgRed, color.Bold)
var BoldCyan = color.New(color.FgCyan, color.Bold)
