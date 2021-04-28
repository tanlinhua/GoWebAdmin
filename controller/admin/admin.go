package admin

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/model"
	"github.com/tanlinhua/go-web-admin/pkg/response"
)

// 登录页面
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "main/login.html", nil)
}

// 退出登录
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "login")
}

// 校验管理员用户名密码
func Check(c *gin.Context) {
	user_name := c.PostForm("user_name")
	password := c.PostForm("password")
	captcha := c.PostForm("captcha")
	fmt.Println(user_name, password, captcha)

	r, id := model.Login(user_name, password)
	if r {
		session := sessions.Default(c)
		session.Set("adminLoginTime", time.Now().Unix())
		session.Set("adminId", id)
		session.Save()
		response.New(c).Success(nil, 0)
		return
	}
	response.New(c).Error(-1, "fail")
}

// 修改密码
func Cpw(c *gin.Context) {
	pwd1 := c.PostForm("pwd1")
	pwd2 := c.PostForm("pwd2")
	pwd3 := c.PostForm("pwd3")
	fmt.Println(pwd1, pwd2, pwd3)

	adminId := sessions.Default(c).Get("adminId")

	r, msg := model.Cpw(adminId.(int), pwd1, pwd2, pwd3)
	if r {
		response.New(c).Success(nil, 0)
		return
	}
	response.New(c).Error(-1, msg)
}

// 后台首页
func Main(c *gin.Context) {
	c.HTML(http.StatusOK, "main/main.html", nil)
}

// 控制台页面
func Console(c *gin.Context) {
	//预览数据展示到页面
	c.HTML(http.StatusOK, "main/console.html", nil)
}
