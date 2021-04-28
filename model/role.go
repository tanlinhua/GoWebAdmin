package model

// 角色模型
type Role struct {
	Id       int    `json:"id"`
	RoleName string `json:"role_name"`
	RoleDesc string `json:"role_desc"`
	PerId    string `json:"per_id"`
}

// 根据adminId得到角色ID
func RoleGetRoleIdByAdminId(admId int) int {
	return 0
}

// 增

// 删

// 改

// 查
