package main

import (
	"fmt"
	"runtime"

	"github.com/tanlinhua/go-web-admin/model"
	"github.com/tanlinhua/go-web-admin/pkg/config"
	"github.com/tanlinhua/go-web-admin/router"
	"github.com/tanlinhua/go-web-admin/service/cron"
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
	go cron.Run()
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
	fmt.Printf("AppName:\t%s\n", AppName)
	fmt.Printf("AppVersion:\t%s\n", AppVersion)
	fmt.Printf("AppMode:\t%s\n", config.AppMode)
	fmt.Printf("GoVersion:\t%s\n", GoVersion)
	fmt.Printf("ApiDoc:\t\t%s\n", ApiDoc)
	fmt.Printf("\n\n")
}
