package model

import (
	"errors"

	"github.com/tanlinhua/go-web-admin/pkg/trace"
	"github.com/tanlinhua/go-web-admin/pkg/validator"
)

// 角色模型
type Role struct {
	Id       int    `json:"id" form:"id"`
	RoleName string `json:"role_name" form:"role_name" validate:"required,min=2,max=40" label:"角色名称"`
	RoleDesc string `json:"role_desc" form:"role_desc" validate:"required,min=2,max=40" label:"角色描述"`
	PerId    string `json:"per_id" form:"per_id"`
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
		return ""
	}
	return result.PerId
}

// 增
func RoleAdd(data *Role) error {
	ok, msg := validator.Validate(data)
	if !ok {
		return errors.New(msg)
	}
	return db.Create(data).Error
}

// 删
func RoleDel(id int) (bool, string) {
	return false, "出于系统考虑,暂不支持后台删除角色,如需删除请联系IT人员."
	// err := db.Delete(&Role{}, id).Error
	// if err != nil {
	// 	return false, err.Error()
	// }
	// return true, "删除成功"
}

// 改
func RoleUpdate(data *Role) (bool, string) {
	ok, msg := validator.Validate(data)
	if !ok {
		return ok, msg
	}

	err := db.Save(data).Error
	if err != nil {
		return false, err.Error()
	}
	return true, "success"
}

// 查
func RoleGet(page, limit, id int, search string) (*[]Role, int64) {
	var total int64
	var data []Role
	Db := db

	if id > 0 {
		Db = Db.Where("id=?", id)
	}
	if len(search) > 0 {
		Db = Db.Where("`role_name` LIKE ?", "%"+search+"%")
	}

	Db.Model(&Role{}).Count(&total) //1.查询总数

	if page > 0 && limit > 0 {
		Db = Db.Limit(limit).Offset((page - 1) * limit)
	}

	err := Db.Find(&data).Error //2.查询数据
	if err != nil {
		trace.Error("RoleGet.err:" + err.Error())
	}

	return &data, total
}
