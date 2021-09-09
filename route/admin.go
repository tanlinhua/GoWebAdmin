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
	e.Use(xss.RemoveXss())                    // xss
	e.Use(sessions.Sessions("cookie", store)) // session
	e.Use(middleware.AdminLog())              // 管理员操作日志
}

// 静态资源
func initAdmResources(e *gin.Engine) {
	// e.Static("assets", "view/static/assets")
	// e.StaticFile("favicon.ico", "view/static/favicon.ico")
	// e.LoadHTMLGlob("view/admin/**/*")

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
}

// 路由配置 -> ADMIN
func initAdmRouter(e *gin.Engine) {
	// 公共路由
	public := e.Group("/admin")
	{
		public.GET("login", admin.AdminLogin)       // 登录页面
		public.POST("check", admin.AdminLoginCheck) // 登录校验
		public.GET("logout", admin.AdminLogout)     // 退出登录
		public.GET("captcha", admin.Captcha)        // 获取图形验证码
	}
	// 鉴权路由
	auth := public
	auth.Use(middleware.CheckSession())
	{
		// other
		auth.GET("google", admin.GenGoogleAuth)       // 生成googleAuth信息
		middleware.RouteRegister(auth, "jason/pprof") // 性能分析

		// 后台首页
		auth.GET("main", admin.AdminMain)       // view
		auth.GET("console", admin.AdminConsole) // 控制台
		auth.POST("cpw", admin.AdminCpw)        // 修改密码

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
