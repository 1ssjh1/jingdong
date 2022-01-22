package utils

import (
	"github.com/gomodule/redigo/redis"
	"jingdong/dao"
)

func SetCk(n string, x string) bool {

	c := dao.PoolInitRedis().Get()
	defer c.Close()
	_, err := c.Do("SET", n, x, "EX", 300)
	if err != nil {
		return false
	}
	return true
}
func GetCk(n string, s string) (bool, string) {
	//取出来rediis 连接
	c := dao.PoolInitRedis().Get()
	//
	v, err := redis.String(c.Do("GET", n))
	if err == redis.ErrNil {
		return false, "验证码超时"
	}
	if v != s {
		return false, "验证码错误"
	}
	return true, "验证码正确"

}
