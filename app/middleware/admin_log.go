package middleware

import (
	"bytes"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/app/model"
	"github.com/tanlinhua/go-web-admin/pkg/utils"
)

// 不写入管理员日志中的uri,比如: 不包含上传相关
var blackUri = []string{
	"/admin/xxx/upload",
}

// 判断uri是否记录日志
func isBlackUri(uri string) bool {
	return utils.In_array(uri, blackUri)
}

// 管理员日志
func AdminLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		if isBlackUri(c.Request.RequestURI) {
			return
		}

		body, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		c.Next()

		// 只记录POST
		if c.Request.Method == "POST" {
			var adminLog model.AdminLog
			session := sessions.Default(c)
			admin_id := session.Get("adminId")

			if utils.Empty(admin_id) {
				adminLog.Uid = 0
			} else {
				adminLog.Uid = admin_id.(int)
			}
			adminLog.Uri = c.Request.RequestURI
			adminLog.Title = getAdminLogTitle(c.Request.RequestURI)
			adminLog.Body, _ = url.QueryUnescape(string(body))
			adminLog.Ip = c.ClientIP()
			adminLog.Add()
		}
	}
}

func getAdminLogTitle(uri string) string {
	// 特殊路由
	if find := strings.Contains(uri, "admin/check"); find {
		return "登录后台"
	}
	if find := strings.Contains(uri, "admin/cpw"); find {
		return "修改登录密码"
	}

	var action string = "action"
	if find := strings.Contains(uri, "add"); find {
		action = "添加"
	} else if find := strings.Contains(uri, "update"); find {
		action = "修改"
	} else if find := strings.Contains(uri, "del"); find {
		action = "删除"
	}

	var model string = "model"
	if find := strings.Contains(uri, "params"); find {
		model = "系统参数"
	} else if find := strings.Contains(uri, "role"); find {
		model = "角色"
	} else if find := strings.Contains(uri, "adm"); find { // 这是个坑 admin/ 包含adm,所以放最后吧
		model = "用户"
	}

	return action + "-" + model
}
