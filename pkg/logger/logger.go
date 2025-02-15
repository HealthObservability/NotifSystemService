package logger

import "github.com/sirupsen/logrus"

var Logger *logger

type logger struct {
	*logrus.Entry
}

func InitLogger(isLocal, isDebug *bool) {
	fieldMap := logrus.FieldMap{
		logrus.FieldKeyTime:  ".timestamp",
		logrus.FieldKeyLevel: "@level",
		logrus.FieldKeyMsg:   "@message",
		logrus.FieldKeyFunc:  "z_caller",
	}

	Logger = &logger{}
	Logger.Entry = logrus.NewEntry(logrus.New())

	Logger.Logger.SetLevel(logrus.InfoLevel)
	Logger.Logger.SetFormatter(&logrus.JSONFormatter{
		//TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: fieldMap,
	})

	if isLocal != nil && *isLocal {
		logInit := logrus.New()
		logInit.Formatter = &logrus.TextFormatter{
			ForceColors:      true,
			DisableTimestamp: false,
			FullTimestamp:    true,
			FieldMap:         fieldMap,
		}
		newEntry := logrus.NewEntry(logInit)
		Logger.Entry = newEntry
	}

	if isDebug != nil && *isDebug || isDebug == nil {
		Logger.Logger.SetLevel(logrus.DebugLevel)
	}
}
