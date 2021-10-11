package main

import (
	"fmt"
	"runtime"

	"github.com/tanlinhua/go-web-admin/app/config"
	"github.com/tanlinhua/go-web-admin/app/model"
	"github.com/tanlinhua/go-web-admin/app/service/cron"
	"github.com/tanlinhua/go-web-admin/route"
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

	go route.InitAdmServer()
	go route.InitApiServer()

	cron.Run()
}

var (
	AppName    = "GoWebAdmin"
	AppVersion = "1.0"
	GoVersion  = runtime.Version()
	ApiDoc     = "http://" + config.AppHost + ":" + config.ApiPort + "/api/doc/index.html"
	Pprof      = "http://" + config.AppHost + ":" + config.AdminPort + "/admin/jason/pprof"
)

func versionInfo() {
	fmt.Printf("AppName:\t%s\n", AppName)
	fmt.Printf("AppVersion:\t%s\n", AppVersion)
	fmt.Printf("AppMode:\t%s\n", config.AppMode)
	fmt.Printf("GoVersion:\t%s\n", GoVersion)
	fmt.Printf("ApiDoc:\t\t%s\n", ApiDoc)
	fmt.Printf("pprof:\t\t%s\n", Pprof)
	fmt.Printf("\n\n")
}
