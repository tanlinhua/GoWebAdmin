package cron

import (
	"fmt"
	"time"
)

var specClean = "0 0 1 * * ?"    // 每天凌晨1点执行一次
var specTest1 = "*/30 * * * * ?" // 每分钟的第30秒执行一次
var specTest2 = "*/60 * * * * ?" // 每分钟的第60秒执行一次

func cleanLog() {
	fmt.Println("每天凌晨1点执行一次", time.Now())
}

func test1() {
	fmt.Println("每分钟的第30秒执行一次!", time.Now())
}

func test2() {
	fmt.Println("每分钟的第60秒执行一次!", time.Now())
}
