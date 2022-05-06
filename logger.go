package emotebox

type Logger interface {
	Log(message string)
	Logf(format string, values ...interface{})
	Warn(message string)
	Warnf(format string, values ...interface{})
	Fatal(message string)
	Fatalf(format string, values ...interface{})
}
