package dao

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// RDB var ctx = context.Background()
// 声明一个全局的rdb变量
var RDB *redis.Client

func RedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	RDB = rdb
	return RDB
}
