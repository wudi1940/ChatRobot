package client

import (
	"ChatRobot/cmd/config"
	"ChatRobot/cmd/utils"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

//https://www.cnblogs.com/xinliangcoder/p/11212573.html

// Create a new instance of the logger. You can have any number of instances.
var logClient *logrus.Logger

func InitLogClient() {
	logClient = createLog(config.InfoLogFileName)
}

func GetLogClient() *logrus.Logger {
	if logClient == nil {
		panic(errors.New("LogClient未初始化"))
	}
	return logClient
}

func LoggerToFile() gin.HandlerFunc {
	ginLog := createLog(config.GinLogFileName)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		//请求方式
		method := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()
		// 打印日志
		ginLog.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   method,
			"req_uri":      reqUrl,
		}).Info()
	}
}

func createLog(fileName string) *logrus.Logger {
	// 写入文件
	src, err := utils.OpenFile(fileName)
	if err != nil {
		fmt.Println("err", err)
	}
	// 实例化
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.Out = src // 设置输出
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 设置 rotatelogs
	loggerWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic(err)
	}

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  loggerWriter,
		logrus.FatalLevel: loggerWriter,
		logrus.DebugLevel: loggerWriter,
		logrus.WarnLevel:  loggerWriter,
		logrus.ErrorLevel: loggerWriter,
		logrus.PanicLevel: loggerWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	logger.AddHook(lfHook)

	return logger
}
