package utils

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisPool *redis.Pool

// PoolInitRedis 之前放在dao 下面的 结果出现循环引用
func PoolInitRedis() *redis.Pool {
	server := Au.Redis.Host + ":" + Au.Redis.Port
	password := Au.Redis.Passwd
	redisPool := &redis.Pool{
		MaxIdle:     4, //空闲数
		IdleTimeout: 240 * time.Second,
		MaxActive:   10, //最大数
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
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
