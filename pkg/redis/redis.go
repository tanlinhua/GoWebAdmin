package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/tanlinhua/go-web-admin/service/config"
)

var Redis *RedisClient

// RedisClient extend client and have itself func
type RedisClient struct {
	*redis.Client
}

// Init the redis client
func NewRedisClient() error {
	if Redis != nil {
		return nil
	}
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPWD,
		DB:       config.RedisDB,
		PoolSize: 10, //连接池大小
		//超时
		DialTimeout:  5 * time.Second, //连接建立超时时间，默认5秒。
		ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
		WriteTimeout: 3 * time.Second, //写超时，默认等于读超时
		PoolTimeout:  4 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。

		//闲置连接检查包括IdleTimeout，MaxConnAge
		IdleCheckFrequency: 60 * time.Second, //闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
		IdleTimeout:        5 * time.Minute,  //闲置超时，默认5分钟，-1表示取消闲置超时检查
		MaxConnAge:         0 * time.Second,  //连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接

		//命令执行失败时的重试策略
		MaxRetries:      0,                      //命令执行失败时，最多重试多少次，默认为0即不重试
		MinRetryBackoff: 8 * time.Millisecond,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		MaxRetryBackoff: 512 * time.Millisecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔

		//仅当客户端执行命令时需要从连接池获取连接时，如果连接池需要新建连接时则会调用此钩子函数
		OnConnect: func(conn *redis.Conn) error {
			// fmt.Printf("创建新的连接: %v\n", conn)
			return nil
		},
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	Redis = &RedisClient{client}
	return nil
}

// init the redis  client
func init() {
	err := NewRedisClient()
	if err != nil {
		fmt.Println("redis初始化失败", err.Error())
	}
}

// get the redis client，if client not initialization
// and create the redis client
func GetRedisClient() (*RedisClient, error) {
	if Redis == nil {
		err := NewRedisClient()
		if err != nil {
			return nil, err
		}
		return Redis, nil
	}
	return Redis, nil
}

// close the redis client
func (redis *RedisClient) CloseRedis() {
	redis.Close()
}

// 设置KEY-VALUE类型缓存
// expTime为过期时间,单位秒,0为永不过期
func (redis *RedisClient) SSet(key string, value interface{}, expTime int32) error {
	return redis.Set(key, value, time.Duration(expTime)*time.Second).Err()
}

// 根据KEY获取VALUE值
func (redis *RedisClient) SGet(key string) string {
	return redis.Get(key).Val()
}

// 在队列尾部插入一个元素
func (redis *RedisClient) ListAdd(key, value string) error {
	return redis.RPush(key, value).Err()
}

// 删除并返回队列中的头元素
func (redis *RedisClient) ListGet(key string) string {
	return redis.LPop(key).Val()
}

// 清空队列
func (redis *RedisClient) ListClear(key string) error {
	return redis.LTrim(key, 1, 0).Err()
}
