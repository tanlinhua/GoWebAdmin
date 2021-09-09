package model

import (
	"fmt"
	"log"
	"time"

	"github.com/tanlinhua/go-web-admin/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

// 初始化数据库
func InitDB() {
	connect := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DbUser, config.DbPassWord, config.DbHost, config.DbPort, config.DbName)

	var sqlLogger logger.Interface = nil
	if config.AppMode == "debug" {
		sqlLogger = logger.Default.LogMode(logger.Info) // logger.Warn 只打印慢查询
	}

	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       connect, // DSN data source name
		DefaultStringSize:         256,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: sqlLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "go_", // 表名前缀，`User`表为`t_users`
			SingularTable: true,  // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
		// 为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）
		// 如果没有这方面的要求，您可以在初始化时禁用它，这将获得大约 30%+ 性能提升
		// SkipDefaultTransaction: true, // 禁用默认事务
	})
	if err != nil {
		log.Panic("连接数据库失败，err：", err)
	}

	sqlDb, e := db.DB()
	if e != nil {
		log.Panic("获取sql.DB失败,error=", e)
	}
	sqlDb.SetMaxIdleConns(10)                  // 设置连接池中的最大闲置连接数
	sqlDb.SetMaxOpenConns(200)                 // 设置数据库的最大连接数量
	sqlDb.SetConnMaxLifetime(10 * time.Second) // 设置连接的最大可复用时间
}
