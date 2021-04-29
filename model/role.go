package model

import "github.com/tanlinhua/go-web-admin/pkg/trace"

// 角色模型
type Role struct {
	Id       int    `json:"id"`
	RoleName string `json:"role_name"`
	RoleDesc string `json:"role_desc"`
	PerId    string `json:"per_id"`
}

// 根据adminId得到PerId
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

// 增

// 删

// 改

// 查
