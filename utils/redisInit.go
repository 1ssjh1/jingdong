package utils

import (
	"fmt"
	//"context"
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

// PoolInitRedis 之前放在dao 下面的 结果出现循环引用
func PoolInitRedis() *redigo.Pool {
	au := Init()
	server := au.Redis.Host + ":" + au.Redis.Port
	password := au.Redis.Passwd
	return &redigo.Pool{
		MaxIdle:     2, //空闲数
		IdleTimeout: 240 * time.Second,
		MaxActive:   3, //最大数
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
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			Pong, err := c.Do("PING")
			fmt.Println(Pong)
			return err
		},
	}
}
