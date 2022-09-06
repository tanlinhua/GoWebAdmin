package model

import (
	"errors"
	"time"

	"github.com/tanlinhua/go-web-admin/app/config"
	"github.com/tanlinhua/go-web-admin/pkg/trace"
	"github.com/tanlinhua/go-web-admin/pkg/utils"
	"github.com/tanlinhua/go-web-admin/pkg/validator"
)

// 管理者模型
type Admin struct {
	Id            int        `json:"id" form:"id" validate:"numeric"`
	UserName      string     `json:"user_name" form:"user_name" validate:"required,min=5,max=32" label:"用户名"`
	Password      string     `json:"password" form:"password" validate:"required,min=6,max=64" label:"密码"`
	Role          int        `json:"role" form:"role" validate:"required,numeric" label:"角色"`
	Pid           int        `json:"pid" form:"pid" validate:"numeric" label:"上级ID"`
	Status        int        `json:"status" form:"status" validate:"required,status" label:"状态"`
	CreatedAt     TimeNormal `json:"created_at" form:"created_at"`
	UpdatedAt     TimeNormal `json:"updated_at" form:"updated_at"`
	LastLoginTime TimeNormal `json:"last_login_time" form:"last_login_time"`
	LastLoginIp   string     `json:"last_login_ip" form:"last_login_ip"`
}

// 增
func AdmAdd(adminId int, data *Admin) error {
	if err := validator.Validate(data); err != nil {
		return err
	}
	exist := AdmExist(data.UserName)
	if exist {
		return errors.New("用户名已存在")
	}
	data.Password = utils.Md5(data.Password)
	data.LastLoginTime = TimeNormal{time.Now()}
	// 如果不是超级管理员新建用户,上级ID只允许是他自己
	if adminId != config.AdminId {
		data.Pid = adminId
	}
	return db.Create(data).Error
}

// 查询用户名是否已存在
func AdmExist(user_name string) bool {
	var tmp Admin
	db.Model(&Admin{}).Select("id").Where("user_name=?", user_name).First(&tmp)
	if tmp.Id > 0 {
		return true
	} else {
		return false
	}
}

// 删
func AdminDel(id int) error {
	if id == config.AdminId {
		return errors.New("超级管理员不允许删除")
	}
	return db.Delete(&Admin{}, id).Error
}

// 改
func AdmUpdate(data *Admin) error {
	var update = make(map[string]interface{})
	if !utils.Empty(data.Password) {
		update["password"] = utils.Md5(data.Password)
	}
	if !utils.Empty(data.Role) {
		update["role"] = data.Role
	}
	if !utils.Empty(data.Pid) {
		update["pid"] = data.Pid
	}
	if data.Status == 0 || data.Status == 1 {
		update["status"] = data.Status //fuck,存在0值必须用map替代
	}
	return db.Model(&Admin{}).Where("id=?", data.Id).Updates(update).Error
}

type AdminGetResult struct {
	Admin
	RoleName string `json:"role_name"`
}

// 查
func AdminGet(adminId, page, limit int, search, role string, startTime, endTime string) (*[]AdminGetResult, int64) {
	var total int64
	var data []AdminGetResult
	Db := db

	Db = Db.Model(&Admin{}).Where("role!=?", 0) //0为内置超级管理员

	if adminId != config.AdminId {
		var find []Admin
		if err := Db.Where("pid=?", adminId).Select("id").Scan(&find).Error; err != nil {
			trace.Error("model.AdminGet.FindChilErr=" + err.Error())
		}
		if len(find) > 0 {
			Db = Db.Where("pid=? or go_admin.id=?", adminId, adminId) // 存在下级用户
		} else {
			Db = Db.Where("go_admin.id=?", adminId) // 不存在下级用户,查出自己
		}
	}

	if len(search) > 0 {
		Db = Db.Where("`user_name` LIKE ?", "%"+search+"%")
	}
	if len(role) > 0 {
		Db = Db.Where("role=?", role)
	}
	if len(startTime) > 0 && len(endTime) > 0 {
		Db = Db.Where("last_login_time BETWEEN ? AND ?", startTime, endTime)
	}

	Db.Count(&total) //1.查询总数

	if page > 0 && limit > 0 {
		Db = Db.Limit(limit).Offset((page - 1) * limit)
	}

	Db = Db.Select("go_admin.id,user_name,role,pid,role_name,created_at,updated_at,last_login_time,last_login_ip,status")
	Db = Db.Joins("left join go_role on go_admin.role=go_role.id")
	Db = Db.Order("go_admin.id asc")

	err := Db.Scan(&data).Error //2.查询数据
	if err != nil {
		trace.Error("AdminGet.err:" + err.Error())
	}

	return &data, total
}

// 管理登录
func AdminLogin(user_name, password string) (bool, int, int, string) {
	result := false
	msg := "用户名或密码错误"
	var admin Admin

	db.Select("id,role,status").Where("user_name=?", user_name).Where("password=?", utils.Md5(password)).First(&admin)
	if admin.Id > 0 {
		result = true
		msg = "登录成功"
		if admin.Status == 0 {
			result = false
			msg = "状态错误"
		}
	}
	return result, admin.Id, admin.Role, msg
}

// 记录最后登录时间及IP
func AdminLoginTimeAndIp(id int, ip string, loginTime time.Time) {
	var admin Admin
	admin.Id = id
	admin.LastLoginIp = ip
	admin.LastLoginTime = TimeNormal{loginTime}
	err := db.Model(&admin).Updates(admin).Error
	if err != nil {
		trace.Error("AdminLoginTimeAndIp.Error=" + err.Error())
	}
}

// 修改密码
func AdminCpw(adminId int, pwd1, pwd2, pwd3 string) error {
	var admin Admin

	if pwd2 != pwd3 {
		return errors.New("新密码与确认密码不一致")
	}
	if len(pwd2) < 6 {
		return errors.New("新密码不能小于6位数")
	}
	db.Select("id").Where("id=?", adminId).Where("password=?", utils.Md5(pwd1)).First(&admin)
	if admin.Id == 0 {
		return errors.New("原密码不正确")
	}
	return db.Model(&Admin{}).Where("id=?", adminId).Update("password", utils.Md5(pwd2)).Error
}

// 检测用户状态
func AdminStatusCheck(id int) error {
	var admin Admin
	if err := db.Select("status").Where("id=?", id).Find(&admin).Error; err != nil {
		return err
	}
	if admin.Status == 1 {
		return nil
	} else {
		return errors.New("账户已禁用")
	}
}

// 根据admId获取roleId
func AdminGetRoleIdByAdmId(admId int) int {
	var admin Admin
	err := db.Select("role").Where("id=?", admId).Find(&admin).Error
	if err != nil {
		return -1
	}
	return admin.Role
}
