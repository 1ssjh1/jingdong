package utils

import (
	"JD/models"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

// MyClaims 仿照官方 将用户基本信息负载在token里
type MyClaims struct {
	models.BasicInfo
	jwt.StandardClaims
}

var secret = []byte("secret-my-sing")

func MakeToken(basicinfo models.BasicInfo) string {
	myClaim := MyClaims{
		basicinfo,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
			Issuer:    "Sianao",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)
	t, _ := token.SignedString(secret)
	return t
}
func ParseToken(tokens string) (*models.BasicInfo, error) {
	if tokens == "" {
		err := errors.New("你是谁 验证信息不知道")
		return nil, err
	}
	token, err := jwt.ParseWithClaims(tokens, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if token.Valid {
		claims, ok := token.Claims.(*MyClaims)
		if ok {
			return &claims.BasicInfo, nil

		}
	}

	//jwt 参考 官方文档 进行 验证错误处理
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorExpired != 0 {
			err = errors.New("您的token 已经过期")
			return nil, err
		}
		if ve.Errors|jwt.ValidationErrorIssuer != 0 {
			err = errors.New("tmd 这根本不是我的呜呜呜")
			return nil, err

		}
		if ve.Errors&jwt.ValidationErrorMalformed|jwt.ValidationErrorNotValidYet != 0 {
			err = errors.New("这是个啥 你给我的是个啥")
			return nil, err
		}
	}
	return nil, err

}
func Transform(User interface{}) (models.BasicInfo, error) {
	Info, ok := User.(*models.BasicInfo)
	if !ok {
		err := errors.New("用户认证失败")
		return models.BasicInfo{}, err
	}
	return *Info, nil
}
