package cron

import (
	"fmt"
	"time"
)

var specSecond = "*/30 * * * * ?" // 每N秒执行一次
var specMinute = "0 */5 * * * ?"  // 每N分钟执行一次
var specHour = "0 0 1 * * ?"      // 每天凌晨1点执行一次	[ "0 0 */2 * * ?" -> 每2个小时执行一次 ]

func test1() {
	fmt.Println("CRON.每N秒执行一次!", specSecond, time.Now())
}

func test2() {
	fmt.Println("CRON.每N分钟执行一次!", specMinute, time.Now())
}

func cleanLog() {
	fmt.Println("CRON.每天凌晨1点执行一次!", specHour, time.Now())
}
