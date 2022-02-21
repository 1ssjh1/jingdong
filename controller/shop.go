package controller

import (
	"JD/dao"
	"JD/models"
	"JD/utils"
	_ "fmt"
	"github.com/gin-gonic/gin"
)

func AllShop(c *gin.Context) {
	var user models.User
	Info, exist := c.Get("Info")
	if !exist {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数缺失",
		})
	}
	BasicInfo, err := utils.Transform(Info)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	user.BasicInfo = BasicInfo
	slice := dao.AllShops()
	if slice == nil {
		c.JSON(200, gin.H{
			"state": false,
		})
		return
	}
	c.JSON(200, gin.H{
		"state": true,
		"msg":   *slice,
	})
	return
}

func Chart(c *gin.Context) {

	var chart models.AddChart
	err := c.ShouldBind(&chart)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数绑定失败",
		})
		return
	}
	Info, exist := c.Get("Info")
	if !exist {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数缺失",
		})
	}
	BasicInfo, err := utils.Transform(Info)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}

	chart.BasicInfo = BasicInfo

	//info := <-s
	msg, err := dao.AddChart(chart)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"state": true,
		"msg":   msg,
	})

}

//c := <-m
//wait.Add(1)

func Update(c *gin.Context) {
	var chart models.ShopChart

	err := c.ShouldBind(&chart)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数绑定失败",
		})
		return
	}
	//chart.Count = c.PostForm("count")
	//chart.ChartId = c.PostForm("chart_id")
	Info, exist := c.Get("Info")
	if !exist {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数缺失",
		})
	}
	BasicInfo, err := utils.Transform(Info)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}

	chart.BasicInfo = BasicInfo

	//这里后面改一下
	ok, state := dao.UpdateChart(chart)
	c.JSON(200, gin.H{
		"state": ok,
		"msg":   state,
	})
	return

}

func AllChart(c *gin.Context) {
	var UserInfo models.Userinfo
	Info, exist := c.Get("Info")
	if !exist {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数缺失",
		})
	}
	BasicInfo, err := utils.Transform(Info)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}

	UserInfo.BasicInfo = BasicInfo
	msg, err := dao.AllChart(UserInfo)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"state": "true",
		"msg":   *msg,
	})
	return

}

func MakeOrder(c *gin.Context) {
	var Oder models.Order
	err := c.ShouldBind(&Oder)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数绑定失败",
		})
		return
	}
	//Oder.ChartId = c.PostFormArray("chart_id")
	Info, exist := c.Get("Info")
	if !exist {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数缺失",
		})
	}
	BasicInfo, err := utils.Transform(Info)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	Oder.BasicInfo = BasicInfo
	ok, state := dao.MakeOrder(Oder)
	c.JSON(200, gin.H{
		"state": ok,
		"msg":   state,
	})
	return
}

func GetCommit(c *gin.Context) {

	var commit models.Commits

	commit.Gid = c.Query("gid")
	Info, exist := c.Get("Info")
	if !exist {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数缺失",
		})
	}
	BasicInfo, err := utils.Transform(Info)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	commit.BasicInfo = BasicInfo
	AllCommit := dao.GetCommit(commit)
	if AllCommit != nil {
		c.JSON(200, gin.H{
			"state":  "true",
			"commit": *AllCommit,
		})
		return
	}
	c.JSON(200, gin.H{
		"state": "false",
		"msg":   "操作失败",
	})

}
func Show(c *gin.Context) {
	msg, err := dao.Class()
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"state": true,
		"msg":   *msg,
	})
}
