package controller

import (
	"JD/utils"
	"github.com/gin-gonic/gin"
)

func News(c *gin.Context) {
	News, err := utils.GetNews()
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "新闻获取失败",
		})
	}
	c.JSON(200, News)

}
