package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/app/model"
	"github.com/tanlinhua/go-web-admin/pkg/response"
)

// 后台用户页面
func AdmView(c *gin.Context) {
	c.HTML(http.StatusOK, "rbac/user.html", nil)
}

// 增加后台用户
func AdmAdd(c *gin.Context) {
	resp := response.New(c)
	var admin model.Admin

	err := c.Bind(&admin)
	if err != nil {
		resp.Error(-1, err.Error())
		return
	}
	adminId, _ := c.Get("admin_id")

	if err := model.AdmAdd(adminId.(int), &admin); err != nil {
		resp.Error(-1, err.Error())
		return
	}
	resp.Success(nil, 0)
}

// 删除后台用户
func AdmDel(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	if err := model.AdminDel(id); err != nil {
		response.New(c).Error(-1, err.Error())
		return
	}
	response.New(c).Success(nil, 0)
}

// 修改后台用户
func AdmUpdate(c *gin.Context) {
	resp := response.New(c)
	var admin model.Admin

	err := c.Bind(&admin)
	if err != nil {
		resp.Error(-1, err.Error())
		return
	}
	if err := model.AdmUpdate(&admin); err != nil {
		resp.Error(-1, err.Error())
	} else {
		resp.Success(nil, 0)
	}
}

// 查询后台用户
func AdmGet(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	search := c.Query("search")
	role := c.Query("role")
	startTime := c.Query("t1")
	endTime := c.Query("t2")

	adminId, _ := c.Get("admin_id")
	datas, total := model.AdminGet(adminId.(int), page, limit, search, role, startTime, endTime)

	response.New(c).Success(datas, total)
}
