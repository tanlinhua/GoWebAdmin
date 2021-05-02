package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/model"
	"github.com/tanlinhua/go-web-admin/pkg/response"
)

// 权限页面
func PermissionView(c *gin.Context) {
	c.HTML(http.StatusOK, "permission/index.html", nil)
}

// 查询权限
func PermissionGet(c *gin.Context) {
	resp := response.New(c)

	//查询指定角色的权限tree
	roleId, err := strconv.Atoi(c.Query("roleid"))
	if err == nil {
		ok, pData := model.PerMenuDataByRoleId(roleId)
		if ok {
			resp.Success(pData, 0)
		} else {
			resp.Error(-1, "fail")
		}
		return
	}
	//查询权限list
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	search := c.Query("search")

	datas, total := model.PermissionGet(page, limit, search)

	resp.Success(datas, total)
}
