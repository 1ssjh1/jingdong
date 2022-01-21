package controller

import (
	"Goto/models"
	"Goto/utils"
	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context){
	number:=c.Param("Phone")
	ok,code :=utils.Sendsms(number)

	if!ok{
		c.JSON(200,gin.H{
			"state":"false",
			"msg":code,
		})
		return
	}
	var codes models.Message
	codes.Send=code
	c.JSON(200,gin.H{
		"state":"true",
		"msg":code,
	})
	return
}



