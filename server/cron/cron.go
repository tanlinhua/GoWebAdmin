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
// Cron表达式几个简单范例：
//      每隔5秒执行一次：*/5 * * * * ?
//      每隔1分钟执行一次：0 */1 * * * ?
//      每天23点执行一次：0 0 23 * * ?
//      每天凌晨1点执行一次：0 0 1 * * ?
//      每月1号凌晨1点执行一次：0 0 1 1 * ?
//      每月最后一天23点执行一次：0 0 23 L * ?
//      每周星期天凌晨1点实行一次：0 0 1 ? * L
//      在26分、29分、33分执行一次：0 26,29,33 * * * ?
//      每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?
