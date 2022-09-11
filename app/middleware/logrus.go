package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"time"

	"github.com/gin-gonic/gin"
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// 自定义日志中间件
func Logger(tag string) gin.HandlerFunc {
	dir := "runtime/log/" + tag + "/"
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

	return func(c *gin.Context) {
		startTime := time.Now()
		body, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) //读过的字节流重新放到body
		c.Next()

		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"path":      path,
			"method":    method,
			"body":      string(body),
			"run_time":  spendTime,
			"ip":        clientIp,
			"resp_size": dataSize,
			"resp_code": statusCode,
			"ua":        userAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
