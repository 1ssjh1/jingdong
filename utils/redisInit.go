package utils

import (
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

var RedisPool *redigo.Pool

// PoolInitRedis 之前放在dao 下面的 结果出现循环引用
func PoolInitRedis() *redigo.Pool {
	server := Au.Redis.Host + ":" + Au.Redis.Port
	password := Au.Redis.Passwd
	redisPool := &redigo.Pool{
		MaxIdle:     4, //空闲数
		IdleTimeout: 240 * time.Second,
		MaxActive:   10, //最大数
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}
	RedisPool = redisPool
	return RedisPool
}
