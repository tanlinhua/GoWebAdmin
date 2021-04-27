package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/tanlinhua/go-web-admin/pkg/trace"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "My7ddamCSmXC",
		DB:       0,
	})
}

func GetListValue(key string) string {
	val, err := rdb.LPop(ctx, key).Result()
	if err != nil {
		trace.Debug("redis错误=>" + "key=" + key + ",errormsg:" + err.Error())
	}
	return val
}

// https://blog.csdn.net/u013804416/article/details/95065229
