package controller

import "github.com/gin-gonic/gin"
import (
	"Goto/models"
	"Goto/service"
)

func Login(c *gin.Context) {
	var u models.Login
	err :=c.ShouldBind(&u)
	if err != nil {
		c.JSON(200,gin.H{
			"state":"false",
			"msg":"参数绑定失败",
		})
		return
	}
	ok,code :=service.Login(u)
	c.JSON(200,gin.H{
		"state":ok,
		"msg":code,
	})
	return
}
