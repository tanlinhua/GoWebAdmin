package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/config"
	"github.com/tanlinhua/go-web-admin/controller/admin"
	"github.com/tanlinhua/go-web-admin/middleware"
)

// 初始化HTTP服务
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
	e.GET("admin/login", admin.Login)   //登录页面
	e.POST("admin/check", admin.Check)  //登录校验
	e.GET("admin/logout", admin.Logout) //退出登录

	auth := e.Group("/admin")
	auth.Use(middleware.CheckSession())
	{
		auth.GET("main", admin.Main)       //后台首页
		auth.GET("console", admin.Console) //首页控制台
		auth.POST("cpw", admin.Cpw)        //修改密码

		auth.GET("params/view", admin.ParamsView)      //参数配置 - view
		auth.POST("params/add", admin.ParamsAdd)       //增
		auth.GET("params/del", admin.ParamsDelete)     //删
		auth.POST("params/update", admin.ParamsUpdate) //改
		auth.GET("params/list", admin.ParamsGet)       //查

		//demo 👇
		auth.GET("form", admin.Form)
		auth.GET("users", admin.Users)
		auth.GET("operaterule", admin.Operaterule)
	}
}
