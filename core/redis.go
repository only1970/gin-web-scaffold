package core

import (
	"context"
	"fmt"
	"gin-web-scaffold/global"
	"time"

	"github.com/go-redis/redis"
)

func InitRedis() *redis.Client {
	// ctx := context.Background()
	db := global.Config.Redis.DB

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.Config.Redis.IP, global.Config.Redis.Port),
		Password: global.Config.Redis.Password, // no password set
		DB:       db,                           // use default DB
		PoolSize: global.Config.Redis.PoolSize, // 连接池
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		global.Log.Fatalln(fmt.Sprintf("[%s] redis 连接失败,err:%s", rdb, err.Error()))
		return nil
	}
	fmt.Printf("%s:%d redis 初始化成功\n", global.Config.Redis.IP, global.Config.Redis.Port)
	return rdb
}
