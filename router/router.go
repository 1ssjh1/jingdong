package router

import (
	"github.com/gin-gonic/gin"
	"jingdong/controller"
	"jingdong/hander"
)

func Entrance() {
	r := gin.Default()
	r.Use(hander.Cors())
	r.GET("/register", controller.SendMessage)
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/logout", controller.Logout)
	UserGroup := r.Group("/user")
	{

		UserGroup.POST("/balance", hander.Auth(), controller.BalanceCharge)
		//Post用于充值
		//GET用于查询
		UserGroup.GET("/balance", hander.Auth(), controller.BalanceGet)

		//用于获取用户订单
		UserGroup.GET("/order", hander.Auth(), controller.Order)
		UserGroup.PUT("/order", hander.Auth(), controller.UpdateOrder)
		UserGroup.DELETE("/order", hander.Auth(), controller.DeleteOrder)
		//Post 提交评论
		UserGroup.POST("/commit", hander.Auth(), controller.Commit)
	}
	ShopCenter := r.Group("/shop")
	{
		//all 显示所有商品
		ShopCenter.GET("/all", hander.Auth(), controller.AllShop)
		//commit 获取评论
		ShopCenter.GET("/commit", hander.Auth(), controller.GetCommit)
		//post 添加商品
		ShopCenter.POST("/chart", hander.Auth(), controller.Chart)
		//Get 获取购物车信息
		ShopCenter.GET("/chart", hander.Auth(), controller.AllChart)
		//update 对购物车信息进行修改
		ShopCenter.PUT("/chart", hander.Auth(), controller.Update)
		//order 生成订单
		ShopCenter.POST("/order", hander.Auth(), controller.MakeOrder)
		//
	}

	Admin := r.Group("/admin")
	{
		Admin.POST("/login", controller.AdminLogin)
	}
	r.Run(":8080")

}
