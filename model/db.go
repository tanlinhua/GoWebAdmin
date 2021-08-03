package model

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tanlinhua/go-web-admin/service/config"
)

var db *gorm.DB

// 初始化数据库
func InitDB() {
	connect := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DbUser, config.DbPassWord, config.DbHost, config.DbPort, config.DbName)

	var err error
	db, err = gorm.Open(config.Db, connect)
	if err != nil {
		log.Panic("连接数据库失败，err：", err)
	}

	logMode := false
	if config.AppMode == "debug" {
		logMode = true
	}

	db.LogMode(logMode)                          //是否打印sql日志
	db.SingularTable(true)                       // 禁用默认表名的复数形式
	db.DB().SetMaxIdleConns(10)                  // SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxOpenConns(100)                 // SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetConnMaxLifetime(10 * time.Second) // SetConnMaxLifetiment 设置连接的最大可复用时间。

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "go_" + defaultTableName
	} //指定表前缀
}
