package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/config"
	"github.com/tanlinhua/go-web-admin/controller/api"
	"github.com/tanlinhua/go-web-admin/middleware"
	"github.com/tanlinhua/go-web-admin/pkg/response"
)

// 初始化API HTTP服务
func InitApiServer() {
	gin.SetMode(config.AppMode)
	engine := gin.New()

	engine.NoRoute(HandleNotFound)
	engine.NoMethod(HandleNotFound)

	initApiMiddleware(engine)
	initApiRouter(engine)

	engine.Run(config.APIPort)
}

// 404
func HandleNotFound(c *gin.Context) {
	response.New(c).Error(404, "资源未找到")
}

// 中间件
func initApiMiddleware(e *gin.Engine) {
	e.Use(gin.Recovery())
	e.Use(middleware.Logger("api"))
}

// 路由配置 -> API
func initApiRouter(e *gin.Engine) {
	e.GET("api/user/login", api.UserLogin)
	e.GET("api/user/reg", api.UserRegister)

	auth := e.Group("/api")
	auth.Use(middleware.CheckJWT())
	{
		auth.POST("user/cpw", api.UserCpw)
	}
}
