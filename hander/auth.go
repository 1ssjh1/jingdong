package hander

import (
	"JD/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		Authorization := c.Request.Header.Get("Authorization")
		fmt.Println(Authorization)
		Info, err := utils.ParseToken(Authorization)
		if err != nil {
			c.JSON(200, gin.H{
				"state": false,
				"msg":   err.Error(),
			})
			c.Abort()
			return
		}
		if Info.Uid == 0 {
			c.Set("Info", Info)
			c.Next()
			cookie, err := c.Cookie("super")
			if err != nil {
				fmt.Println(err)
			}
			if cookie == "" {
				c.JSON(200, gin.H{
					"msg":  "你还没有登录",
					"code": false,
				})
				c.Abort()
				return
			}
		} else {
			cookie, err := c.Cookie("Userinfo")
			if err != nil {
				fmt.Println(err)
			}
			if cookie == "" {
				c.JSON(200, gin.H{
					"msg":  "你还没有登录",
					"code": false,
				})
				c.Abort()
				return
			}
			c.Set("Info", Info)
			c.Next()
		}
	}
}
