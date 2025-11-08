package dao

import (
	"context"
	"fmt"
	"log"
	"project-user/config"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	rdb *redis.Client
}

var Rc *RedisCache

func init() {
	rdb := redis.NewClient(config.C.InitRedisOptions())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		// 如果连接失败，打印错误日志并可以终止程序
		log.Printf("无法连接到Redis: %v", err)
		// log.Fatalf 会打印日志并调用 os.Exit(1) 终止程序
	} else {
		fmt.Println("成功连接到Redis！")
	}
	Rc = &RedisCache{
		rdb: rdb,
	}
}

func (rc *RedisCache) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	result, err := rc.rdb.Get(ctx, key).Result()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (rc *RedisCache) Put(key, value string, expire time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := rc.rdb.Set(ctx, key, value, expire).Err()
	if err != nil {
		return err
	}
	return nil
}
