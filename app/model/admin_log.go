package model

import (
	"time"

	"github.com/tanlinhua/go-web-admin/app/config"
	"github.com/tanlinhua/go-web-admin/pkg/trace"
)

// 管理员日志模型
type AdminLog struct {
	Id        int        `json:"id"`
	Uid       int        `json:"uid"`
	Uri       string     `json:"uri"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Ip        string     `json:"ip"`
	CreatedAt TimeNormal `json:"created_at"`
}

// 增
func (s *AdminLog) Add() {
	err := db.Create(s).Error
	if err != nil {
		trace.Error("model.admin_log.Add.err=" + err.Error())
	}
}

type AdminLogResult struct {
	AdminLog
	UserName string `json:"user_name"`
}

// 查
func AdminLogGet(adminId, page, limit int, title, name, ip, startTime, endTime string) (*[]AdminLogResult, int64) {
	var total int64
	var data []AdminLogResult
	Db := db

	Db = Db.Model(&AdminLog{})

	if adminId != config.AdminId {
		Db = Db.Where("go_admin_log.uid=? or go_admin.pid=?", adminId, adminId) // 非超级管理只能查看自己或下级的操作日志
	}
	if len(title) > 0 {
		Db = Db.Where("`title` LIKE ?", "%"+title+"%")
	}
	if len(name) > 0 {
		Db = Db.Where("go_admin.user_name = ?", name)
	}
	if len(ip) > 0 {
		Db = Db.Where("`ip` = ?", ip)
	}
	if len(startTime) > 0 && len(endTime) > 0 {
		Db = Db.Where("go_admin_log.created_at BETWEEN ? AND ?", startTime, endTime)
	}

	Db = Db.Select("go_admin_log.id, user_name, uri, title, body, ip, go_admin_log.created_at")
	Db = Db.Joins("left join go_admin on go_admin_log.uid=go_admin.id")
	Db = Db.Order("go_admin_log.id desc")

	Db.Count(&total)

	if page > 0 && limit > 0 {
		Db = Db.Limit(limit).Offset((page - 1) * limit)
	}

	err := Db.Scan(&data).Error
	if err != nil {
		trace.Error("AdminGet.err:" + err.Error())
	}

	return &data, total
}

// 清理记录 beforeDay(-7) 天之前的数据
func AdminLogClean(beforeDay int) {
	before := time.Now().AddDate(0, 0, beforeDay).Format("2006-01-02 15:04:05")
	trace.Info("AdminLogClean.before=" + before)
	err := db.Where("created_at < ?", before).Delete(&AdminLog{}).Error
	if err != nil {
		trace.Error("AdminLogClean.err=" + err.Error())
	}
}
