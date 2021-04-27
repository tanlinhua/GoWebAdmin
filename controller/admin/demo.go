package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 控制台页面
func Form(c *gin.Context) {
	c.HTML(http.StatusOK, "demo/form.html", nil)
}

// 用户组
func Users(c *gin.Context) {
	c.HTML(http.StatusOK, "demo/users.html", nil)
}

// 权限页面
func Operaterule(c *gin.Context) {
	c.HTML(http.StatusOK, "demo/operaterule.html", nil)
}
