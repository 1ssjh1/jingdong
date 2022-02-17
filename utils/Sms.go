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
	au := Init()
	credential := common.NewCredential(au.Sms.ID, au.Sms.Key)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)
	request := sms.NewSendSmsRequest()
	request.PhoneNumberSet = common.StringPtrs([]string{Number})
	fmt.Println(Number)
	//短信模板
	request.SmsSdkAppId = common.StringPtr(au.Sms.Smsid)
	request.SignName = common.StringPtr(au.Sms.Signname)
	request.TemplateId = common.StringPtr(au.Sms.Templeid)
	//生成验证码
	code := strconv.FormatInt(time.Now().UnixNano()%1000000, 10)
	fmt.Println("\n", code)
	fmt.Println(Number)
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
	for _, v := range response.Response.SendStatusSet {
		fmt.Println(v.Code)
		if *v.Code == sms.FAILEDOPERATION_FAILRESOLVEPACKET {
			fmt.Println("请求包解析失败")
			err := errors.New("短信请求失败了")
			return err
		}
		if *v.Code == sms.UNSUPPORTEDOPERATION_CONTAINDOMESTICANDINTERNATIONALPHONENUMBER {
			fmt.Println("短信号码错误 再试试吧")
			err := errors.New("短信号码错误，你莫呼我")
			return err
		}
		if *v.Code == sms.INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER {
			fmt.Println("手机号格式错误")
			err := errors.New("手机号格式错误")
			return err
		}
		if *v.Code == sms.MISSINGPARAMETER_EMPTYPHONENUMBERSET {
			fmt.Println("传入的号码列表为空")
			err := errors.New("你没输入号码")
			return err
		}

		if *v.Code == sms.LIMITEXCEEDED_PHONENUMBERTHIRTYSECONDLIMIT {
			fmt.Println("30秒只能发送一条验证码")
			err := errors.New("发的太快了 歇会儿再试试")
			return err
		}
		if *v.Code == sms.LIMITEXCEEDED_PHONENUMBERONEHOURLIMIT {
			fmt.Println("半小时只能发送五条短信")
			err := errors.New("发的太快了 歇会儿再试试")
			return err
		}
		if *v.Code == sms.LIMITEXCEEDED_PHONENUMBERDAILYLIMIT {
			err := errors.New("今天发的太多了 明天再来吧")
			return err
			//fmt.Println("今天没短信了明天再来吧")
		}
		fmt.Println(*response.Response.SendStatusSet[0].Code)
	}

	SetCk(Number, code)

	return nil
}
