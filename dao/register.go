package dao

import (
	"fmt"
)

type User struct {
	name string
	password string
}
func Register(name string,word string,number string)(bool,string) {
	var U User
	stm, err := DB.Prepare("select  name from user_info where name = ?" )
	if err != nil {
		fmt.Println(err)
		return false ,"注册失败"
	}
	defer stm.Close()
	rows ,err :=stm.Query(&name)
	if err!=nil {
		fmt.Println(err)
		return false ,"注册失败"
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&U.name)
		if name ==U.name {
			return false ,"用户名已存在 要不登录试试"
		}
	}
	stm, err = DB.Prepare("select  number from user_info where number = ?" )
	if err != nil {
		fmt.Println(err)
		return false ,"注册失败"
	}
	defer stm.Close()
	rows ,err =stm.Query(&number)
	if err!=nil {
		fmt.Println(err)
		return false ,"注册失败"
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&U.name)
		if number ==U.name {
			return false ,"手机号已经被注册"
		}
	}
	stm, err = DB.Prepare("insert into user_info(name,word,number) values(?,?,?);")
	if err != nil {
		fmt.Println(err)
		return false ,"注册失败"
	}
	_, err = stm.Exec(name,word,number)
	if err != nil {
		fmt.Println(err)
		return false ,"注册失败"
	}
	return true,"注册成功"

}

