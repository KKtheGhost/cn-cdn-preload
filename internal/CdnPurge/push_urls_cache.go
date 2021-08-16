package CdnPurge

import (
	"fmt"
	"time"

	aliyuncdn "github.com/aliyun/alibaba-cloud-sdk-go/services/cdn"
	tencentcdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */

/** 阿里云预热函数 */
func AliyunPushObject(ObjectPath string, Output bool) {
	client, err := aliyuncdn.NewClientWithAccessKey("cn-hangzhou", config.Aliyun.AccessKey, config.Aliyun.AccessToken)

	request := aliyuncdn.CreatePushObjectCacheRequest()
	request.Scheme = "https"

	request.ObjectPath = ObjectPath
	request.Area = "domestic"

	response, err := client.PushObjectCache(request)
	if err != nil {
		loger.Println(err.Error())
	}
	loger.Printf("Aliyun PushCache: HttpCode is %#v, requestId is %#v, pushTaskId is %#v.\n", response.GetHttpStatus(), response.RequestId, response.PushTaskId)
	if Output {
		fmt.Printf("%v Aliyun PushCache Done! HttpCode is %#v, requestId is %#v, pushTaskId is %#v.\n", time.Now().Format("2006/01/02 15:04:05"), response.GetHttpStatus(), response.RequestId, response.PushTaskId)
	}
}

/** 腾讯云预热函数 */
func TencentPushObject(ObjectPaths []string, Output bool) {
	credential := common.NewCredential(
		config.Tencent.SecretId,
		config.Tencent.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cdn.tencentcloudapi.com"
	client, _ := tencentcdn.NewClient(credential, "", cpf)

	request := tencentcdn.NewPushUrlsCacheRequest()

	//fmt.Println(ObjectPaths)
	request.Urls = common.StringPtrs(ObjectPaths)
	request.Area = common.StringPtr("mainland")

	response, err := client.PushUrlsCache(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		loger.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		loger.Println(err.Error())
	}
	loger.Printf("Tencent PushCache: RequestId is %#v, pushTaskId is %#v.\n", *response.Response.TaskId, *response.Response.RequestId)
	if Output {
		fmt.Printf("%v Tencent PushCache Done! RequestId is %#v, pushTaskId is %#v.\n", time.Now().Format("2006/01/02 15:04:05"), *response.Response.TaskId, *response.Response.RequestId)
	}
}
