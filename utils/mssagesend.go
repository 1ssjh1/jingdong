package utils

import (
	"Goto/models"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"strconv"
	"time"
)
var Massage models.Message

func Sendsms(n string) (bool,string) {

	credential := common.NewCredential("AKIDQfeOs64mEQN5XhbEcjZ8RSDhurYz93g1", "6wWv4wYtymKbJkwykGVTAajbCWZLq42r")
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)
	request := sms.NewSendSmsRequest()
	request.PhoneNumberSet = common.StringPtrs([]string{n} )
	request.SmsSdkAppId = common.StringPtr("1400624150")
	request.SignName = common.StringPtr("郑金坤个人生活之旅")
	request.TemplateId = common.StringPtr("1282576")
	code:= strconv.FormatInt(time.Now().UnixNano()%1000000, 10)
	fmt.Println("\n" ,code)
	request.TemplateParamSet = common.StringPtrs([]string{ string(code), "5" })
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return false,"短信发送失败请重试"
	}
	if err != nil {
		panic(err)
	}
	Massage.Send=code
	Massage.Sendtime=time.Now()
	fmt.Printf("%s", response.ToJsonString())
	return true, code
}
func ConformMassage(code  string) (bool,string) {
	if code!=Massage.Send {
		return false,"验证码错误"
	}
	if !time.Now().After(Massage.Sendtime.Add(time.Minute*5)){
		return false,"验证码超时"
	}
	return true,"验证码校验成功"
}
