package controller

import (
	"Goto/models"
	"Goto/service"
	"github.com/gin-gonic/gin"
)


func Register(c *gin.Context)  {
	//接受参数
	var register models.Register
	err :=c.ShouldBind(&register)
	if err!=nil {
		c.JSON(200,gin.H{
			"state":false,
			"msg":"参数绑定失败",
		})
	}
	//参数传递 进行校验
	ok,code :=service.Register(register)
	if!ok{
		c.JSON(200,gin.H{
			"state":"false",
			"msg":code,
		})
		return
	}
	c.JSON(200,gin.H{
		"state":"ture",
		"msg":code,
	})
	return

}
