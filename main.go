package main

import (
	"fmt"
	"runtime"

	"github.com/tanlinhua/go-web-admin/model"
	"github.com/tanlinhua/go-web-admin/pkg/config"
	"github.com/tanlinhua/go-web-admin/router"
	"github.com/tanlinhua/go-web-admin/server/cron"
)

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
)

func versionInfo() {
	fmt.Printf("AppName:\t%s\n", AppName)
	fmt.Printf("AppVersion:\t%s\n", AppVersion)
	fmt.Printf("AppMode:\t%s\n", config.AppMode)
	fmt.Printf("GoVersion:\t%s\n", GoVersion)
}
