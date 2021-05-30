package main

import (
	"fmt"
	"runtime"

	"github.com/tanlinhua/go-web-admin/model"
	"github.com/tanlinhua/go-web-admin/pkg/config"
	"github.com/tanlinhua/go-web-admin/pkg/cron"
	"github.com/tanlinhua/go-web-admin/router"
)

// @title GoWeb
// @version 1.0
// @description golang web template.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	versionInfo()

	model.InitDB()
	go cron.Work()
	go router.InitAdmServer()
	router.InitApiServer()
}

var (
	AppName    = "GoWebAdmin"
	AppVersion = "1.0"
	GoVersion  = runtime.Version()
	ApiDoc     = "http://host:port/api/doc/index.html"
)

func versionInfo() {
	fmt.Printf("%c[1;40;35mAppName:\t%s %c[0m\n", 0x1B, AppName, 0x1B)
	fmt.Printf("%c[1;40;35mAppVersion:\t%s %c[0m\n", 0x1B, AppVersion, 0x1B)
	fmt.Printf("%c[1;40;35mAppMode:\t%s %c[0m\n", 0x1B, config.AppMode, 0x1B)
	fmt.Printf("%c[1;40;35mGoVersion:\t%s %c[0m\n", 0x1B, GoVersion, 0x1B)
	fmt.Printf("%c[1;40;35mApiDocs:\t%s %c[0m\n", 0x1B, ApiDoc, 0x1B)
}
