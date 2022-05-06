package example

import "fmt"

type Logger struct{}

func (l *Logger) Log(message string) {
	fmt.Println(message)
}

func (l *Logger) Logf(format string, values ...interface{}) {
	fmt.Printf(format, values...)
}

func (l *Logger) Warn(message string) {
	l.Log(message)
}

func (l *Logger) Warnf(format string, values ...interface{}) {
	l.Logf(format, values...)
}

func (l *Logger) Fatal(message string) {
	l.Log(message)
}

func (l *Logger) Fatalf(format string, values ...interface{}) {
	l.Logf(format, values...)
}
