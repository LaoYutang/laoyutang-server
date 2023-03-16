package logger

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// 初始化方法
func init() {
	// 创建日志文件夹
	mkdirErr := os.MkdirAll("./logs", os.ModePerm)
	if mkdirErr != nil {
		panic(mkdirErr)
	}

	// 获取文件writer
	file := "./logs/" + time.Now().Format("20060102") + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}

	// 设置logrus
	logrus.SetOutput(io.MultiWriter(os.Stdout, logFile))
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05", PrettyPrint: true})
	logrus.SetReportCaller(false)
}
