package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/pkg/middleware"
)

// 用户登录
func UserLogin(c *gin.Context) {
	u := c.Query("u")
	id, _ := strconv.Atoi(c.Query("id"))
	_, token := middleware.GetJWT(u, id)
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
