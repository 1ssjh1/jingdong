package dao

import (
	"fmt"
	"jingdong/models"
)

func Login(u models.Login) (bool, string) {
	stm, err := DB.Prepare("select  word from user_info where name = ?")
	if err != nil {
		fmt.Println(err)
		return false, "登录失败"
	}
	defer stm.Close()

	rows, err := stm.Query(u.Username)
	if err != nil {
		fmt.Println(err)
		return false, "登录失败"
	}
	var tmp models.Login
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&tmp.Password)
	}
	if tmp.Password == "" {
		return false, "你还没注册登录个屁"
	}
	if tmp.Password != u.Password {
		return false, "密码错误"

	}
	return true, "密码正确"
}
func AdminLogin(name string, word string) (bool, string) {
	stm, err := DB.Prepare("select  word from admin_info where name = ?")
	if err != nil {
		return false, "登录失败"
	}
	rows, err := stm.Query(name)
	if err != nil {
		return false, "登录失败"
	}
	var s string
	for rows.Next() {
		rows.Scan(&s)
	}
	if s != word {
		return false, "密码错误"
	}
	return true, "密码正确"

}
