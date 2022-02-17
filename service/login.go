package service

import (
	"JD/dao"
	"JD/hander"
	"JD/models"
	"github.com/gin-gonic/gin"
)

func Login(u models.Login, c *gin.Context) (bool, error, string) {
	Info, err := dao.Login(u)
	if err != nil {
		return false, err, ""
	}

	//发现这里只用返回两个参数的
	token := hander.Login(c, *Info)
	return true, nil, token

}
func Admin(name string, word string) (bool, string) {
	ok, state := dao.AdminLogin(name, word)

	return ok, state
}
