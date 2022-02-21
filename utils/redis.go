package utils

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

func SetConform(n string, x string) bool {
	//参照 官方示例 写个redis 短信验证 过期检验 很简单 很烂
	c := RedisPool.Get()
	defer c.Close()
	_, err := c.Do("SET", n, x, "EX", 300)
	if err != nil {
		return false
	}
	return true
}
func GetConform(Number string, ConformCode string) error {
	//
	c := RedisPool.Get()
	//
	defer c.Close()
	ok, err := redis.Bool(c.Do("EXIST", Number))
	if !ok {
		err = errors.New("验证码过期或未发送")
		return err
	}
	v, err := redis.String(c.Do("GET", Number))
	if err == redis.ErrNil {
		err = errors.New("验证码超时")
		return err
	}
	if v != ConformCode {
		err = errors.New("验证码错误")
		return err
	}
	return nil
}

func SetToken(token string) bool {
	c := RedisPool.Get()
	defer c.Close()
	//与token 验证时间一致
	_, err := redis.String(c.Do("SET", token, "", "EX", 1800))
	if err != nil {
		return false
	}
	return true

}
func ConformToken(token string) bool {
	c := RedisPool.Get()
	defer c.Close()
	ok, err := redis.Bool(c.Do("EXISTS", token))
	if err != nil {
		return false
	}
	return ok
}
func DeleteToken(token string) bool {
	c := RedisPool.Get()
	defer c.Close()
	ok, err := redis.Bool(c.Do("DEL", token))
	if err != nil {
		return false
	}
	return ok
}
