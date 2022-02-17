package hander

import (
	"JD/models"
	"JD/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context, u models.BasicInfo) string {
	c.SetCookie("Userinfo", u.Username, 3600, "/", "sanser.ltd", false, false)
	token := utils.MakeToken(u)
	return token

}
