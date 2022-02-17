package router

import (
	"JD/controller"
	"JD/hander"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Entrance() {
	r := gin.Default()
	//使用中间件 获取用户部分状态
	r.Use(hander.Cors())
	r.Use(static.Serve("/", static.LocalFile("./front", false)))
	//静态文件加载 但是貌似 有一丢丢慢
	r.Static("/static", "/www/static")
	//注册短信发送接口
	r.GET("/register", controller.SendMessage)
	//注册接口
	r.POST("/register", controller.Register)
	//登录接口
	r.POST("/login", controller.Login)
	//找回密码
	r.POST("/find", controller.Find)
	//登出接口
	r.GET("/logout", controller.Logout)
	r.GET("/show", controller.Show)
	UserGroup := r.Group("/user")
	{
		//用户主界面
		UserGroup.GET("/", hander.Auth(), controller.Info)
		//更新用户信息
		UserGroup.PUT("/image", hander.Auth(), controller.ImageUser)
		//Post用于充值
		UserGroup.POST("/balance", hander.Auth(), controller.BalanceCharge)
		//GET用于查询
		UserGroup.GET("/balance", hander.Auth(), controller.BalanceGet)
		//用于获取用户订单
		UserGroup.GET("/order", hander.Auth(), controller.Order)
		//更新用户订单
		UserGroup.PUT("/order", hander.Auth(), controller.UpdateOrder)
		//删除订单
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

	admin := r.Group("/admin")
	{

		admin.POST("/login", controller.RootLogin)
		//展示所有订单 也写个分类吧
		admin.GET("/order", hander.Auth(), controller.RootAll)
		//更新订单
		admin.PUT("/order", hander.Auth(), controller.UpdateONeOrder)
		//删除订单
		admin.DELETE("/order", hander.Auth(), controller.DeleteUserOrder)
		//增加商品
		admin.POST("/goods", hander.Auth(), controller.AddGoods)
		//获取商品
		admin.GET("/goods", hander.Auth(), controller.AllShop)
		//更新商品信息
		admin.PUT("/goods", hander.Auth(), controller.UpdateGoods)
		//删除商品
		admin.DELETE("/goods", hander.Auth(), controller.DeleteGoods)
		//登出
		admin.GET("/logout", hander.Auth(), controller.RootLogout)
	}
	r.Run(":8080")
	//r.RunTLS(":443", "test.pem", "test.key")

}
