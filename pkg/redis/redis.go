package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/tanlinhua/go-web-admin/app/config"
)

// http://doc.redisfans.com

var Handler *RedisClient

type RedisClient struct {
	*redis.Client
}

func init() {
	err := new()
	if err != nil {
		fmt.Println("redis初始化失败", err.Error())
	}
}

func new() error {
	if Handler != nil {
		return nil
	}
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr, //HOST
		Password: config.RedisPWD,  //密码
		DB:       config.RedisDB,   //DB
		PoolSize: 10,               //连接池大小

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
	Handler = &RedisClient{client}
	return nil
}

// 获取redis操作对象
func GetClient() (*RedisClient, error) {
	if Handler == nil {
		err := new()
		if err != nil {
			return nil, err
		}
		return Handler, nil
	}
	return Handler, nil
}

func (rdb *RedisClient) CloseRedis() {
	rdb.Close()
}

//----------------------------------------------------------------------
// 常用功能封装
// 未封装的使用示例:
// import rdb "github.com/tanlinhua/go-web-admin/pkg/redis"
// rdb.Handler.ZRange
//----------------------------------------------------------------------

//----------------------------------------------------------------------
// KEY命名规则:简洁,高效,可维护; 单词与单词之间以 : 隔开
// user:token:id => set user:token:id:1 tokenValue
//----------------------------------------------------------------------

//----------------------------------------------------------------------
// Tag.字符串 (string)
//----------------------------------------------------------------------

// 设置KEY-VALUE类型缓存
// expTime为过期时间,单位秒,0为永不过期
func (rdb *RedisClient) SSet(key string, value interface{}, expTime int32) error {
	return rdb.Set(key, value, time.Duration(expTime)*time.Second).Err()
}

// 根据KEY获取VALUE值
func (rdb *RedisClient) SGet(key string) string {
	return rdb.Get(key).Val()
}

// 自增,返回当前的值
func (rdb *RedisClient) Inc(key string) (int64, error) {
	return rdb.Incr(key).Result()
}

// 自减,返回当前的值
func (rdb *RedisClient) Dec(key string) (int64, error) {
	return rdb.Decr(key).Result()
}

//----------------------------------------------------------------------
// Tag.哈希 (hash)
//----------------------------------------------------------------------

// 添加hash key&value值
func (rdb *RedisClient) HashSet(key string, fields map[string]interface{}) (string, error) {
	return rdb.HMSet(key, fields).Result()
}

// 获取hash表中所有的key&value
// HGET -> 获取hash表中单个key的值
// HMGET -> 获取hash表中多个key的值
func (rdb *RedisClient) HashGet(key string) map[string]string {
	return rdb.HGetAll(key).Val()
}

//----------------------------------------------------------------------
// Tag.列表 (list)
//----------------------------------------------------------------------

// 在队列尾部插入一个元素
func (rdb *RedisClient) ListAdd(key, value string) error {
	return rdb.RPush(key, value).Err()
}

// 删除并返回队列中的头元素
func (rdb *RedisClient) ListGet(key string) string {
	return rdb.LPop(key).Val()
}

// 清空队列
func (rdb *RedisClient) ListClear(key string) error {
	return rdb.LTrim(key, 1, 0).Err()
}

//----------------------------------------------------------------------
// Tag.集合 (set)
//----------------------------------------------------------------------

//----------------------------------------------------------------------
// Tag.有序集合 (sorted set)
//----------------------------------------------------------------------

//----------------------------------------------------------------------
// Tag.事务操作 [保证单个客户端的多个操作是原子的]
//----------------------------------------------------------------------

//----------------------------------------------------------------------
// Tag.redis.opt
//----------------------------------------------------------------------

// 保存缓存数据到文件
func (rdb *RedisClient) Save() (string, error) {
	return rdb.BgSave().Result()
}

// 设置key的过期时间,秒
func (rdb *RedisClient) SetKeyExpire(key string, second int) {
	rdb.Expire(key, time.Duration(second)*time.Second)
}

//----------------------------------------------------------------------
// Tag.脚本
//----------------------------------------------------------------------

// 脚本
func (rdb *RedisClient) Script(script string, keys []string, args ...interface{}) (interface{}, error) {
	scripter := redis.NewScript(script)
	cmd := scripter.Run(Handler, keys, args)
	return cmd.Result()
}
