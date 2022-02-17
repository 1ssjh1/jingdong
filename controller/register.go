package controller

import (
	"JD/models"
	"JD/service"
	"JD/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	//接受参数
	var register models.Register

	err := c.ShouldBind(&register)
	fmt.Println(err)
	fmt.Println(register)
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
			"state": false,
			"msg":   code.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"state": true,
		"msg":   "注册成功",
	})
	return

}

func SendMessage(c *gin.Context) {
	number := c.Query("Phone")
	fmt.Println(number)
	err := utils.Sendsms(number)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"state": true,
		"msg":   "短信发送成功",
	})
}
