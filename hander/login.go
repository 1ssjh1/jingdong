package hander

import (
	"github.com/gin-gonic/gin"
	"jingdong/models"
	"jingdong/utils"
)

func Login(c *gin.Context, u models.Login) string {
	c.SetCookie("Userinfo", u.Username, 3600, "/", "localhost", false, false)
	code := utils.MakeToken(u.Username)
	return code

}
