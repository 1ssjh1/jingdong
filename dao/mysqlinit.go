package dao

import (
	"JD/utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func MysqlInit() *sql.DB {
	a := utils.Init()
	Dns := a.Mysql.User + ":" + a.Mysql.Word + "@tcp(110.42.216.125)/" + a.Mysql.Bases
	db, err := sql.Open("mysql", Dns)

	//err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxIdleConns(10)
	//defer db.Close()
	DB = db
	return DB

}
