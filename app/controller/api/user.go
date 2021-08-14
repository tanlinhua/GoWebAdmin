package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/app/middleware"
	"github.com/tanlinhua/go-web-admin/app/model"
	"github.com/tanlinhua/go-web-admin/pkg/response"
)

// @Tags 用户模块
// @Summary 用户登录
// @accept application/json
// @Produce application/json
// @Param json_data body model.User true "用户名/密码"
// @Response 200 {object} response.ResultData "code=0成功,否则失败,data.token为jwt"
// @Router /api/user/login [post]
func UserLogin(c *gin.Context) {
	rsp := response.New(c)

	var user model.User

	if err := c.BindJSON(&user); err != nil {
		rsp.Error(-1, err.Error())
		return
	}

	uId, msg := user.Login()
	if uId <= 0 {
		rsp.Error(-1, msg)
		return
	}

	ok, content := middleware.GetJWT(user.UserName, uId)
	if ok {
		rsp.Success(gin.H{"token": content}, 0)
	} else {
		rsp.Error(-1, content)
	}
}

// @Tags 用户模块
// @Summary 用户注册
// @accept application/json
// @Produce application/json
// @Param json_data body model.User true "用户名/密码/手机号/设备类型/设备型号 ..."
// @Response 200 {object} response.ResultData "code=0成功,否则失败"
// @Router /api/user/reg [post]
func UserRegister(c *gin.Context) {
	c.String(http.StatusOK, "UserRegister")
}

// @Tags 用户模块
// @Summary 修改用户密码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param json_data body model.User true "当前密码/新密码/确认密码"
// @Response 200 {object} response.ResultData "code=0成功,否则失败"
// @Router /api/user/cpw [post]
func UserCpw(c *gin.Context) {
	//jwt获取用户ID
	c.String(http.StatusOK, "UserCpw")
}
