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

	result, count := model.ParamsGet(page, limit, search)

	obj := response.Success("success", count, result)
	c.JSON(http.StatusOK, obj)
}

// 修改配置数据
func ParamsUpdate(c *gin.Context) {
	var rsp map[string]interface{}

	id, _ := strconv.Atoi(c.PostForm("id"))
	value := c.PostForm("value")

	result, msg := model.ParamsUpdate(id, value)

	if result {
		rsp = response.Success(msg, 0, nil)
	} else {
		rsp = response.Error(msg, 0, nil)
	}
	c.JSON(http.StatusOK, rsp)
}

// 删除配置数据
func ParamsDelete(c *gin.Context) {
	c.JSON(http.StatusOK, response.Error("暂不支持后台删除系统配置参数", 0, nil))
}

// 新增配置数据
func ParamsAdd(c *gin.Context) {
	c.JSON(http.StatusOK, response.Error("暂不支持后台添加系统配置参数", 0, nil))
}
