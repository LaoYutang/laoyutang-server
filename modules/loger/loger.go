package loger

import (
	"log"
	"os"
	"time"
)

// 初始化方法
func init() {
	mkdirErr := os.MkdirAll("./logs", os.ModePerm)
	if mkdirErr != nil {
		panic(mkdirErr)
	}
	file := "./logs/" + time.Now().Format("20060102") + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[sys]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}
