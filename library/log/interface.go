package log

type Level uint8

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
)

const DefaultLevel = INFO

type Logger interface {
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warning(string, ...interface{})
	Error(string, ...interface{})
	SetLevel(Level)
}

var (
	logInstance Logger
)

func init() {
	SetLogger(newGoLoggingLogger())
	SetLevel(DefaultLevel)
}

func SetLogger(logger Logger) {
	logInstance = logger
}

func Debug(message string, param ...interface{}) {
	logInstance.Debug(message, param...)
}

func Info(message string, param ...interface{}) {
	logInstance.Info(message, param...)
}

func Warning(message string, param ...interface{}) {
	logInstance.Warning(message, param...)
}

func Error(message string, param ...interface{}) {
	logInstance.Error(message, param...)
}

func SetLevel(level Level) {
	logInstance.SetLevel(level)
}
