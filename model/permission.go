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
type PerMenuData struct {
}

// 获取权限内的菜单数据
func PerMenuDataByAdmId(adminId int) {
	//获取角色ID,获取该角色拥有的菜单权限,赋值给页面进行{{range $i, $v := .slice}} {{end}}
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
