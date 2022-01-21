package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func MysqlInit() *sql.DB {
	Dns := "sianao:Simple2002@tcp(110.42.216.125:3306)/sql_test"
	db ,err :=sql.Open("mysql",Dns)
	if err != nil {
		fmt.Println(err)
	}
	//err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxIdleConns(10)
	//defer db.Close()
	DB=db
	return DB


}



