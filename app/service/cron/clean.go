package cron

import (
	"github.com/tanlinhua/go-web-admin/app/model"
	"github.com/tanlinhua/go-web-admin/pkg/trace"
)

var SpecSecondLoop = "*/30 * * * * ?"  // 每N秒执行一次
var SpecMinuteLoop = "0 */5 * * * ?"   // 每N分钟执行一次
var SpecHourLoop = "0 0 */2 * * ?"     // 每N小时执行一次
var SpecCleanTrashData = "0 0 1 * * ?" // 每天凌晨1点执行一次

func test2() {
	trace.Debug("CRON.每N分钟执行一次!")
}

func test3() {
	trace.Debug("CRON.每N小时执行一次!")
}

// 清理冗余数据
func cleanTrashData() {
	trace.Info("清理冗余数据")
	model.AdminLogClean(-7)
}
