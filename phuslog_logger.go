package phuslog_logger

import (
	"reflect"

	"github.com/goccy/go-json"
	"github.com/phuslu/log"
)

type Option func(l *log.Logger)

func WithLogger(logs *log.Logger) Option {
	return func(l *log.Logger) {
		l.Level = logs.Level
		l.Caller = logs.Caller
		l.TimeField = logs.TimeField
		l.TimeFormat = logs.TimeFormat
		l.Writer = logs.Writer
		l.Context = logs.Context
	}
}

func WithConsoleLogger() Option {
	return func(l *log.Logger) {
		l.Level = log.InfoLevel
		l.Caller = 2
		l.TimeField = ""
		l.TimeFormat = ""
		l.Writer = &log.ConsoleWriter{}
	}
}

type phuslogLogger struct {
	lg log.Logger
}

func logs(entry *log.Entry, msg string, keysAndValues []interface{}) {
	if len(keysAndValues) > 1 {
		key := keysAndValues[0]
		value := keysAndValues[1]
		switch reflect.TypeOf(value).Kind() {
		case reflect.Struct:
		case reflect.Ptr:
			data, err := json.Marshal(value)
			if err == nil {
				entry.Str(key.(string), string(data)).Msg(msg)
			}
		default:
			entry.KeysAndValues(keysAndValues).Msg(msg)
		}
	} else {
		entry.KeysAndValues(keysAndValues).Msg(msg)
	}
}

func (p *phuslogLogger) Debug(msg string, keysAndValues ...interface{}) {
	p.lg.Debug().KeysAndValues(keysAndValues).Msg(msg)
}

func (p *phuslogLogger) Info(msg string, keysAndValues ...interface{}) {
	logs(p.lg.Info(), msg, keysAndValues)
}

func (p *phuslogLogger) Warn(msg string, keysAndValues ...interface{}) {
	logs(p.lg.Warn(), msg, keysAndValues)
}

func (p *phuslogLogger) Error(msg string, keysAndValues ...interface{}) {
	logs(p.lg.Error(), msg, keysAndValues)
}

func (p *phuslogLogger) Fatal(msg string, keysAndValues ...interface{}) {
	logs(p.lg.Fatal(), msg, keysAndValues)
}

func NewPhuslogLogger(options ...Option) Logger {
	dftLogger := log.Logger{}

	for _, option := range options {
		option(&dftLogger)
	}

	if len(options) == 0 {
		dftLogger = log.Logger{
			Level:      log.InfoLevel,
			Caller:     2,
			TimeField:  "",
			TimeFormat: "",
			Writer:     &log.ConsoleWriter{},
		}
	}

	log.DefaultLogger = dftLogger

	return &phuslogLogger{
		lg: log.DefaultLogger,
	}
}
