package CdnPurge

import (
	"os"

	"gopkg.in/ini.v1"
)

func ConfigLoader() (config FullConfig) {
	var tmpConfig FullConfig
	raw_config, err := ini.Load("/etc/CdnPurge/config.ini")
	if err != nil {
		loger.Printf("Fail to read file: %v, please touch a config file according to README.md", err)
		os.Exit(1)
	}

	tmpConfig.LogDir = raw_config.Section("Common").Key("LogDir").String()
	tmpConfig.Aliyun.AccessKey = raw_config.Section("Aliyun").Key("AccessKey").String()
	tmpConfig.Aliyun.AccessToken = raw_config.Section("Aliyun").Key("AccessToken").String()
	tmpConfig.Tencent.SecretId = raw_config.Section("Tencent").Key("SecretId").String()
	tmpConfig.Tencent.SecretKey = raw_config.Section("Tencent").Key("SecretKey").String()

	return tmpConfig
}
