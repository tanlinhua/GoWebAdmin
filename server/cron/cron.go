package cron

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func Work() {
	c := cron.New()

	spec1 := "*/30 * * * * ?" //cron表达式，每30秒一次
	c.AddFunc(spec1, func() {
		fmt.Println("30s", time.Now())
	})

	spec3 := "*/60 * * * * ?" //cron表达式，每60秒一次
	c.AddFunc(spec3, func() {
		fmt.Println("60s", time.Now())
	})

	c.Start()
	select {} //阻塞主线程停止
}

// https://www.cnblogs.com/dubinyang/p/12327675.html
