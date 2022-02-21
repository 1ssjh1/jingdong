package dao

import (
	"JD/utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func MysqlInit() *sql.DB {

	Dns := utils.Au.Mysql.User + ":" + utils.Au.Mysql.Word + "@tcp(110.42.216.125)/" + utils.Au.Mysql.Bases
	db, err := sql.Open("mysql", Dns)

	//err = db.Ping()
	if err != nil {
	}
	db.SetMaxIdleConns(10)
	//defer db.Close()
	DB = db
	return DB

}
