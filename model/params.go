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
func ParamsGet(page, limit int, search string) (*[]SysParams, int) {
	var total int
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
		trace.Error("GetParamsList.err:" + err.Error())
	}

	return &data, total
}

// 修改系统配置数据
func ParamsUpdate(id int, value string) (bool, string) {
	err := db.Model(&SysParams{}).Where("id=?", id).Update("value", value).Error
	if err != nil {
		return false, err.Error()
	}
	return true, "修改成功"
}

// 增加系统配置
func ParamsAdd(key, value, remarks string) (bool, string) {

	return false, "fail"
}
