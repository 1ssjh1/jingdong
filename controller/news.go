package controller

import (
	"JD/utils"
	"github.com/gin-gonic/gin"
)

func News(c *gin.Context) {
	News := utils.GetNews()
	c.JSON(200, gin.H{
		"state": true,
		"msg":   News,
	})
}
