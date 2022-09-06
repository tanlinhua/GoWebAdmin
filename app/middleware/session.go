package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/app/config"
	"github.com/tanlinhua/go-web-admin/app/model"
	"github.com/tanlinhua/go-web-admin/pkg/response"
	"github.com/tanlinhua/go-web-admin/pkg/utils"
)

// 登录鉴权
func CheckSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		// 1. 校验登录是否超时
		timeOutCtr := true                         //false：无权限控制，true：启用超时·权限控制
		timeOut := int64(2 * 3600)                 //无操作超时时间:N*小时
		nowTime := time.Now().Unix()               //当前时间
		loginTime := session.Get("adminLoginTime") //登录时间
		if loginTime == nil {
			loginTime = int64(0)
		}
		calcTime := nowTime - loginTime.(int64) //登录时间差
		if calcTime > timeOut && timeOutCtr {
			timeOutHandler(c)
			return
		} else {
			session.Set("adminLoginTime", time.Now().Unix())
			session.Save()
		}

		// 2.校验后台操作权限
		adminId := session.Get("adminId")
		if utils.Empty(adminId) {
			timeOutHandler(c)
			return
		}
		if adminId != config.AdminId { //管理员ID不受权限约束
			ok, msg := checkAdminPermission(adminId.(int), c.Request.RequestURI, c.Request.Method)
			if !ok {
				response.New(c).Error(-1, msg)
				c.Abort()
				return
			}
		}

		// 3.校验用户状态
		if err := model.AdminStatusCheck(adminId.(int)); err != nil {
			response.New(c).Error(-1, err.Error())
			c.Abort()
			return
		}

		c.Set("admin_id", adminId)
		c.Set("adminName", session.Get("adminName"))
		c.Next()
	}
}

// 验证管理用户是否有该uri操作权限
func checkAdminPermission(adminId int, uri string, method string) (bool, string) {
	newUri := utils.Strstr(uri, "?", true)

	// 兼容vue通过角色id获取权限ids
	if find := strings.Contains(uri, "admin/role/get?id="); find {
		return true, "SUCCESS"
	}

	ok := model.PerCheck(adminId, newUri, method) //校验权限
	if ok {
		return true, "SUCCESS"
	}
	return false, "权限不足"
}

func timeOutHandler(c *gin.Context) {
	if c.Request.RequestURI == "/admin/main" {
		c.Redirect(http.StatusFound, "/admin/login")
	} else {
		c.String(http.StatusBadRequest, "权限不足! 😎")
	}
	c.Abort()
}
