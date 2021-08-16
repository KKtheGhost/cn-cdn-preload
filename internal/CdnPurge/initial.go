package CdnPurge

import (
	"fmt"
	"os"
	"time"
)

//全局基础变量
var LogFilePath, UrlFile string
var config FullConfig
var UrlList []string
var CurrentQuota CdnQuota

//全局结构体
type FullConfig struct {
	Aliyun  AliyunConfig
	Tencent TencentConfig
	LogDir  string
}

type AliyunConfig struct {
	AccessKey   string
	AccessToken string
}

type TencentConfig struct {
	SecretId  string
	SecretKey string
}

type CdnQuota struct {
	PreloadQuota  int //预热条目配额
	PreloadRemain int //剩余预热条目配额
	UrlQuota      int //刷新条目配额
	UrlRemain     int //剩余刷新条目配额
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//包体初始化
func init() {
	config = ConfigLoader()
	if Exists(config.LogDir) {
		fmt.Println(time.Now().Format("2006/01/02 15:04:05"), "Detect log file dir, start working.")
	} else {
		os.MkdirAll(config.LogDir, os.ModePerm)
	}
	LogerInit()
}
