package middleware

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 登录鉴权
func CheckSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		timeOutCtr := true                         //false：无权限控制，true：启用超时·权限控制
		timeOut := int64(24 * 3600)                // 无操作超时时间:N*小时
		nowTime := time.Now().Unix()               //当前时间
		loginTime := session.Get("adminLoginTime") //登录时间
		if loginTime == nil {
			loginTime = int64(0)
		}
		calcTime := nowTime - loginTime.(int64) //登录时间差

		if calcTime > timeOut && timeOutCtr {
			c.Redirect(http.StatusFound, "login")
		} else {
			session.Set("adminLoginTime", time.Now().Unix())
			session.Save()
		}
	}
}
