package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"jingdong/dao"
)

func SetCk(n string, x string) bool {
	ctx := context.Background()
	if err := dao.RedisClient(); err != nil {
		return false
	}
	err := dao.RDB.Set(ctx, n, x, 300).Err()
	if err != nil {
		fmt.Println(err)
	}
	return true
}
func GetCk(n string, s string) (bool, string) {
	ctx := context.Background()
	if err := dao.RedisClient(); err != nil {
		return false, "错误"
	}
	val, err := dao.RDB.Get(ctx, n).Result()
	if err == redis.Nil {
		return false, "参数不存在"
	}
	if n != val {
		return false, "验证码错误"
	}
	if val != s {
		return false, "验证码错误"
	}

	return true, "验证成功"

}
