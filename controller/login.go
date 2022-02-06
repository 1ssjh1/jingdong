package controller

import (
	"github.com/gin-gonic/gin"
)
import (
	"jingdong/models"
	"jingdong/service"
)

func Login(c *gin.Context) {
	var u models.Login
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(200, gin.H{
			"state": "false",
			"msg":   "参数绑定失败",
		})
		return
	}
	ok, state, code := service.Login(u, c)
	if ok {
		//c.SetCookie("userinfo", u.Username, 3600, "/", "localhost", false, true)
		c.JSON(200, gin.H{
			"state": ok,
			"msg":   state,
			"Token": code,
		})
		return
	}
	c.JSON(200, gin.H{
		"state": ok,
		"msg":   state,
	})
	return
}
func AdminLogin(c *gin.Context) {
	var admin models.Admin
	err := c.ShouldBind(&admin)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "false",
			"state": "参数绑定失败",
		})
	}
	ok, state := service.Admin(admin.Name, admin.Password)
	if !ok {
		c.JSON(200, gin.H{
			"code": ok,
			"msg":  state,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": ok,
		"msg":  state,
	})

}
func Logout(context *gin.Context) {
	cookie, err := context.Cookie("Userinfo")
	if err != nil {
		context.JSON(200, gin.H{
			"state": "false",
			"msg":   "退出登录失败",
		})
	}
	if cookie == "" {
		context.JSON(200, gin.H{
			"state": "false",
			"msg":   "你丫还没登录呢",
		})
	}
	context.SetCookie("Userinfo", "", 0, "/", "localhost", false, false)
	context.JSON(200, gin.H{
		"state": "ture",
		"msg":   "退出登录成功",
	})
}
