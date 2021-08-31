package admin

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/app/config"
	"github.com/tanlinhua/go-web-admin/app/model"
	"github.com/tanlinhua/go-web-admin/pkg/captcha"
	"github.com/tanlinhua/go-web-admin/pkg/google"
	"github.com/tanlinhua/go-web-admin/pkg/response"
	"github.com/tanlinhua/go-web-admin/pkg/utils"
)

// 登录页面
func AdminLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "main/login.html", gin.H{"LoginAuth": config.LoginAuth})
}

// 后台首页
func AdminMain(c *gin.Context) {
	adminName, _ := c.Get("adminName")
	adminId, _ := c.Get("admin_id")
	menuData := model.PerMenuDataByAdmId(adminId.(int))

	c.HTML(http.StatusOK, "main/main.html", gin.H{"adminName": adminName, "menuData": menuData})
}

// 控制台页面
func AdminConsole(c *gin.Context) {
	//根据角色ID,查询所属预览数据展示到页面
	c.HTML(http.StatusOK, "main/console.html", nil)
}

// 验证码
type CaptchaResult struct {
	Id     string `json:"id"`
	Base64 string `json:"base64"`
}

type googleAuth struct {
	Secret    string `json:"secret"`
	QrCodeUrl string `json:"qrCodeUrl"`
}

// 生成google authenticator信息并存入数据库
func GenGoogleAuth(c *gin.Context) {
	var save model.SysParams
	resp := response.New(c)

	value := model.ParamsGetValueByKey("GoogleAuthenticator")
	if !utils.Empty(value) {
		response.New(c).Error(-1, "请勿重复请求")
		return
	}

	secret := google.NewGoogleAuth().GetSecret()
	url := google.NewGoogleAuth().GetQrcodeUrl("Go.Admin.Auth", secret)
	data := googleAuth{Secret: secret, QrCodeUrl: url}
	saveValue, _ := utils.Json_encode(data)

	save.Type = 0
	save.Key = "GoogleAuthenticator"
	save.Value = saveValue
	save.Remarks = "Google身份验证器"

	err := save.Add()
	if err != nil {
		resp.Error(-1, err.Error())
	}
	resp.Success(nil, 0)
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

// 退出登录
func AdminLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "login")
}

// 校验管理员用户名密码
func AdminLoginCheck(c *gin.Context) {
	capt_ok := false
	resp := response.New(c)

	user_name := c.PostForm("user_name")
	password := c.PostForm("password")
	code := c.PostForm("code")
	cid := c.PostForm("cid")

	if utils.Empty(user_name) || utils.Empty(password) || utils.Empty(code) || utils.Empty(cid) {
		resp.Error(-1, "检查输入")
		return
	}

	if config.LoginAuth == 1 {
		value := model.ParamsGetValueByKey("GoogleAuthenticator")
		if utils.Empty(value) {
			resp.Error(-1, "GoogleAuthenticator信息不存在")
			return
		}
		authJson, _ := utils.Json_decode(value)
		secret := authJson["secret"]
		capt_ok = google.NewGoogleAuth().VerifyCode(secret.(string), code)
	} else {
		capt_ok = captcha.CaptchaVerify(cid, code)
	}

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
