package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jingdong/dao"
	"jingdong/models"
	"jingdong/service"
	"jingdong/utils"
)

func BalanceGet(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"state": "false",
			"msg":   "参数绑定错误",
		})
	}
	cookie, _ := c.Cookie("Userinfo")

	ok := utils.ParseToken(user.Token, cookie)
	if !ok {
		c.JSON(200, gin.H{
			"state": ok,
			"msg":   "身份验证失败",
		})
		return
	}
	ok, state := service.BalanceGet(user.Username)
	if !ok {
		c.JSON(200, gin.H{
			"msg":   state,
			"state": ok,
		})
		return
	}
	c.JSON(200, gin.H{
		"state":   ok,
		"msg":     "查找成功",
		"balance": state,
	})

}
func BalanceCharge(c *gin.Context) {
	var user models.Balance
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"state": "false",
			"msg":   "参数绑定失败",
		})
		return
	}
	cookie, _ := c.Cookie("Userinfo")
	ok := utils.ParseToken(user.Token, cookie)
	if !ok {
		c.JSON(200, gin.H{
			"msg":  "用户信息不匹配",
			"code": "false",
		})
		return
	}

	ok, state := service.BalanceCharge(user)
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
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"state": "false",
			"msg":   "参数错误",
		})
		return
	}
	cookie, _ := c.Cookie("Userinfo")

	ok := utils.ParseToken(user.Token, cookie)
	if !ok {
		c.JSON(200, gin.H{
			"state": ok,
			"msg":   "身份验证失败",
		})
		return
	}
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
	cookie, _ := c.Cookie("Userinfo")
	err := c.ShouldBind(&order)
	if err != nil {
		c.JSON(200, gin.H{
			"state": "false",
			"msg":   "参数绑定失败",
		})
		return
	}
	fmt.Println(order)
	ok := utils.ParseToken(order.Token, cookie)
	if !ok {
		c.JSON(200, gin.H{
			"state": ok,
			"msg":   "身份验证失败",
		})
		return
	}
	ok, state := dao.UpdateOrder(order)
	c.JSON(200, gin.H{
		"state": ok,
		"msg":   state,
	})

}
func DeleteOrder(c *gin.Context) {
	var order models.UpdateOrder
	cookie, _ := c.Cookie("Userinfo")
	err := c.ShouldBind(&order)
	if err != nil {
		c.JSON(200, gin.H{
			"state": "false",
			"msg":   "参数绑定失败",
		})
		return
	}
	fmt.Println(order)
	ok := utils.ParseToken(order.Token, cookie)
	if !ok {
		c.JSON(200, gin.H{
			"state": ok,
			"msg":   "身份验证失败",
		})
		return
	}
	ok, state := dao.DeleteOrder(order)
	c.JSON(200, gin.H{
		"state": ok,
		"msg":   state,
	})
}
func Commit(c *gin.Context) {
	var Commit models.Commit
	err := c.ShouldBind(&Commit)
	if err != nil {
		c.JSON(200, gin.H{
			"state": "false",
			"msg":   "参数绑定失败",
		})
		return
	}
	cookie, _ := c.Cookie("Userinfo")
	ok := utils.ParseToken(Commit.Token, cookie)
	if !ok {
		c.JSON(200, gin.H{
			"state": ok,
			"msg":   "身份验证失败",
		})
		return
	}
	ok, msg := dao.Commit(Commit)
	c.JSON(200, gin.H{
		"state": ok,
		"msg":   msg,
	})

}
