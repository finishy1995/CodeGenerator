package log

import (
	"github.com/op/go-logging"
	"os"
)

type goLoggingLogger struct {
	log *logging.Logger
}

const (
	ModuleName = "CodeGenerator"
)

var (
	levelMap = map[Level]logging.Level{
		DEBUG:   logging.DEBUG,
		INFO:    logging.INFO,
		WARNING: logging.WARNING,
		ERROR:   logging.ERROR,
	}
)

func newGoLoggingLogger() *goLoggingLogger {
	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} [%{level:.4s}] ▶ %{color:reset} %{message}`,
	)
	backend := logging.NewBackendFormatter(logging.NewLogBackend(os.Stdout, "", 0), format)
	log := &goLoggingLogger{
		log: logging.MustGetLogger(ModuleName),
	}
	logging.SetBackend(backend)

	return log
}

func (g *goLoggingLogger) Debug(s string, i ...interface{}) {
	g.log.Debugf(s, i...)
}

func (g *goLoggingLogger) Info(s string, i ...interface{}) {
	g.log.Infof(s, i...)
}

func (g *goLoggingLogger) Warning(s string, i ...interface{}) {
	g.log.Warningf(s, i...)
}

func (g *goLoggingLogger) Error(s string, i ...interface{}) {
	g.log.Errorf(s, i...)
}

func (g *goLoggingLogger) SetLevel(level Level) {
	if l, ok := levelMap[level]; ok {
		logging.SetLevel(l, ModuleName)
	} else {
		logging.SetLevel(levelMap[DefaultLevel], ModuleName)
	}
}
