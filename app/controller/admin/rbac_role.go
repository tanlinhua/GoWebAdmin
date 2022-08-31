package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/app/model"
	"github.com/tanlinhua/go-web-admin/pkg/response"
)

// 角色页面
func RoleView(c *gin.Context) {
	c.HTML(http.StatusOK, "rbac/role.html", nil)
}

// 增加角色
func RoleAdd(c *gin.Context) {
	resp := response.New(c)
	var role model.Role
	if err := c.Bind(&role); err != nil {
		resp.Error(-1, err.Error())
		return
	}
	if err := model.RoleAdd(&role); err != nil {
		resp.Error(-1, err.Error())
		return
	}
	resp.Success(nil, 0)
}

// 删除角色
func RoleDel(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	ok, msg := model.RoleDel(id)
	if ok {
		response.New(c).Success(nil, 0)
	} else {
		response.New(c).Error(-1, msg)
	}
}

// 修改角色
func RoleUpdate(c *gin.Context) {
	resp := response.New(c)
	var role model.Role

	err := c.Bind(&role)
	if err != nil {
		resp.Error(-1, err.Error())
		return
	}
	ok, msg := model.RoleUpdate(&role)
	if ok {
		resp.Success(nil, 0)
	} else {
		resp.Error(-1, msg)
	}
}

// 查询角色
func RoleGet(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Query("id")) //查询指定角色的权限tree
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	search := c.Query("search")

	datas, total := model.RoleGet(page, limit, roleId, search)

	response.New(c).Success(datas, total)
}
