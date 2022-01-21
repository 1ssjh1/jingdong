package router

import (
	"github.com/gin-gonic/gin"
	"jingdong/controller"
	"jingdong/hander"
)

func Entrance()  {
	r :=gin.Default()
	r.Use(hander.Cors())
	r.POST("/register-message",controller.SendMessage)
	r.POST("/register",controller.Register)
	r.POST("/login",controller.Login)
	r.Run(":80")

}