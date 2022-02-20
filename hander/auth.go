package hander

import (
	"JD/utils"
	"github.com/gin-gonic/gin"
)

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		Authorization := c.Request.Header.Get("Authorization")
		//fmt.Println(Authorization)
		ok := utils.ConformToken(Authorization)
		if !ok {
			c.JSON(200, gin.H{
				"state": false,
				"msg":   "你还没有登录",
			})
			c.Abort()
			return
		}
		Info, err := utils.ParseToken(Authorization)
		if err != nil {

			c.JSON(200, gin.H{
				"state": false,
				"msg":   err.Error(),
			})
			c.Abort()
		}
		c.Set("Info", Info)
		c.Next()
	}
}
