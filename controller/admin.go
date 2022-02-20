package controller

import (
	"JD/dao"
	"JD/models"
	"JD/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

func RootLogin(c *gin.Context) {
	var admin models.Admin
	err := c.ShouldBind(&admin)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数绑定失败",
		})
	}
	ok, state := dao.AdminLogin(admin.Name, admin.Password)

	if !ok {
		c.JSON(200, gin.H{
			"state": ok,
			"msg":   state,
		})
		return
	}
	var BasicInfo models.BasicInfo
	BasicInfo.Uid = 0
	BasicInfo.Username = admin.Name
	token := utils.MakeToken(BasicInfo)
	ok = utils.SetToken(token)
	if !ok {
		c.JSON(20, gin.H{
			"state": false,
			"msg":   "登录失败",
		})
		return

	}
	c.JSON(200, gin.H{
		"code":  ok,
		"msg":   state,
		"token": token,
	})
	return

}
func RootLogout(c *gin.Context) {
	Authorization := c.Request.Header.Get("Authorization")
	ok := utils.DeleteToken(Authorization)
	if !ok {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "退出登录失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"state": true,
		"msg":   "退出登录成功",
	})
	return
}
func RootAll(c *gin.Context) {
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
	if BasicInfo.Uid != 0 {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "权限不够",
		})
		return
	}
	Allorder, err := dao.GetAllOrder()
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"state": true,
		"msg":   *Allorder,
	})

}
func UpdateONeOrder(c *gin.Context) {
	var update models.UpdateUserOrder
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
	if BasicInfo.Uid != 0 {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "权限不够",
		})
		return
	}
	err = c.ShouldBind(&update)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
	}
	msg, err := dao.OrderChange(update)
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
func DeleteUserOrder(c *gin.Context) {
	var update models.UpdateUserOrder
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
	if BasicInfo.Uid != 0 {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "权限不够",
		})
		return
	}
	err = c.ShouldBind(&update)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
	}
	msg, err := dao.DeleteUserOrder(update)
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
}
func AddGoods(c *gin.Context) {
	var add models.GoodsAdd
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
	if BasicInfo.Uid != 0 {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "权限不够",
		})
		return
	}
	err = c.ShouldBind(&add)
	if err != nil {
		err = errors.New("参数绑定失败")
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	msg, err := dao.AddGoods(add, c)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"state": false,
		"msg":   msg,
	})
}

func UpdateGoods(c *gin.Context) {
	var update models.UpdateGoods
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
	if BasicInfo.Uid != 0 {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "权限不够",
		})
		return
	}
	err = c.ShouldBind(&update)
	if err != nil {
		err = errors.New("参数绑定失败")
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	msg, err := dao.UpdateGoods(update, c)
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
}
func DeleteGoods(c *gin.Context) {
	Gid := c.PostForm("Gid")
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
	if BasicInfo.Uid != 0 {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "权限不够",
		})
		return
	}
	msg, err := dao.DeleteGoods(Gid)
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

}
