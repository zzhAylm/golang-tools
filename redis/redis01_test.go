package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func TestFunRedis(t *testing.T) {
	// 创建redis客户端实例
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis地址
		Password: "",               // Redis密码
		DB:       0,                // Redis数据库索引
	})
	result, err := rdb.Set(
		context.TODO(),
		"zzh",
		"zzh",
		time.Second,
	).Result()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("success", result)

	// 使用完后关闭连接
	defer rdb.Close()

	FunRedis()
}
