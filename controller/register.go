package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jingdong/models"
	"jingdong/service"
	"jingdong/utils"
)

func Register(c *gin.Context) {
	//接受参数
	var register models.Register
	err := c.ShouldBind(&register)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数绑定失败",
		})
		return
	}
	//参数传递 进行校验
	ok, code := service.Register(register)
	if !ok {
		c.JSON(200, gin.H{
			"state": "false",
			"msg":   code,
		})
		return
	}
	c.JSON(200, gin.H{
		"state": "ture",
		"msg":   code,
	})
	return

}

func SendMessage(c *gin.Context) {
	number := c.Query("Phone")
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
