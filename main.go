package main

import (
	"JD/dao"
	"JD/router"
	"JD/utils"
)

func main() {
	utils.PoolInitRedis()
	dao.MysqlInit()
	router.Entrance()

}
