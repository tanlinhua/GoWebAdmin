package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/model"
	"github.com/tanlinhua/go-web-admin/pkg/response"
)

// 角色页面
func RoleView(c *gin.Context) {
	c.HTML(http.StatusOK, "role/index.html", nil)
}

// 增加角色
func RoleAdd(c *gin.Context) {

}

// 删除角色
func RoleDel(c *gin.Context) {

}

// 修改角色
func RoleUpdate(c *gin.Context) {

}

// 查询角色
func RoleGet(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	search := c.Query("search")

	datas, total := model.RoleGet(page, limit, search)

	response.New(c).Success(datas, total)
}
