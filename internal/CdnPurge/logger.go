package CdnPurge

import (
	"log"
	"os"
	//"time"
)

var loger *log.Logger

//初始化一个loger方法用于log输出。
func LogerInit() {
	LogFilePath = config.LogDir + "/" + "cdn.log"
	logFile, err := os.OpenFile(LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	loger = log.New(logFile, "[Cloud_CDN_Purge] ", log.LstdFlags|log.Lshortfile|log.LUTC)
	// 将文件设置为loger作为输出
	return
}
