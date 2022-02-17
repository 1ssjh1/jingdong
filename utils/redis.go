package utils

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

func SetCk(n string, x string) bool {

	c := PoolInitRedis().Get()
	defer c.Close()
	_, err := c.Do("SET", n, x, "EX", 300)
	if err != nil {
		return false
	}
	return true
}
func GetCk(n string, s string) (bool, error) {
	//取出来rediis 连接
	c := PoolInitRedis().Get()
	//
	defer c.Close()
	v, err := redis.String(c.Do("GET", n))
	if err == redis.ErrNil {
		err = errors.New("验证码超时")
		return false, err
	}
	if v != s {
		err = errors.New("验证码错误")
		return false, err
	}
	return true, nil
}
