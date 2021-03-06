package cron

import (
	"github.com/robfig/cron/v3"
)

func Run() {
	c := cron.New(cron.WithSeconds())

	c.AddFunc(SpecMinuteLoop, test2)
	c.AddFunc(SpecHourLoop, test3)
	c.AddFunc(SpecCleanTrashData, cleanTrashData)

	c.Start()
	select {} //阻塞主线程停止
}

// 1.Cron表达式几个简单范例：
//      每隔5秒执行一次：*/5 * * * * ?
//      每隔1分钟执行一次：0 */1 * * * ?
//      每天23点执行一次：0 0 23 * * ?
//      每天凌晨1点执行一次：0 0 1 * * ?
//      每月1号凌晨1点执行一次：0 0 1 1 * ?
//      每月最后一天23点执行一次：0 0 23 L * ?
//      每周星期天凌晨1点实行一次：0 0 1 ? * L
//      在26分、29分、33分执行一次：0 26,29,33 * * * ?
//      每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?

// 2.更多应用场景自行搜索了解: cron表达式
