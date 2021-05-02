package model

import "github.com/tanlinhua/go-web-admin/pkg/trace"

// 角色模型
type Role struct {
	Id       int    `json:"id"`
	RoleName string `json:"role_name"`
	RoleDesc string `json:"role_desc"`
	PerId    string `json:"per_id"`
}

// 根据adminId获取PerId
func RoleGetPerIdsByAdminId(admId int) string {
	type result struct{ PerId string }
	var r result
	err := db.Table("go_role").Select("go_role.per_id").
		Joins("left join go_admin on go_admin.role=go_role.id").
		Where("go_admin.id=?", admId).Scan(&r).Error
	if err != nil {
		trace.Error("RoleGetPerIdsByAdminId.Error=" + err.Error())
	}
	return r.PerId
}

// 根据roleId获取PerId
func RoleGetPerIdsByRoleId(roleId int) string {
	var result Role
	err := db.Where("id=?", roleId).Find(&result).Error
	if err != nil {
		trace.Error("RoleGetPerIdsByRoleId.Error=" + err.Error())
		return "1"
	}
	return result.PerId
}

// 增

// 删

// 改

// 查
func RoleGet(page, limit int, search string) (*[]Role, int) {
	var total int
	var data []Role
	Db := db

	Db.Model(&Role{}).Count(&total) //1.查询总数

	if len(search) > 0 {
		Db = Db.Where("`role_name` LIKE ?", "%"+search+"%")
	}

	if page > 0 && limit > 0 {
		Db = Db.Limit(limit).Offset((page - 1) * limit)
	}

	err := Db.Find(&data).Error //2.查询数据
	if err != nil {
		trace.Error("RoleGet.err:" + err.Error())
	}

	return &data, total
}
