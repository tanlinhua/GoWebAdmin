package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// uber-zap -> https://blog.csdn.net/weixin_43881017/article/details/111277435

// 自定义日志中间件
func Logger(tag string) gin.HandlerFunc {
	filePath := "runtime/log/" + tag + "/"
	// linkName := "latest_log_" + tag + ".log" // 最新日志的软链接

	logWriter, _ := retalog.New(
		filePath+"%Y%m%d.log",
		retalog.WithMaxAge(7*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
		// retalog.WithLinkName(linkName),
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

	return func(c *gin.Context) {
		startTime := time.Now()
		body, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) //读过的字节流重新放到body
		c.Next()

		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
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
			"HostName":    hostName,
			"RunTime":     spendTime,
			"IP":          clientIp,
			"ReqBody":     string(body),
			"ReqMethod":   method,
			"ReqURI":      path,
			"RspDataSize": dataSize,
			"RspStatus":   statusCode,
			"UA":          userAgent,
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
