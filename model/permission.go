package model

import "fmt"

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
	fmt.Println("adminId =", adminId, ",uri =", uri, ",method =", method)
	//临时测试
	if uri == "/admin/console" || uri == "/admin/main" {
		return true
	}
	return false
}

// 增

// 删

// 改

// 查
