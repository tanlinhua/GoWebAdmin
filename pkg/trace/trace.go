package trace

import (
	"time"

	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func Error(msg string) {
	writer(msg, "error")
}

func Debug(msg string) {
	writer(msg, "debug")
}

func Info(msg string) {
	writer(msg, "info")
}

func Warn(msg string) {
	writer(msg, "warn")
}

func writer(msg string, level string) {
	filePath := "runtime/" + level + "/"

	logWriter, _ := retalog.New(
		filePath+"%Y%m%d.log",
		retalog.WithMaxAge(7*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.AddHook(Hook)

	entry := logger.WithFields(logrus.Fields{
		"trace": msg,
	})

	switch level {
	case "debug":
		entry.Debug()
	case "error":
		entry.Error()
	case "info":
		entry.Info()
	case "warn":
		entry.Warn()
	default:
		entry.Warn()
	}
}
