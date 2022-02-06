package controller

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"jingdong/dao"
	"jingdong/models"
	"jingdong/utils"
)

func AllShop(c *gin.Context) {
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
	if ok {

		slice := dao.AllShops()
		c.JSON(200, gin.H{
			"state": true,
			"msg":   slice,
		})
		return
	}

	c.JSON(200, gin.H{
		"state": ok,
		"msg":   "身份验证失败",
	})
	return
}
func Chart(c *gin.Context) {
	var Chart models.ShopChart
	err := c.ShouldBind(&Chart)
	if err != nil {
		c.JSON(200, gin.H{
			"state": "false",
			"msg":   "参数绑定错误",
		})
		return
	}
	cookie, _ := c.Cookie("Userinfo")
	ok := utils.ParseToken(Chart.Token, cookie)
	if ok {

		ok, state := dao.AddChart(Chart)
		if !ok {
			c.JSON(200, gin.H{
				"state": ok,
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
	c.JSON(200, gin.H{
		"state": ok,
		"msg":   "身份验证失败",
	})
	return

}
func Update(c *gin.Context) {
	var chart models.ShopChart
	err := c.ShouldBind(&chart)
	if err != nil {
		c.JSON(200, gin.H{
			"state": "false",
			"msg":   "参数绑定错误",
		})
		return
	}
	cookie, _ := c.Cookie("Userinfo")
	ok := utils.ParseToken(chart.Token, cookie)
	if ok {
		ok, state := dao.UpdateChart(chart)

		c.JSON(200, gin.H{
			"state": ok,
			"msg":   state,
		})
		return

	}
	c.JSON(200, gin.H{
		"state": "false",
		"msg":   "身份验证失败",
	})
	return

}

func AllChart(c *gin.Context) {
	var UserInfo models.Userinfo
	err := c.ShouldBind(&UserInfo)
	if err != nil {
		c.JSON(200, gin.H{
			"state": "false",
			"msg":   "参数绑定错误",
		})
		return
	}
	cookie, _ := c.Cookie("Userinfo")
	ok := utils.ParseToken(UserInfo.Token, cookie)
	if ok {

		ok, state := dao.AllChart(UserInfo)
		if !ok {
			c.JSON(200, gin.H{
				"state": ok,
				"msg":   "查询失败",
			})
		}
		c.JSON(200, gin.H{
			"state": "true",
			"msg":   state,
		})
		return
	}
	c.JSON(200, gin.H{
		"state": ok,
		"msg":   "身份验证失败",
	})
	return
}

func MakeOrder(c *gin.Context) {
	var Oder models.Order
	err := c.ShouldBind(&Oder)
	if err != nil {
		c.JSON(200, gin.H{
			"state": "false",
			"msg":   "参数绑定错误",
		})
	}
	cookie, _ := c.Cookie("Userinfo")
	ok := utils.ParseToken(Oder.Token, cookie)
	if ok {
		ok, state := dao.MakeOrder(Oder)
		c.JSON(200, gin.H{
			"state": ok,
			"msg":   state,
		})
		return
	}
	c.JSON(200, gin.H{
		"state": ok,
		"msg":   "身份验证失败",
	})
	return
}
func GetCommit(c *gin.Context) {

	var (
		commit models.Commits
	)
	err := c.ShouldBind(&commit)
	if err != nil {
		c.JSON(200, gin.H{
			"state": "false",
			"msg":   "参数错误",
		})
		return
	}
	cookie, _ := c.Cookie("Userinfo")
	ok := utils.ParseToken(commit.Token, cookie)
	if !ok {
		c.JSON(200, gin.H{
			"state": "false",
			"msg":   "身份验证错误",
		})
		return
	}
	Allcommit := dao.GetCommit(commit)
	if Allcommit != nil {
		c.JSON(200, gin.H{
			"state":  "true",
			"commit": *Allcommit,
		})
		return
	}
	c.JSON(200, gin.H{
		"state": "false",
		"msg":   "操作失败",
	})

}
