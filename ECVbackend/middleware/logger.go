package middleware

import (
	"ecvbackend/model"
	"ecvbackend/pkg"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var GlobalCount int64
var ListenPort string

// 日志记录到文件
func AccessLoggerToFileAndDB() gin.HandlerFunc {
	sec, err := setting.Cfg.GetSection("log")
	fileName := sec.Key("ACCESS_LOG").String()
	src, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("err", err)
	}
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{})
	return func(c *gin.Context) {
		GlobalCount = GlobalCount + 1
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		// 日志格式
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIp,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
			"port":         ListenPort,
		}).Info()
		model.Visit{
			StartTime:   startTime,
			EndTime:     endTime,
			LatencyTime: latencyTime,
			ReqMethod:   reqMethod,
			ReqUri:      reqUri,
			StatusCode:  statusCode,
			ClientIp:    clientIp,
		}.Insert()
	}
}
