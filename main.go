package main

import (
	"Goto/router"
	"Goto/dao"
)
func main() {
	dao.MysqlInit()
	router.Entrance()

}
