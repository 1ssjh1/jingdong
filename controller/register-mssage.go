package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jingdong/utils"
)

func SendMessage(c *gin.Context) {
	number := c.Query("number")
	//number:=c.PostForm("number")
	fmt.Println(number)
	ok, code := utils.Sendsms(number)
	if !ok {
		c.JSON(200, gin.H{
			"state": "false",
			"msg":   code,
		})
		return
	}

	c.JSON(200, gin.H{
		"state": "true",
		"msg":   "短信发送成功",
	})
}
