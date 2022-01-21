package router

import (
	"Goto/controller"
	"Goto/hander"
	"github.com/gin-gonic/gin"
)

func Entrance()  {
	r :=gin.Default()
	r.Use(hander.Cors())
	r.POST("/register-message",controller.SendMessage)
	r.POST("/register",controller.Register)
	r.POST("/login",controller.Login)
	r.Run()

}