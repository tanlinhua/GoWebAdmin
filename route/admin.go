package route

import (
	"html/template"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/app/config"
	"github.com/tanlinhua/go-web-admin/app/controller/admin"
	"github.com/tanlinhua/go-web-admin/app/middleware"
	"github.com/tanlinhua/go-web-admin/app/view"
	"github.com/tanlinhua/go-web-admin/public"
)

// 初始化Admin HTTP服务
func InitAdmServer() {
	gin.SetMode(config.AppMode)
	engine := gin.New()

	// 不使用任何代理，禁用此功能Engine.SetTrustedProxies(nil)
	// 然后Context.ClientIP()将直接返回远程地址，以避免一些不必要的计算。
	engine.SetTrustedProxies(nil)

	initAdmResources(engine)
	initAdmMiddleware(engine)
	initAdmRouter(engine)

	engine.Run(":" + config.AdminPort)
}

// 中间件
func initAdmMiddleware(e *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))                            // sessionStore-cookie存储
	store.Options(sessions.Options{MaxAge: 0, Path: "/", HttpOnly: true}) // cookie相关设置,MaxAge=0关闭浏览器则会话结束

	var xss middleware.XssMw

	e.Use(gin.Recovery())                     // 如果存在恐慌(panics)，中间件恢复(recovers)写入500
	e.Use(middleware.Logger("admin"))         // 自定义日志记录&切割
	e.Use(middleware.IpLimiter())             // IP请求限制器
	e.Use(middleware.Secure())                // Secure
	e.Use(xss.RemoveXss())                    // xss
	e.Use(sessions.Sessions("cookie", store)) // session
	e.Use(middleware.AdminLog())              // 管理员操作日志
}

// 静态资源
func initAdmResources(e *gin.Engine) {
	// Layui (打包到程序中)
	tpl := template.Must(template.New("").ParseFS(view.Admin, "admin/**/*"))
	e.SetHTMLTemplate(tpl)

	s := e.Group("public")
	s.Use(middleware.StaticFileHandler()) // 静态资源缓存
	{
		s.StaticFS("", http.FS(public.Static))
	}

	e.GET("favicon.ico", func(c *gin.Context) {
		file, _ := public.Static.ReadFile("static/favicon.ico")
		c.Data(http.StatusOK, "image/x-icon", file)
	})

	// vue.1 (不打包到程序中)
	// e.Static("static/js", "vue/admin/static/js")
	// e.Static("static/css", "vue/admin/static/css")
	// e.Static("static/gif", "vue/admin/static/gif")
	// e.Static("static/png", "vue/admin/static/png")
	// e.StaticFile("/favicon.ico", "vue/admin/favicon.ico")
	// e.LoadHTMLGlob("vue/admin/index.html")
	// e.GET("admin", func(c *gin.Context) {
	// 	c.HTML(200, "index.html", nil)
	// })

	// vue.2 (todo: 打包到程序中,嵌入太多go:embed,没太大含义)
	// tpl := template.Must(template.New("").ParseFS(vue.HTML, "*")) // vue.HTML dist根目录
	// e.SetHTMLTemplate(tpl)
	// s := e.Group("static")
	// s.Use(middleware.StaticFileHandler())
	// {
	// 	s.StaticFS("", http.FS(vue.Static)) vue.HTML dist/static根目录
	// }
}

// 路由配置 -> ADMIN
func initAdmRouter(e *gin.Engine) {
	// 公共路由
	public := e.Group("/admin")
	{
		public.GET("login", admin.Login)             // 登录页面
		public.POST("check", admin.LoginCheck)       // 登录校验
		public.GET("logout", admin.Logout)           // 退出登录
		public.GET("captcha", admin.Captcha)         // 获取图形验证码
		public.GET("sys/params", admin.SystemParams) // 一些公共的系统参数
		public.GET("ga/gen", admin.GenGoogleAuth)    // 生成googleAuth信息
	}
	// 鉴权路由
	auth := public
	auth.Use(middleware.CheckSession())
	{
		// other
		middleware.RouteRegister(auth, "jason/pprof") // 性能分析

		// 后台首页-控制台
		auth.GET("main", admin.Main)                         // view
		auth.GET("console", admin.Console)                   // 控制台view
		auth.GET("console/get", admin.Dashboard)             // 查
		auth.POST("cpw", admin.Cpw)                          // 修改密码
		auth.POST("message/update", admin.UpdateMainMessage) // 修改通知消息
		auth.GET("message/get", admin.GetMainMessage)        // 获取通知消息

		// 权限配置-后台用户管理 Manager
		auth.GET("adm/view", admin.AdmView)      // view
		auth.POST("adm/add", admin.AdmAdd)       // 增
		auth.POST("adm/del", admin.AdmDel)       // 删
		auth.POST("adm/update", admin.AdmUpdate) // 改
		auth.GET("adm/get", admin.AdmGet)        // 查
		// 权限配置-角色管理
		auth.GET("role/view", admin.RoleView)      // view
		auth.POST("role/add", admin.RoleAdd)       // 增
		auth.POST("role/del", admin.RoleDel)       // 删
		auth.POST("role/update", admin.RoleUpdate) // 改
		auth.GET("role/get", admin.RoleGet)        // 查
		// 权限配置-权限管理
		auth.GET("per/view", admin.PermissionView) // view
		auth.GET("per/get", admin.PermissionGet)   // 查

		// 系统配置-参数配置
		auth.GET("params/view", admin.ParamsView)      // view
		auth.POST("params/add", admin.ParamsAdd)       // 增
		auth.POST("params/del", admin.ParamsDelete)    // 删
		auth.POST("params/update", admin.ParamsUpdate) // 改
		auth.GET("params/get", admin.ParamsGet)        // 查
		// 系统配置-操作日志
		auth.GET("adminlog/view", admin.AdminLogView) // view
		auth.GET("adminlog/get", admin.AdminLogGet)   // 查
	}
}
