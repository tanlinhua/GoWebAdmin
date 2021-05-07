package model

import (
	"strconv"

	"github.com/tanlinhua/go-web-admin/pkg/config"
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
	Uri      string    `json:"uri"`
	Icon     string    `json:"icon"`
	Checked  bool      `json:"checked"`  //tree是否选中(也意味着是否有权限)
	Spread   bool      `json:"spread"`   //tree是否展开
	Children []PerData `json:"children"` //子菜单
}

// 权限tree
type TreeList struct {
	Id       int       `json:"id"`
	Name     string    `json:"title"`
	Uri      string    `json:"uri"`
	Icon     string    `json:"icon"`
	Checked  bool      `json:"checked"`
	Spread   bool      `json:"spread"`
	Children []PerData `json:"children"`
}

// 获取后台用户权限内的菜单数据
func PerMenuDataByAdmId(adminId int) []TreeList {
	roleId := AdminGetRoleIdByAdmId(adminId) //获取角色ID
	if -1 == roleId {
		return nil
	}
	ok, data := PerMenuDataByRoleId(roleId)
	if !ok {
		return nil
	}
	for index1, item1 := range data {
		for index2, item2 := range item1.Children {
			for _, item3 := range item2.Children {
				if adminId == config.AdminId {
					data[index1].Checked = true
					data[index1].Children[index2].Checked = true
				} else if item3.Checked {
					data[index1].Checked = true
					data[index1].Children[index2].Checked = true
				}
			}
		}
	}
	return data
}

// 获取指定角色ID的菜单数据
func PerMenuDataByRoleId(roleId int) (bool, []TreeList) {
	var tree []TreeList

	// 查询角色ids
	ids := RoleGetPerIdsByRoleId(roleId)
	idsArr := utils.Explode(",", ids)

	err := db.Model(&Permission{}).Select("id,name,uri,icon").Where("pid=?", 0).Scan(&tree).Error // 一级
	if err != nil {
		return false, nil
	}
	for index1, item1 := range tree {
		if find := utils.In_array(strconv.Itoa(item1.Id), idsArr); find {
			tree[index1].Checked = true
		}
	}
	for idx1, itm1 := range tree {
		err := db.Model(&Permission{}).Select("id,name,uri,icon").Where("pid=?", itm1.Id).Scan(&tree[idx1].Children).Error // 二级
		if err != nil {
			return false, nil
		}
		for idx2, itm2 := range tree[idx1].Children {
			tree[idx1].Checked = false
			tree[idx1].Spread = true
			if find := utils.In_array(strconv.Itoa(itm2.Id), idsArr); find {
				tree[idx1].Children[idx2].Checked = true
			}

			err := db.Model(&Permission{}).Select("id,name,uri,icon").Where("pid=?", itm2.Id).Scan(&tree[idx1].Children[idx2].Children).Error // 三级
			if err != nil {
				return false, nil
			}
			for idx3, itm3 := range tree[idx1].Children[idx2].Children {
				tree[idx1].Children[idx2].Checked = false
				tree[idx1].Children[idx2].Spread = true
				if find := utils.In_array(strconv.Itoa(itm3.Id), idsArr); find {
					tree[idx1].Children[idx2].Children[idx3].Checked = true
				}
			}
		}
	}
	return true, tree
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
func PermissionAdd() {}

// 删
func PermissionDel() {}

// 改
func PermissionUpdate() {}

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
