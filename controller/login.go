package controller

import (
	"JD/dao"
	"JD/models"
	"JD/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Login(c *gin.Context) {
	var u models.Login
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数绑定失败",
		})
		return
	}
	Info, err := dao.Login(u)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	token := utils.MakeToken(*Info)
	ok := utils.SetToken(token)
	if !ok {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "登录失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"state": true,
		"msg":   "登录成功",
		"Token": token,
	})
	return

}
func Logout(context *gin.Context) {
	Authorization := context.Request.Header.Get("Authorization")
	ok := utils.DeleteToken(Authorization)
	if !ok {

		context.JSON(200, gin.H{
			"state": false,
			"msg":   "退出登录失败",
		})
		return
	}
	context.JSON(200, gin.H{
		"state": true,
		"msg":   "退出登录成功",
	})
	return
}
func Find(c *gin.Context) {
	var Forget models.Register
	err := c.ShouldBind(&Forget)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   "参数绑定失败",
		})
		return
	}
	err = utils.GetConform(Forget.Number, Forget.Code)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	err = dao.Find(Forget)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"state": true,
		"msg":   "密码找回成功",
	})
	return
}
func Oauth(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://github.com/login/oauth/authorize?client_id=a3112bb967a7bbe3bcf1&redirect_uri=https://sanser.ltd/callback")
}
func Callback(c *gin.Context) {
	code := c.Query("code")
	fmt.Println(code)
	reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", "a3112bb967a7bbe3bcf1", "82625129d028e98a671c52c81bd9e45b4b574705", code)
	req, err := http.NewRequest(http.MethodPost, reqURL, nil)
	if err != nil {

	}
	req.Header.Set("Accept", "application/json")
	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	defer res.Body.Close()
	info, _ := ioutil.ReadAll(res.Body)
	var token models.Token
	err = json.Unmarshal(info, &token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
	req, err = http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "token "+token.AccessToken)
	res, err = httpClient.Do(req)
	defer res.Body.Close()
	info, _ = ioutil.ReadAll(res.Body)
	//fmt.Println(string(info))
	var basicinfo models.HubBasicInfo
	err = json.Unmarshal(info, &basicinfo)
	if err != nil {

	}
	user, err := dao.HubLogin(basicinfo)
	if err != nil {
		c.JSON(200, gin.H{
			"state": false,
			"err":   err,
		})
		return
	}
	newtoken := utils.MakeToken(user)
	ok := utils.SetToken(newtoken)
	if !ok {
		c.JSON(200, gin.H{
			"state": false,
		})
		return
	}
	c.JSON(200, gin.H{
		"state": true,
		"msg":   "登录成功",
		"token": newtoken,
	})
	return
}
