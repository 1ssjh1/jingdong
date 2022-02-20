package main

import (
	"JD/dao"
	"JD/router"
	"JD/utils"
)

func main() {
	utils.Init()
	utils.PoolInitRedis()
	dao.MysqlInit()
	router.Entrance()

}
