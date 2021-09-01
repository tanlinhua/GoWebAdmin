package model

import (
	"strconv"

	"github.com/tanlinhua/go-web-admin/app/config"
	"github.com/tanlinhua/go-web-admin/pkg/trace"
	"github.com/tanlinhua/go-web-admin/pkg/utils"
)

// 权限模型
type Permission struct {
	Id       int          `json:"id"`
	Name     string       `json:"title"` // title for layui
	Pid      int          `json:"pid"`
	Uri      string       `json:"uri"`
	Method   string       `json:"method"`
	Icon     string       `json:"icon"`
	Level    int          `json:"level"`
	Checked  bool         `json:"checked"`           // tree是否选中,也意味着是否有权限
	Spread   bool         `json:"spread"`            // tree是否展开
	Children []Permission `json:"children" gorm:"-"` // 子菜单
}

// 后台首页,获取后台用户权限内的菜单数据
func PerMenuDataByAdmId(adminId int) []Permission {
	roleId := AdminGetRoleIdByAdmId(adminId) //获取角色ID
	if -1 == roleId {
		return nil
	}
	ok, data := PerMenuDataByRoleId(roleId)
	if !ok {
		return nil
	}

	// 处理后台菜单显示
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

// 获取指定角色ID的菜单TREE数据
func PerMenuDataByRoleId(roleId int) (bool, []Permission) {
	var root []Permission
	var lowLevel int = 3

	ids := utils.Explode(",", RoleGetPerIdsByRoleId(roleId)) // 查询角色ids
	db.Find(&root)
	db.Model(&Permission{}).Select(`max(level)`).Scan(&lowLevel)
	tree := getTreeMenu(root, 0, ids, lowLevel)
	return true, tree
}

// GET TREE DATA
func getTreeMenu(menuList []Permission, pid int, ids []string, lowLevel int) []Permission {
	var tree []Permission
	for _, v := range menuList {
		if v.Pid == pid {
			checked := false
			if v.Level == lowLevel {
				if utils.In_array(strconv.Itoa(v.Id), ids) {
					checked = true
				}
			}
			node := Permission{
				Id:      v.Id,
				Name:    v.Name,
				Pid:     v.Pid,
				Uri:     v.Uri,
				Method:  v.Method,
				Icon:    v.Icon,
				Level:   v.Level,
				Checked: checked,
				Spread:  true,
			}
			node.Children = getTreeMenu(menuList, v.Id, ids, lowLevel)
			tree = append(tree, node)
		}
	}
	return tree
}

// 校验权限
func PerCheck(adminId int, uri string, method string) bool {
	ids := RoleGetPerIdsByAdminId(adminId) // 获取角色ID所拥有的权限ids
	pid := PerIdByUriMethod(uri, method)   // 根据uri及method获取对应权限id

	idsArr := utils.Explode(",", ids)
	ok := utils.In_array(strconv.Itoa(pid), idsArr) // 判断权限id是否存在ids中
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
func PermissionGet(page, limit int, search string) (*[]Permission, int64) {
	var total int64
	var data []Permission
	Db := db

	if len(search) > 0 {
		Db = Db.Where("`name` LIKE ?", "%"+search+"%")
	}

	Db.Model(&Permission{}).Count(&total) // 1.查询总数

	if page > 0 && limit > 0 {
		Db = Db.Limit(limit).Offset((page - 1) * limit)
	}

	err := Db.Find(&data).Error // 2.查询数据
	if err != nil {
		trace.Error("Permission.err:" + err.Error())
	}

	return &data, total
}
