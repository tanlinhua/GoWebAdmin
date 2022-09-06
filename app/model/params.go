package model

import (
	"github.com/tanlinhua/go-web-admin/pkg/trace"
)

// 系统配置模型
type SysParams struct {
	Id      int    `json:"id"`
	Type    int    `json:"type"`
	Key     string `json:"key"`
	Value   string `json:"value"`
	Remarks string `json:"remarks"`
}

// 获取系统配置数据
func ParamsGet(page, limit int, search string) (*[]SysParams, int64) {
	var total int64
	var data []SysParams
	Db := db

	Db.Model(&SysParams{}).Count(&total) //1.查询总数

	if len(search) > 0 {
		Db = Db.Where("`key` LIKE ?", "%"+search+"%")
	}

	if page > 0 && limit > 0 {
		Db = Db.Limit(limit).Offset((page - 1) * limit)
	}
	Db = Db.Where("type=?", 1)

	err := Db.Find(&data).Error //2.查询数据
	if err != nil {
		trace.Error("ParamsGet.err:" + err.Error())
	}

	return &data, total
}

// 修改系统配置数据
func ParamsUpdate(id int, value string) error {
	return db.Model(&SysParams{}).Where("id=?", id).Update("value", value).Error
}

// 修改系统配置数据
func ParamsUpdateByKey(key string, value string) error {
	return db.Model(&SysParams{}).Where("`key`=?", key).Update("value", value).Error
}

// 增加系统配置
func (s *SysParams) Add() error {
	return db.Create(s).Error
}

// 查询指定key的value值
func ParamsGetValueByKey(key string) string {
	var row SysParams
	err := db.Model(&SysParams{}).Select("value").Where("`key`=?", key).Scan(&row).Error
	if err != nil {
		trace.Error("ParamsGetValueByKey.err=" + err.Error() + ",key=" + key)
	}
	return row.Value
}
