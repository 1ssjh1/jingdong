package controller

import (
	"JD/dao"
	"JD/models"
	"JD/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u models.Login
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数绑定失败",
		})
		return
	}
	Info, err := dao.Login(u)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	token := utils.MakeToken(*Info)
	ok := utils.SetToken(token)
	if !ok {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "登录失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"state": true,
		"msg":   "登录成功",
		"Token": token,
	})
	return

}
func Logout(context *gin.Context) {
	Authorization := context.Request.Header.Get("Authorization")
	ok := utils.DeleteToken(Authorization)
	if !ok {

		context.JSON(200, gin.H{
			"state": false,
			"msg":   "退出登录失败",
		})
		return
	}
	context.JSON(200, gin.H{
		"state": true,
		"msg":   "退出登录成功",
	})
	return
}
func Find(c *gin.Context) {
	var Forget models.Register
	err := c.ShouldBind(&Forget)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数绑定失败",
		})
		return
	}
	err = utils.GetConform(Forget.Number, Forget.Code)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	err = dao.Find(Forget)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"state": true,
		"msg":   "密码找回成功",
	})
	return
}
