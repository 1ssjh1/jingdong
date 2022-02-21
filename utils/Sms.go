package utils

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	oerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

//var Massage models.Message

func Sendsms(Number string) error {

	///腾讯云短信发送
	// 密钥

	credential := common.NewCredential(Au.Sms.ID, Au.Sms.Key)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)
	request := sms.NewSendSmsRequest()
	request.PhoneNumberSet = common.StringPtrs([]string{Number})
	//短信模板
	request.SmsSdkAppId = common.StringPtr(Au.Sms.Smsid)
	request.SignName = common.StringPtr(Au.Sms.Signname)
	request.TemplateId = common.StringPtr(Au.Sms.Templeid)
	//生成验证码
	code := strconv.FormatInt(time.Now().UnixNano()%1000000, 10)
	request.TemplateParamSet = common.StringPtrs([]string{code, "5"})
	response, err := client.SendSms(request)
	if err, ok := err.(*oerr.TencentCloudSDKError); ok {
		//原本以为  错误信息是在这里处理 结果是在 response里处理
		fmt.Printf("An API error has returned: %s", err)
		err := errors.New("短信发送失败")
		return err
	}
	if err != nil {
		panic(err)
	}
	//错误处理 参照文档列出来的部分错误 进行解析 返回
	for _, v := range response.Response.SendStatusSet {
		if *v.Code == sms.FAILEDOPERATION_FAILRESOLVEPACKET {
			err := errors.New("短信请求失败了")
			return err
		}
		if *v.Code == sms.UNSUPPORTEDOPERATION_CONTAINDOMESTICANDINTERNATIONALPHONENUMBER {
			err := errors.New("短信号码错误，你莫呼我")
			return err
		}
		if *v.Code == sms.INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER {
			err := errors.New("手机号格式错误")
			return err
		}
		if *v.Code == sms.MISSINGPARAMETER_EMPTYPHONENUMBERSET {
			err := errors.New("手机号为空")
			return err
		}

		if *v.Code == sms.LIMITEXCEEDED_PHONENUMBERTHIRTYSECONDLIMIT {
			err := errors.New("发的太快了 歇会儿再试试")
			return err
		}
		if *v.Code == sms.LIMITEXCEEDED_PHONENUMBERONEHOURLIMIT {
			err := errors.New("发的太快了 歇会儿再试试")
			return err
		}
		if *v.Code == sms.LIMITEXCEEDED_PHONENUMBERDAILYLIMIT {
			err := errors.New("今天发的太多了 明天再来吧")
			return err
		}
	}

	SetConform(Number, code)

	return nil
}
