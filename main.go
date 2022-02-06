package main

import (
	"jingdong/dao"
	"jingdong/router"
)

func main() {
	dao.PoolInitRedis()
	dao.MysqlInit()
	router.Entrance()

}
