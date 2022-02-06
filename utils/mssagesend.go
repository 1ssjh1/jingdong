package utils

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"strconv"
	"time"
)

//var Massage models.Message

func Sendsms(n string) (bool, string) {

	credential := common.NewCredential()
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)
	request := sms.NewSendSmsRequest()
	request.PhoneNumberSet = common.StringPtrs([]string{n})
	fmt.Println(n)
	request.SmsSdkAppId = common.StringPtr("1400624150")
	request.SignName = common.StringPtr("郑金坤个人生活之旅")
	request.TemplateId = common.StringPtr("1282576")
	code := strconv.FormatInt(time.Now().UnixNano()%1000000, 10)
	fmt.Println("\n", code)
	fmt.Println(n)
	request.TemplateParamSet = common.StringPtrs([]string{code, "5"})
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return false, "短信发送失败请重试"
	}
	if err != nil {
		panic(err)
	}

	SetCk(n, code)
	fmt.Printf("%s"+
		"\n", response.ToJsonString())
	return true, code
}
