package utils

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

func SetCk(n string, x string) bool {
	//参照 官方示例 写个redis 短信验证 过期检验 很简单 很烂
	c := PoolInitRedis().Get()
	defer c.Close()
	_, err := c.Do("SET", n, x, "EX", 300)
	if err != nil {
		return false
	}
	return true
}
func GetCk(n string, s string) error {
	//
	c := PoolInitRedis().Get()
	//
	defer c.Close()
	v, err := redis.String(c.Do("GET", n))
	if err == redis.ErrNil {
		err = errors.New("验证码超时")
		return err
	}
	if v != s {
		err = errors.New("验证码错误")
		return err
	}
	return nil
}
