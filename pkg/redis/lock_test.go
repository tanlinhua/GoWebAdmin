package redis

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// var redisClient = redis.NewClient(&redis.Options{
// 	Addr:     "localhost:6379",
// 	Password: "", // no password set
// 	DB:       0,  // use default DB
// })

var redisClient = Handler

func TestRedisLock_TryLock(t *testing.T) {
	timeNow := time.Now()
	lock := NewRedisLock(redisClient.Client, "test-try-lock", "myLock")
	wg := new(sync.WaitGroup)
	wg.Add(50)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			success, err := lock.TryLock()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			if !success {
				fmt.Println("TryLock Fail")
			} else {
				defer func() {
					lock.Unlock()
					fmt.Println("release the lock")
				}()
				fmt.Println("TryLock Success")
			}
		}()
	}
	wg.Wait()
	deltaTime := time.Since(timeNow)
	fmt.Println(deltaTime)
}

func TestRedisLock_Lock(t *testing.T) {
	timeNow := time.Now()
	lock := NewRedisLock(redisClient.Client, "test-lock", "myLock")
	wg := new(sync.WaitGroup)
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			err := lock.Lock()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("get lock")
			defer func() {
				lock.Unlock()
				fmt.Println("release the lock")
			}()

			time.Sleep(time.Second * 2)
			defer wg.Done()
		}()
	}
	wg.Wait()
	deltaTime := time.Since(timeNow)
	fmt.Println(deltaTime)
}

func TestRedisLock_LockWithTimeout(t *testing.T) {
	timeNow := time.Now()
	lock := NewRedisLock(redisClient.Client, "test-lock-with-timeout", "myLock")
	wg := new(sync.WaitGroup)
	wg.Add(20)

	for i := 0; i < 20; i++ {
		go func() {
			err := lock.LockWithTimeout(3 * time.Second)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("get lock")
			defer func() {
				lock.Unlock()
				fmt.Println("release the lock")
			}()
			time.Sleep(time.Second * 4)
			defer wg.Done()
		}()
	}
	wg.Wait()
	deltaTime := time.Since(timeNow)
	fmt.Println(deltaTime)
}

func TestRedisLock_SpinLock(t *testing.T) {
	timeNow := time.Now()
	lock := NewRedisLock(redisClient.Client, "test-spain-lock", "myLock")
	wg := new(sync.WaitGroup)
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			err := lock.SpinLock(5)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("get lock")
			defer lock.Unlock()
			defer func() {
				fmt.Println("release the lock")
			}()
			defer wg.Done()
		}()
	}
	wg.Wait()
	deltaTime := time.Since(timeNow)
	fmt.Println(deltaTime)
}
