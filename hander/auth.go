package hander

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("Userinfo")
		if err != nil {
			fmt.Println(err)
		}
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "验证错误",
				"code": "false",
			})
			c.Abort()
			return
		}
		if cookie == "" {
			c.JSON(200, gin.H{
				"msg":  "你还没有登录",
				"code": "false",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
