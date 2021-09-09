package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type RedisLock struct {
	key         string
	value       string
	redisClient *redis.Client
	expiration  time.Duration
	cancelFunc  context.CancelFunc
}

const (
	PubSubPrefix        = "{redis_lock}_"
	DefaultExpiration   = 30
	DefaultSpinInterval = 100
)

/*
https://github.com/Spongecaptain/redisLock

构造分布式锁实例
利用 NewRedisLock 以及 NewRedisLockWithExpireTime 函数能够构造出一个分布式锁实例
NewRedisLockWithExpireTime 的区别在于其能够自定义锁的过期时间。
NewRedisLock 方法接收的 key 决定了分布式锁的粒度，value 决定了只有 value 值相同才能够进行解锁。
*/

func NewRedisLock(redisClient *redis.Client, key string, value string) *RedisLock {
	return &RedisLock{
		key:         key,
		value:       value,
		redisClient: redisClient,
		expiration:  time.Duration(DefaultExpiration) * time.Second,
	}
}

func NewRedisLockWithExpireTime(redisClient *redis.Client, key string, value string, expiration time.Duration) *RedisLock {
	return &RedisLock{
		key:         key,
		value:       value,
		redisClient: redisClient,
		expiration:  expiration}
}

// TryLock 仅尝试一次锁的获取，如果失败，那么不会阻塞，直接返回。
func (lock *RedisLock) TryLock() (bool, error) {
	success, err := lock.redisClient.SetNX(lock.key, lock.value, lock.expiration).Result()
	if err != nil {
		return false, err
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	lock.cancelFunc = cancelFunc
	lock.renew(ctx)
	return success, nil
}

// Lock 会不断尝试索取分布式锁，这会导致调用此方法的协程阻塞。
func (lock *RedisLock) Lock() error {
	for {
		success, err := lock.TryLock()
		if err != nil {
			return err
		}
		if success {
			return nil
		}
		if !success {
			err := lock.subscribeLock()
			if err != nil {
				return err
			}
		}
	}
}

// Unlock 方法用于解锁，由于涉及网络通信，解锁可能失败， error!=nil 意味着解锁失败。
func (lock *RedisLock) Unlock() error {
	script := redis.NewScript(fmt.Sprintf(
		`if redis.call("get", KEYS[1]) == "%s" then return redis.call("del", KEYS[1]) else return 0 end`,
		lock.value))
	runCmd := script.Run(lock.redisClient, []string{lock.key})
	res, err := runCmd.Result()
	if err != nil {
		return err
	}
	if tmp, ok := res.(int64); ok {
		if tmp == 1 {
			lock.cancelFunc() //cancel renew goroutine
			err := lock.publishLock()
			if err != nil {
				return err
			}
			return nil
		}
	}
	err = fmt.Errorf("unlock script fail: %s", lock.key)
	return err
}

// LockWithTimeout 方法会在获取锁资源成功或者超时后返回。
func (lock *RedisLock) LockWithTimeout(d time.Duration) error {
	timeNow := time.Now()
	for {
		success, err := lock.TryLock()
		if err != nil {
			return err
		}
		if success {
			return nil
		}
		deltaTime := d - time.Since(timeNow)
		if !success {
			err := lock.subscribeLockWithTimeout(deltaTime)
			if err != nil {
				return err
			}
		}
	}
}

// 支持指定次数地进行自旋式的锁获取。
func (lock *RedisLock) SpinLock(times int) error {
	for i := 0; i < times; i++ {
		success, err := lock.TryLock()
		if err != nil {
			return err
		}
		if success {
			return nil
		}
		time.Sleep(time.Millisecond * DefaultSpinInterval)
	}
	return fmt.Errorf("max spin times reached")
}

// subscribeLock 阻塞直到锁被释放
func (lock *RedisLock) subscribeLock() error {
	pubSub := lock.redisClient.Subscribe(getPubSubTopic(lock.key))
	_, err := pubSub.Receive()
	if err != nil {
		return err
	}
	<-pubSub.Channel()
	return nil
}

// subscribeLock 阻塞直到锁被释放或超时
func (lock *RedisLock) subscribeLockWithTimeout(d time.Duration) error {
	timeNow := time.Now()
	pubSub := lock.redisClient.Subscribe(getPubSubTopic(lock.key))
	_, err := pubSub.ReceiveTimeout(d)
	if err != nil {
		return err
	}
	deltaTime := time.Since(timeNow) - d
	select {
	case <-pubSub.Channel():
		return nil
	case <-time.After(deltaTime):
		return fmt.Errorf("timeout")
	}
}

// publishLock 发布关于锁被释放的消息
func (lock *RedisLock) publishLock() error {
	err := lock.redisClient.Publish(getPubSubTopic(lock.key), "release lock").Err()
	if err != nil {
		return err
	}
	return nil
}

// renew 更新锁的到期时间，并且可以在调用 Unlock 时取消
func (lock *RedisLock) renew(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(lock.expiration / 3)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				lock.redisClient.Expire(lock.key, lock.expiration).Result()
			}
		}
	}()
}

// getPubSubTopic key -> PubSubPrefix + key
func getPubSubTopic(key string) string {
	return PubSubPrefix + key
}
