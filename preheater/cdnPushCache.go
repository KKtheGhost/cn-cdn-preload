package main

import (
	"CdnPurge"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	UrlFilePath        string
	CspPlatform        string
	HelpFlag           bool
	isOutput           bool
	Quota              int
	ObjectPaths        []string
	ObjectPathSliceAli string
	ObjectPathSliceTC  []string
	UrlCount           int
)

func init() {
	flag.StringVar(&UrlFilePath, "u", "", "The url file to purge.")
	flag.StringVar(&CspPlatform, "p", "aliyun", "The CSP platform to select.")
	flag.BoolVar(&isOutput, "o", false, "Whether to output the result to screen.")
	flag.BoolVar(&HelpFlag, "h", false, "CdnPushCache Help manual.")
}

func main() {
	flag.Parse()

	if HelpFlag {
		fmt.Println("Usage of /tmp/go-build224457796/b001/exe/cdnPushCache:\n  -o    Whether to output the result to screen.\n  -p string\n        The CSP platform to select. (default \"aliyun\")\n  -u string\n        The url file to purge.\n  -h    CdnPushCache Help manual.")
		os.Exit(0)
	}

	ObjectPaths = CdnPurge.ReadUrls(UrlFilePath)
	ObjectPathSliceAli = ""
	ObjectPathSliceTC = make([]string, 0)
	UrlCount = 0

	switch CspPlatform {
	case "aliyun":
		Quota = CdnPurge.AliyunDescribeRefreshQuota().PreloadRemain
		if Quota < len(ObjectPaths) {
			fmt.Println(time.Now().Format("2006/01/02 15:04:05"), "CDN push cache quota is less than target url, job aborted.")
			os.Exit(0)
		} else {
			fmt.Println(time.Now().Format("2006/01/02 15:04:05"), "CDN push cache quota is efficient:", Quota, ", continue >>>")
		}
		for _, ObjectUrl := range ObjectPaths {
			UrlCount = UrlCount + 1
			if UrlCount <= 9 {
				ObjectPathSliceAli = ObjectPathSliceAli + ObjectUrl + "\r\n"
			} else {
				ObjectPathSliceAli = ObjectPathSliceAli + ObjectUrl + "\r\n"

				CdnPurge.AliyunPushObject(ObjectPathSliceAli[:len(ObjectPathSliceAli)-2], isOutput)
				time.Sleep(3 * time.Second)
				//单次推送完毕后指标归零
				UrlCount = 0
				ObjectPathSliceAli = ""
			}
			//fmt.Println(UrlIndex,ObjectUrl)
		}
		//最后一轮请求。因为无法塞满UrlCount上线，需要最外面单独执行一次
		CdnPurge.AliyunPushObject(ObjectPathSliceAli[:len(ObjectPathSliceAli)-2], isOutput)
		time.Sleep(3 * time.Second)

	case "tencent":
		Quota = CdnPurge.TencentDescribeRefreshQuota().PreloadRemain
		if Quota < len(ObjectPaths) {
			fmt.Println(time.Now().Format("2006/01/02 15:04:05"), "CDN push cache quota is less than target url, job aborted.")
			os.Exit(0)
		} else {
			fmt.Println(time.Now().Format("2006/01/02 15:04:05"), "CDN push cache quota is efficient:", Quota, ", continue >>>")
		}
		for _, ObjectUrl := range ObjectPaths {
			UrlCount = UrlCount + 1
			if UrlCount <= 9 {
				ObjectPathSliceTC = append(ObjectPathSliceTC, ObjectUrl)
			} else {
				ObjectPathSliceTC = append(ObjectPathSliceTC, ObjectUrl)

				CdnPurge.TencentPushObject(ObjectPathSliceTC, isOutput)
				time.Sleep(3 * time.Second)
				//单次推送完毕后指标归零
				UrlCount = 0
				ObjectPathSliceTC = make([]string, 0)
			}
			//fmt.Println(UrlIndex,ObjectUrl)
		}
		//最后一轮请求。因为无法塞满UrlCount上线，需要最外面单独执行一次
		CdnPurge.TencentPushObject(ObjectPathSliceTC, isOutput)
		time.Sleep(3 * time.Second)
	default:
		fmt.Println(time.Now().Format("2006/01/02 15:04:05"), "Illegal Cloud Service Provider name, Aborted.")
	}

}
