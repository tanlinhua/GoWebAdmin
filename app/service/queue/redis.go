package queue

import (
	"fmt"
	"time"

	"github.com/tanlinhua/go-web-admin/pkg/redis"
	"github.com/tanlinhua/go-web-admin/pkg/utils"
	"github.com/tidwall/gjson"
)

type CallBack func(*RedisQueue)

type RedisQueue struct {
	Name     string      `json:"name"`     // 队列名称
	Data     interface{} `json:"data"`     // 数据传递
	Interval int64       `json:"interval"` // 间隔时间
	Fire     CallBack    `json:"-"`        // 回调函数
}

func (t *RedisQueue) Push() error {
	queue := "Queue:Name:" + t.Name
	runing := "Queue:Running:" + t.Name

	if tmp, err := utils.Json_encode(t); err == nil {
		if err := redis.Handler.ListAdd(queue, tmp); err != nil {
			return fmt.Errorf("添加RedisList失败: %v", err.Error())
		}
	} else {
		return fmt.Errorf("转json失败: %v", err.Error())
	}

	if redis.Handler.SGet(runing) != "true" {
		go t.worker(queue, runing)
	}
	return nil
}

func (t *RedisQueue) worker(queue, runing string) {
	redis.Handler.SSet(runing, "true", 0)
	defer redis.Handler.SSet(runing, "false", 0)

	for {
		time.Sleep(time.Duration(t.Interval) * time.Second)

		task := redis.Handler.ListGet(queue)
		if utils.Empty(task) {
			break
		}
		t.Data = gjson.Get(task, "data").Value()
		t.Fire(t)
	}
}

/*
// Test👇

func cb(t *RedisQueue) {
	fmt.Printf("cb.t --->%+v\n", t)
}

func Test(key string) {
	var t RedisQueue

	t.Name = key
	t.Fire = cb
	t.Interval = 1

	datas := []string{`{"k1","v1"}`, `{"k2","v2"}`, `{"k3","v3"}`}

	for index, item := range datas {
		t.Data = item
		err := t.Push()
		fmt.Println("Test.Push --->", index, err)
	}
}

queue.Test("Q1")
queue.Test("Q2")
time.Sleep(time.Second * 30) // 正常情况主程序不会退出,所以也不需要sleep
return
*/
