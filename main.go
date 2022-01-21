package main

import (
	"jingdong/dao"
	"jingdong/router"
)
func main() {
	dao.MysqlInit()
	router.Entrance()

}
