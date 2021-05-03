package model

import (
	"strconv"

	"github.com/tanlinhua/go-web-admin/pkg/trace"
	"github.com/tanlinhua/go-web-admin/pkg/utils"
)

// 权限模型
type Permission struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Pid    int    `json:"pid"`
	Uri    string `json:"uri"`
	Method string `json:"method"`
	Level  int    `json:"level"`
}

// 权限菜单
type PerData struct {
	Id       int       `json:"id"`
	Name     string    `json:"title"`
	Checked  bool      `json:"checked"`
	Spread   bool      `json:"spread"`
	Children []PerData `json:"children"`
}

// 获取后台用户权限内的菜单数据
func PerMenuDataByAdmId(adminId int) {
	//获取角色ID
	//获取该角色拥有的菜单权限 PerMenuDataByRoleId ↓
	//赋值给页面进行{{range $i, $v := .slice}} {{end}}
}

// 获取指定角色ID的菜单数据
func PerMenuDataByRoleId(roleId int) (bool, *[]PerData) {
	var menu []PerData

	// 查询角色ids
	ids := RoleGetPerIdsByRoleId(roleId)
	idsArr := utils.Explode(",", ids)

	err := db.Model(&Permission{}).Select("id,name").Where("pid=?", 0).Scan(&menu).Error
	if err != nil {
		return false, nil
	}
	for index1, item1 := range menu {
		if find := utils.In_array(strconv.Itoa(item1.Id), idsArr); find {
			menu[index1].Checked = true
		}
	}
	for index2, item2 := range menu {
		err := db.Model(&Permission{}).Select("id,name").Where("pid=?", item2.Id).Scan(&menu[index2].Children).Error
		if err != nil {
			return false, nil
		}
		for index3, item3 := range menu[index2].Children {
			menu[index2].Checked = false
			menu[index2].Spread = true
			if find := utils.In_array(strconv.Itoa(item3.Id), idsArr); find {
				menu[index2].Children[index3].Checked = true
			}
		}
	}
	return true, &menu
}

// 校验权限
func PerCheck(adminId int, uri string, method string) bool {
	ids := RoleGetPerIdsByAdminId(adminId) //获取角色ID所拥有的权限ids
	pid := PerIdByUriMethod(uri, method)   //根据uri及method获取对应权限id

	idsArr := utils.Explode(",", ids)
	ok := utils.In_array(strconv.Itoa(pid), idsArr) //判断权限id是否存在ids中
	if ok {
		return true
	} else {
		return false
	}
}

// 获取权限id
func PerIdByUriMethod(uri, method string) int {
	var per Permission
	err := db.Model(&per).Select("id").Where("uri=?", uri).Where("method=?", method).Find(&per).Error
	if err != nil {
		trace.Error("PerIdByUriMethod.Error=" + err.Error())
	}
	return per.Id
}

// 增
// 删
// 改
// 查
func PermissionGet(page, limit int, search string) (*[]Permission, int) {
	var total int
	var data []Permission
	Db := db

	Db.Model(&Permission{}).Count(&total) //1.查询总数

	if len(search) > 0 {
		Db = Db.Where("`name` LIKE ?", "%"+search+"%")
	}

	if page > 0 && limit > 0 {
		Db = Db.Limit(limit).Offset((page - 1) * limit)
	}

	err := Db.Find(&data).Error //2.查询数据
	if err != nil {
		trace.Error("Permission.err:" + err.Error())
	}

	return &data, total
}
