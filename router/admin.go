package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/config"
	"github.com/tanlinhua/go-web-admin/controller/admin"
	"github.com/tanlinhua/go-web-admin/middleware"
)

// 初始化Admin HTTP服务
func InitAdmServer() {
	gin.SetMode(config.AppMode)
	engine := gin.New()

	initAdmMiddleware(engine)
	initAdmResources(engine)
	initAdmRouter(engine)

	engine.Run(config.AdminPort)
}

// 中间件
func initAdmMiddleware(e *gin.Engine) {
	e.Use(gin.Recovery())                        // 如果存在恐慌(panics)，中间件恢复(recovers)写入500
	e.Use(middleware.Logger("admin"))            // 自定义日志记录&切割
	store := cookie.NewStore([]byte("secret"))   // sessionStore-cookie存储
	e.Use(sessions.Sessions("mysession", store)) // session
}

// 静态资源
func initAdmResources(e *gin.Engine) {
	e.Static("assets", "view/static/assets")
	e.StaticFile("favicon.ico", "view/static/favicon.ico")
	e.LoadHTMLGlob("view/admin/**/*")
}

// 路由配置 -> ADMIN
func initAdmRouter(e *gin.Engine) {
	e.GET("admin/login", admin.AdminLogin)       //登录页面
	e.GET("admin/captcha", admin.Captcha)        //获取图形验证码
	e.POST("admin/check", admin.AdminLoginCheck) //登录校验
	e.GET("admin/logout", admin.AdminLogout)     //退出登录

	auth := e.Group("/admin")
	auth.Use(middleware.CheckSession())
	{
		auth.GET("main", admin.AdminMain)       //后台首页
		auth.GET("console", admin.AdminConsole) //首页控制台
		auth.POST("cpw", admin.AdminCpw)        //修改密码

		// 参数配置
		auth.GET("params/view", admin.ParamsView)      //view
		auth.POST("params/add", admin.ParamsAdd)       //增
		auth.GET("params/del", admin.ParamsDelete)     //删
		auth.POST("params/update", admin.ParamsUpdate) //改
		auth.GET("params/list", admin.ParamsGet)       //查

		// 角色管理

		// 权限管理

		// demo 👇
		auth.GET("demo1", admin.Form)
		auth.GET("demo2", admin.Users)
		auth.GET("demo3", admin.Operaterule)
	}
}
