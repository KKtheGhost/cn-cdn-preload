package CdnPurge

import (
	//"fmt"
	"os"
	"strconv"
	cdn20180510  "github.com/alibabacloud-go/cdn-20180510/client"
	openapi  "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
    cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
)


/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */

/** 阿里云CDN配额获取部分  */
func AliyunCreateClient (accessKeyId *string, accessKeySecret *string) (_result *cdn20180510.Client, _err error) {
  config := &openapi.Config{
    // 您的AccessKey ID
    AccessKeyId: accessKeyId,
    // 您的AccessKey Secret
    AccessKeySecret: accessKeySecret,
  }
  // 访问的域名
  config.Endpoint = tea.String("cdn.aliyuncs.com")
  _result = &cdn20180510.Client{}
  _result, _err = cdn20180510.NewClient(config)
  return _result, _err
}

func _AliyunDescribeRefreshQuota (args []*string) (_result *cdn20180510.DescribeRefreshQuotaResponse,_err error) {
  client, _err := AliyunCreateClient(tea.String(config.Aliyun.AccessKey), tea.String(config.Aliyun.AccessToken))
  if _err != nil {
	loger.Println(_err)
    os.Exit(0)
  }

  describeRefreshQuotaRequest := &cdn20180510.DescribeRefreshQuotaRequest{}
  // 复制代码运行请自行打印 API 的返回值
  ApiFeedBack, _ := client.DescribeRefreshQuota(describeRefreshQuotaRequest)
  if _err != nil {
	loger.Println(_err)
    os.Exit(0)
  }
  return ApiFeedBack,_err
}

func AliyunDescribeRefreshQuota() (result CdnQuota){
  QuotaResult, err := _AliyunDescribeRefreshQuota(tea.StringSlice(os.Args[1:]))
  if err != nil {
    panic(err)
  }
  CurrentQuota.PreloadQuota,_=strconv.Atoi(*QuotaResult.Body.PreloadQuota)
  CurrentQuota.PreloadRemain,_=strconv.Atoi(*QuotaResult.Body.PreloadRemain)
  CurrentQuota.UrlQuota,_=strconv.Atoi(*QuotaResult.Body.UrlQuota)
  CurrentQuota.UrlRemain,_=strconv.Atoi(*QuotaResult.Body.UrlRemain)
  return CurrentQuota
}

/** 腾讯云CDN配额获取部分  */
func TencentDescribeRefreshQuota() (result CdnQuota){

	credential := common.NewCredential(
		config.Tencent.SecretId,
		config.Tencent.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cdn.tencentcloudapi.com"
	client, _ := cdn.NewClient(credential, "", cpf)

	requestPush := cdn.NewDescribePushQuotaRequest()
	responsePush, err := client.DescribePushQuota(requestPush)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		loger.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	//UrlPush[0]代表"Area": "mainland"
	CurrentQuota.PreloadQuota=int(*responsePush.Response.UrlPush[0].Total)
	CurrentQuota.PreloadRemain=int(*responsePush.Response.UrlPush[0].Available)

	requestPurge := cdn.NewDescribePurgeQuotaRequest()
	responsePurge, err := client.DescribePurgeQuota(requestPurge)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		loger.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	//UrlPush[0]代表"Area": "mainland"
	CurrentQuota.UrlQuota=int(*responsePurge.Response.UrlPurge[0].Total)
	CurrentQuota.UrlRemain=int(*responsePurge.Response.UrlPurge[0].Available)

	return CurrentQuota
} 