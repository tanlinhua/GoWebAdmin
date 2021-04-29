package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 权限页面
func PermissionView(c *gin.Context) {
	c.HTML(http.StatusOK, "permission/index.html", nil)
}

// 增加权限
func PermissionAdd(c *gin.Context) {

}

// 删除权限
func PermissionDel(c *gin.Context) {

}

// 修改权限
func PermissionUpdate(c *gin.Context) {

}

// 查询权限
func PermissionGet(c *gin.Context) {

}
