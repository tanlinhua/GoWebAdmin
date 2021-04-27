package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/middleware"
)

// 用户登录
func UserLogin(c *gin.Context) {
	_, token := middleware.GetJWT("Test", "1")
	c.String(http.StatusOK, token)
}

// 用户注册
func UserRegister(c *gin.Context) {
	c.String(http.StatusOK, "UserRegister")
}

// 修改用户密码
func UserCpw(c *gin.Context) {
	c.String(http.StatusOK, "UserCpw")
}
