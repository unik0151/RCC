package middleware

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"miniapp/conf"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

/**
日志三件套
logrus 可以格式化日志，输出多个等级
lfshook 添加钩子函数
rotatelogs
*/

func LoggerTOFile() gin.HandlerFunc {

	logFileName := conf.LOG_NAME
	logFilePath := conf.LOG_FILE_PATH
	if _, err := os.Stat(logFileName); err != nil {
		os.Mkdir(logFilePath, os.ModePerm)
	}
	fileName := filepath.Join(logFilePath, logFileName)

	//写入日志文件
	//dist, errs := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	//if errs != nil {
	//	panic(errs)
	//}
	logger := logrus.New()

	//logger.Out = dist
	logger.SetLevel(logrus.InfoLevel)

	//用于日志分割
	// rotate: failed to create new symlink: symlink miniapp.log.20250402.log log\miniapp.log.20250402.log_symlink: A required privilege is not held by the client
	// window会出现上面报错，因为计算机账户只能管理员有权限创建软连接
	logWriter, errs := rotatelogs.New(
		fileName+".%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),
		//最多保留七天
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		//最大保存树
		//rotatelogs.WithRotationCount(5),
	)
	if errs != nil {
		panic(errs)
	}

	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.PanicLevel: logWriter,
		logrus.DebugLevel: logWriter,
	}
	//格式化日志
	lfsHook := lfshook.NewHook(writerMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(lfsHook)
	return func(c *gin.Context) {
		startTime := time.Now()
		//处理请求
		c.Next()
		endTime := time.Now()

		sub := endTime.Sub(startTime)

		reqMethod := c.Request.Method
		// /user/get
		reqUri := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求客户端ip
		clientIP := c.ClientIP()
		//请求客户端 浏览器类型
		clientUserAgent := c.Request.UserAgent()

		logger.WithFields(logrus.Fields{
			"req_start_time":    startTime.Format("2006-01-02 15:04:05"),
			"client_ip":         clientIP,
			"client_user_agent": clientUserAgent,
			"req_method":        reqMethod,
			"sub":               sub.String(),
			"req_uri":           reqUri,
			"status_code":       statusCode,
		}).Info()

	}

}
