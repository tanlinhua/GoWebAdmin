package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/config"
	"github.com/tanlinhua/go-web-admin/controller/api"
	"github.com/tanlinhua/go-web-admin/middleware"
)

func InitApiServer() {
	gin.SetMode(config.AppMode)
	engine := gin.New()

	initApiMiddleware(engine)
	initApiRouter(engine)

	engine.Run(config.APIPort)
}

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
