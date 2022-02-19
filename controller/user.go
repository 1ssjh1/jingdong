package controller

import (
	"JD/dao"
	"JD/models"
	"JD/utils"

	"github.com/gin-gonic/gin"
)

func BalanceGet(c *gin.Context) {
	var user models.User
	Info, exist := c.Get("Info")
	if !exist {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数缺失",
		})
	}
	BasicInfo, err := utils.Transform(Info)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	user.BasicInfo = BasicInfo
	balance, err := dao.GetBalance(user.Username)
	if err != nil {
		c.JSON(200, gin.H{
			"msg":   err.Error(),
			"state": false,
		})
		return
	}
	c.JSON(200, gin.H{
		"state":   true,
		"msg":     "查找成功",
		"balance": *balance,
	})

}
func BalanceCharge(c *gin.Context) {
	var user models.Balance
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数绑定失败",
		})
		return
	}
	Info, exist := c.Get("Info")
	if !exist {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数缺失",
		})
	}
	BasicInfo, err := utils.Transform(Info)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	user.BasicInfo = BasicInfo

	ok, state := dao.ChargeBalance(user)
	if !ok {
		c.JSON(200, gin.H{
			"state": "false",
			"msg":   state,
		})
		return
	}
	c.JSON(200, gin.H{
		"state": ok,
		"msg":   state,
	})
	return

}
func Order(c *gin.Context) {
	var user models.User
	Info, exist := c.Get("Info")
	if !exist {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数缺失",
		})
	}
	BasicInfo, err := utils.Transform(Info)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	user.BasicInfo = BasicInfo
	ok, info := dao.AllOrder(user)
	if ok {
		c.JSON(200, *info)
		return
	}

	c.JSON(200, gin.H{
		"state": "false",
	})
}
func UpdateOrder(c *gin.Context) {
	var order models.UpdateOrder
	err := c.ShouldBind(&order)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数绑定失败",
		})
		return
	}
	Info, exist := c.Get("Info")
	if !exist {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数缺失",
		})
	}
	BasicInfo, err := utils.Transform(Info)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	order.BasicInfo = BasicInfo
	err = dao.UpdateOrder(order)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"state": true,
		"msg":   "确认收货成功",
	})

}
func DeleteOrder(c *gin.Context) {
	var order models.UpdateOrder
	err := c.ShouldBind(&order)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数绑定失败",
		})
		return
	}
	Info, exist := c.Get("Info")
	if !exist {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数缺失",
		})
	}
	BasicInfo, err := utils.Transform(Info)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	order.BasicInfo = BasicInfo
	err = dao.DeleteOrder(order)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"state": true,
		"msg":   "订单销毁成功",
	})
}
func Commit(c *gin.Context) {
	var Commit models.Commit
	err := c.ShouldBind(&Commit)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数绑定失败",
		})
		return
	}
	Info, exist := c.Get("Info")
	if !exist {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数缺失",
		})
	}
	BasicInfo, err := utils.Transform(Info)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	Commit.BasicInfo = BasicInfo
	ok, msg := dao.Commit(Commit)
	c.JSON(200, gin.H{
		"state": ok,
		"msg":   msg,
	})
}

// ImageUser 用户头像修改
func ImageUser(c *gin.Context) {
	var User models.UserImage
	err := c.ShouldBind(&User)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数绑定失败",
		})
		return
	}
	Info, exist := c.Get("Info")
	if !exist {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数缺失",
		})
	}
	BasicInfo, err := utils.Transform(Info)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}

	User.BasicInfo = BasicInfo
	url, err := utils.SaveFile(User.Image, c)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	msg, err := dao.SaveFile(url, User.BasicInfo)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"state": true,
		"msg":   msg,
	})
	return
}
func Info(c *gin.Context) {
	var User models.BasicInfo
	//从中间件的set 中取出token的荷载信息 //
	//原本这玩意是放在body里 后面又说放在head里

	Info, exist := c.Get("Info")
	if !exist {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数缺失",
		})
	}
	BasicInfo, err := utils.Transform(Info)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	User = BasicInfo
	UserInfo, err := dao.MyInfo(User)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"state": true,
		"msg":   *UserInfo,
	})
}
