package logger

import (
	"github.com/sirupsen/logrus"
)

var entry *logrus.Entry

var fieldMap = logrus.FieldMap{
	logrus.FieldKeyTime:  ".timestamp",
	logrus.FieldKeyLevel: "@level",
	logrus.FieldKeyMsg:   "@message",
	logrus.FieldKeyFunc:  "z_caller",
}

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	if entry == nil {
		panic("logger not initialized")
	}

	return &Logger{entry}
}

func InitDefaultLogger(opts ...func(*logrus.Logger)) {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: fieldMap,
	}

	for _, opt := range opts {
		opt(log)
	}

	entry = logrus.NewEntry(log)
}
