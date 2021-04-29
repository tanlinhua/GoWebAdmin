package model

import (
	"time"

	"github.com/tanlinhua/go-web-admin/pkg/trace"
	"github.com/tanlinhua/go-web-admin/pkg/utils"
)

// 管理者模型
type Admin struct {
	Id            int       `json:"id" validate:"numeric"`
	UserName      string    `json:"user_name" validate:"required,min=5,max=32" label:"用户名"`
	Password      string    `json:"password" validate:"required,min=6,max=64" label:"密码"`
	Role          int       `json:"role" validate:"omitempty,numeric" label:"角色ID"`
	Status        string    `json:"status" validate:"omitempty,status" label:"状态码"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	LastLoginTime time.Time `json:"last_login_time"`
	LastLoginIp   string    `json:"last_login_ip"`
}

// 保存前置操作
func (m *Admin) BeforeSave() {
	m.Password = utils.Md5(m.Password)
}

// 增

// 删

// 改

// 查

// 管理登录
func AdminLogin(user_name, password string) (bool, int) {
	result := false
	var admin Admin

	db.Select("id").Where("user_name=?", user_name).Where("password=?", utils.Md5(password)).First(&admin)
	if admin.Id > 0 {
		result = true
	}
	return result, admin.Id
}

// 记录最后登录时间及IP
func AdminLoginTimeAndIp(id int, ip string, loginTime time.Time) {
	var admin Admin
	admin.Id = id
	admin.LastLoginIp = ip
	admin.LastLoginTime = loginTime
	err := db.Model(&admin).Updates(admin).Error
	if err != nil {
		trace.Error("AdminLoginTimeAndIp.Error=" + err.Error())
	}
}

// 修改密码
func AdminCpw(adminId int, pwd1, pwd2, pwd3 string) (result bool, msg string) {
	var admin Admin

	if pwd2 != pwd3 {
		return false, "新密码与确认密码不一致"
	}
	if len(pwd2) < 6 {
		return false, "新密码不能小于6位数"
	}
	db.Select("id").Where("id=?", adminId).Where("password=?", utils.Md5(pwd1)).First(&admin)
	if admin.Id == 0 {
		return false, "原密码不正确"
	}
	err := db.Model(&Admin{}).Where("id=?", adminId).Update("password", utils.Md5(pwd2)).Error
	if err != nil {
		return false, err.Error()
	}
	return true, "修改成功"
}
