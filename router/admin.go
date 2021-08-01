package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/controller/admin"
	"github.com/tanlinhua/go-web-admin/pkg/config"
	"github.com/tanlinhua/go-web-admin/service/middleware"
)

// 初始化Admin HTTP服务
func InitAdmServer() {
	gin.SetMode(config.AppMode)
	engine := gin.New()

	initAdmMiddleware(engine)
	initAdmResources(engine)
	initAdmRouter(engine)

	engine.Run(":" + config.AdminPort)
}

// 中间件
func initAdmMiddleware(e *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))                            // sessionStore-cookie存储
	store.Options(sessions.Options{MaxAge: 0, Path: "/", HttpOnly: true}) //cookie相关设置,MaxAge=0关闭浏览器则会话结束

	var xss middleware.XssMw

	e.Use(gin.Recovery())                     // 如果存在恐慌(panics)，中间件恢复(recovers)写入500
	e.Use(middleware.Logger("admin"))         // 自定义日志记录&切割
	e.Use(middleware.IpLimiter())             // IP请求限制器
	e.Use(xss.RemoveXss())                    // xss
	e.Use(sessions.Sessions("cookie", store)) // session
}

// 静态资源
func initAdmResources(e *gin.Engine) {
	e.Static("assets", "view/static/assets")
	e.StaticFile("favicon.ico", "view/static/favicon.ico")
	e.LoadHTMLGlob("view/admin/**/*")
}

// 路由配置 -> ADMIN
func initAdmRouter(e *gin.Engine) {
	e.GET("admin/google/auth", admin.GenGoogleAuth) //生成googleAuth信息
	e.GET("admin/captcha", admin.Captcha)           //获取图形验证码

	e.GET("admin/login", admin.AdminLogin)       //登录页面
	e.POST("admin/check", admin.AdminLoginCheck) //登录校验
	e.GET("admin/logout", admin.AdminLogout)     //退出登录

	auth := e.Group("/admin")
	auth.Use(middleware.CheckSession())
	{
		// 后台首页
		auth.GET("main", admin.AdminMain)       //view
		auth.GET("console", admin.AdminConsole) //控制台
		auth.POST("cpw", admin.AdminCpw)        //修改密码

		// 后台用户管理
		auth.GET("adm/view", admin.AdmView)      //view
		auth.POST("adm/add", admin.AdmAdd)       //增
		auth.GET("adm/del", admin.AdmDel)        //删
		auth.POST("adm/update", admin.AdmUpdate) //改
		auth.GET("adm/get", admin.AdmGet)        //查

		// 角色管理
		auth.GET("role/view", admin.RoleView)      //view
		auth.POST("role/add", admin.RoleAdd)       //增
		auth.GET("role/del", admin.RoleDel)        //删
		auth.POST("role/update", admin.RoleUpdate) //改
		auth.GET("role/get", admin.RoleGet)        //查

		// 权限管理
		auth.GET("per/view", admin.PermissionView) //view
		auth.GET("per/get", admin.PermissionGet)   //查

		// 参数配置
		auth.GET("params/view", admin.ParamsView)      //view
		auth.POST("params/add", admin.ParamsAdd)       //增
		auth.GET("params/del", admin.ParamsDelete)     //删
		auth.POST("params/update", admin.ParamsUpdate) //改
		auth.GET("params/get", admin.ParamsGet)        //查
	}
}
