package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/model"
	"github.com/tanlinhua/go-web-admin/pkg/response"
)

// 参数配置-view
func ParamsView(c *gin.Context) {
	c.HTML(http.StatusOK, "params/index.html", nil)
}

// 查询配置数据
func ParamsGet(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	search := c.Query("search")

	datas, total := model.ParamsGet(page, limit, search)

	response.New(c).Success(datas, total)
}

// 修改配置数据
func ParamsUpdate(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	value := c.PostForm("value")

	r, msg := model.ParamsUpdate(id, value)

	if r {
		response.New(c).Success(nil, 0)
	} else {
		response.New(c).Error(-1, msg)
	}
}

// 删除配置数据
func ParamsDelete(c *gin.Context) {
	response.New(c).Error(-1, "暂不支持后台删除系统配置参数")
}

// 新增配置数据
func ParamsAdd(c *gin.Context) {
	response.New(c).Error(-1, "暂不支持后台添加系统配置参数")
}
