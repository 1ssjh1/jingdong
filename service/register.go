package service

import (
	"Goto/dao"
	"Goto/models"
	"Goto/utils"
	"time"
)


func Register(c models.Register) (bool ,string) {

	ok,state :=utils.ConformMassage(c.Code)
	if !ok{
		return false,state
	}
	ok,state=dao.Register(c.Username,c.Password,c.Number)
	return ok,state
}
func Confirmtime(t time.Time,s string,r string)(bool,string)  {
	if time.Now().After(t.Add(time.Minute*5)) {
		return false,"验证码超时"
	}
	if s!=r {
		return false,"验证码错误"
	}
	return true,"校验成功"
}