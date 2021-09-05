package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/app/model"
	"github.com/tanlinhua/go-web-admin/pkg/response"
)

// 后台用户页面
func AdminLogView(c *gin.Context) {
	c.HTML(http.StatusOK, "system/adminlog.html", nil)
}

// 查询操作日志
func AdminLogGet(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	title := c.Query("title")
	name := c.Query("name")
	ip := c.Query("ip")
	startTime := c.Query("t1")
	endTime := c.Query("t2")

	adminId, _ := c.Get("admin_id")
	datas, total := model.AdminLogGet(adminId.(int), page, limit, title, name, ip, startTime, endTime)

	response.New(c).Success(datas, total)
}
