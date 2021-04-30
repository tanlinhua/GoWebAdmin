package admin

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/model"
	"github.com/tanlinhua/go-web-admin/pkg/captcha"
	"github.com/tanlinhua/go-web-admin/pkg/response"
	"github.com/tanlinhua/go-web-admin/pkg/utils"
)

// Admin模型

// 验证码
type CaptchaResult struct {
	Id     string `json:"id"`
	Base64 string `json:"base64"`
}

// 登录页面
func AdminLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "main/login.html", nil)
}

// 生成图形验证码
func Captcha(c *gin.Context) {
	resp := response.New(c)
	id, b64, err := captcha.CaptchaMake()
	if err != nil {
		resp.Error(-1, err.Error())
		return
	}
	capt := CaptchaResult{Id: id, Base64: b64}
	resp.Success(capt, 0)
}

// 后台首页
func AdminMain(c *gin.Context) {
	adminName, _ := c.Get("adminName")
	admin_id, _ := c.Get("admin_id")
	model.PerMenuDataByAdmId(admin_id.(int))
	c.HTML(http.StatusOK, "main/main.html", gin.H{"adminName": adminName})
}

// 控制台页面
func AdminConsole(c *gin.Context) {
	//根据角色ID,查询所属预览数据展示到页面
	c.HTML(http.StatusOK, "main/console.html", nil)
}

// 退出登录
func AdminLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "login")
}

// 校验管理员用户名密码
func AdminLoginCheck(c *gin.Context) {
	resp := response.New(c)

	user_name := c.PostForm("user_name")
	password := c.PostForm("password")
	code := c.PostForm("code")
	cid := c.PostForm("cid")

	if utils.Empty(user_name) || utils.Empty(password) || utils.Empty(code) || utils.Empty(cid) {
		resp.Error(-1, "检查输入")
		return
	}

	capt_ok := captcha.CaptchaVerify(cid, code)
	if !capt_ok {
		resp.Error(-1, "验证码错误")
		return
	}

	ok, id, msg := model.AdminLogin(user_name, password)
	if ok {
		model.AdminLoginTimeAndIp(id, c.ClientIP(), time.Now()) //记录最后登录时间及IP
		session := sessions.Default(c)
		session.Set("adminLoginTime", time.Now().Unix())
		session.Set("adminName", user_name)
		session.Set("adminId", id)
		session.Save()
		resp.Success(nil, 0)
		return
	}
	resp.Error(-1, msg)
}

// 修改密码
func AdminCpw(c *gin.Context) {
	pwd1 := c.PostForm("pwd1")
	pwd2 := c.PostForm("pwd2")
	pwd3 := c.PostForm("pwd3")

	adminId := sessions.Default(c).Get("adminId")

	r, msg := model.AdminCpw(adminId.(int), pwd1, pwd2, pwd3)
	if r {
		response.New(c).Success(nil, 0)
		return
	}
	response.New(c).Error(-1, msg)
}

// 后台用户页面
func AdmView(c *gin.Context) {
	c.HTML(http.StatusOK, "adm/index.html", nil)
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
	ok, msg := model.AdmAdd(&admin)
	if ok {
		resp.Success(nil, 0)
	} else {
		resp.Error(-1, msg)
	}
}

// 删除后台用户
func AdmDel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	ok, msg := model.AdminDel(id)
	if ok {
		response.New(c).Success(nil, 0)
	} else {
		response.New(c).Error(-1, msg)
	}
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
	ok, msg := model.AdmUpdate(&admin)
	if ok {
		resp.Success(nil, 0)
	} else {
		resp.Error(-1, msg)
	}
}

// 查询后台用户
func AdmGet(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	search := c.Query("search")

	datas, total := model.AdminGet(page, limit, search)

	response.New(c).Success(datas, total)
}
