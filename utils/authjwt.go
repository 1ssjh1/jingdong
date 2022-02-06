package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var secret = []byte("secret-my-sing")

func MakeToken(username string) string {
	myClaim := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "Sianao",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)
	t, _ := token.SignedString(secret)
	fmt.Println(t)
	return t
}
func ParseToken(tokens, name string) bool {
	if tokens == "" {
		return false
	}
	token, _ := jwt.ParseWithClaims(tokens, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	claims, ok := token.Claims.(*MyClaims)
	if ok {
		fmt.Println(claims.Username)
		if token.Valid {

			fmt.Println(claims.Username)
			if name == claims.Username {
				return true
			}
		}
	}
	return false
	//claims, ok := token.Claims.(*MyClaims)
	//ve, ok := err.(*jwt.ValidationError)
	//if ok {
	//	if ve.Errors&jwt.ValidationErrorExpired != 0 {
	//		fmt.Println("超时了哦")
	//	}
	//
	//	return false
	//}
	//return false
}
