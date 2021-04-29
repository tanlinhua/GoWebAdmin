package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

}
