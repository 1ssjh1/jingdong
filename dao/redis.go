package dao

import (
	"fmt"
	//"context"
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

func PoolInitRedis() *redigo.Pool {
	server := "localhost:6379"
	password := ""
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
