package route

import (
	"os"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/tanlinhua/go-web-admin/docs"

	"github.com/gin-gonic/gin"

	"github.com/tanlinhua/go-web-admin/app/config"
	"github.com/tanlinhua/go-web-admin/app/controller/api"
	"github.com/tanlinhua/go-web-admin/app/middleware"
	"github.com/tanlinhua/go-web-admin/pkg/response"
)

// 初始化API HTTP服务
func InitApiServer() {
	gin.SetMode(config.AppMode)
	engine := gin.New()

	engine.NoRoute(HandleNotFound)
	engine.NoMethod(HandleNotFound)

	initApiMiddleware(engine)
	initSwagger(engine)
	initApiRouter(engine)

	engine.Run(":" + config.APIPort)
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
	if config.SwaggerIsOpen != 1 {
		os.Setenv(disablingKey, "true") // 禁用swagger
	}
	// http://host:port/api/doc/index.html
	e.GET("/api/doc/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, disablingKey))
}

// 路由配置 -> API
func initApiRouter(e *gin.Engine) {
	e.POST("api/user/login", api.UserLogin)
	e.POST("api/user/reg", api.UserRegister)

	e.POST("api/test/upload", api.TestUpload)

	auth := e.Group("/api")
	auth.Use(middleware.CheckJWT())
	{
		auth.POST("user/cpw", api.UserCpw)
	}
}
