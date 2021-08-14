package cron

import (
	"fmt"
	"time"
)

var specClean = "0 0 1 * * ?"    // 每天凌晨1点执行一次
var specTest1 = "*/20 * * * * ?" // 每N秒执行一次
var specTest2 = "0 */2 * * * ?"  // 每N分钟执行一次

func cleanLog() {
	fmt.Println("每天凌晨1点执行一次", specClean, time.Now())
}

func test1() {
	fmt.Println("每N秒执行一次!", specTest1, time.Now())
}

func test2() {
	fmt.Println("每N分钟执行一次!", specTest2, time.Now())
}
