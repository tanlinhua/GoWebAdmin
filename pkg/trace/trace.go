package trace

import (
	"time"

	rotate "github.com/lestrrat-go/file-rotatelogs"
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
	dir := "runtime/" + level + "/"
	p := dir + "%Y%m%d.log"

	logWriter, _ := rotate.New(
		p,
		rotate.WithMaxAge(7*24*time.Hour),
		rotate.WithRotationTime(24*time.Hour), // 日志切割时间
		rotate.WithRotationSize(10*1024*1024), // 日志切割大小M
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	// &logrus.TextFormatter
	hook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.AddHook(hook)

	// fuck new obj

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
