package service

import (
	"github.com/gin-gonic/gin"
	"jingdong/dao"
	"jingdong/hander"
	"jingdong/models"
)

func Login(u models.Login, c *gin.Context) (bool, string, string) {
	ok, sate := dao.Login(u)
	if ok {
		code := hander.Login(c, u)
		return true, sate, code
	}
	return false, sate, ""
}
func Admin(name string, word string) (bool, string) {
	ok, state := dao.AdminLogin(name, word)
	return ok, state
}
