package route

import (
	"os"

	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/app/config"
	"github.com/tanlinhua/go-web-admin/app/controller/api"
	"github.com/tanlinhua/go-web-admin/app/middleware"
	_ "github.com/tanlinhua/go-web-admin/docs"
	"github.com/tanlinhua/go-web-admin/pkg/response"
)

// 初始化API HTTP服务
func InitApiServer() {
	gin.SetMode(config.AppMode)
	engine := gin.New()
	engine.SetTrustedProxies(nil)

	engine.NoRoute(HandleNotFound)
	engine.NoMethod(HandleNotFound)

	initApiMiddleware(engine)
	initSwagger(engine)
	initApiRouterV1(engine)

	engine.Run(":" + config.ApiPort)
}

// 404
func HandleNotFound(c *gin.Context) {
	response.New(c).Error(404, "not found")
}

// 中间件
func initApiMiddleware(e *gin.Engine) {
	var xss middleware.XssMw

	e.Use(gin.Recovery())
	e.Use(middleware.Logger("api")) // 自定义日志记录&切割
	e.Use(middleware.IpLimiter())   // IP请求限制器
	e.Use(xss.RemoveXss())          // xss
}

// 初始化swagger
func initSwagger(e *gin.Engine) {
	disablingKey := "GO_API_SWAGGER_DISABLE"
	if config.SwaggerOpen != 1 {
		os.Setenv(disablingKey, "true") // 禁用swagger
	}
	// http://host:port/api/doc/index.html
	e.GET("/api/doc/*any", ginSwagger.DisablingWrapHandler(files.Handler, disablingKey))
}

// 路由配置 -> API
func initApiRouterV1(e *gin.Engine) {
	// 公共路由
	public := e.Group("/api/v1")
	{
		public.POST("user/login", api.UserLogin)
		public.POST("user/reg", api.UserRegister)

		public.POST("/upload", api.TestUpload)    // 测试!
		public.Static("upload", "runtime/upload") // test for api/v1/upload
	}
	// 鉴权路由
	auth := public
	auth.Use(middleware.CheckJWT())
	{
		auth.POST("user/cpw", api.UserCpw)
	}
}

/*
RESTful API 设计指南👇
https://www.ruanyifeng.com/blog/2014/05/restful_api.html

1.路由定义
auth.POST("/user", api.AddUser)
auth.DELETE("/user/:id", api.DeleteUser) // 单用户来调用接口可以通过jwt或session中的内容来获取
auth.PUT("/user/:id", api.UpdateUser)
auth.GET("/user", api.GetsUser)
auth.GET("/user/:id", api.GetUser)

2.参数获取 (c *gin.Context)
api参数 -> id := c.Param("id") // 获取user/:id中id的值
url参数 -> page := c.Query("page") // 获取?后面的参数
表单参数 -> c.PostForm & c.DefaultPostForm
jsonBody绑定	-> c.ShouldBindJSON
表单body绑定	-> c.Bind
uri数据绑定	-> r.GET("/:user/:password") // c.ShouldBindUri

tips: Shouldxxx和Bindxxx区别就是Bindxxx会在head中添加400的返回信息，而Shouldxxx不会
更多可以查看 c.Bind IDE的提示阅读gin源码
*/
